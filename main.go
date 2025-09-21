package main
// go er each file e ekta kore package takhbe 
// main go file er package hote hobe main 
// fmt holo formet package eta holo go er built in package 
// jkn kono variable declare kori tokn sei variable ta ram e store hoyea takhe 
// Data Type 1 Numeric 2 Boolean 3 String
// Interger Number      Floating Point Number 
// 10 						10.05 float32 float64
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
import "fmt"

func main(){
	// // var age int = 22
	// age :=18
	// // var age =10
    // if age>18 {
	// 	fmt.Println("you are eligble to be married")
	// } else if age<18 {
	// 	fmt.Println("You are not eligible to be marriend, but you can love someone ")
	// }else if age == 18{
	// 	fmt.Println("You are just a teenager,not eligble to be married")
	// }
	age := 20
	sex:="male"
	if age>60 || sex =="male" {
		fmt.Println("You are ready to be married!")
	}
}