package main

import (
	"fmt"
	"strconv"
)

func main() {

	// balikannya ada errornya, jika konversinya menyebabkan error
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

}
