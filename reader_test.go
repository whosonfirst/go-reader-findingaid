package findingaid

import (
	"context"
	"fmt"
	"github.com/whosonfirst/go-reader"
	"os"
	"testing"
)

//

func TestSQLiteFindingAid(t *testing.T) {

	ctx := context.Background()

	cwd, err := os.Getwd()

	if err != nil {
		t.Fatalf("Failed to determine current working directory")
	}

	template := fmt.Sprintf("fs://%s/fixtures/{repo}/data", cwd)

	reader_uri := fmt.Sprintf("findingaid://sqlite?dsn=fixtures/sfomuseum-data-maps.db&template=%s", template)

	r, err := reader.NewReader(ctx, reader_uri)

	if err != nil {
		t.Fatalf("Failed to create new reader, %v", err)
	}

	uri := "1746160269"

	fh, err := r.Read(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to read %s, %v", uri, err)
	}

	fh.Close()
}

func TestHTTPFindingAid(t *testing.T) {

	ctx := context.Background()

	reader_uri := "findingaid://https/static.sfomuseum.org/findingaid?template=https://raw.githubusercontent.com/sfomuseum-data/{repo}/main/data/"

	r, err := reader.NewReader(ctx, reader_uri)

	if err != nil {
		t.Fatalf("Failed to create new reader, %v", err)
	}

	uri := "102527513"

	fh, err := r.Read(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to read %s, %v", uri, err)
	}

	fh.Close()
}
