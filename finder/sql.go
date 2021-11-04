package finder

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	wof_reader "github.com/whosonfirst/go-reader"
	_ "github.com/whosonfirst/go-reader-http"
	"net/url"
)

// WHOSONFIRST_DATA_TEMPLATE is a URL template for the root `data` directory in Who's On First data repositories.
const WHOSONFIRST_DATA_TEMPLATE string = "https://raw.githubusercontent.com/whosonfirst-data/{repo}/master/data/"

// type SQLiteFinder implements the `whosonfirst/go-reader` interface for use with Who's On First finding aids.
type SQLiteFinder struct {
	Finder
	// A SQLite `sql.DB` instance containing Who's On First finding aid data.
	db *sql.DB
}

func init() {
	ctx := context.Background()
	wof_reader.RegisterReader(ctx, "sqlite", NewSQLiteFinder)
	wof_reader.RegisterReader(ctx, "sqlite3", NewSQLiteFinder)
}

// NewSQLiteFinder will return a new `whosonfirst/go-reader.` instance for reading Who's On First
// documents by first resolving a URL using a Who's On First finding aid.
func NewSQLiteFinder(ctx context.Context, uri string) (Finder, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URL, %w", err)
	}

	q := u.Query()

	dsn := q.Get("dsn")

	db, err := sql.Open("sqlite3", dsn)

	if err != nil {
		return nil, fmt.Errorf("Failed to open database, %w", err)
	}

	f := &SQLiteFinder{
		db: db,
	}

	return f, nil
}

// GetRepo returns the name of the repository associated with this ID in a Who's On First finding
// aid.
func (r *SQLiteFinder) GetRepo(ctx context.Context, id int64) (string, error) {

	var repo string

	q := "SELECT s.name FROM catalog c, sources s WHERE c.id = ? AND c.repo_id = s.id"

	row := r.db.QueryRowContext(ctx, q, id)
	err := row.Scan(&repo)

	if err != nil {
		return "", fmt.Errorf("Failed to scan row, %w", err)
	}

	return repo, nil
}
