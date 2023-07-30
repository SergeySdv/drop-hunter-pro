package defi

import (
	"github.com/hardstylez72/cry/internal/lib"
	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
)

var SupportedTokens = lib.NewSet[v1.Token]()
var SupportedNetworks = lib.NewSet[v1.Network]()

func init() {
	for k := range v1.Token_name {
		SupportedTokens.Set(v1.Token(k))
	}
	for k := range v1.Network_name {
		SupportedNetworks.Set(v1.Network(k))
	}
}

func TokenSupported(t v1.Token) bool {
	return SupportedTokens.Get(t)
}

func TokenFromSting(t string) (v1.Token, bool) {
	id, exist := v1.Token_value[t]
	if exist {
		return v1.Token(id), true
	}
	return -1, false
}

func NetworkSupported(t v1.Network) bool {
	return SupportedNetworks.Get(t)
}
