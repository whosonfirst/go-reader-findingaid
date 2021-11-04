package finder

import (
	"context"
	"fmt"
	wof_reader "github.com/whosonfirst/go-reader"
	"gocloud.dev/docstore"
	"net/url"
)

// type DocstoreFinder implements the `whosonfirst/go-reader` interface for use with Who's On First finding aids.
type DocstoreFinder struct {
	Finder
	// A Docstore `sql.DB` instance containing Who's On First finding aid data.
	collection *docstore.Collection
}

func init() {

	ctx := context.Background()

	wof_reader.RegisterReader(ctx, "awsdynamodb", NewDocstoreFinder)

	for _, scheme := range docstore.DefaultURLMux().CollectionSchemes() {

		err := wof_reader.RegisterReader(ctx, scheme, NewDocstoreFinder)

		if err != nil {
			panic(err)
		}
	}
}

// NewDocstoreFinder will return a new `whosonfirst/go-reader.` instance for reading Who's On First
// documents by first resolving a URL using a Who's On First finding aid.
func NewDocstoreFinder(ctx context.Context, uri string) (Finder, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URL, %w", err)
	}

	collection, err := docstore.OpenCollection(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to open collection, %w", err)
	}

	f := &DocstoreFinder{
		collection: collection,
	}

	return f, nil
}

// getRepo returns the name of the repository associated with this ID in a Who's On First finding
// aid.
func (r *DocstoreFinder) GetRepo(ctx context.Context, id int64) (string, error) {

	var repo string

	return repo, nil
}
