package main

func main() {
	/* 	var wg sync.WaitGroup // Wait Group es un contador de gournetes ( serian Hilos )

	   	for i := 0; i <= 5; i++ {
	   		wg.Add(1)

	   		n := i

	   		go func() {
	   			// Decrementa los gournetes cuando se termina | Defer garantiza que se llamen aunque haya un panic
	   			defer wg.Done()

	   			fmt.Printf("Gournetes  %d ejecutandose \n", n)
	   		}()

	   	}

	   	wg.Wait() // Lo que hace es bloquear cuando el contador termina en 0
	   	fmt.Println("Todos los gournetes se terminaron")

	*/

	ejemploSelect()
}
