package lihaorong

import (
	"fmt"
)

func GetApp() {

	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
}
