package main

import "fmt"

func main(){
	result:=""
for i:='a';i<='z';i++{
	if i%2==0{
		result+=string(i-32)
	}else{
		result+=string(i)
	}
}

fmt.Println(result)

}