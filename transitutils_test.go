package transitutils_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/russolsen/transit"
	"github.com/tmc/transitutils"
)

func TestToGo(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{`[]`, `[]`},
		{`["^ ", "~:x/y", "foo"]`, `{":x/y":"foo"}`},
		{`["^ ", "~:x/y", "foo", "~:x/z", []]`, `{":x/y":"foo",":x/z":[]}`},
		{`["^ ", "~:x/z", [["^ ", "~:x/y", "foo"]]]`, `{":x/z":[{":x/y":"foo"}]}`},
	}
	for _, tt := range cases {
		buf := strings.NewReader(tt.in)
		decoder := transit.NewDecoder(buf)
		value, err := decoder.Decode()
		if err != nil {
			t.Fatal(err)
		}
		v, err := transitutils.ToGo(value)
		if err != nil {
			t.Fatal(err)
		}
		j, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		got := string(j)
		if tt.out != got {
			t.Errorf("got '%s', expected '%s'", got, tt.out)
		}
	}
}
