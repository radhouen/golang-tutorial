package ptest

/*
Export members
A utility package is supposed to provide some variables to a package who imports it. Like export syntax in JavaScript,
Go exports a variable if a variable name starts with Uppercase. All other variables not starting with an uppercase letter
is private to the package.I am going to use variable word from now on in this article, to describe an export member but
export members can be of any type like constant, map, function, struct, array, slice etc
*/

func Add(a int , b int) int  {
	return  a+b
}
