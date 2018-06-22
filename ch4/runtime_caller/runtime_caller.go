package main

import (
	"fmt"
	"runtime"
)

func a(){
	_, file, line, _ := runtime.Caller(1)
	fmt.Println("===",file,line)
}

func main() {
	a()
}