package encoding_test

import (
	"bytes"
	"testing"

	"apigo.cc/go/cast"
	"apigo.cc/go/encoding"
)

// --- Hex ---
func TestHex(t *testing.T) {
	data := []byte("hello go")
	encoded := encoding.HexToString(data)
	decoded, err := encoding.UnHex([]byte(encoded))
	if err != nil || !bytes.Equal(data, decoded) {
		t.Fatal("Hex roundtrip failed")
	}

	if !bytes.Equal(cast.As(encoding.UnHexFromString(encoded)), data) {
		t.Error("UnHexFromString with cast.As failed")
	}

	if len(cast.As(encoding.UnHexFromString("!@#$"))) != 0 {
		t.Error("UnHexFromString should return empty for invalid hex chars with cast.As")
	}
}

// --- Base64 ---
func TestBase64(t *testing.T) {
	data := []byte("hello world")
	
	// Standard
	enc := encoding.Base64ToString(data)
	dec, err := encoding.UnBase64([]byte(enc))
	if err != nil || !bytes.Equal(data, dec) {
		t.Error("Base64 roundtrip failed")
	}
	
	if !bytes.Equal(cast.As(encoding.UnBase64FromString(enc)), data) {
		t.Error("UnBase64FromString with cast.As failed")
	}
	
	// URL
	uEnc := encoding.UrlBase64ToString([]byte("hello/world+"))
	uDec, _ := encoding.UnUrlBase64([]byte(uEnc))
	if !bytes.Equal([]byte("hello/world+"), uDec) {
		t.Error("UrlBase64 roundtrip failed")
	}

	if !bytes.Equal(cast.As(encoding.UnUrlBase64FromString(uEnc)), []byte("hello/world+")) {
		t.Error("UnUrlBase64FromString with cast.As failed")
	}
}

// --- Web ---
func TestWebEncoding(t *testing.T) {
	// URL
	data := []byte("a b+c")
	enc := encoding.UrlEncode(data)
	if !bytes.Equal(cast.As(encoding.UnUrlEncode(enc)), data) {
		t.Error("UrlEncode roundtrip failed")
	}
	if len(cast.As(encoding.UnUrlEncode("%ZZ"))) != 0 {
		t.Error("UnUrlEncode should return empty for invalid input with cast.As")
	}

	// HTML
	htmlData := []byte("<script>")
	escaped := encoding.HtmlEscape(htmlData)
	if encoding.HtmlUnescape(escaped) != string(htmlData) {
		t.Error("HtmlEscape roundtrip failed")
	}
}

// --- Utf8 ---
func TestUtf8(t *testing.T) {
	if !encoding.Utf8Valid([]byte("你好")) {
		t.Error("Valid UTF-8 should pass")
	}
	if encoding.Utf8Valid([]byte{0xff, 0xff}) {
		t.Error("Invalid UTF-8 should fail")
	}
}
