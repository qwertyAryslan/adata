package usecase

import (
	"adata_crystal_spring/internal/domain"
	"adata_crystal_spring/pkg/http"
	"encoding/json"
	"fmt"
)

type TokenUsecase struct {
	httpClient *http.HttpClient
}

func NewTokenUsecase(client *http.HttpClient) *TokenUsecase {
	return &TokenUsecase{
		httpClient: client,
	}
}

func (u *TokenUsecase) GetBasicData(iinBin string) (string, error) {
	token, err := u.getToken(iinBin)
	if err != nil {
		return "", err
	}

	basicData, err := u.getBasicData(token)
	if err != nil {
		return "", err
	}

	return basicData, nil
}

func (u *TokenUsecase) getToken(iinBin string) (string, error) {
	url := fmt.Sprintf("http://10.1.22.179:8070/company/token?iinBin=%s", iinBin)
	body, err := u.httpClient.Get(url)
	if err != nil {
		return "", err
	}

	var tokenResponse domain.TokenResponse
	if err := json.Unmarshal([]byte(body), &tokenResponse); err != nil {
		return "", err
	}

	return tokenResponse.Token, nil
}

func (u *TokenUsecase) getBasicData(token string) (string, error) {
	url := fmt.Sprintf("http://10.1.22.179:8070/basic?token=%s", token)
	body, err := u.httpClient.Get(url)
	if err != nil {
		return "", err
	}

	return body, nil
}
