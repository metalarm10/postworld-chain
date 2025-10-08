package integration

import (
	"testing"

	"github.com/metalarm10/postworld-chain/tests/integration/indexer"
)

func TestKVIndexer(t *testing.T) {
	indexer.TestKVIndexer(t, CreateEvmd)
}
