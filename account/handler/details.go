package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msjai/auth-full-stack/account/model"
	"github.com/msjai/auth-full-stack/account/model/apperrors"
)

// omitempty must be listed first (tags evaluated sequentially, it seems)
type detailsReq struct {
	Name    string `json:"name" binding:"omitempty,max=50"`
	Email   string `json:"email" binding:"required,email"`
	Website string `json:"website" binding:"omitempty,url"`

	FirstName   string `json:"firstname" binding:"omitempty,max=50"`
	LastName    string `json:"lastname" binding:"omitempty,max=50"`
	Telegram    string `json:"telegram" binding:"omitempty,max=45"`
	PhoneNumber string `json:"phonenumber" binding:"omitempty,max=50"`
	Country     string `json:"country" binding:"omitempty,max=50"`
	City        string `json:"city" binding:"omitempty,max=50"`
	Location    string `json:"location" binding:"omitempty,max=300"`
	Lang        string `json:"lang" binding:"omitempty,max=10"`
}

// Details handler
func (h *Handler) Details(c *gin.Context) {
	authUser := c.MustGet("user").(*model.User)

	var req detailsReq

	if ok := bindData(c, &req); !ok {
		return
	}

	// Should be returned with current imageURL
	u := &model.User{
		UID:         authUser.UID,
		Name:        req.Name,
		Email:       req.Email,
		Website:     req.Website,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Telegram:    req.Telegram,
		PhoneNumber: req.PhoneNumber,
		Country:     req.Country,
		City:        req.City,
		Location:    req.Location,
		Lang:        req.Lang,
	}

	ctx := c.Request.Context()
	err := h.UserService.UpdateDetails(ctx, u)

	if err != nil {
		log.Printf("Failed to update user: %v\n", err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
