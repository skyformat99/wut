package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	now := time.Now().Unix()
	unix := now
	flag.Parse() // we may add flags later.
	var err error
	if flag.NArg() > 0 {
		unix, err = strconv.ParseInt(flag.Args()[0], 10, 64)
		if err != nil {
			fmt.Printf("Invalid unix time: %v\n", flag.Args()[0])
			os.Exit(-1)
		}
	}
	t := time.Unix(unix, 0)

	fmt.Printf("%v is %v\n", unix, t)

}
