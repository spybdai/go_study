package performance

import (
	"testing"
)

// go 1.10

func BenchmarkStringCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0

	b.ResetTimer()
}
