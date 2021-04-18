package main

import (
	"fmt"
)

func getServiceName() string {
	return "bookdir"
}

func main() {
	fmt.Println(getServiceName())
}
