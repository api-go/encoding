package encoding

import (
	"bytes"
	"testing"
)

func TestIntEncoderInstance(t *testing.T) {
	enc := DefaultIntEncoder
	val := uint64(123456789)
	encoded := enc.EncodeInt(val)
	decoded := enc.DecodeInt(encoded)
	if val != decoded {
		t.Errorf("IntEncoder instance failed: expected %d, got %d", val, decoded)
	}

	// NewIntEncoder errors
	_, err := NewIntEncoder("12", 5) // 字符集不足
	if err == nil {
		t.Error("expected error for insufficient digits")
	}

	_, err = NewIntEncoder("112345", 6) // 重复字符
	if err == nil {
		t.Error("expected error for repeated digits")
	}

	// DecodeInt(nil)
	if DecodeInt(nil) != 0 {
		t.Error("DecodeInt(nil) != 0")
	}

	// FillInt
	buf := []byte(EncodeInt(123))
	filled := FillInt(buf, 10)
	if len(filled) != 10 {
		t.Errorf("FillInt failed: expected len 10, got %d", len(filled))
	}
	if !bytes.HasPrefix(filled, buf) {
		t.Error("FillInt prefix mismatch")
	}

	// ExchangeInt
	buf = []byte("abcde")
	exchanged := ExchangeInt(buf)
	if len(exchanged) != len(buf) {
		t.Error("ExchangeInt len mismatch")
	}
	if bytes.Equal(exchanged, buf) {
		t.Error("ExchangeInt should change data")
	}

	// HashInt
	hashed := HashInt(buf, []byte("key"))
	if len(hashed) == 0 {
		t.Error("HashInt failed")
	}
}
