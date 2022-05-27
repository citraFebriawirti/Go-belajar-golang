// There are several ways to create a variable at Golang, at this program we'll try two of them

package main

import "fmt"

func main(){
	// First way
	var nama string

	nama = "Nobita"
	fmt.Println(nama)

	nama = "Doraemon"
	fmt.Println(nama)

	// Second way
	var negara = "Jepang"
	fmt.Println(negara)

	var umur = 15
	fmt.Println(umur)
}