package unmarshal

func parseArray(p *peekReader) ([]interface{}, bool) {
	p.MoveByte()

	res := make([]interface{}, 0)

	for {
		skipSpace(p)
		b, ok := p.PeekByte()
		if !ok {
			return nil, false
		}

		if b == ']' {
			p.MoveByte()
			return res, true
		} else {
			subval, ok := parseValue(p)
			if !ok {
				return nil, false
			}
			res = append(res, subval)
			skipSpace(p)
			b2, ok := p.PeekByte()
			if !ok {
				return nil, false
			}
			if b2 == ',' {
				p.MoveByte()
				continue
			} else if b2 == ']' {
				p.MoveByte()
				return res, true
			}
		}
	}
}
