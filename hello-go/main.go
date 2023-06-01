package main

import (
	"github.com/go-lang-treinamento/hello-go/hello-go/matematica"
	"fmt"
)

func main()  {
	fmt.Println("Hello Go");
	resultado := matematica.Soma(1,2)
	fmt.Printf("%v", resultado)
	fmt.Printf("%T", resultado)
}