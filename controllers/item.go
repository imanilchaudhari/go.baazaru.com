package controllers

import (
	"net/http"

	"baazaru.com/components/view"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

// Displays the specific item
func ItemView(w http.ResponseWriter, r *http.Request) {
	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	slug := params.ByName("slug")

	v := view.New(r)
	v.Vars["slug"] = slug
	v.Name = "item_view"
	v.Render(w)
}
