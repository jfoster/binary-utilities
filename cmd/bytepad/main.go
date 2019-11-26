package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/jfoster/bintools"
)

func main() {
	size := flag.Int64("size", 0, "Size in bytes of the new padded file")
	data := flag.Int64("data", 0x00, "A single byte of data used for padding")

	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		log.Fatal(errors.New("Incorrect number of arguments specified"))
	}

	file, err := os.OpenFile(args[0], os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Printf("opened file %s", file.Name())

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	diff := *size - stat.Size()

	if diff < 0 {
		log.Fatal(errors.New("diff cannot be negative"))
	}

	pad := bintools.Bytes(byte(*data), diff)

	n, err := file.Write(pad)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("%s padded with %d bytes of 0x%.2X", file.Name(), n, *data)
}
