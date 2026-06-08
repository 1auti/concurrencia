package main

import (
	"fmt"
	"time"
)

func ejemploSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Esto seria un gour que es lenta porque tiene mas trabajo
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "respuestade servicio 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "respuesta servicio 2"
	}()

	for i := 0; i < 2; i++ {
		/* EL SELECT es el swith de los canales pero tiene una diferencia en el comportamineto que es interesante
		 * Un Switch evalua casos de arriba hacia abajo y entra al primero que matchea
			* Un Select espera hasta que alguna de sus operaciones termine para avanzar. Si varias estan listas, eligue al azar...
		*/
		select {
		case msg1 := <-ch1:
			fmt.Println("Recibi:", msg1)

		case msg2 := <-ch2:
			fmt.Println("Recibi: ", msg2)

		case <-time.After(3 * time.Second): // Devuelve un canal que dispara a los 3 seg | ES MUY UTIL API REST QUE LLAMA A OTRO SERVICIO
			fmt.Println("Timeout: el servicio tardo demasiado")

		// case <- done :
		// return // Otro gourn cerro o envio por donde para avisar frenar

		default: // SI no hay un canal que esta lista -> respuesta default
			fmt.Println("No hay nada todovia para procesar") // IMPORTANTE CON DEFAULT EL SELECT NO ESPERA NADA | SIN DEFAULT EL SELECET ESPERA
		}
	}
}
