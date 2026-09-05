package encoding

import (
	"bytes"
	"testing"
)

func TestEncoding(t *testing.T) {
	data := []byte("hello world")

	// Hex
	encoded := Hex(data)
	decoded, err := UnHex(encoded)
	if err != nil || !bytes.Equal(decoded, data) {
		t.Errorf("Hex failed: got %v, error: %v", decoded, err)
	}

	// Base64
	enc := Base64(data)
	dec, err := UnBase64(enc)
	if err != nil || !bytes.Equal(dec, data) {
		t.Errorf("Base64 failed: got %v, error: %v", dec, err)
	}

	// Base64Raw
	rawEnc := Base64Raw(data)
	rawDec, err := UnBase64(rawEnc)
	if err != nil || !bytes.Equal(rawDec, data) {
		t.Errorf("Base64Raw failed: got %v, error: %v", rawDec, err)
	}

	// URLBase64
	uData := []byte("https://apigo.cc?a=1&b=2")
	uEnc := URLBase64(uData)
	uDec, _ := UnURLBase64(uEnc)
	if !bytes.Equal(uDec, uData) {
		t.Errorf("URLBase64 failed: got %v", uDec)
	}

	// URLBase64Raw
	uRawEnc := URLBase64Raw(uData)
	uRawDec, err := UnURLBase64(uRawEnc)
	if err != nil || !bytes.Equal(uRawDec, uData) {
		t.Errorf("URLBase64Raw failed: got %v, error: %v", uRawDec, err)
	}

	// URLEncode
	urlEnc := URLEncode(data)
	urlDec, err := UnURLEncode(urlEnc)
	if err != nil || !bytes.Equal(urlDec, data) {
		t.Errorf("URLEncode failed: got %v, error: %v", urlDec, err)
	}

	// HTMLEscape
	htmlData := "<div>hello</div>"
	escaped := HTMLEscape(htmlData)
	if HTMLUnescape(escaped) != htmlData {
		t.Errorf("HTMLEscape failed")
	}

	// UTF8Valid
	if !UTF8Valid("你好") {
		t.Error("UTF8Valid failed")
	}
}

func TestSortJoin(t *testing.T) {
	m := map[string]any{
		"b": 2,
		"a": 1,
		"c": "3",
	}
	res := SortJoin(m, "&", "=", true)
	if res != "a=1&b=2&c=3" {
		t.Errorf("SortJoin map failed: %s", res)
	}

	type User struct {
		ID   int
		Name string
	}
	u := User{ID: 1, Name: "sam"}
	resStruct := SortJoin(u, ";", ":", false)
	if resStruct != "id:1;name:sam" {
		t.Errorf("SortJoin struct failed: %s", resStruct)
	}
}

func TestIntEncoder(t *testing.T) {
	val := uint64(123456789)
	encoded := EncodeInt(val)
	decoded := DecodeInt(encoded)
	if val != decoded {
		t.Errorf("IntEncoder failed: expected %d, got %d", val, decoded)
	}
}
