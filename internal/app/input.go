package app

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

func GetInput() byte {
	// opens the terminal
	f, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	// returns an error if not logged correctly
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// creates a file descriptor to read from
	fd := int(f.Fd())

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		log.Fatal(err)
	}
	defer term.Restore(fd, oldState)

	// max 3 byte input for special characters
	buf := make([]byte, 3)
	n, err := f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	// if the buffer length is 3, it is an arrow char
	if n == 3 {
		fmt.Println("Arrow key:", buf[2])
	} else {
		fmt.Println("Key:", buf[0])
	}

	return buf[2]
}
