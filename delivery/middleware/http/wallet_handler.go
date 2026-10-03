package http

import (
	"expense-manager-golang/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	WalletUsecase domain.WalletUsecase
}

func NewWalletHandler(wu domain.WalletUsecase) *WalletHandler {
	return &WalletHandler{WalletUsecase: wu}
}

// struct for validation input when create a new wallet
// initial balance can be set to 0, but cant be negative
type createWalletRequest struct {
	Name           string  `json:"name" binding:"required"`
	InitialBalance float64 `json:"initial_balance" binding:"gte=0"`
}

func (h *WalletHandler) Create(c *gin.Context) {
	var req createWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//extract user id from middleware
	userIDVal, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := userIDVal.(uint)

	//call usecase
	err := h.WalletUsecase.CreateWallet(c.Request.Context(), userID, req.Name, req.InitialBalance)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "wallet created!"})
}

func (h *WalletHandler) GetMyWallets(c *gin.Context) {
	//extract user id from middleware
	userIDVal, exist := c.Get("user_id")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := userIDVal.(uint)

	//call usecase
	wallets, err := h.WalletUsecase.GetMyWallets(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cant retrieve wallet data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": wallets})
}
