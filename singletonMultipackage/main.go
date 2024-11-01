// main.go
package main

import (
	"fmt"

	"singletonMultipackage/package1"

	"singletonMultipackage/package2"

	"singletonMultipackage/singleton"
)

func main() {
	// Access singleton from the main package
	instance := singleton.GetInstance()
	fmt.Println("Main package using singleton with name:", instance.Name)

	// Use singleton in package1 and package2
	package1.UseSingleton()
	package2.UseSingleton()
}
