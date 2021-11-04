package findingaid

import (
	"context"
	"fmt"
	"github.com/jtacoma/uritemplates"
	"gocloud.dev/docstore"
	wof_reader "github.com/whosonfirst/go-reader"
	_ "github.com/whosonfirst/go-reader-http"
	wof_uri "github.com/whosonfirst/go-whosonfirst-uri"
	"io"
	"net/url"
)

// type DocstoreFindingAid implements the `whosonfirst/go-reader` interface for use with Who's On First finding aids.
type DocstoreFindingAid struct {
	FindingAid
	// A Docstore `sql.DB` instance containing Who's On First finding aid data.
	collection *doctore.Collection
}

func init() {
	ctx := context.Background()
	// wof_reader.Register(ctx, "findingaid", NewDocstoreFindingAid)
}

// NewDocstoreFindingAid will return a new `whosonfirst/go-reader.` instance for reading Who's On First
// documents by first resolving a URL using a Who's On First finding aid.
func NewDocstoreFindingAid(ctx context.Context, uri string) (wof_reader., error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URL, %w", err)
	}

	collection, err := docstore.OpenCollection(ctx, uri)
	

	if err != nil {
		return nil, fmt.Errorf("Failed to open collection, %w", err)
	}

	f := &DocstoreFindingAid{
		collection: collection,
	}

	return f, nil
}

// getRepo returns the name of the repository associated with this ID in a Who's On First finding
// aid.
func (r *DocstoreFindingAid) GetRepo(ctx context.Context, id int64) (string, error) {

	var repo string


	return repo, nil
}
