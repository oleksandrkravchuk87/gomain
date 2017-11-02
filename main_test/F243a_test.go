package main

import (
	"testing"

	"strings"

	"main/functions"
)



func TestF243a(t *testing.T) {
	result := functions.F243a(25)
	if !strings.EqualFold(result, "3^2 +4^2 = 25") {
		t.Errorf("err", "3^2 +4^2 = 25", result)
	}
	result = functions.F243a(10000)
	if !strings.EqualFold(result, "28^2 +96^2 = 10000") {
		t.Errorf("err", "28^2 +96^2 = 10000", result)
	}
	result = functions.F243a(42)
	if !strings.EqualFold(result, "the entered value is not a sum of two squares") {
		t.Errorf("err", "the entered value is not a sum of two squares", result)
	}
	result = functions.F243a(0)
	if !strings.EqualFold(result, "the entered value is not a sum of two squares") {
		t.Errorf("err", "the entered value is not a sum of two squares", result)
	}
	result = functions.F243a(-10)
	if !strings.EqualFold(result, "the entered value is not a sum of two squares") {
		t.Errorf("err", "the entered value is not a sum of two squares", result)
	}
}


