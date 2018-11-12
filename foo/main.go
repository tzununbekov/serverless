package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%s\n", os.Getenv("FUNCTION"))
}
