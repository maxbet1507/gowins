package main

import (
	"fmt"
	"time"

	"github.com/maxbet1507/gowins/SingleInstanceApplication"
)

func main() {
	sia, stop := SingleInstanceApplication.Check("test")
	defer stop()

	fmt.Println(sia)
	if sia {
		time.Sleep(10 * time.Second)
	}
}
