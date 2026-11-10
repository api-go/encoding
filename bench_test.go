package encoding_test

import (
	"testing"

	"github.com/api-go/encoding"
)

func BenchmarkHexEncode(b *testing.B) {
	data := []byte("hello go performance")
	for i := 0; i < b.N; i++ {
		_ = encoding.Hex(data)
	}
}

func BenchmarkBase64Encode(b *testing.B) {
	data := []byte("hello go performance")
	for i := 0; i < b.N; i++ {
		_ = encoding.Base64(data)
	}
}

func BenchmarkIntEncoder(b *testing.B) {
	enc := encoding.DefaultIntEncoder
	val := uint64(123456789)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = enc.EncodeInt(val)
	}
}
