package unmarshal

import "testing"

func TestParseNull(t *testing.T) {
	i, err := parse(" null  ")
	if err != nil {
		t.Errorf("unexpected")
	}

	if i != nil {
		t.Errorf("unexpected")
	}

	_, err = parse(" nul l  ")
	if err == nil {
		t.Errorf("unexpected")
	}
}
