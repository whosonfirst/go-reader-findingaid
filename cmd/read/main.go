// read is a command-line tool to resolve one or more URIs, using a Who's On First finding aid and read their corresponding Who's On First documents, outputting each to STDOUT.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-reader"
	_ "github.com/whosonfirst/go-reader-findingaid"
	"io"
	"log"
	"os"
)

func main() {

	reader_uri := flag.String("reader-uri", "", "A valid whosonfirst/go-reader-findingaid URI")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Resolve one or more URIs, using a Who's On First finding aid and read their corresponding Who's On First documents, outputting each to STDOUT.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s uri(N) uri(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	uris := flag.Args()

	ctx := context.Background()

	r, err := reader.NewReader(ctx, *reader_uri)

	if err != nil {
		log.Fatalf("Failed to create new reader, %v", err)
	}

	for _, path := range uris {

		fh, err := r.Read(ctx, path)

		if err != nil {
			log.Fatalf("Failed to read '%s', %v", path, err)
		}

		defer fh.Close()

		_, err = io.Copy(os.Stdout, fh)

		if err != nil {
			log.Fatalf("Failed to copy contents of '%s' to STDOUT, %v", path, err)
		}
	}

}
