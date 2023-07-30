package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	authTypeGoogle = "Google"
)

type GoogleOauthUserData struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type GoogleAuth struct {
	Oauth2Config  *oauth2.Config
	UserInfoURL   string
	UserRedirects UserRedirects
	httClient     *http.Client
	SessionCookie SessionCookieConfig
	DBFn          DBFn
}

type CustomUser struct {
	Id     string `db:"id"`
	Email  string `db:"email"`
	Access bool   `db:"access"`
}

type DBFn = func(ctx context.Context, u CustomUser) (CustomUser, error)

func NewGoogleOAuth2Controller(cfg Config, sessionCookie SessionCookieConfig, DBFn DBFn) *GoogleAuth {
	return &GoogleAuth{
		Oauth2Config: &oauth2.Config{
			RedirectURL:  cfg.RedirectURL,
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Scopes:       cfg.Scopes,
			Endpoint:     google.Endpoint,
		},
		UserRedirects: cfg.UserRedirects,
		UserInfoURL:   cfg.UserInfoURL,
		httClient:     &http.Client{},
		SessionCookie: sessionCookie,
		DBFn:          DBFn,
	}
}

func (a *GoogleAuth) HandleLogin(w http.ResponseWriter, r *http.Request) {

	uniq := uuid.New().String()
	rm[uniq] = true
	url := a.Oauth2Config.AuthCodeURL(uniq)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *GoogleAuth) HandleCallback(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	u, err := a.GetUser(ctx, r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		http.Redirect(w, r, a.UserRedirects.OnFailure, http.StatusTemporaryRedirect)
		return
	}

	if !u.Email.Valid {
		http.Redirect(w, r, a.UserRedirects.OnFailure, http.StatusTemporaryRedirect)
		return
	}

	lu, err := a.DBFn(r.Context(), CustomUser{
		Id:     uuid.New().String(),
		Email:  u.Email.String,
		Access: false,
	})
	if err != nil {
		http.Redirect(w, r, a.UserRedirects.OnFailure, http.StatusTemporaryRedirect)
		return
	}

	token, err := GenerateToken(&User{
		ExternalUser: *u,
		UserId:       lu.Id,
	})
	if err != nil {
		http.Redirect(w, r, a.UserRedirects.OnFailure, http.StatusTemporaryRedirect)
		return
	}

	SetSessionCookie(w, token, a.SessionCookie)

	http.Redirect(w, r, a.UserRedirects.OnSuccess, http.StatusTemporaryRedirect)
	return
}

func (a *GoogleAuth) GetUser(ctx context.Context, state string, code string) (*ExternalUser, error) {
	if !rm[state] {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := a.Oauth2Config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, buildGoogleGetUserUrl(token.AccessToken, a.UserInfoURL), nil)
	if err != nil {
		return nil, fmt.Errorf("can not get user from google: %s", err.Error())
	}

	response, err := a.httClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("can not get user from google: authorization error")
	}

	user := new(GoogleOauthUserData)

	dec := json.NewDecoder(response.Body)
	if err := dec.Decode(user); err != nil {
		return nil, fmt.Errorf("can not decode response body: %s", err.Error())
	}

	return convertGoogleUser(user, authTypeGoogle), nil
}

func buildGoogleGetUserUrl(token, baseUrl string) string {
	return baseUrl + "?access_token=" + token
}

func convertGoogleUser(googleUser *GoogleOauthUserData, authTypeGoogleId string) *ExternalUser {

	u := ExternalUser{
		AuthType:   authTypeGoogleId,
		ExternalId: googleUser.ID,
		Email:      NullString{},
		Name:       NullString{},
	}

	if googleUser.Email != "" {
		u.Email.String = googleUser.Email
		u.Email.Valid = true
	}

	if googleUser.Name != "" {
		u.Name.String = googleUser.FamilyName + " " + googleUser.Name
		u.Name.Valid = true
	}

	return &u
}
