package unmarshal

func parseFalse(p *peekReader) bool {
	if b, ok := p.MoveByte(); !ok || b != 'f' {
		return false
	}
	if b, ok := p.MoveByte(); !ok || b != 'a' {
		return false
	}
	if b, ok := p.MoveByte(); !ok || b != 'l' {
		return false
	}
	if b, ok := p.MoveByte(); !ok || b != 's' {
		return false
	}
	if b, ok := p.MoveByte(); !ok || b != 'e' {
		return false
	}
	return true
}

func parseTrue(p *peekReader) bool {
	if b, ok := p.MoveByte(); !ok || b != 't' {
		return false
	}
	if b, ok := p.MoveByte(); !ok || b != 'r' {
		return false
	}
	if b, ok := p.MoveByte(); !ok || b != 'u' {
		return false
	}
	if b, ok := p.MoveByte(); !ok || b != 'e' {
		return false
	}
	return true
}
