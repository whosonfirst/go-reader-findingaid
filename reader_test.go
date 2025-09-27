package findingaid

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/whosonfirst/go-reader/v2"
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

	exists, err := r.Exists(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to determine if %s exists, %v", uri, err)
	}

	if !exists {
		t.Fatalf("Expected %s to exists", uri)
	}
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

	exists, err := r.Exists(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to determine if %s exists, %v", uri, err)
	}

	if !exists {
		t.Fatalf("Expected %s to exists", uri)
	}
}

func TestMultiFindingAid(t *testing.T) {

	ctx := context.Background()

	reader_uri := "findingaid://multi?resolver=https%3A%2F%2Fstatic.sfomuseum.org%2Ffindingaid%3Ftemplate%3Dhttps%3A%2F%2Fraw.githubusercontent.com%2Fsfomuseum-data%2F%7Brepo%7D%2Fmain%2Fdata%2F&resolver=https%3A%2F%2Fdata.whosonfirst.org%2Ffindingaid%3Ftemplate%3Dhttps%3A%2F%2Fraw.githubusercontent.com%2Fwhosonfirst-data%2F%7Brepo%7D%2Fmaster%2Fdata%2F"

	r, err := reader.NewReader(ctx, reader_uri)

	if err != nil {
		t.Fatalf("Failed to create new reader, %v", err)
	}

	uri := "85865975"

	fh, err := r.Read(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to read %s, %v", uri, err)
	}

	fh.Close()

	exists, err := r.Exists(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to determine if %s exists, %v", uri, err)
	}

	if !exists {
		t.Fatalf("Expected %s to exists", uri)
	}
}
