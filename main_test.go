package main

import "testing"

type teste struct {
	data   []int
	answer int
}

func TestSomador(t *testing.T) {
	teste := somador(1, 2)
	resultado := 3

	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

func TestSomadorComTabela(t *testing.T) {
	testes := []teste{
		{data: []int{1, 2, 3}, answer: 6},
		{data: []int{4, 5, 6}, answer: 15},
	}

	for _, v := range testes {
		resultado := somatorio(v.data...)

		if v.answer != resultado {
			t.Error("Expected:", v.answer, "Got:", resultado)
		}
	}
}

func TestSomadorComTabelaStructAnonimo(t *testing.T) {
	var testes = []struct {
		data   []int
		answer int
	}{
		{data: []int{1, 2, 3}, answer: 6},
		{data: []int{4, 5, 6}, answer: 15},
	}

	for _, v := range testes {
		resultado := somatorio(v.data...)

		if v.answer != resultado {
			t.Error("Expected:", v.answer, "Got:", resultado)
		}
	}
}
