package zksyncera

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func Test_makeDataForBurnLiquidity(t *testing.T) {
	addr := common.HexToAddress("0x4a6e7c137a6691d55693ca3bc7e5c698d9d43815")
	expected := "0x0000000000000000000000004a6e7c137a6691d55693ca3bc7e5c698d9d438150000000000000000000000000000000000000000000000000000000000000001"
	actualB := makeDataForBurnLiquidity(addr, false)
	actual := "0x" + common.Bytes2Hex(actualB)
	assert.Equal(t, expected, actual)
}
