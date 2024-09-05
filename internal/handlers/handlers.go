package handlers

import (
	"net/http"
)

/*
-- ==========================================
-- -- server functions					-- --
-- ==========================================
*/

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static")

	// TODO: modular renderPage function (renders views? renders raw html passed in?)

	// TODO: 3 options for routing (writing http...)
	// err := ...
	// w.Write()
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
