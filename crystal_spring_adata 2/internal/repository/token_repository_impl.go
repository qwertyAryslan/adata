package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"get_adata/internal/entity"
	"net/http"
)

type tokenRepository struct {
	baseURL string
}

func NewTokenRepository(baseURL string) TokenRepository {
	return &tokenRepository{baseURL: baseURL}
}

func (r *tokenRepository) GetToken(ctx context.Context, iinBin string) (string, error) {
	url := fmt.Sprintf("%s/company/token?iinBin=%s", r.baseURL, iinBin)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token: %s", resp.Status)
	}

	var tokenResp entity.Token
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", err
	}

	return tokenResp.Value, nil
}

func (r *tokenRepository) GetCompanyData(ctx context.Context, token string) (entity.Company, error) {
	url := fmt.Sprintf("%s/basic?token=%s", r.baseURL, token)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return entity.Company{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return entity.Company{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entity.Company{}, fmt.Errorf("failed to get company data: %s", resp.Status)
	}

	var company entity.Company
	if err := json.NewDecoder(resp.Body).Decode(&company); err != nil {
		return entity.Company{}, err
	}

	return company, nil
}
