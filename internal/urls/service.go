package urls

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	repo "github.com/usrbinbryce/chop/internal/adapters/postgresql/sqlc"
)

type Service interface {
	CreateURL(ctx context.Context, urlToCreate createUrlRequest) (repo.Url, error)
}

type svc struct {
	repo *repo.Queries
	db   *pgxpool.Pool
}

func NewService(repo *repo.Queries, db *pgxpool.Pool) Service {
	return &svc{repo: repo, db: db}
}

// @TODO: reminder to check for existing shortened URLs before inserting into the DB
func (s *svc) CreateURL(ctx context.Context, urlToCreate createUrlRequest) (repo.Url, error) {
	if urlToCreate.ShortCode == "" {
		return repo.Url{}, ErrShortCodeRequired
	}
	if urlToCreate.Destination == "" {
		return repo.Url{}, ErrDestinationRequired
	}

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Url{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.repo.WithTx(tx)

	url, err := qtx.CreateURL(ctx, repo.CreateURLParams{
		ShortCode:   urlToCreate.ShortCode,
		Destination: urlToCreate.Destination,
	})
	if err != nil {
		return repo.Url{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return repo.Url{}, err
	}

	return url, nil
}
