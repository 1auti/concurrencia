package main

import (
	"fmt"
	"sync"
	"time"
)

// WORKER es un goroutine que procesa los trabajos del channels jobs y mandar los resultadosal channel results
func worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	// jobs <- chan int  lo que hace es recibir ( SOLO LECTURA )
	// result chan <- int es que envia ( SOLO ESCRITURA )

	defer wg.Done()

	for job := range jobs {
		fmt.Printf("El worker %d procesando el jobs %d\n", id, job)
		time.Sleep(500 * time.Millisecond)

		result <- job * 2 // Mandamos los resultados
	}
}

func ejemplo5() {
	const numJObs = 10
	const numWorkers = 3 // solo procesa 3 goroutines pricesando en paralelo

	jobs := make(chan int, numJObs)
	results := make(chan int, numJObs)

	// Lanzamos los pool de los workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Mandamos los trabajos al channel
	for j := 1; j <= numJObs; j++ {
		jobs <- j
	}

	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()
}
