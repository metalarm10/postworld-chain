package vm

import (
	"github.com/metalarm10/postworld-chain/testutil/integration/evm/network"
	evmante "github.com/metalarm10/postworld-chain/x/vm/ante"

	storetypes "cosmossdk.io/store/types"
)

func (s *EvmAnteTestSuite) TestBuildEvmExecutionCtx() {
	network := network.New(s.create, s.options...)

	ctx := evmante.BuildEvmExecutionCtx(network.GetContext())

	s.Equal(storetypes.GasConfig{}, ctx.KVGasConfig())
	s.Equal(storetypes.GasConfig{}, ctx.TransientKVGasConfig())
}
