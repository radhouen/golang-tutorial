package main

import (
	"fmt"
	"lab1/src/ptest"
)

/*
Importing a package
Now, we need an executable package which will consume our greet package.
Let’s create an app directory inside src and create entry.go file with main
package declaration and main function. Note here, Go packages do not have an
entry file naming system like index.js in Node. For an executable package,
a file with main function is the entry file for execution.
To import a package, we use import syntax followed by the package name.
💡 Unlike other programming languages, a package name can also be a subpath
like some-dir/greet and Go will automatically resolve the path to the greet package
for us as you will see in the nested package topic ahead.
Go first searches for package directory inside GOROOT/src directory and if it doesn’t
find the package, then it looks for GOPATH/src. Since, fmt package is part of Go’s standard
library which is located in GOROOT/src, it is imported from there. But since Go cannot find
greet package inside GOROOT, it will lookup inside GOPATH/src and we have it there.
 */
func main()  {
	fmt.Println(ptest.Add(5,7))
	fmt.Println(ptest.Min(5,8))
}
