package main

import (
	"fmt"

	belajar "github.com/Go-Master-Code/belajar"
	//repos "github.com/Go-Master-Code/repos"
	repos "github.com/Go-Master-Code/repos/v2" //module dengan major changes
)

func main() {
	belajar.LogInfo("Ini adalah method dari module belajar")
	repos.LogInfo("Berhasil import method dari module repos")
	//fmt.Println(repos.SayHello()) //method v1.5.0
	fmt.Println(repos.SayHello("Hasil major changes"))    //method tambahan dari versi modul v2.0.0
	fmt.Println(repos.SayHello("Hasil major changes 2x")) //method tambahan dari versi modul v2.0.0
}
