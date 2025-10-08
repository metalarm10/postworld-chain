package ante

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/metalarm10/postworld-chain/postworldd/tests/integration"
	"github.com/metalarm10/postworld-chain/tests/integration/ante"
)

func TestEvmUnitAnteTestSuite(t *testing.T) {
	suite.Run(t, ante.NewEvmUnitAnteTestSuite(integration.CreateEvmd))
}
