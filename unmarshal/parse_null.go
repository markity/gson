package unmarshal

func parseNull(p *peekReader) bool {
	if b, ok := p.MoveByte(); !ok || b != 'n' {
		return false
	}
	if b, ok := p.MoveByte(); !ok || b != 'u' {
		return false
	}
	if b, ok := p.MoveByte(); !ok || b != 'l' {
		return false
	}
	if b, ok := p.MoveByte(); !ok || b != 'l' {
		return false
	}
	return true
}
