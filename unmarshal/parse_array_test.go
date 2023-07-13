package unmarshal

import (
	"fmt"
	"testing"
)

func TestParseArray(t *testing.T) {
	res, err := parse("[ 1,  2  , 3, 3.4 ]")
	if err != nil {
		t.Errorf("unexpected error: %v, input is %v", err, "[ 1,  2  , 3, 3.4 ]")
		return
	}
	if fmt.Sprint(res) != "[1 2 3 3.4]" {
		t.Errorf("unexpected result: %v, input is %v", res, "[ 1,  2  , 3, 3.4 ]")
	}
}
