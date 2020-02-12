package main

import "testing"

func TestInitializeAddress(t *testing.T)  {

	got := InitializeAddress("Tabarka", 8112, "Jendouba", "Tunisia")
	want := Address{"Tabarka", 8112, "Jendouba", "Tunisia"}

	if got != want {
		t.Errorf("the want value %v is equal to the got value %v = %v", want, got, want == got)
	}



}


