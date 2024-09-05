package handlers

import (
	"fmt"
	"net/http"
)

/*
-- ==========================================
-- -- server functions					-- --
-- ==========================================
*/

func Home(w http.ResponseWriter, r *http.Request) {
	// TODO: modular renderPage function (renders views? renders raw html passed in?)
	// err := ...
	// w.Write()
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
