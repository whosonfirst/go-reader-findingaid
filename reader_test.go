package findingaid

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"testing"

	"github.com/whosonfirst/go-reader/v2"
)

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

	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Debug("Verbose logging enabled")

	ctx := context.Background()

	reader_q := url.Values{}
	reader_q.Add("resolver", "https://static.sfomuseum.org/findingaid?template=https://raw.githubusercontent.com/sfomuseum-data/{repo}/main/data/")
	reader_q.Add("resolver", "https://data.whosonfirst.org/findingaid?template=https://raw.githubusercontent.com/whosonfirst-data/{repo}/master/data/")

	reader_u := url.URL{}
	reader_u.Scheme = "findingaid"
	reader_u.Host = "multi"
	reader_u.RawQuery = reader_q.Encode()

	reader_uri := reader_u.String()
	slog.Debug("Create reader", "uri", reader_uri)

	r, err := reader.NewReader(ctx, reader_uri)

	if err != nil {
		t.Fatalf("Failed to create new reader, %v", err)
	}

	tests := []string{
		"85865975",
		"1159396133",
	}

	for _, uri := range tests {

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
}
