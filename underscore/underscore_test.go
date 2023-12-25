package underscore_test

import (	
	"testing"

	"github.com/dmcclung/twg/underscore"
)

func TestCamel(t *testing.T) {
	testCases := map[string]struct {
			arg string
			want string
		}{
			"camel case": { "thisIsCamelCase", "this_is_camel_case" },
			"two capitals": { "aFZ", "a_f_z" },
			"all lower case": { "alllowercase", "alllowercase" },
			"snake case": { "snake_case_already", "snake_case_already" },
		}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := underscore.Camel(tc.arg); got != tc.want {
				t.Errorf("Camel(%s), want %s, got %s", tc.arg, tc.want, got)
			}
		})
	}
}