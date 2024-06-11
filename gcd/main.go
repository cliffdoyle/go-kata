package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
args:= os.Args[1:]
if len(args) != 2{
	return
}
args1,_:=strconv.Atoi(args[0])
args2,_:=strconv.Atoi(args[1])

gcd:=1

for i:=1;i<=args1 && i <= args2;i++{
	if args1 % i ==0 && args2 % i==0{
		gcd=i
	}

}
fmt.Println(gcd)


}