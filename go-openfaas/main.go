package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Os env \"function\" value is %q\n", os.Getenv("FUNCTION"))
}
