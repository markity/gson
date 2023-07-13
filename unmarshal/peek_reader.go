package unmarshal

type peekReader struct {
	Data   []byte
	CurPos int
	Length int
}

func newPeekReader(data []byte) *peekReader {
	if len(data) == 0 {
		return nil
	}
	return &peekReader{
		Data:   data,
		CurPos: 0,
		Length: len(data),
	}
}

func (pr *peekReader) MoveByte() (byte, bool) {
	if pr.CurPos < pr.Length {
		pr.CurPos++
		return pr.Data[pr.CurPos-1], true
	}

	return 0, false
}

func (pr *peekReader) PeekByte() (byte, bool) {
	if pr.CurPos < pr.Length {
		return pr.Data[pr.CurPos], true
	}

	return 0, false
}

func (pr *peekReader) IsEnd() bool {
	return pr.CurPos == pr.Length
}
