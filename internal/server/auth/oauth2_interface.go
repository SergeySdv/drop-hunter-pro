package auth

import (
	"context"
	"net/http"
	"strings"
)

type ExternalUser struct {
	AuthType   string `json:"authType"`
	ExternalId string `json:"externalId"`

	Email NullString `json:"email"`
	Login NullString `json:"login"`
	Name  NullString `json:"name"`
}

type NullString struct {
	Valid  bool
	String string
}

type OAuth2 interface {
	GetUser(ctx context.Context, state string, code string) (*ExternalUser, error)
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleCallback(w http.ResponseWriter, r *http.Request)
}

func (t *NullString) UnmarshalJSON(data []byte) error {

	t.String = string(data)
	t.String = strings.Replace(t.String, "\"", "", -1)
	if t.String != "" {
		t.Valid = true
	}

	return nil
}
