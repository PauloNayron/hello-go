package main

import (
	"errors"
	"fmt"
)

func main() {
	/*// exemplo de consulta com erro
	res, err := http.Get("http://gooogle.com.br")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res.Header)
	*/

	/*// exemplo com a func soma
	res, err := soma(10,10)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
	*/

	// blank identifier
	res, _ := soma(7,2)
	fmt.Println(res)
}

func soma(x int, y int) (int, error) {
	res := x + y
	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}
	return res, nil
}