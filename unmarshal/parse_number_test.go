package unmarshal

import "testing"

func TestParseNumber(t *testing.T) {
	tests := []struct {
		input      string
		expectedOK bool
		expect     float64
	}{
		{input: "  -0.314  ", expectedOK: true, expect: -0.314},
		{input: "  +30.5  ", expectedOK: true, expect: -0.314},
		{input: "  -60  ", expectedOK: true, expect: -0.314},

		{input: "  -+0.314  ", expectedOK: false},
		{input: "  x-0.314x  ", expectedOK: false},
		{input: "  -0.314x  ", expectedOK: false},
		{input: "  -0x.314x  ", expectedOK: false},
		{input: "  -0.31x4x  ", expectedOK: false},
		{input: "  -0.314x  ", expectedOK: false},
		{input: "  --3  ", expectedOK: false},
	}

	for _, v := range tests {
		val, err := parse(v.input)
		if !v.expectedOK && err == nil {
			t.Errorf("expected not ok, but got no err: input is %v", v.input)
		}
		if v.expectedOK && err != nil {
			t.Errorf("expected ok, but got err: input is %v", v.input)
		}
		if v.expectedOK && err != nil {
			if val.(float64) != v.expect {
				t.Errorf("expected %v, but got %v: input is %v", v.expect, val, v.input)
			}
		}
	}

}
