package findingaid

import (
	"context"
)

type FindingAid interface {
	GetRepo(context.Context, int64) (string, error)
}
