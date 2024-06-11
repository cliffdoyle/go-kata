package main

import "fmt"

func main(){
	result:=""
for i:='z';i>='a';i--{
	if i%2==1{
		result+=string(i-32)
	}else{
		result+=string(i)
	}
}

fmt.Println(result)

}