package line

import "context"

type LineRepository interface {
	Create(ctx context.Context, line *Line) error
	Get(ctx context.Context, id string) (*Line, error)
	Update(ctx context.Context, line *Line) error
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context, filter *LineFilter) ([]*Line, error)
}
