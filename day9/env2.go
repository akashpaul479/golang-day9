package day9

import (
	"fmt"
	"os"
)

func Env2() {
	os.Clearenv()
	fmt.Println("All environment variables cleared")
}
