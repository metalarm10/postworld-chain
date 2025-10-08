package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/metalarm10/postworld-chain/tests/integration/wallets"
)

func TestLedgerTestSuite(t *testing.T) {
	s := wallets.NewLedgerTestSuite(CreateEvmd)
	suite.Run(t, s)
}
