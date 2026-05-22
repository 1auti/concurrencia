package main

import "fmt"

func ejemplo3() {
	ch := make(chan int, 3) // Channel con buffer 3 | PUEDE TENER HASTA 3 ELEMENTOS SIN BLOQUEAR!!!!

	// Estos 3 envios no se bloquean ( PORQUE HAY ESPACIO )
	ch <- 1
	ch <- 2
	ch <- 3

	// Esto seria bloqueado porque el buffer esta lleno
	// ch <- 4
	//
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
