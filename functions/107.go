
package functions

import (
	"fmt"
	"math"

	"strconv"
)

func F107(m int) string {
	var (
		k int
	)

		fmt.Println("m =  ", m)
		if (m <= 1) {
			return "no m values can be found for " + strconv.Itoa(m)
		} else {
			for k = 0; ; {
				if math.Pow(4, float64(k+1)) >= float64(m) {
					break
				} else {
					k++
				}
			}
			fmt.Println(k)
			return ("the biggest integer k, fulfilling the condition  4^k < m is " + strconv.Itoa(k))
		}
	}


