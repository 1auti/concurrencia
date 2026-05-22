package main

import "fmt"

func ejemplo2() {
	ch := make(chan string) // Creamos un canal de string

	go func() {
		ch <- "primer mensaje" // se tendria que bloquear cuando se lee
		ch <- "segundo mensaje"
		ch <- "tercer mensaje"

		close(ch) // cierra el canal | indicando que no hay mas datos
	}()

	for msg := range ch {
		fmt.Println("recibido:", msg)
	}

	fmt.Println("channel cerrado, fin")
}
