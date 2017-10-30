package run

import (
	"testing"
)

func TestRun(t *testing.T) {
	t.Skip("skipping test")
}

func BenchmarkRun(b *testing.B) {
	Run("./bench.sh")
}
