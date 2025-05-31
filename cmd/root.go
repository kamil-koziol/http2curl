package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/kamil-koziol/http2curl/pkg/http2curl"
)

func Run() int {
	output := flag.String("o", "-", "Output file, use '-' for stdout")
	input := flag.String("i", "-", "Input file, use '-' for stdin")
	args := flag.String("args", "", "Args that will be passed to CURL")

	flag.Parse()

	var inputReader io.Reader
	if *input == "-" {
		inputReader = os.Stdin
	} else {
		file, err := os.Open(*input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
			return 1
		}
		defer file.Close()
		inputReader = file
	}

	var outWriter io.Writer
	if *output == "-" {
		outWriter = os.Stdout
	} else {
		file, err := os.Create(*output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create output file: %v\n", err)
			return 1
		}
		defer file.Close()
		outWriter = file
	}

	curl, err := http2curl.Convert(inputReader, *args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to convert .http to curl: %s\n", err)
		return 1
	}

	fmt.Fprintln(outWriter, curl)
	return 0
}
