package api

import "net/http"
type CatalogAPI struct {
	
}
func CatalogAPI(router *http.ServeMux) {
	router.Handle("/items", func (w http.ResponseWriter, r *http.Request)  {
		handleGetAllItem(w, r)
	})
}