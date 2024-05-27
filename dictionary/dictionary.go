package main

const (
	errNaoEncontrado      = ErrDicionario("não foi possível encontrar a palavra que você procura")
	errPalavraExistente   = ErrDicionario("não foi possível adicionar a palavra pois ela já existe")
	errPalavraInexistente = ErrDicionario("não foi possível atualizar a palavra pois ela não existe")
)

type ErrDicionario string

type Dicionario map[string]string

func (e ErrDicionario) Error() string {
	return string(e)
}

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", errNaoEncontrado
	}

	return definicao, nil
}

func (d Dicionario) Adiciona(palavra, definicao string) error {
	_, err := d.Busca(palavra)
	switch err {
	case errNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return errPalavraExistente
	default:
		return err

	}

	return nil
}

func (d Dicionario) Atualiza(palavra, definicao string) error {
	_, err := d.Busca(palavra)
	switch err {
	case errNaoEncontrado:
		return errPalavraExistente
	case nil:
		d[palavra] = definicao
	default:
		return err
	}

	return nil
}

func (d Dicionario) Deleta(palavra string) {
	delete(d, palavra)
}
