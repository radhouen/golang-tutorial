package main

import "testing"


func TestHello(t *testing.T)  {
	getFunctionResult := Hello()
	waitingRestult :="Hello World !"
	/* nb+
	rsum := Sum(1,2)
	rsum1  := Sum(1,9)
	rsum2  := Sum(79,9)
	if rsum == 3 && rsum1 == 10 && rsum2 == 88 {

	} else {

	}*/
	if getFunctionResult != waitingRestult {
		t.Errorf("the value of waiting result %v is different than %v", waitingRestult, getFunctionResult)
	}
 }
