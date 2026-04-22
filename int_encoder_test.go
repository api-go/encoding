package encoding_test

import (
	"bytes"
	"testing"

	"apigo.cc/go/encoding"
)

// --- 正常逻辑测试 ---
func TestIntEncoderNormal(t *testing.T) {
	enc := encoding.DefaultIntEncoder
	nums := []uint64{0, 1, 100, 123456789, 18446744073709551615}
	for _, n := range nums {
		if enc.DecodeInt(enc.EncodeInt(n)) != n {
			t.Errorf("Roundtrip failed for %d", n)
		}
	}
}

// --- 边界与异常测试 ---
func TestIntEncoderEdge(t *testing.T) {
	// 初始化异常
	_, err := encoding.NewIntEncoder("12", 5) // 字符集不足
	if err == nil { t.Error("NewIntEncoder failed to detect short digits") }
	
	_, err = encoding.NewIntEncoder("112345", 2) // 重复字符
	if err == nil { t.Error("NewIntEncoder failed to detect repeated digits") }

	// 空输入
	if encoding.DecodeInt(nil) != 0 { t.Error("DecodeInt(nil) != 0") }
	
	// 填充逻辑边界
	buf := encoding.EncodeInt(123)
	filled := encoding.DefaultIntEncoder.FillInt(buf, 2) // 长度小于原长
	if len(filled) != len(buf) {
		t.Error("FillInt should not truncate data")
	}
}

// --- 混淆与哈希测试 ---
func TestIntEncoderAdvanced(t *testing.T) {
	enc := encoding.DefaultIntEncoder
	buf := enc.EncodeInt(12345)
	
	// Exchange
	exchanged := enc.ExchangeInt(buf)
	if len(exchanged) != len(buf) {
		t.Fatal("Exchange length mismatch")
	}
	
	// Hash
	h1 := enc.HashInt(buf, []byte("key"))
	h2 := enc.HashInt(buf, []byte("key"))
	if !bytes.Equal(h1, h2) {
		t.Error("HashInt non-deterministic")
	}
}
