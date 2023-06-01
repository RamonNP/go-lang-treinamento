package main

import (
	"fmt"
	"log"

	"github.com/go-lang-treinamento/hello-go/hello-go/matematica"
)

func main() {
	fmt.Println("Hello Go")
	resultado, err := matematica.Soma(11, 2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%v", resultado)
	fmt.Printf("%T", resultado)
}
