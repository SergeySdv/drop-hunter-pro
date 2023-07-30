package orbiter

import (
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

var supportedTokens = lib.NewSet[v1.Token]()
var supportedNetworks = lib.NewSet[v1.Network]()

func tokenSupported(t v1.Token) bool {
	return supportedTokens.Get(t)
}

func tokenFromSting(t string) (v1.Token, bool) {
	id, exist := v1.Token_value[t]
	if exist {
		return v1.Token(id), true
	}
	return -1, false
}

func networkSupported(t v1.Network) bool {
	return supportedNetworks.Get(t)
}
