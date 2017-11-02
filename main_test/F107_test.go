package main_test

import (
	"testing"

	"strings"

	"main/functions"
)



func TestF107(t *testing.T) {
	result := functions.F107(25)
	if !strings.EqualFold(result, "the biggest integer k, fulfilling the condition  4^k < m is 2") {
		t.Errorf("err", "Wanted: %v, Got: %v", "the biggest integer k, fulfilling the condition  4^k < m is 2", result)
	}
	result = functions.F107(2)
	if !strings.EqualFold(result, "the biggest integer k, fulfilling the condition  4^k < m is 0") {
		t.Errorf("err", "Wanted: %v, Got: %v", "the biggest integer k, fulfilling the condition  4^k < m is 0", result)
	}
	result = functions.F107(4)
	if !strings.EqualFold(result, "the biggest integer k, fulfilling the condition  4^k < m is 0") {
		t.Errorf("err", "the biggest integer k, fulfilling the condition  4^k < m is 0", result)
	}
	result = functions.F107(0)
	if !strings.EqualFold(result, "no m values can be found for 0") {
		t.Errorf("err", "no m values can be found for 0", result)
	}
	result = functions.F107(-10)
	if !strings.EqualFold(result, "no m values can be found for -10") {
		t.Errorf("err", "no m values can be found for -10", result)
	}
}

