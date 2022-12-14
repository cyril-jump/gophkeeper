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
	UpdateCardDataByID(ctx context.Context, userID string, data domain.CardData) error
}

func (r *Repo) UpdateCardDataByIDDB(ctx context.Context, userID string, data domain.CardData) error {

	return r.provider.UpdateCardDataByID(ctx, userID, data)
}
