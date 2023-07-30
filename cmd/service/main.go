package main

import (
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/cors"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	paycli "github.com/hardstylez72/cry-pay/proto/gen/go/v1"
	"github.com/hardstylez72/cry/internal/exchange/pub"
	"github.com/hardstylez72/cry/internal/lib"
	log "github.com/hardstylez72/cry/internal/log"
	"github.com/hardstylez72/cry/internal/orbiter"
	"github.com/hardstylez72/cry/internal/pay"
	core "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/hardstylez72/cry/internal/process"
	"github.com/hardstylez72/cry/internal/server/access"
	"github.com/hardstylez72/cry/internal/server/auth"
	"github.com/hardstylez72/cry/internal/server/config"
	"github.com/hardstylez72/cry/internal/server/proxy"
	"github.com/hardstylez72/cry/internal/server/repository"
	"github.com/hardstylez72/cry/internal/server/repository/pg"
	"github.com/hardstylez72/cry/internal/server/service/v1"
	"github.com/hardstylez72/cry/internal/server/ui"
	"github.com/hardstylez72/cry/internal/settings"
	"github.com/hardstylez72/cry/internal/snapshot"
	"github.com/hardstylez72/cry/internal/tg"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const (
	project = "server"
)

type (
	services struct {
		profileService    *v1.ProfileService
		helperService     *v1.HelperService
		withdrawerService *v1.WithdrawerService
		flowService       *v1.FlowService
		processService    *v1.ProcessService
		settingsService   *v1.SettingsService
		swap1inchService  *v1.Swap1inchService
		orbiterService    *v1.OrbiterService

		processRepository    repository.ProcessRepository
		flowRepository       repository.FlowRepository
		profileRepository    repository.ProfileRepository
		settingsRepository   repository.SettingsRepository
		withdrawerRepository repository.WithdrawerRepository
		userRepository       repository.UserRepository
		statRepository       repository.StatRepository
		userSettingsService  *settings.Service

		dispatcher *process.Dispatcher

		payService *pay.Service
	}
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	l, err := log.NewLogger(project, "")
	if err != nil {
		panic(errors.Wrap(err, "log.NewLogger"))
	}
	defer l.Sync()

	s, err := initServices(ctx, cfg)
	if err != nil {
		panic(err)
	}

	fatal := make(chan error)
	go func() {
		err := ListenGRPC(ctx, cfg.GRPCAddr, s)
		if err != nil {
			fatal <- err
		}
	}()
	go func() {
		err := ListenGW(ctx, cfg, s)
		if err != nil {
			fatal <- err
		}
	}()

	if cfg.Env == config.Prod {
		go func() {
			err := ui.ListenStatic(cfg.StaticPort, "ui/build")
			if err != nil {
				fatal <- errors.Wrap(err, "ListenStatic")
			}
		}()

		go func() {
			err = proxy.ListenProxy(cfg.GWAddr, cfg.StaticPort, cfg.ProxyPort)
			if err != nil {
				fatal <- errors.Wrap(err, "startProxyServer")
			}
		}()
	}

	select {
	case err := <-fatal:
		panic(err)
	}
}

func ListenGW(ctx context.Context, cfg *config.Config, s *services) error {

	conn, err := grpc.Dial(cfg.GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption("application/base64", &CryptoMarshaller{}),
	)

	if err := core.RegisterProfileServiceHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err := core.RegisterHelperServiceHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err := core.RegisterWithdrawerServiceHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err := core.RegisterFlowServiceHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err := core.RegisterProcessServiceHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err := core.RegisterSettingsServiceHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err := core.RegisterSwap1InchServiceHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err := core.RegisterOrbiterServiceHandler(ctx, mux, conn); err != nil {
		return err
	}

	ga := auth.NewGoogleOAuth2Controller(auth.Config{
		RedirectURL:  cfg.App.Schema + "://" + cfg.App.Domain + cfg.App.Port + "/api/gw/google/oauth/callback",
		ClientID:     cfg.Auth.GoogleClientId,
		ClientSecret: cfg.Auth.GoogleClientSecret,
		Scopes:       []string{"email"},
		UserInfoURL:  "https://www.googleapis.com/oauth2/v2/userinfo",
		UserRedirects: auth.UserRedirects{
			OnSuccess: cfg.Auth.RedirectOnSuccess,
			OnFailure: cfg.Auth.RedirectOnFailure,
		},
	}, auth.SessionCookieConfig{
		Name:   cfg.Auth.CookieName,
		Domain: cfg.App.Domain,
		Path:   "/",
		MaxAge: 60 * 60 * 24 * 30,
		Secure: false,
	}, func(ctx context.Context, u auth.CustomUser) (auth.CustomUser, error) {
		userDB, alreadyRegistred, err := s.userRepository.GetOrCreateUser(ctx, &repository.User{
			Id:     u.Id,
			Email:  u.Email,
			Access: u.Access,
		})
		if err != nil {
			return u, err
		}

		res, err := s.payService.AccountExist(ctx, &paycli.AccountExistReq{Id: userDB.Id})
		if err != nil {
			return u, err
		}
		if !res.Exist {
			_, err = s.payService.CreateAccount(ctx, &paycli.CreateAccountReq{Id: userDB.Id, Login: userDB.Email})
			if err != nil {
				return u, err
			}
		}

		if !alreadyRegistred {
			_, _ = s.userSettingsService.ResolveSettingsForUser(ctx, userDB.Id)
		}

		return auth.CustomUser{
			Id:     userDB.Id,
			Email:  userDB.Email,
			Access: userDB.Access,
		}, nil
	})
	h := http.Handler(mux)
	if cfg.Env == config.Local {
		h = cors.Handler(cors.Options{
			AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:     []string{"*"},
			AllowedOrigins:     []string{"http://localhost:*"},
			AllowCredentials:   true,
			Debug:              false,
			OptionsPassthrough: false,
			ExposedHeaders:     []string{"tz"},
		})(h)
	}
	h = Auth(h, ga)

	log.Log.Info("Listening grpc-gw server on ", cfg.GWAddr)

	go func() {

		go func() {
			for {
				stat := s.dispatcher.GetStat()
				gougeActiveProcesses.Set(stat.ActiveProcesses.Float())
				gougeActiveProfiles.Set(stat.ActiveProfiles.Float())
				gougeActiveTasks.Set(stat.ActiveTasks.Float())
				time.Sleep(time.Second * 10)
			}
		}()

		http.Handle("/metrics", promhttp.Handler())
		err = http.ListenAndServe(cfg.PrometheusPort, nil)
	}()

	err = http.ListenAndServe(cfg.GWAddr, h)
	if err != nil {
		return err
	}

	return nil
}
func ListenGRPC(ctx context.Context, port string, s *services) error {

	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(auth.AuthGRPC),
			grpc_auth.UnaryServerInterceptor(access.Access(s.userRepository)),
			otelgrpc.UnaryServerInterceptor(),
		)),
	)

	core.RegisterProfileServiceServer(server, s.profileService)
	core.RegisterHelperServiceServer(server, s.helperService)
	core.RegisterWithdrawerServiceServer(server, s.withdrawerService)
	core.RegisterFlowServiceServer(server, s.flowService)
	core.RegisterProcessServiceServer(server, s.processService)
	core.RegisterSettingsServiceServer(server, s.settingsService)
	core.RegisterSwap1InchServiceServer(server, s.swap1inchService)
	core.RegisterOrbiterServiceServer(server, s.orbiterService)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	log.Log.Info("Listening grpc server on ", port)

	return server.Serve(lis)
}

