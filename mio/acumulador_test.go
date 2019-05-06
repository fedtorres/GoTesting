package mio

import (
	"fmt"
	"testing"
)

func TestAcumulado(t *testing.T) {
	lista := []int {6, 3, 1, 5, 7, 0, 2, 9, 8, 4}
	err := Acumular(lista)
	if err != nil {
		t.Error("No se esperaba un error.")
	}
	if lista[9] != 45 {
		t.Error("No se acumuló según lo esperado")
	}
}

func TestAcumularListaNula(t *testing.T) {
	err := Acumular(nil)
	if err == nil {
		t.Error("No se esperaba un error.")
	}
	errorEsperado := "La lista está vacía"
	if err.Error() != errorEsperado {
		t.Error(fmt.Sprintf("Mensaje de error esperado %s, mensaje de error recibido: %s", errorEsperado, err))
	}
}

func BenchmarkAcumular10Elementos(b *testing.B) {
	lista := generarLista(10)
	for n := 0; n < b.N; n++ {
		Acumular(lista)
	}
}

func BenchmarkAcumular100Elementos(b *testing.B) {
	lista := generarLista(100)
	for n := 0; n < b.N; n++ {
		Acumular(lista)
	}
}

func BenchmarkAcumular1000Elementos(b *testing.B) {
	lista := generarLista(1000)
	for n := 0; n < b.N; n++ {
		Acumular(lista)
	}
}

func BenchmarkAcumular10000Elementos(b *testing.B) {
	lista := generarLista(10000)
	for n := 0; n < b.N; n++ {
		Acumular(lista)
	}
}

func generarLista(n int) []int {
	var lista = make([]int, n)
	for i := 0; i < n; i++ {
		lista[i] = i
	}
	return lista
}

/*user = &User {
	ID: 1234567,
}
user.Get()
//Create WaitGroup
//Create Channel
go categories.Get()
go currencies.Get()
//Read channel
return &MyMl{}*/