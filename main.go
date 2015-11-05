package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
)

func main() {
	n := &Normalizer{
		Reader: os.Stdin,
		Writer: os.Stdout,
	}

	flag.IntVar(&n.Indentation, "i", 4, "number of spaces for indentation")
	flag.Parse()

	if err := n.Normalize(); err != nil {
		fmt.Fprintln(os.Stderr, color.RedString("error: %s", err.Error()))
	}
}

type Normalizer struct {
	Reader      io.Reader
	Writer      io.Writer
	Indentation int
}

func (n *Normalizer) Normalize() error {
	inBytes, err := ioutil.ReadAll(n.Reader)
	if err != nil {
		return err
	}

	var inData interface{}
	if err := json.Unmarshal(inBytes, &inData); err != nil {
		return err
	}

	indent := make([]byte, n.Indentation)
	for i := 0; i < n.Indentation; i++ {
		indent[i] = ' '
	}

	outBytes, err := json.MarshalIndent(inData, "", string(indent))
	if err != nil {
		return err
	}

	if _, err = n.Writer.Write(outBytes); err != nil {
		return err
	}

	return nil
}
