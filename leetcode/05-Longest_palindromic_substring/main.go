package main

import (
	"fmt"
	lon "longest/Manacher"
)

func main() {
	strin := "wekababcucbabawke"
	fmt.Println(lon.Manacher(strin))
}
