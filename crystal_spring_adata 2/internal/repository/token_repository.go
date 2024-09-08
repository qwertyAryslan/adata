package repository

import (
	"context"
	"get_adata/internal/entity"
)

type TokenRepository interface {
	GetToken(ctx context.Context, iinBin string) (string, error)
	GetCompanyData(ctx context.Context, token string) (entity.Company, error)
}
