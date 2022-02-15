package solution_test

import (
	"strings"
	"testing"
)

func solution_test(t *testing.T) {
	expected := string([]rune{72, 101, 108, 108, 111, 32, 128506, 65039, 32, 33})
	adminText2 := string("Hello :world_map: !")

	if !strings.EqualFold(adminText2, expected) {
		t.Errorf("Unexpected result:\n\tExpected: %q\n\tGot: %q", expected, adminText2)
	}
}