package encoding

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/taikoxyz/taiko-mono/packages/taiko-client/bindings/minimal"
)

var (
	// Pacaya fork
	TaikoInboxAlethiaABI *abi.ABI
)

func init() {
	var err error

	if TaikoInboxAlethiaABI, err = minimal.IInboxMetaData.GetAbi(); err != nil {
		log.Crit("Get TaikoL1 ABI error", "error", err)
	}
}
