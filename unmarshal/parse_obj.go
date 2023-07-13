package unmarshal

const (
	DictWantingAny = iota
	DictWantingColon
	DictWantingValue
)

// 第一个字符是{
// {}
// {"hello":1.23}
func parseObj(p *peekReader) (map[string]interface{}, bool) {
	p.MoveByte()

	res := make(map[string]interface{})

	state := DictWantingAny
	var _key string

	for {
		skipSpace(p)
		b, ok := p.PeekByte()
		if !ok {
			return nil, false
		}

		switch state {
		case DictWantingAny:
			if b == '}' {
				p.MoveByte()
				return res, true
			}

			if b != '"' {
				return nil, false
			}

			key, ok := parseString(p)
			if !ok {
				return nil, false
			}

			state = DictWantingColon
			_key = key
			continue
		case DictWantingColon:
			if b != ':' {
				return nil, false
			}

			state = DictWantingValue
			p.MoveByte()
			continue
		case DictWantingValue:
			val, ok := parseValue(p)
			if !ok {
				return nil, false
			}

			res[_key] = val

			// 消耗掉可能的逗号
			skipSpace(p)
			b2, ok := p.PeekByte()
			if !ok {
				return nil, false
			}

			if b2 == '}' {
				p.MoveByte()
				return res, true
			} else if b2 == ',' {
				p.MoveByte()
				state = DictWantingAny
				continue
			}
		}
	}
}
