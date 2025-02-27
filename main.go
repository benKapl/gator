package main

import (
	"fmt"

	"github.com/benKapl/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	cfg.SetUser("Ben")

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", cfg)

}
