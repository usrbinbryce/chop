package urls

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	repo "github.com/usrbinbryce/chop/internal/adapters/postgresql/sqlc"
)

type Service interface {
	CreateURL(ctx context.Context, urlToCreate createURLPayload) (repo.Url, error)
	GetURLFromShortCode(ctx context.Context, shortCode string) (repo.Url, error)
}

type svc struct {
	repo *repo.Queries
	db   *pgxpool.Pool
}

func NewService(repo *repo.Queries, db *pgxpool.Pool) Service {
	return &svc{repo: repo, db: db}
}

// @TODO: reminder to check for existing shortened URLs before inserting into the DB
func (s *svc) CreateURL(ctx context.Context, urlToCreate createURLPayload) (repo.Url, error) {
	if err := validateURLPayload(urlToCreate); err != nil {
		return repo.Url{}, err
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

func (s *svc) GetURLFromShortCode(ctx context.Context, shortCode string) (repo.Url, error) {
	if shortCode == "" {
		return repo.Url{}, ErrShortCodeRequired
	}

	return s.repo.GetURLFromShortCode(ctx, shortCode)
}

func validateURLPayload(pl createURLPayload) error {
	if pl.ShortCode == "" {
		return ErrShortCodeRequired
	}

	if pl.Destination == "" {
		return ErrDestinationRequired
	}

	return nil
}
