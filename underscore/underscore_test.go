package underscore_test

import (
	"fmt"
	"testing"

	"github.com/dmcclung/twg/underscore"
)

func TestCamel(t *testing.T) {
	testCases := []struct {
			arg string
			want string
		}{
			{
				"thisIsCamelCase", "this_is_camel_case",
			},
			{
				"aFuncZ", "a_func_z",
			},		
			{
				"alllowercase", "alllowercase",
			},
			{
				"snake_case_already", "snake_case_already",
			},
		}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(tc.arg), func(t *testing.T) {
			if got := underscore.Camel(tc.arg); got != tc.want {
				t.Errorf("Camel(%s), want %s, got %s", tc.arg, tc.want, got)
			}
		})
	}
}