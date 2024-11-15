package main

import (
	"github.com/ParsaImi/test-go/encrypt"
	"fmt"
	"github.com/ParsaImi/test-go/decrypt"


)

func main (){
	encrypted := encrypt.Nimbus("imicorp")
	fmt.Println(decrypt.Nimbus(encrypted))
}