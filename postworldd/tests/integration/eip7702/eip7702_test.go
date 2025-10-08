package eip7702

import (
	"testing"

	"github.com/metalarm10/postworld-chain/postworldd/tests/integration"
	"github.com/metalarm10/postworld-chain/tests/integration/eip7702"
)

func TestEIP7702IntegrationTestSuite(t *testing.T) {
	eip7702.TestEIP7702IntegrationTestSuite(t, integration.CreateEvmd)
}
