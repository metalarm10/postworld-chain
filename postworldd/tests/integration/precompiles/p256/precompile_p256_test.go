package p256

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/metalarm10/postworld-chain/postworldd/tests/integration"
	"github.com/metalarm10/postworld-chain/tests/integration/precompiles/p256"
)

func TestP256PrecompileTestSuite(t *testing.T) {
	s := p256.NewPrecompileTestSuite(integration.CreateEvmd)
	suite.Run(t, s)
}

func TestP256PrecompileIntegrationTestSuite(t *testing.T) {
	p256.TestPrecompileIntegrationTestSuite(t, integration.CreateEvmd)
}
