package testdata

import (
	contractutils "github.com/metalarm10/postworld-chain/contracts/utils"
	evmtypes "github.com/metalarm10/postworld-chain/x/vm/types"
)

func LoadBankCallerContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("BankCaller.json")
}
