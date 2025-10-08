package ante

import (
	"testing"

	"github.com/metalarm10/postworld-chain/postworldd/tests/integration"
	"github.com/metalarm10/postworld-chain/tests/integration/ante"
)

func TestAnte_Integration(t *testing.T) {
	ante.TestIntegrationAnteHandler(t, integration.CreateEvmd)
}

func BenchmarkAnteHandler(b *testing.B) {
	// Run the benchmark with a mock EVM app
	ante.RunBenchmarkAnteHandler(b, integration.CreateEvmd)
}

func TestValidateHandlerOptions(t *testing.T) {
	ante.RunValidateHandlerOptionsTest(t, integration.CreateEvmd)
}
