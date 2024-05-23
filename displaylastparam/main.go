package main

import (
	"fmt"
	"os"
)


func main(){
args:=os.Args[1:]

var lastparam string

for _,arg:= range args{
	lastparam=arg

}

fmt.Println(lastparam)

}