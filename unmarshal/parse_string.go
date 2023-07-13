package unmarshal

import "unicode/utf8"

const (
	StringWantingAny int = iota
	StringWantingSpecial
)

// 第一个字节是存在的
func parseUtf8Character(p *peekReader) (rune, bool) {
	b, _ := p.PeekByte()

	bs := make([]byte, 0, 6)

	switch {
	case b&0b10000000 == 0b00000000:
		p.MoveByte()
		bs = append(bs, b)
	case b&0b11100000 == 0b10000000:
		p.MoveByte()
		bs = append(bs, b)

		b2, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b2)
	case b&0b11110000 == 0b11100000:
		p.MoveByte()
		bs = append(bs, b)

		b2, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b2)

		b3, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b3)
	case b&0b11111000 == 0b11110000:
		p.MoveByte()
		bs = append(bs, b)

		b2, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b2)

		b3, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b3)

		b4, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b4)
	case b&0b11111100 == 0b11111000:
		p.MoveByte()
		bs = append(bs, b)

		b2, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b2)

		b3, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b3)

		b4, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b4)

		b5, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b5)
	case b&0b11111110 == 0b11111100:
		p.MoveByte()
		bs = append(bs, b)

		b2, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b2)

		b3, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b3)

		b4, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b4)

		b5, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b5)

		b6, ok := p.MoveByte()
		if !ok {
			return 0, false
		}
		bs = append(bs, b6)
	}

	r, _ := utf8.DecodeRune(bs)

	return r, true
}

func parseString(p *peekReader) (string, bool) {
	p.MoveByte()

	state := StringWantingAny

	res := ""

	for {
		b, ok := p.PeekByte()

		if !ok {
			return "", false
		}

		switch state {
		case StringWantingAny:
			if b == '\\' {
				state = StringWantingSpecial
				p.MoveByte()
				continue
			} else if b == '"' {
				p.MoveByte()
				return res, true
			} else {
				rune, ok := parseUtf8Character(p)
				if !ok {
					return "", false
				}
				res += string(rune)
				continue
			}
		case StringWantingSpecial:
			switch b {
			case '"':
				res += string(`"`)
			case '\\':
				res += string(`\`)
			case 'r':
				res += string('\r')
			case 'b':
				res += string('\b')
			case 't':
				res += string('\t')
			case 'f':
				res += string('\f')
			case 'n':
				res += string('\n')
			case '/':
				res += string('/')
			// case 'U':
			//	panic("unspported")
			// TODO: 支持UNICODE翻译
			default:
				return "", false
			}
			state = StringWantingAny
			p.MoveByte()
			continue
		}
	}

}
