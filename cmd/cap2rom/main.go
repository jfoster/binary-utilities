package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	utils "github.com/jfoster/binary-utilities"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		log.Fatal(errors.New("Incorrect number of arguments specified"))
	}

	capfile, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer capfile.Close()

	log.Printf("opened file %s", capfile.Name())

	stat, err := capfile.Stat()
	if err != nil {
		log.Fatal(err)
	}

	ret, err := capfile.Seek(utils.PrevPowerOfTwo(stat.Size()), 2)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("seeking to offset %d", ret)

	romfile, err := os.Create(strings.TrimSuffix(capfile.Name(), filepath.Ext(capfile.Name())) + ".rom")
	if err != nil {
		log.Fatal(err)
	}
	defer romfile.Close()

	log.Printf("created file %s", romfile.Name())

	written, err := io.Copy(romfile, capfile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%d bytes copied to %s", written, romfile.Name())
}
