package main

import (
	"fmt"
)

func main() {
	x := 25
	fmt.Println(x)
	fmt.Println(&x)
	y := &x
	fmt.Println(y)  // imprime la dirección de x
	fmt.Println(*y) // da el valor de la direccion y

}

func cambiarValor(a int) {
	fmt.Println(&a)
	a = 36
}
