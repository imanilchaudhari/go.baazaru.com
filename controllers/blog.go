package controllers

// Displays the default home page
import (
	"net/http"

	"baazaru.com/components/view"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func BlogIndex(w http.ResponseWriter, r *http.Request) {
	// Display the view
	v := view.New(r)
	v.Name = "blog_index"
	v.Caching = false
	v.Render(w)
}

func BlogView(w http.ResponseWriter, r *http.Request) {
	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)
	slug := params.ByName("slug")

	// Display the view
	v := view.New(r)
	v.Vars["slug"] = slug
	v.Name = "blog_view"
	v.Caching = false
	v.Render(w)
}
