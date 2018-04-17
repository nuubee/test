package main

import (
	"fmt"
)

type error interface {
	Error() string
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, error.New("math:error")
	}
}

func main() {

}
