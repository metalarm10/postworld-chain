package postworldd

import (
	cmn "github.com/metalarm10/postworld-chain/precompiles/common"
	evmtypes "github.com/metalarm10/postworld-chain/x/vm/types"
)

type BankKeeper interface {
	evmtypes.BankKeeper
	cmn.BankKeeper
}
