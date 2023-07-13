package unmarshal

import "testing"

func TestParseBool(t *testing.T) {
	i, err := parse("  true   ")
	if err != nil {
		t.Errorf("unexpected")
	}

	b, ok := i.(bool)
	if !ok {
		t.Errorf("unexpected")
	}

	if !b {
		t.Errorf("unexpected")
	}

	_, err = parse("  tru e   ")
	if err == nil {
		t.Errorf("unexpected")
	}

	i, err = parse("  false   ")
	if err != nil {
		t.Errorf("unexpected")
	}
	if err != nil {
		t.Errorf("unexpected")
	}
	b, ok = i.(bool)
	if !ok {
		t.Errorf("unexpected")
	}
	if b {
		t.Errorf("unexpected")
	}

}
