package mempool

import (
	"testing"

	"github.com/metalarm10/postworld-chain/postworldd/tests/integration"

	"github.com/stretchr/testify/suite"

	"github.com/metalarm10/postworld-chain/tests/integration/mempool"
)

func TestMempoolIntegrationTestSuite(t *testing.T) {
	suite.Run(t, mempool.NewMempoolIntegrationTestSuite(integration.CreateEvmd))
}
