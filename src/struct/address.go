package main


/*
int (int32 , int64 ....
float ( float32, float64
string
rune
array
slice
function
map
...
 */

/*
street
Postal Code
town
country
 */
//===> Address
type Address struct {
	street string
	PostalCode int
	town string
	country string
}

func InitializeAddress(street string, pCode int, town string, country string ) (addr Address) {

	addr = Address{street, pCode, town ,country}
	return
}
