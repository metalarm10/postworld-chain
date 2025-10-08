package contracts

import (
	contractutils "github.com/metalarm10/postworld-chain/contracts/utils"
	evmtypes "github.com/metalarm10/postworld-chain/x/vm/types"
)

func LoadIcs20CallerContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ICS20Caller.json")
}
