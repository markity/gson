package unmarshal

import (
	"errors"
)

func skipSpace(p *peekReader) {
	for {
		b, ok := p.PeekByte()
		if !ok {
			return
		}

		if b == ' ' || b == '\t' || b == '\n' {
			p.MoveByte()
			continue
		}
		return
	}
}

func IsByteAsciiNumberNo0(b byte) bool {
	switch b {
	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}

func isByteAsciiNumber(b byte) bool {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	}
	return false
}

func parseValue(p *peekReader) (interface{}, bool) {
	v, ok := p.PeekByte()
	if !ok {
		return nil, false
	}

	switch {
	// null, ok
	case v == 'n':
		if !parseNull(p) {
			return nil, false
		}
		return nil, true
	// obj, ok
	case v == '{':
		return parseObj(p)
	// false, ok
	case v == 'f':
		if !parseFalse(p) {
			return nil, false
		}
		return false, true
	// true, ok
	case v == 't':
		if !parseTrue(p) {
			return nil, false
		}
		return true, true
	// array, ok
	case v == '[':
		return parseArray(p)
	// string, ok
	case v == '"':
		return parseString(p)
	// number, ok
	case isByteAsciiNumber(v) || v == '-' || v == '+':
		return parseNumber(p)
	}

	return nil, false

}

func parse(jsonData string) (interface{}, error) {
	// 跳过空白字符
	pr := newPeekReader([]byte(jsonData))
	skipSpace(pr)
	if pr.IsEnd() {
		return nil, errors.New("format error: no value")
	}

	i, ok := parseValue(pr)
	if !ok {
		return nil, errors.New("format error")
	}

	skipSpace(pr)
	if !pr.IsEnd() {
		return nil, errors.New("format error")
	}

	return i, nil
}
