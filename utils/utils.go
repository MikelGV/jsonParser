package utils

import (
	"log"
	"os"
)

type OpenHandler struct {
    FileName string
}


func (oh OpenHandler) OpenFile() (*os.File, error) {
    file, err := os.Open(oh.FileName)

    if err != nil {
        log.Fatal(err)
    }

    return file, nil
}


func IsEmptyStdin() (bool, error) {
    fi, err := os.Stdin.Stat()

    if err != nil {
        return false, err
    }

    return fi.Mode()&os.ModeNamedPipe == 0, nil
}


