package main

import (
	"log"

	"github.com/MikelGV/jsonParser/cli"
	"github.com/MikelGV/jsonParser/utils"
)

func main() {
    ci := cli.New()

    if ci.FileName == "" {
        isEmpty, err := utils.IsEmptyStdin()
        
        if err != nil {
            log.Fatal(err)
        }

        if isEmpty {
            log.Fatal("There is no file or named pipe.")
        }

        // Here should go all the logic for the parser
    }
}
