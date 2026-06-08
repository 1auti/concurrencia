package main

import (
	"fmt"
	"sync"
)

// Ejecutar este comando :  go run -race .
func ejemploSinMutex() {
	contador := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador++ // ← 1000 goroutines tocan esto a la vez
		}()
	}

	wg.Wait()
	fmt.Println(contador) // ¿Sale 1000?
}

/*
*  REGLAS A LA HORA DE USAR EL MUTEX:
  - 1 ) Siempre defer mu.Unlock() inmediatamente después del mu.Lock(). Si los separás, algún día un return o panic en el medio dejará el mutex bloqueado para siempre (deadlock).
    2 ) El mutex va con el dato que protege. No lo pongas en otro lado. Lo ideal es que vivan en el mismo struct.
    3 ) Nunca copiés un mutex. Si lo pasás a una función, pasalo por puntero. Go hasta tiene una regla del linter para esto: sync.Mutex no debe copiarse después de primer uso.
*/
func ejemploConMutex() {
	contador := 0
	var wg sync.WaitGroup
	var mu sync.Mutex // el guardián del contador

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Lock: "entro a la sección crítica, nadie más puede entrar"
			// Si otra goroutine ya tiene el lock, esta se BLOQUEA hasta que lo libere.
			mu.Lock()
			// defer garantiza que siempre desbloqueemos, incluso ante un panic.
			// Es el equivalente de tu try/finally { lock.unlock(); } en Java.
			defer mu.Unlock()

			contador++ // ahora sí: solo una goroutine a la vez llega acá
		}()
	}

	wg.Wait()
	fmt.Println(contador) // siempre 1000
}

// POr lo general en PROD EL MUTEX VA ASOCIADO A UN STRUCT
type SafeCounter struct {
	mu    sync.Mutex // no exportado nadie de afuera toca el mutex
	value int
}

// Estandarizamos que increment es la unica manera de leerlo
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Asi estandarizams que la funcion Value es la unica manera de leerlo
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

/** Hay casos donde tenemos que leer mas que escribir. Imaginate que tenemos un Cache o un Config lo que vamos a hacer ahi es leer no escribir
* Ahi nos sirve mas un propiedad llamada RWMutex
*
* Múltiples goroutines pueden tener RLock al mismo tiempo.
// Se bloquea solo si alguien tiene un Lock completo (escritura).
mu.RLock()
defer mu.RUnlock()
// ... leer

// Lock completo: bloquea tanto lectores como escritores.
mu.Lock()
defer mu.Unlock()
// ... escribir
//
//
//
// ANALOGIA !!!!!!!!!!!!!
//
//  RWMutex es como una sala de reuniones con una regla — pueden entrar muchas personas a leer los documentos al mismo tiempo, pero si alguien necesita modificarlos,
//  todos salen y espera a estar solo. El sync.Mutex básico en cambio solo deja entrar a una persona a la vez, aunque solo vaya a leer.
*/
