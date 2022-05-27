package main

import "fmt"

func main(){
	fmt.Print("Nama Depan : ")
	var fname string
	fmt.Scan(&fname)

	fmt.Print("Nama Belakang : ")
	var lname string
	fmt.Scan(&lname)
	fmt.Println(fname+" "+lname)
}