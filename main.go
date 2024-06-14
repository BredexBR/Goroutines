package main

import (
	"fmt"  //funcionalidades para formatação de entrada e saída.
	"time" //funcionalidades relacionadas à manipulação de tempo.
)

func worker(workerID int, data chan int) { // recebe cada goroutine e exibe worker e seus valores.
	for x := range data {
		fmt.Printf("Worker %d recebeu %d\n", workerID, x)
		time.Sleep(time.Second) //espera 1 segundo
	}
}

func main() { //passar valores para as goroutines
	canal := make(chan int)
	qtdWorkers := 10 //goroutines

	for i := range qtdWorkers { // cria 10 goroutines
		go worker(i, canal)
	}

	for i := range 100 { //enviar 100 valores para o canal
		canal <- i
	}
}
