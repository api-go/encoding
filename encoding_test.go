package encoding_test

import (
	"bytes"
	"testing"

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

	if !bytes.Equal(encoding.MustUnHexFromString(encoded), data) {
		t.Error("MustUnHexFromString failed")
	}

	if len(encoding.MustUnHexFromString("!@#$")) != 0 {
		t.Error("MustUnHexFromString should return empty for invalid hex chars")
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
	
	if !bytes.Equal(encoding.MustUnBase64FromString(enc), data) {
		t.Error("MustUnBase64FromString failed")
	}
	
	// URL
	uEnc := encoding.UrlBase64ToString([]byte("hello/world+"))
	uDec, _ := encoding.UnUrlBase64([]byte(uEnc))
	if !bytes.Equal([]byte("hello/world+"), uDec) {
		t.Error("UrlBase64 roundtrip failed")
	}
}

// --- Web ---
func TestWebEncoding(t *testing.T) {
	// URL
	data := []byte("a b+c")
	enc := encoding.UrlEncode(data)
	if !bytes.Equal(encoding.MustUnUrlEncode(enc), data) {
		t.Error("UrlEncode roundtrip failed")
	}
	if len(encoding.MustUnUrlEncode("%ZZ")) != 0 {
		t.Error("MustUnUrlEncode should return empty for invalid input")
	}

	// HTML
	htmlData := []byte("<script>")
	escaped := encoding.HtmlEscape(htmlData)
	if encoding.MustUnHtmlEscape(escaped) != string(htmlData) {
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
