package main

import(
	"tasks/homework"
     "fmt"
	 
)

func main(){
	//fmt.Println(homework.Kabisa(400))
    //fmt.Println(homework.IsPrime(233))
    var a int
    fmt.Scan(&a)
    result := homework.IsDiv8(a)
    fmt.Println(result)
}