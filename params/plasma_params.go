package params

import (
	"math/big"

	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/crypto"
)

var (
	NullAddress = common.Address{0x0000000000000000000000000000000000000000}
	NullKey, _  = crypto.HexToECDSA("00")

	SubmitBlockGasPrice        = big.NewInt(1e9)
	SubmitBlockGasLimit uint64 = 4000000

	RequestTxGasPrice        = big.NewInt(1e9)
	RequestTxGasLimit uint64 = 100000
)
