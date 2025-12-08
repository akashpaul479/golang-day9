package day9

import (
	"fmt"
	"os"
)

func Env() {
	os.Setenv("GEEKS", "geeks")
	fmt.Println("GEEKS", os.Getenv("GEEKS"))

	os.Unsetenv("GEEKS")
	value, ok := os.LookupEnv("GEEKS")
	fmt.Println("GEEKS", value, "Is present", ok)
}
