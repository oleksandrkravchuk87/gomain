package functions

import (
	"fmt"
	"math"

	"strconv"
)

func F243b(x int) ([]string, string) {
	var (
		h, k, n int

		found      bool
	)


	s := make([]string, 5)
	fmt.Println("x = ", x)

	n = int(math.Sqrt(float64(x )))
	found = false
	k, h = 1, n

	for k < x && h > 0 && k<=h {
		if x == k*k+h*h {
			found = true
			s = append(s, strconv.Itoa(h)+"^2 +" + strconv.Itoa(k)+"^2 = " +  strconv.Itoa(h*h+k*k))
			k = k + 1
		} else if (k*k + h*h) < x {
			k = k + 1
		} else {
			h = h - 1
		}

	}
	if !found {
		return nil, "the entered value is not a sum of two squares"
	} else {
		return s,""
	}
}

