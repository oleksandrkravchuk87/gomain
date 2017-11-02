package functions

import (
	"fmt"
	"math"
	"strconv"

)

func F243a(x int) string {
	var (
		h, k, n int
		found bool
	)

	fmt.Println("x = ", x)

	n = int(math.Sqrt(float64(x)))
	found = false
	h, k = 1, n

	for h < x && k > 0 {
		if x == h*h+k*k {
			found = true
			break
		} else if (h*h + k*k) < x {
			h = h + 1
		} else {
			k = k - 1
		}

	}
	if found {
		return  strconv.Itoa(h)+"^2 +" + strconv.Itoa(k)+"^2 = " +  strconv.Itoa(h*h+k*k)
	} else {
		return "the entered value is not a sum of two squares"
	}
}

