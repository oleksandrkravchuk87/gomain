package main

import (
	"testing"

	"strings"

	"main/functions"
)



func TestF243b(t *testing.T) {
	_, str := functions.F243b(0)
	if !strings.EqualFold(str, "the entered value is not a sum of two squares") {
		t.Errorf("err", "the entered value is not a sum of two squares", str)
	}
	_, str = functions.F243b(1)
	if !strings.EqualFold(str, "the entered value is not a sum of two squares") {
		t.Errorf("err", "the entered value is not a sum of two squares", str)
	}

	s := make([]string, 5)
	s = append(s, "30^2 +10^2 = 1000")
	s = append(s, "26^2 +18^2 = 1000")

	r,_ := functions.F243b(1000)
	sliceEquals := isSliceEquals(s, r)
	if !sliceEquals {
		t.Errorf("err", "30^2 +10^2 = 1000 26^2 +18^2 = 1000", r)
	}


	_, str = functions.F243b(-1)
	if !strings.EqualFold(str, "the entered value is not a sum of two squares") {
		t.Errorf("err", "the entered value is not a sum of two squares", str)
	}

}

func isSliceEquals(s, r []string) bool {

	if s == nil && r == nil {
		return true;
	}

	if s == nil || r == nil {
		return false;
	}

	if len(s) != len(r) {
		return false
	}

	for i := range s {
		if s[i] != r[i] {
			return false
		}
	}

	return true
}
