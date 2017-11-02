package main

import (
	"main/functions"
	"fmt"
	"os"
	"strconv"

)

func main() {

	var (err error
		v string
		m int)

	if (len(os.Args) > 1) {
		v = os.Args[1]
	} else {fmt.Println("enter an integer positive value")
		fmt.Scan(&v)}
	m, err = strconv.Atoi(v)
	if err != nil {
		fmt.Println("Error")
	} else {
		slice, str := functions.F243b(m)
		if slice != nil {
			for _, element := range slice {
				fmt.Println(element)
			}
		} else {
			fmt.Println(str)
		}
	}


}
