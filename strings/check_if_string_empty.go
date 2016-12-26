// Compile and run as,
// export GOPATH=$PWD
// go build check_if_string_empty.go
// ./check_if_string_empty

package main

import (
	"fmt"
)

func main() {
	var myname string //declare string variable
	myname = "Lynxbee"

	// check if myname string is empty
	if myname == "" { 
		fmt.Println("first method to check string is empty")
	}

	if len(myname) == 0 { 
		fmt.Println("second method to check string is empty")
	}

	// check if myname string is not empty 
	if myname != "" { 
		fmt.Println("first method to check string is not empty")
	}

	if len(myname) != 0 { 
		fmt.Println("second method to check string is not empty")
	}
}
