package encoding

import (
	"crypto/hmac"
	"crypto/sha512"
	"errors"

	"apigo.cc/go/cast"
)

// IntEncoder 提供整数与字节切片之间的自定义进制转换
type IntEncoder struct {
	radix     uint8
	digits    string
	decodeMap [256]int
}

// EncodeInt 将整数转换为字符串
func (enc *IntEncoder) EncodeInt(u uint64) string {
	return string(enc.AppendInt(nil, u))
}

// AppendInt 将整数追加到已有字节切片中
func (enc *IntEncoder) AppendInt(buf []byte, u uint64) []byte {
	if buf == nil {
		buf = make([]byte, 0)
	}
	radix := uint64(enc.radix)
	for u >= radix {
		q := u / radix
		buf = append(buf, enc.digits[uint(u-q*radix)])
		u = q
	}
	buf = append(buf, enc.digits[uint(u)])
	return buf
}

// FillInt 使用循环字符序列填充数据至指定长度
func (enc *IntEncoder) FillInt(data any, length int) []byte {
	buf := cast.To[[]byte](data)
	currLen := len(buf)
	if currLen >= length {
		return buf
	}
	if cap(buf) < length {
		newBuf := make([]byte, currLen, length)
		copy(newBuf, buf)
		buf = newBuf
	}
	buf = buf[:length]
	radix := int(enc.radix)
	for i := currLen; i < length; i++ {
		buf[i] = enc.digits[i%radix]
	}
	return buf
}

// ExchangeInt 对数据进行位置交替重排
func (enc *IntEncoder) ExchangeInt(data any) []byte {
	buf := cast.To[[]byte](data)
	size := len(buf)
	if size <= 1 {
		return buf
	}

	buf2 := make([]byte, size)
	buf2_i := 0
	buf2_ai := 0
	buf2_ri := size - 1
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			buf2[buf2_i] = buf[buf2_ri]
			buf2_ri--
		} else {
			buf2[buf2_i] = buf[buf2_ai]
			buf2_ai++
		}
		buf2_i++
	}
	return buf2
}

// HashInt 对数据进行 HMAC-SHA512 哈希
func (enc *IntEncoder) HashInt(data any, key any) []byte {
	hash := hmac.New(sha512.New, cast.To[[]byte](key))
	hash.Write(cast.To[[]byte](data))
	return hash.Sum([]byte{})
}

// DecodeInt 从数据解码为整数
func (enc *IntEncoder) DecodeInt(data any) uint64 {
	buf := cast.To[[]byte](data)
	radix := uint64(enc.radix)
	if buf == nil {
		return 0
	}
	var n uint64 = 0
	for i := len(buf) - 1; i >= 0; i-- {
		p := enc.decodeMap[buf[i]]
		if p >= 0 {
			n = n*radix + uint64(p)
		}
	}
	return n
}

// NewIntEncoder 创建一个新的整数编码器
func NewIntEncoder(digits string, radix uint8) (*IntEncoder, error) {
	if len(digits) < int(radix) {
		return nil, errors.New("int encoder digits length is less than radix")
	}

	e := IntEncoder{digits: digits, radix: radix, decodeMap: [256]int{}}
	for i := 0; i < 256; i++ {
		e.decodeMap[i] = -1
	}

	for i := 0; i < int(radix); i++ {
		d := digits[i]
		if e.decodeMap[d] != -1 {
			return nil, errors.New("int encoder digits is repeated: " + digits)
		}
		e.decodeMap[d] = i
	}

	return &e, nil
}

// 默认编码器实例
var DefaultIntEncoder, _ = NewIntEncoder("9ukH1grX75TQS6LzpFAjIivsdZoO0mc8NBwnyYDhtMWEC2V3KaGxfJRPqe4lbU", 62)
var OrderedIntEncoder, _ = NewIntEncoder("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", 62)

// EncodeInt 使用默认编码器将整数转换为字符串
func EncodeInt(u uint64) string {
	return DefaultIntEncoder.EncodeInt(u)
}

// AppendInt 使用默认编码器将整数追加到已有字节切片中
func AppendInt(buf []byte, u uint64) []byte {
	return DefaultIntEncoder.AppendInt(buf, u)
}

// DecodeInt 使用默认编码器从数据解码为整数
func DecodeInt(data any) uint64 {
	return DefaultIntEncoder.DecodeInt(data)
}

// ExchangeInt 使用默认编码器对数据进行位置交替重排
func ExchangeInt(data any) []byte {
	return DefaultIntEncoder.ExchangeInt(data)
}

// HashInt 对数据进行 HMAC-SHA512 哈希
func HashInt(data any, key any) []byte {
	return DefaultIntEncoder.HashInt(data, key)
}

// FillInt 使用默认编码器填充数据至指定长度
func FillInt(data any, length int) []byte {
	return DefaultIntEncoder.FillInt(data, length)
}
