package main

import (
	"fmt"
	"time"

	"github.com/maxbet1507/gowins/SingleInstanceApplication"
)

func main() {
	err, stop := SingleInstanceApplication.Check("test")
	defer stop()

	fmt.Println(err)

	if err == nil {
		time.Sleep(10 * time.Second)
	}
}
