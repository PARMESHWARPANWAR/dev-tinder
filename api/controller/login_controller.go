package controller

import (
	"net/http"

    "golang.org/x/crypto/bcrypt"

	"github.com/PARMESHWARPANWAR/dev-tinder/bootstrap"
	"github.com/PARMESHWARPANWAR/dev-tinder/domain"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context){
	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message:err.Error()})
		return 
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return 
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message:"Invalid credentials"})
		return 
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret,lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return 
	}

	// Set access token cookie
    c.SetCookie(
        "access_token",         // name
        accessToken,            // value
        lc.Env.AccessTokenExpiryHour * 3600,  // expiration time in seconds
        "/",                    // path
        lc.Env.ServerAddress,   // domain
        true,                   // secure
        true,                   // httpOnly
    )

    // Set refresh token cookie
    c.SetCookie(
        "refresh_token",        // name
        refreshToken,           // value
        lc.Env.RefreshTokenExpiryHour * 3600, // expiration time in seconds
        "/",                    // path
        lc.Env.ServerAddress,   // domain
        true,                   // secure
        true,                   // httpOnly
    )

    // Send success response
    c.JSON(http.StatusOK, domain.SuccessResponse{Message:"Login successful!"})
}