package main

import (
	"fmt"
	"github.com/Nikolay-Ishev/Proxiad-GoAcademy/tree/master/01-intro-lab/stringutil"
	"rsc.io/quote"

)

func main() {
	s := "Hello from Golang!"
	fmt.Println(s)
	goquote := quote.Go()
	fmt.Println(goquote)
	fmt.Println(stringutil.Reverse(goquote))
}