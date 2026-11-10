package encoding

import (
	"errors"

	"github.com/api-go/cast"
	"github.com/api-go/rand"
)

// IntEncoder 提供整数与字节切片之间的自定义进制转换
type IntEncoder struct {
	radix     uint8
	digits    string
	decodeMap [256]int
}

// EncodeInt 将整数转换为字节切片
func (enc *IntEncoder) EncodeInt(u uint64) []byte {
	return enc.AppendInt(nil, u)
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

// FillInt 使用随机字符序列填充数据至指定长度，高并发下具备零锁极速随机性能，就地修改并返回
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
		// 使用 go/rand 零锁随机性能从 digits 中抽取字符填充
		rnd := rand.FastInt(0, radix-1)
		buf[i] = enc.digits[rnd]
	}
	return buf
}

// FoldInt 使用编码器的字符集模加法，将超出 length 的尾部字节折叠到前 length 字节中，就地修改并截断返回
func (enc *IntEncoder) FoldInt(buf []byte, length int) []byte {
	size := len(buf)
	if size <= length {
		return buf
	}
	radix := int(enc.radix)
	for i := length; i < size; i++ {
		p1 := enc.decodeMap[buf[i%length]]
		p2 := enc.decodeMap[buf[i]]
		if p1 < 0 {
			p1 = 0
		}
		if p2 < 0 {
			p2 = 0
		}
		newP := (p1 + p2) % radix
		buf[i%length] = enc.digits[newP]
	}
	return buf[:length]
}

// ExchangeInt 对数据进行位置交替重排，就地修改并返回
func (enc *IntEncoder) ExchangeInt(data any) []byte {
	buf := cast.To[[]byte](data)
	size := len(buf)
	if size <= 1 {
		return buf
	}

	var buf2 []byte
	if size <= 20 {
		var tmp [20]byte // 栈分配：极致性能，零垃圾回收负担
		buf2 = tmp[:size]
	} else {
		buf2 = make([]byte, size)
	}

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

	// 将打乱后的数据拷贝回原切片，实现“原地修改”
	copy(buf, buf2)
	return buf
}

// HashInt 双向非线性渐进式置换混淆，原地修改并返回，所有值保持在 radix 字符集内
func (enc *IntEncoder) HashInt(buf []byte) []byte {
	if len(buf) == 0 {
		return buf
	}

	// 第一轮：正向传播（左向右）
	prevP := (len(buf) * 17) % int(enc.radix)
	for i, c := range buf {
		p := enc.decodeMap[c]
		if p < 0 {
			p = 0
		}
		p = (prevP + p) % int(enc.radix)
		buf[i] = enc.digits[p]
		prevP = OrderedIntEncoder.decodeMap[DefaultIntEncoder.digits[p]]
	}

	// 第二轮：反向传播（右向左），回传差异实现全局雪崩扩散
	prevP = (len(buf) * 31) % int(enc.radix)
	for i := len(buf) - 1; i >= 0; i-- {
		p := enc.decodeMap[buf[i]]
		if p < 0 {
			p = 0
		}
		p = (prevP + p) % int(enc.radix)
		buf[i] = enc.digits[p]
		prevP = OrderedIntEncoder.decodeMap[DefaultIntEncoder.digits[p]]
	}

	return buf
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

// EncodeInt 使用默认编码器将整数转换为字节切片
func EncodeInt(u uint64) []byte {
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

// HashInt 对数据进行渐进式置换混淆
func HashInt(buf []byte) []byte {
	return DefaultIntEncoder.HashInt(buf)
}

// FillInt 使用默认编码器填充数据至指定长度
func FillInt(data any, length int) []byte {
	return DefaultIntEncoder.FillInt(data, length)
}
