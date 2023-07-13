package unmarshal

import (
	"math/big"
	"strings"
)

// 一些特殊例子
// 0000.2不合法
// 000002不合法
// 2.00000合法

// 3.1415926
const (
	NumberStart int = iota
	NumberWantingIntegerPart
	NumberWantingPointOrEnd
	NumberWantingFloatPart
)

func parseNumber(p *peekReader) (float64, bool) {
	firstByte, ok := p.PeekByte()
	if !ok {
		return 0, false
	}

	// 保存符号
	var sign byte = '+'

	// 保存每个位的string表示
	var num_diget_string = make([]string, 0)

	if firstByte == '-' {
		p.MoveByte()
		sign = '-'

	} else if (firstByte) == '+' {
		p.MoveByte()
	}

	state := NumberStart

	// 循环读取更多数字, 直到读完Reader或者遇到其它符号
	// 30
	// 0.12
	// +30
	// +0.12
	// +3.140000
	for {
		b, ok := p.PeekByte()
		if !ok {
			break
		}

		switch state {
		case NumberStart:
			// 0.23
			if b == '0' {
				state = NumberWantingPointOrEnd
				num_diget_string = append(num_diget_string, "0")
				p.MoveByte()
				continue
			} else if IsByteAsciiNumberNo0(b) {
				// 那么就是数字
				state = NumberWantingIntegerPart
				num_diget_string = append(num_diget_string, string(b))
				p.MoveByte()
				continue
			} else {
				// 比如+x
				return 0, false
			}
		case NumberWantingPointOrEnd:
			if b == '.' {
				num_diget_string = append(num_diget_string, ".")
				state = NumberWantingFloatPart
				p.MoveByte()
				continue
			} else {
				return 0, true
			}
		case NumberWantingIntegerPart:
			if b == '.' {
				num_diget_string = append(num_diget_string, ".")
				state = NumberWantingFloatPart
				p.MoveByte()
				continue
			} else if isByteAsciiNumber(b) {
				p.MoveByte()
				num_diget_string = append(num_diget_string, string(b))
				continue
			} else {
				goto out
			}
		case NumberWantingFloatPart:
			// +3.56
			if isByteAsciiNumber(b) {
				num_diget_string = append(num_diget_string, string(b))
				state = NumberWantingFloatPart
				p.MoveByte()
				continue
			} else {
				goto out
			}
		}
	}

out:

	bigSum := big.NewFloat(0)
	var stringNum string
	if sign == '+' {
		stringNum = strings.Join(num_diget_string, "")
	} else {
		stringNum = strings.Join(num_diget_string, "")
		stringNum = "-" + stringNum
	}

	f, ok := bigSum.SetString(stringNum)
	if !ok {
		panic("unexpected error")
	}

	sum, _ := f.Float64()
	return sum, true
}

// const (
// 	NumberNone int = iota
// 	NumberWantingPointOrEnd
// 	NumberWantingAny
// 	NumberWaintingAnyNum
// )

// func ParseNumber(p *PeekReader) (float64, bool) {
// 	firstByte, ok := p.PeekByte()
// 	if !ok {
// 		return 0, false
// 	}

// 	// 保存符号
// 	var sign byte = '+'

// 	// 保存每个位的string表示
// 	var num_diget_string = make([]string, 0)

// 	if firstByte == '-' {
// 		p.MoveByte()
// 		sign = '-'

// 	} else if (firstByte) == '+' {
// 		p.MoveByte()
// 	}

// 	state := NumberNone

// 	// 循环读取更多数字, 直到读完Reader或者遇到其它符号
// 	// 30
// 	// 0.12
// 	// +30
// 	// +0.12
// 	// +3.140000
// 	for {
// 		b, ok := p.PeekByte()
// 		if !ok {
// 			break
// 		}

// 		switch state {
// 		case NumberNone:
// 			// 0.23
// 			if b == '0' {
// 				state = NumberWantingPointOrEnd
// 				p.MoveByte()
// 				continue
// 			} else if IsByteAsciiNumberNo0(b) {
// 				// 那么就是数字
// 				state = NumberWantingAny
// 				num_diget_string = append(num_diget_string, string(b))
// 				p.MoveByte()
// 				continue
// 			} else {
// 				// 比如+x
// 				return 0, false
// 			}
// 		case NumberWantingPointOrEnd:
// 			if b == '.' {
// 				num_diget_string = append(num_diget_string, ".")
// 				state = NumberWaintingAnyNum
// 				p.MoveByte()
// 				continue
// 			} else {
// 				p.MoveByte()
// 				return 0, true
// 			}
// 		case NumberWantingAny:
// 			// +3.56
// 			if b == '.' {
// 				num_diget_string = append(num_diget_string, ".")
// 				state = NumberWaintingAnyNum
// 				p.MoveByte()
// 				continue
// 			} else if IsByteAsciiNumber(b) {
// 				num_diget_string = append(num_diget_string, string(b))
// 				state = NumberWantingAny
// 				p.MoveByte()
// 				continue
// 			} else {
// 				goto out
// 			}
// 		case NumberWaintingAnyNum:
// 			if !IsByteAsciiNumber(b) {
// 				goto out
// 			}
// 			num_diget_string = append(num_diget_string, string(b))
// 			p.MoveByte()
// 		}
// 	}

// out:

// 	bigSum := big.NewFloat(0)
// 	var stringNum string
// 	if sign == '+' {
// 		stringNum = strings.Join(num_diget_string, "")
// 	} else {
// 		stringNum = strings.Join(num_diget_string, "")
// 		stringNum = "-" + stringNum
// 	}

// 	f, ok := bigSum.SetString(stringNum)
// 	if !ok {
// 		panic("unexpected error")
// 	}

// 	sum, _ := f.Float64()
// 	return sum, true
// }
