package math

import (
	"fmt"
	"reflect"
)
type Student struct {
	name string
	username string
	age int
}
func Sum(a [5]int) (sum int) {
	fmt.Println(reflect.TypeOf(a))
	for i:=0; i < len(a); i++  {
		sum+= a[i]
	}
	return

}

func IntializeStudent(str string, str2 string, a int) (s Student)  {
s = Student{str, str2, a}
return
}
