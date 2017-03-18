package main

import (
	"github.com/jonas747/dca"
	"os"
	"io"
)

func main() {
    encodeSession := dca.EncodeFile("12.mp3", dca.StdEncodeOptions)
    // Make sure everything is cleaned up, that for example the encoding process if any issues happened isnt lingering around
    defer encodeSession.Cleanup()

    output, err := os.Create("output.dca")
    if err != nil {
        // Handle the error
    }

    io.Copy(output, encodeSession)
}
