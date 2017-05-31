package main

import (
	"github.com/goadesign/goa"
	"github.com/rymccue/fifa-tracker/app"
	"github.com/rymccue/fifa-tracker/repositories"
)

// MatchController implements the match resource.
type MatchController struct {
	*goa.Controller
}

// NewMatchController creates a match controller.
func NewMatchController(service *goa.Service) *MatchController {
	return &MatchController{Controller: service.NewController("MatchController")}
}

// Create runs the create action.
func (c *MatchController) Create(ctx *app.CreateMatchContext) error {
	data := ctx.Payload
	m, err := repositories.CreateMatch(data.HomeTeam, data.AwayTeam, data.HomeUser, data.AwayUser, data.HomeGoals, data.AwayGoals)
	if err != nil {
		return ctx.InternalServerError()
	}
	res := &app.FtMatch{
		HomeTeam:  m.HomeTeam,
		AwayTeam:  m.AwayTeam,
		HomeUser:  m.HomeUserID,
		AwayUser:  m.AwayUserID,
		HomeGoals: m.HomePoints,
		AwayGoals: m.AwayPoints,
	}
	return nil
}

// Delete runs the delete action.
func (c *MatchController) Delete(ctx *app.DeleteMatchContext) error {
	m, err := repositories.GetMatchByID(ctx.MatchID)
	if err != nil {
		return ctx.InternalServerError()
	}
	err = repositories.DeleteMatch(m)
	if err != nil {
		return ctx.InternalServerError()
	}
	return nil
}

// List runs the list action.
func (c *MatchController) List(ctx *app.ListMatchContext) error {
	matches, err := repositories.GetAllMatches(1, 100)
	if err != nil {
		return ctx.InternalServerError()
	}
	resp := app.FtMatchCollection{}
	for _, m := range matches {
		resp = append(resp, &app.FtMatch{
			AwayGoals: m.AwayPoints,
			HomeGoals: m.HomePoints,
			HomeTeam:  m.HomeTeam,
			AwayTeam:  m.AwayTeam,
			HomeUser:  m.HomeUserID,
			AwayUser:  m.AwayUserID,
		})
	}
	return ctx.OK(resp)
}

// Show runs the show action.
func (c *MatchController) Show(ctx *app.ShowMatchContext) error {
	m, err := repositories.GetMatchByID(ctx.MatchID)
	if err != nil {
		return ctx.InternalServerError()
	}
	res := &app.FtMatch{
		HomeTeam:  m.HomeTeam,
		AwayTeam:  m.AwayTeam,
		HomeUser:  m.HomeUserID,
		AwayUser:  m.AwayUserID,
		HomeGoals: m.HomePoints,
		AwayGoals: m.AwayPoints,
	}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *MatchController) Update(ctx *app.UpdateMatchContext) error {
	// MatchController_Update: start_implement

	// Put your logic here

	// MatchController_Update: end_implement
	res := &app.FtMatch{}
	return ctx.OK(res)
}
