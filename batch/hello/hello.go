package hello

import (
	"fmt"
)

func Run(s *string) {
	fmt.Printf("hello %v\n", *s)	
}