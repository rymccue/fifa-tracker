package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("FIFA Tracker", func() {
	Title("The Freshworks FIFA Tracker")
	Description("An API to keep track of FIFA games played")
	Contact(func() {
		Name("Ryan McCue")
		Email("ryanm@freshworks.io")
	})
	Host("localhost:8080")
	Scheme("http")
	BasePath("/")
	Consumes("application/json")
	Produces("application/json")
})
