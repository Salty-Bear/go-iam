package resource

import (
	"context"

	"github.com/melvinodsa/go-iam/sdk"
)

type service struct {
	s Store
}

func NewService(s Store) Service {
	return service{s: s}
}

func (s service) Search(ctx context.Context, query sdk.ResourceQuery) (*sdk.ResourceList, error) {
	return s.s.Search(ctx, query)
}

func (s service) Get(ctx context.Context, id string) (*sdk.Resource, error) {
	if len(id) == 0 {
		return nil, ErrResourceNotFound
	}
	return s.s.Get(ctx, id)
}

func (s service) Create(ctx context.Context, resource *sdk.Resource) error {
	return s.s.Create(ctx, resource)
}

func (s service) Update(ctx context.Context, resource *sdk.Resource) error {
	return s.s.Update(ctx, resource)
}
