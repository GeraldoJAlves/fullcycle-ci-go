package main

import "testing"

func TestSomaWithNegatives(t *testing.T) {
	total := soma(15, -15)

	if total != 0 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}
