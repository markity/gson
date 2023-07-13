package unmarshal

import (
	"fmt"
	"testing"
)

func TestParseObj(t *testing.T) {
	d, err := parse(`  {  "hello" :  "world"    }   `)
	if err != nil {
		t.Errorf("unexpected")
	}

	m, ok := d.(map[string]interface{})
	if !ok {
		t.Errorf("unexpected")
	}

	if fmt.Sprint(m) != "map[hello:world]" {
		t.Errorf("unexpected result %v", fmt.Sprint(m))
	}

}
