package matematica

import "errors"

var A string = "SHOW"

func Soma(a int, b int) (int, error) {
	res := a + b
	if res > 10 {
		return 0, errors.New("Soma maior que 10 ")
	}

	return a + b, nil
}
