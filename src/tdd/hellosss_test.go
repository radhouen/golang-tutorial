package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestAdd(t *testing.T)  {
	got:= add(2,2)
	want:= 4
	if got != want {
		t.Errorf("Please verify  your function add", want)

	}

}


