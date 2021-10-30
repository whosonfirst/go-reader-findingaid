# go-reader-findingaid

Go package implementing the whosonfirst/go-reader interface for use with Who's On First "finding aids".

## Important

This package targets "version 2" of the `whosonfirst/go-whosonfirst-findingaid` package which is not finalized yet:

https://github.com/whosonfirst/go-whosonfirst-findingaid/tree/v2

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/whosonfirst/go-reader-findingaid.svg)](https://pkg.go.dev/github.com/whosonfirst/go-reader-findingaid)

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

	reader_uri := flag.String("reader-uri", "", "A valid whosonfirst/go-reader-findingaid URI.")

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

For a complete working example see [cmd/read](cmd/read/main.go).

## Finding aids

TBW

## Finding aid URIs

Finding aid URIs take the form of:

```
findingaid://?{QUERY_PARAMETERS}
```

Valid finding aid query parameters are:

| Name | Type | Notes | Required
| --- | --- | --- | --- |
| dsn | string | A valid `matt/go-sqlite3` DSN string | yes |
| template | string | A valid `jtacoma/uritemplates` for resolving final reader URIs. If empty the default URI template mapping to `whosonfirst-data/whosonfirst-data` repositories will be used. | no |

For example:

```
findingaid://?dsn=/usr/local/data/findingaids/wof.db
```

## Tools

```
$> make cli
go build -mod vendor -o bin/read cmd/read/main.go
```

### read

Resolve one or more URIs, using a Who's On First finding aid and read their corresponding Who's On First documents, outputting each to `STDOUT`.

```
> ./bin/read -h
Resolve one or more URIs, using a Who's On First finding aid and read their corresponding Who's On First documents, outputting each to STDOUT.
Usage:
	 ./bin/read uri(N) uri(N)
  -reader-uri string
    	A valid whosonfirst/go-reader-findingaid URI
```

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