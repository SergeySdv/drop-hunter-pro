package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testUsers = []User{
	{
		ExternalUser: ExternalUser{
			AuthType:   "1",
			ExternalId: "2",
			Email: NullString{
				Valid:  true,
				String: "2",
			},
			Login: NullString{},
			Name: NullString{
				Valid:  false,
				String: "",
			},
		},
	},
}

func TestSessionToken(t *testing.T) {
	for _, expectedUser := range testUsers {
		token, err := GenerateToken(&expectedUser)
		assert.NoError(t, err)

		actualUser, _, err := ExtractUserAndTokenFromStringToken(token)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, *actualUser)
	}
}
