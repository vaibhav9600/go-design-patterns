package package2

import (
	"fmt"

	"singletonMultipackage/singleton"
)

func UseSingleton() {
	// Access the singleton instance
	instance := singleton.GetInstance()
	fmt.Println("Package2 using singleton with name:", instance.Name)
}
