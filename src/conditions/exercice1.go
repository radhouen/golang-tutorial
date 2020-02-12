package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(displayDistribution("windows"))
	fmt.Println(displayDistribution("Linux"))
	fmt.Println(displayDistribution("unix"))
	fmt.Println(displayDistribution("mac-os"))
}

func displayDistribution(dist string) (distro string) {
	// write code here
	if strings.Compare(dist, "windows") == 0 {
		distro = " It is a windows system"
	}
	if strings.Compare(dist, "Linux") == 0 {
		distro = " It is a Linux system"
	}
	if strings.Compare(dist, "unix") == 0 {
		distro = " It is a unix system"
	}
	if strings.Compare(dist, "mac-os") == 0 {
		distro = " It is a mac-os system"
	}
	return
}
