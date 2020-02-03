package math

import (
	"fmt"
	"reflect"
	"testing"
)



func TestSum(t *testing.T) {

	numbers := [5]int{1, 2, 2, 4, 5}

	got := Sum(numbers)
	want := 14

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestIntializeStruct(t *testing.T) {
	want := Student{"radhouen", "assakra", 27}
	got  := IntializeStudent("radhouen","assakra",27)
	fmt.Println(reflect.DeepEqual(want, got))
	if want != got {
		t.Errorf("intialize new struct is failed because %v is different to %v", got, want)
	}
}

func TestDiv(t *testing.T) {
	got := Div(5, 5)
	want := 1
	if want != got {
		t.Errorf("got %d want %d given", got, want)
	}
}
