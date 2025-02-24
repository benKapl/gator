package main

import (
	"fmt"

	"github.com/benKapl/gator/internal/config"
	// "os"
)

func main() {
	fmt.Println(config.GetConfigFilePath())

}
