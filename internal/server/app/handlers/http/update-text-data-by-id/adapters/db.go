package adapters

import (
	"context"

	"github.com/cyril-jump/gophkeeper/internal/server/app/domain"
)

type Repo struct {
	provider Provider
}

func New(p Provider) *Repo {
	return &Repo{
		provider: p,
	}
}

type Provider interface {
	UpdateTextDataByID(ctx context.Context, userID int, data domain.TextData) error
}

func (r *Repo) UpdateTextDataByIDDB(ctx context.Context, userID int, data domain.TextData) error {

	return r.provider.UpdateTextDataByID(ctx, userID, data)
}
