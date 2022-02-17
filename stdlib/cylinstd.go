package main

import "C"

import (
	"fmt"
	"time"
)

//export printi
func printi(x int32) {
	fmt.Print(x)
}

//export utime
func utime() int32 {
	return int32(time.Now().Unix())
}

func main() {}
