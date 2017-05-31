package main

import (
	"github.com/goadesign/goa"
	"github.com/rymccue/fifa-tracker/app"
	"github.com/rymccue/fifa-tracker/repositories"
	"github.com/rymccue/fifa-tracker/utils"
)

// AuthenticationController implements the authentication resource.
type AuthenticationController struct {
	*goa.Controller
}

// NewAuthenticationController creates a authentication controller.
func NewAuthenticationController(service *goa.Service) *AuthenticationController {
	return &AuthenticationController{Controller: service.NewController("AuthenticationController")}
}

// Login runs the login action.
func (c *AuthenticationController) Login(ctx *app.LoginAuthenticationContext) error {
	data := ctx.Payload
	u, err := repositories.GetUserByEmail(data.Email)
	if err != nil {
		return ctx.BadRequest(err)
	}
	if !utils.VerifyPassword(data.Password, u) {
		return ctx.BadRequest(goa.ErrBadRequest("Invalid password"))
	}
	signedToken, err := utils.GenerateJWTToken(u.ID)
	if err != nil {
		return ctx.InternalServerError()
	}

	res := &app.FtToken{
		Token: signedToken,
	}
	return ctx.OK(res)
}

// Register runs the register action.
func (c *AuthenticationController) Register(ctx *app.RegisterAuthenticationContext) error {
	data := ctx.Payload
	u, err := repositories.CreateUser(data.Email, data.Password, data.FirstName, data.LastName)
	if err != nil {
		return ctx.InternalServerError()
	}
	signedToken, err := utils.GenerateJWTToken(u.ID)
	if err != nil {
		return ctx.InternalServerError()
	}

	res := &app.FtToken{
		Token: signedToken,
	}
	return ctx.OK(res)
}
