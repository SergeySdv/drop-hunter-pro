package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var key = []byte("dc094250d93c4524802b87202f647a6c")

func TestCrypto(t *testing.T) {
	data1 := "123456789fjfjdf!@#$#@^"
	userId := "dc094250d93c-4524-802b-87202f647a6c"
	ss := "1473263457"

	encrypted, err := Encrypt(userId, ss, []byte(data1))
	assert.NoError(t, err)
	data2, err := Decrypt(userId, ss, encrypted)
	assert.NoError(t, err)
	assert.Equal(t, string(data2), data1)
}

//
//func TestCrypto2(t *testing.T) {
//	data1 := "123456789fjfjdf!@#$#@^"
//
//	encrypted, err := AesEncrypt([]byte(data1), key)
//	assert.NoError(t, err)
//	data2, err := AesDecrypt(encrypted, key)
//	assert.NoError(t, err)
//	assert.Equal(t, string(data2), data1)
//}
