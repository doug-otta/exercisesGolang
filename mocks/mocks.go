package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const ultimaPalavra = "Vai!"
const inicioContagem = 3

// Sleeper te permite definir pausas
type Sleeper interface {
	Pausa()
}

// SleeperPadrao é uma impçementação de Sleeper com um atraso pré-definido
type SleeperPadrao struct{}

type SleeperSpy struct {
	Chamadas int
}

// Vai pausar a execução pela Duração definida
func (d *SleeperPadrao) Pausa() {
	time.Sleep(1 * time.Second)
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

// Contagem imprime uma contagem de 3 para a saída com um atraso determinado por Sleeper
func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}
	sleeper.Pausa()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}
