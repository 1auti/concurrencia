package main

import (
	"fmt"
	"time"
)

func ejemplo4() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "mensaje de voz"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "mensaje de voz"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Recibiendo mensaje  de ch 1", msg1)
		case msg2 := <-ch2:
			fmt.Println("Recibiendo mensaje de ch2", msg2)
		}
	}

	// Tambien se puede hacer un select con timeout
	select {
	case resultado := <-ch1:
		fmt.Println("resultado:", resultado)
	case <-time.After(3 * time.Second):
		// time.After retorna un channel que recibe después de N segundos
		fmt.Println("timeout — la operación tardó demasiado")
	}
}
