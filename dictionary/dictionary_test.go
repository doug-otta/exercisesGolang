package main

import (
	"testing"
)

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é apenas um teste"}

	t.Run("palavra conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é apenas um teste"

		comparaString(t, resultado, esperado)
	})

	t.Run("palavra deseconhecida", func(t *testing.T) {
		_, resultado := dicionario.Busca("desconhecida")

		comparaErro(t, resultado, errNaoEncontrado)
	})
}

func TestAdiciona(t *testing.T) {

	t.Run("palavra nova", func(t *testing.T) {
		dicionario := Dicionario{}
		palavra := "teste"
		definicao := "isso é apenas um teste"

		err := dicionario.Adiciona(palavra, definicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})

	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Adiciona(palavra, "teste novo")

		comparaErro(t, err, errPalavraExistente)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		novaDefinicao := "nova definição"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Atualiza(palavra, novaDefinicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, novaDefinicao)
	})

	t.Run("palavra nova", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{}

		err := dicionario.Atualiza(palavra, definicao)

		comparaErro(t, err, errPalavraInexistente)
	})
}

func TestDeleta(t *testing.T) {
	palavra := "teste"
	dicionario := Dicionario{palavra: "definição de teste"}

	dicionario.Deleta(palavra)

	_, err := dicionario.Busca(palavra)
	if err != errNaoEncontrado {
		t.Errorf("espera-se que '%s' seja deletado", palavra)
	}
}

func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()

	resultado, err := dicionario.Busca(palavra)

	if err != nil {
		t.Fatal("deveria ter encontrado a palavra adicionada: ", err)
	}

	if definicao != resultado {
		t.Errorf("resultado: '%s', esperado '%s'", resultado, definicao)
	}
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("Resultado erro '%s', Esperado '%s'", resultado, esperado)
	}
}

func comparaString(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("RESULTADO: '%s', ESPERADO: '%s', dado '%s'", resultado, esperado, "teste")
	}
}
