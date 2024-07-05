package main

import (
	"fmt"
	"test-task/infra"
)

func main() {
	i := infra.New("/.env")
	fmt.Print(i.Config())
}
