package auth

import (
	"net/http"
	"time"
)

var (
	rm = map[string]bool{} // todo: remove!!
)

type controller struct {
	google OAuth2
}

type Oauth struct {
	Google              Config
	SessionCookieConfig SessionCookieConfig
}

type SessionCookieConfig struct {
	Name   string
	Domain string
	Path   string
	MaxAge int
	Secure bool
}

type Config struct {
	RedirectURL   string
	ClientID      string
	ClientSecret  string
	Scopes        []string
	UserInfoURL   string
	UserRedirects UserRedirects
}

type UserRedirects struct {
	OnSuccess string
	OnFailure string
}

func SetSessionCookie(w http.ResponseWriter, token string, cookie SessionCookieConfig) {
	http.SetCookie(w, &http.Cookie{
		Name:    cookie.Name,
		Value:   token,
		MaxAge:  cookie.MaxAge,
		Expires: time.Now().Add(time.Hour * 24 * 30),
		Path:    cookie.Path,
		Secure:  cookie.Secure,
	})
}
