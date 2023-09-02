package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ronaudinho/gpt/gp"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("no file to read")
	}
	b, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}
	song := &gp.Song{}
	song.ReadGP5(b)
	b, _ = json.MarshalIndent(song, "", "  ")
	fmt.Println(string(b))
}
