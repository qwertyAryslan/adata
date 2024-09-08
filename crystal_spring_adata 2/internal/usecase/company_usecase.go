package usecase

import (
	"context"
	"get_adata/internal/entity"
	"get_adata/internal/repository"
)

type CompanyUseCase interface {
	FetchCompanyData(ctx context.Context, iinBin string) (entity.Company, error)
}

type companyUseCase struct {
	tokenRepo repository.TokenRepository
}

func NewCompanyUseCase(tr repository.TokenRepository) CompanyUseCase {
	return &companyUseCase{
		tokenRepo: tr,
	}
}

func (uc *companyUseCase) FetchCompanyData(ctx context.Context, iinBin string) (entity.Company, error) {
	token, err := uc.tokenRepo.GetToken(ctx, iinBin)
	if err != nil {
		return entity.Company{}, err
	}

	companyData, err := uc.tokenRepo.GetCompanyData(ctx, token)
	if err != nil {
		return entity.Company{}, err
	}

	return companyData, nil
}
