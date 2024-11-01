// package1/package1.go
package package1

import (
	"fmt"

	"singletonMultipackage/singleton"
)

func UseSingleton() {
	// Access the singleton instance
	instance := singleton.GetInstance()
	fmt.Println("Package1 using singleton with name:", instance.Name)
}
