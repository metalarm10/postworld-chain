package testutil

import (
	anteinterfaces "github.com/metalarm10/postworld-chain/ante/interfaces"
	"github.com/metalarm10/postworld-chain/x/vm/statedb"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewStateDB returns a new StateDB for testing purposes.
func NewStateDB(ctx sdk.Context, evmKeeper anteinterfaces.EVMKeeper) *statedb.StateDB {
	return statedb.New(ctx, evmKeeper, statedb.NewEmptyTxConfig())
}