func initServices(ctx context.Context, cfg *config.Config) (*services, error) {

	tp, err := tracerProvider(cfg.JaegerUrl, false)
	if err == nil {
		otel.SetTracerProvider(tp)
	}

	tr := otel.Tracer("")
	_, span := tr.Start(ctx, "initServices")
	defer span.End()

	conn, err := pg.NewPGConnection(cfg.Database.Conn)
	if err != nil {
		return nil, err
	}
	mconn, err := pg.WrapPgConnWithSqlx(conn)
	if err != nil {
		return nil, err
	}

	if err := pg.RunMigrations(conn, cfg.MigrationsDir, false); err != nil {
		return nil, err
	}

	payService, err := pay.NewService(&pay.Config{
		Host: cfg.PayServiceGRPCAddr,
	})
	if err != nil {
		return nil, err
	}

	listener := pub.NewIListener()
	listener.Listen(ctx)

	rep := repository.NewPGRepository(mconn)

	statRepository := repository.StatRepository(rep)
	userRepository := repository.UserRepository(rep)
	processRepository := repository.ProcessRepository(rep)
	flowRepository := repository.FlowRepository(rep)
	profileRepository := repository.NewProfileRepositoryCrypto(repository.ProfileRepository(rep), userRepository, cfg.Lazanya)
	settingsRepository := repository.SettingsRepository(rep)
	settingsService := settings.NewService(settingsRepository)
	withdrawerRepository := repository.NewWithdrawerRepositoryCrypto(repository.WithdrawerRepository(rep), userRepository, cfg.Lazanya)

	tgBot, err := tg.NewBot(ctx, &tg.Config{
		Token:          cfg.TelegramToken,
		UserRepository: userRepository,
	})
	if err != nil {
		return nil, errors.Wrap(err, "tg.NewBot")
	}

	runner := process.NewRunner(processRepository, profileRepository, withdrawerRepository, tgBot, userRepository)
	snapshotService := snapshot.NewService(&snapshot.Config{
		Host: cfg.SnapshotHost,
	})
	orbiterService, err := orbiter.NewService()
	if err != nil {
		return nil, err
	}

	dispatcher := process.NewDispatcher(
		processRepository,
		runner,
		snapshotService,
		settingsService,
		payService,
		orbiterService,
	)
	go dispatcher.RunDispatcher(ctx)

	return &services{
		profileService:       v1.NewProfileService(profileRepository, settingsService),
		helperService:        v1.NewHelperService(settingsService, profileRepository, userRepository, payService, statRepository),
		withdrawerService:    v1.NewWithdrawerService(withdrawerRepository, userRepository, profileRepository),
		flowService:          v1.NewFlowService(flowRepository),
		processService:       v1.NewProcessService(processRepository, dispatcher, flowRepository, settingsService),
		settingsService:      v1.NewSettingsService(settingsService),
		swap1inchService:     v1.NewSwap1inchService(),
		processRepository:    processRepository,
		flowRepository:       flowRepository,
		profileRepository:    profileRepository,
		settingsRepository:   settingsRepository,
		withdrawerRepository: withdrawerRepository,
		userRepository:       userRepository,
		dispatcher:           dispatcher,
		payService:           payService,
		userSettingsService:  settingsService,
		orbiterService:       v1.NewOrbiterService(orbiterService),
		statRepository:       statRepository,
	}, nil
}

