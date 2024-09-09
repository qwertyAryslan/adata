package handler

import (
	"net/http"

	"adata_crystal_spring/internal/usecase"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	tokenUsecase *usecase.TokenUsecase
}

func NewTokenHandler(usecase *usecase.TokenUsecase) *TokenHandler {
	return &TokenHandler{
		tokenUsecase: usecase,
	}
}

// GetBasicData godoc
// @Summary Get basic data
// @Description Get basic data by iinBin
// @Tags basic
// @Accept  json
// @Produce  json
// @Param iinBin path string true "IIN or BIN"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /basic/{iinBin} [get]
func (h *TokenHandler) GetBasicData(c *gin.Context) {
	iinBin := c.Param("iinBin")
	data, err := h.tokenUsecase.GetBasicData(iinBin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
