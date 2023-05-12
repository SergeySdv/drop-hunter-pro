package main

import (
	"context"
	log "crypto_scripts/internal/log"
	"crypto_scripts/internal/server/config"
	core "crypto_scripts/internal/server/pb/gen/proto/go/v1"
	"crypto_scripts/internal/server/process"
	"crypto_scripts/internal/server/proxy"
	"crypto_scripts/internal/server/repository"
	"crypto_scripts/internal/server/repository/pg"
	"crypto_scripts/internal/server/service/v1"
	"crypto_scripts/internal/server/ui"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/cors"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	project = "server"
)

type (
	services struct {
		r                 *repository.PGRepository
		profileService    *v1.ProfileService
		helperService     *v1.HelperService
		withdrawerService *v1.WithdrawerService
		flowService       *v1.FlowService
		processService    *v1.ProcessService
		settingsService   *v1.SettingsService
	}
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	s, err := initServices(ctx, cfg)
	if err != nil {
		panic(err)
	}

	l, err := log.NewLogger(project, "")
	if err != nil {
		panic(errors.Wrap(err, "log.NewLogger"))
	}
	defer l.Sync()

	fatal := make(chan error)
	go func() {
		err := ListenGRPC(ctx, cfg.GRPCAddr, s)
		if err != nil {
			fatal <- err
		}
	}()
	go func() {
		err := ListenGW(ctx, cfg)
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

func ListenGW(ctx context.Context, cfg *config.Config) error {

	conn, err := grpc.Dial(cfg.GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	mux := runtime.NewServeMux()

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

	log.Log.Info("Listening grpc-gw server on ", cfg.GWAddr)

	err = http.ListenAndServe(cfg.GWAddr, h)
	if err != nil {
		return err
	}

	return nil
}
func ListenGRPC(ctx context.Context, port string, s *services) error {

	server := grpc.NewServer(
	//grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
	//	grpc_auth.UnaryServerInterceptor(auth.AuthGRPC),
	//)),
	)

	core.RegisterProfileServiceServer(server, s.profileService)
	core.RegisterHelperServiceServer(server, s.helperService)
	core.RegisterWithdrawerServiceServer(server, s.withdrawerService)
	core.RegisterFlowServiceServer(server, s.flowService)
	core.RegisterProcessServiceServer(server, s.processService)
	core.RegisterSettingsServiceServer(server, s.settingsService)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	log.Log.Info("Listening grpc server on ", port)

	return server.Serve(lis)
}

func initServices(ctx context.Context, cfg *config.Config) (*services, error) {

	conn, err := pg.NewPGConnection(cfg.Database.Conn)
	if err != nil {
		return nil, err
	}
	mconn, err := pg.WrapPgConnWithSqlx(conn)
	if err != nil {
		return nil, err
	}

	if err := pg.RunMigrations(conn, "./internal/server/migrations", false); err != nil {
		return nil, err
	}

	rep := repository.NewPGRepository(mconn)

	runner := process.NewRunner(rep, true)
	dispatcher := process.NewDispatcher(rep, runner)
	go func() {
		for {

			if err := dispatcher.Run(ctx); err != nil {
				println(err)
			}
			time.Sleep(time.Second)
		}
	}()

	return &services{
		r:                 rep,
		profileService:    v1.NewProfileService(rep, rep),
		helperService:     v1.NewHelperService(rep, rep),
		withdrawerService: v1.NewWithdrawerService(rep),
		flowService:       v1.NewFlowService(rep),
		processService:    v1.NewProcessService(rep, dispatcher),
		settingsService:   v1.NewSettingsService(rep),
	}, nil
}