var (
	gougeActiveProcesses = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   "cry",
		Subsystem:   "backend",
		Name:        "active_processes",
		Help:        "The total number of active processes",
		ConstLabels: nil,
	})

	gougeActiveProfiles = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   "cry",
		Subsystem:   "backend",
		Name:        "active_profiles",
		Help:        "The total number of active profiles",
		ConstLabels: nil,
	})

	gougeActiveTasks = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   "cry",
		Subsystem:   "backend",
		Name:        "active_tasks",
		Help:        "The total number of active tasks",
		ConstLabels: nil,
	})
)

func Auth(h http.Handler, ga *auth.GoogleAuth) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/api/google/login" {
			ga.HandleLogin(w, r)
			return
		}
		if r.URL.Path == "/api/gw/google/oauth/callback" {
			ga.HandleCallback(w, r)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func tracerProvider(url string, disable bool) (*tracesdk.TracerProvider, error) {
	var err error
	var exp tracesdk.SpanExporter
	if disable {
		exp = tracetest.NewInMemoryExporter()
	} else {
		exp, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
		if err != nil {
			return nil, err
		}
	}

	// Create the Jaeger exporter
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.CFG.JaegerServiceName),
			attribute.String("env", string(config.CFG.Env)),
		)),
	)
	return tp, nil
}

type CryptoMarshaller struct{}

func (*CryptoMarshaller) ContentType(_ interface{}) string {
	return "application/base64"
}

// Marshal marshals "v" into JSON
func (j *CryptoMarshaller) Marshal(v interface{}) ([]byte, error) {

	p, ok := v.(proto.Message)
	if ok {
		bjson, err := protojson.MarshalOptions{Multiline: true, UseEnumNumbers: false, UseProtoNames: false, EmitUnpopulated: true}.Marshal(p)

		if err != nil {
			return nil, err
		}

		encryptedB, err := lib.AesEncrypt(bjson, []byte(config.CFG.WebIV), []byte(config.CFG.WebKey))
		if err != nil {
			return nil, err
		}
		return encryptedB, nil
	} else {
		println(1)
	}

	return nil, errors.New("CryptoMarshaller.Marshal pizda")
}

// Unmarshal unmarshals JSON data into "v".
func (j *CryptoMarshaller) Unmarshal(data []byte, v interface{}) error {

	p, ok := v.(proto.Message)
	if !ok {
		return errors.New("CryptoMarshaller.Unmarshal pizda")
	}

	decrypted, err := lib.AesDecrypt(data, []byte(config.CFG.WebIV), []byte(config.CFG.WebKey))
	if err != nil {
		return err
	}
	err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(decrypted, p)
	if err != nil {
		return err
	}
	return nil
}

func (d *Decoder) Decode(val interface{}) error {

	b, err := io.ReadAll(d.reader)
	if err != nil {
		return err
	}

	m := &CryptoMarshaller{}

	if err := m.Unmarshal(b, val); err != nil {
		return err
	}

	return nil
}

type Decoder struct {
	reader io.Reader
}

// NewDecoder returns a Decoder which reads JSON stream from "r".
func (j *CryptoMarshaller) NewDecoder(r io.Reader) runtime.Decoder {
	return &Decoder{r}
}

// NewEncoder returns an Encoder which writes JSON stream into "w".
func (j *CryptoMarshaller) NewEncoder(w io.Writer) runtime.Encoder {
	return json.NewEncoder(w)
}

// Delimiter for newline encoded JSON streams.
func (j *CryptoMarshaller) Delimiter() []byte {
	return []byte("\n")
}
