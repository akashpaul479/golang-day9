package day9

import (
	"fmt"
	"os"
)

func Env1() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
