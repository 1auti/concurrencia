package main

import (
	"fmt"
	"time"
)

/**
 *
 * Cuándo elegir channel vs WaitGroup: si solo necesitás esperar → WaitGroup.
 * Si necesitás pasar datos/resultados entre goroutines → channel. Y se pueden combinar: WaitGroup para saber cuándo cerrar el channel, channel para los datos.
 * Eso lo vas a ver muy seguido.
 */

func ejercicio1() {
	urls := []string{
		"http://api.com/productos",
		"http://api.com/usuarios",
		"http://api.com/orders",
	}

	resultados := make(chan string)

	for _, url := range urls {
		go func(url string) {
			time.Sleep(500 * time.Millisecond)       // esto es simplemente para simular un trabajo que no hay
			resultados <- fmt.Sprintf("OK: %s", url) // enviamos los resultaods

			/**
			 * Caundo leemos recursos que no sabemos cuando pueden terner lo que agregamos es un close(resultados)
				*  SEIMPRE CIERRA EL EMISOR NUNCA EL RECEPTOR
			*/
		}(url)
	}

	for i := 0; i < len(resultados); i++ {
		fmt.Println(<-resultados)
	}

	fmt.Println("Todos las url fueron descargadas")
}

func procesar(url string) {
	fmt.Printf("Empezando la descarga %s\n ", url)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Terminando %s\n", url)
}
