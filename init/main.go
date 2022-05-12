package main

import (
	"fmt"

	"github.com/tzuhsitseng/labs-go/init/lib"
)

func init() {
	fmt.Println("init main")
}

func main() {
	fmt.Println("do main")
	lib.Do()
}
