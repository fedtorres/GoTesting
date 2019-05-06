package mio

import (
	"errors"
)

func Acumular(lista []int) error {
	if lista == nil {
		return errors.New("La lista está vacía")
	}
	for i := 1; i < (len(lista)); i++ {
		lista[i] = lista[i-1] + lista[i]
	}
	return nil
}
