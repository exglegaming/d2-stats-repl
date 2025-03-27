package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Print("Thank you for using the Destiny 2 Stats app!... Have a good day!")
	os.Exit(0)
	return nil
}
