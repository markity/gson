package unmarshal

import "testing"

func TestParseString(t *testing.T) {
	tests := []struct {
		input      string
		expectedOK bool
		expect     string
	}{
		{input: ` " -0.314 " `, expectedOK: true, expect: " -0.314 "},
		{input: ` " -0.314 \r " `, expectedOK: true, expect: " -0.314 \r "},
		{input: ` " \"-0.314 \r " `, expectedOK: true, expect: " \"-0.314 \r "},
		{input: ` " \"-0.314 \b " `, expectedOK: true, expect: " \"-0.314 \b "},
		{input: ` " \"-0.314 \r " `, expectedOK: true, expect: " \"-0.314 \r "},
		{input: ` " \"-0.314 \t \t" `, expectedOK: true, expect: " \"-0.314 \r \t"},
		{input: ` " \"-0.314 \n \t" `, expectedOK: true, expect: " \"-0.314 \n \t"},
		{input: ` " \"-0.314 \\ \/ \t \f" `, expectedOK: true, expect: " \"-0.314 \\ / \t \f"},
		{input: ` " \"-0.314 \\ \/ \t \f `, expectedOK: false},
		{input: ` " 测试utf8 " `, expectedOK: true, expect: " 测试utf8 "},
		{input: `    `, expectedOK: false},
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
			if val.(string) != v.expect {
				t.Errorf("expected %v, but got %v: input is %v", v.expect, val, v.input)
			}
		}
	}
}
