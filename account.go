package main

import (
	"github.com/goadesign/goa"
	"github.com/rymccue/fifa-tracker/app"
	"github.com/rymccue/fifa-tracker/repositories"
)

// AccountController implements the account resource.
type AccountController struct {
	*goa.Controller
}

// NewAccountController creates a account controller.
func NewAccountController(service *goa.Service) *AccountController {
	return &AccountController{Controller: service.NewController("AccountController")}
}

// Delete runs the delete action.
func (c *AccountController) Delete(ctx *app.DeleteAccountContext) error {
	u, err := repositories.GetUserByID(ctx.AccountID)
	if err != nil {
		return ctx.NotFound()
	}
	err = repositories.DeleteUser(u)
	if err != nil {
		return ctx.InternalServerError()
	}
	return nil
}

// List runs the list action.
func (c *AccountController) List(ctx *app.ListAccountContext) error {
	users, err := repositories.GetAllUsers(1, 100)
	if err != nil {
		return ctx.InternalServerError()
	}
	resp := app.FtAccountCollection{}
	for _, u := range users {
		resp = append(resp, &app.FtAccount{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
		})
	}
	return ctx.OK(resp)
}

// Show runs the show action.
func (c *AccountController) Show(ctx *app.ShowAccountContext) error {
	u, err := repositories.GetUserByID(ctx.AccountID)
	if err != nil {
		return ctx.NotFound()
	}
	res := &app.FtAccount{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *AccountController) Update(ctx *app.UpdateAccountContext) error {
	u, err := repositories.GetUserByID(ctx.AccountID)
	if err != nil {
		return ctx.NotFound()
	}
	data := ctx.Payload
	if data.FirstName != nil {
		u.FirstName = *data.FirstName
	}
	if data.LastName != nil {
		u.LastName = *data.LastName
	}
	if data.Email != nil {
		u.Email = *data.Email
	}
	err = repositories.SaveUser(u)
	if err != nil {
		return ctx.BadRequest(err)
	}
	return nil
}
