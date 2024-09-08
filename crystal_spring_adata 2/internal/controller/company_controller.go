package controller

import (
	"context"
	"encoding/json"
	"get_adata/internal/usecase"
	"net/http"
)

type CompanyController struct {
	companyUseCase usecase.CompanyUseCase
}

func NewCompanyController(cu usecase.CompanyUseCase) *CompanyController {
	return &CompanyController{companyUseCase: cu}
}

func (cc *CompanyController) GetCompanyData(w http.ResponseWriter, r *http.Request) {
	iinBin := r.URL.Query().Get("iinBin")
	if iinBin == "" {
		http.Error(w, "iinBin is required", http.StatusBadRequest)
		return
	}

	companyData, err := cc.companyUseCase.FetchCompanyData(context.Background(), iinBin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companyData)
}
