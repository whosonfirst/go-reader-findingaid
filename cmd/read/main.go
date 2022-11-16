// read is a command-line tool to resolve one or more URIs, using a Who's On First finding aid and read their corresponding Who's On First documents,
// outputting each to STDOUT. For example:
//
//	$> ./bin/read -reader-uri 'findingaid://awsdynamodb/{TABLE}?partition_key=id&region={REGION}&credentials={CREDENTIALS}' \
//		-data-template 'https://raw.githubusercontent.com/sfomuseum-data/{repo}/main/data/' \
//		1762946673 | jq '.properties["wof:name"]'
//	"flight officer wings: Big Sky Airlines"
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-reader"
	"github.com/whosonfirst/go-reader-findingaid"
	"io"
	"log"
	"net/url"
	"os"
)

func main() {

	reader_uri := flag.String("reader-uri", "", "A valid whosonfirst/go-reader-findingaid URI")

	data_template := flag.String("data-template", "", fmt.Sprintf("A valid URI template to use for resolving final reader URIs. Default is %s", findingaid.WHOSONFIRST_DATA_TEMPLATE))

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Resolve one or more URIs, using a Who's On First finding aid and read their corresponding Who's On First documents, outputting each to STDOUT.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s uri(N) uri(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	uris := flag.Args()

	ctx := context.Background()

	if *data_template != "" {

		u, err := url.Parse(*reader_uri)

		if err != nil {
			log.Fatalf("Failed to parse reader URI, %v", err)
		}

		q := u.Query()
		q.Del("template")
		q.Set("template", *data_template)

		u.RawQuery = q.Encode()
		*reader_uri = u.String()
	}

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
