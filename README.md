# go-reader-findingaid

Go package implementing the whosonfirst/go-reader interface for use with Who's On First "finding aids".

## Important

This is work in progress and targets "version 2" of the `whosonfirst/go-whosonfirst-findingaid` package which is also not finalized yet:

https://github.com/whosonfirst/go-whosonfirst-findingaid/tree/v2

## Documentation

Documentation is incomplete at this time.

## Example

_Error handling omitted for brevity._

```
package main

import (
	"context"
	"flag"
	"github.com/whosonfirst/go-reader"
	_ "github.com/whosonfirst/go-reader-findingaid"
	"io"
	"os"
)

func main() {

	reader_uri := flag.String("reader-uri", "", "A valid whosonfirst/go-reader URI.")

	flag.Parse()

	ctx := context.Background()

	r, _ := reader.NewReader(ctx, *reader_uri)

	for _, path := range flag.Args() {

		fh, _ := r.Read(ctx, path)
		defer fh.Close()

		io.Copy(os.Stdout, fh)
	}
}

```

## Tools

### read

For example:

```
$> ./bin/read \
	-reader-uri 'findingaid://?dsn=/usr/local/data/findingaids/wof.db' \
	102527513 \
	
| jq '.["properties"]["wof:name"]'

"San Francisco International Airport"
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-reader
* https://github.com/whosonfirst/go-whosonfirst-reader-http
* https://github.com/whosonfirst/go-whosonfirst-findingaid/tree/v2