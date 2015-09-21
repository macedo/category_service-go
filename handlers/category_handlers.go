package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/macedo/category_service-go/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/macedo/category_service-go/models"
)

func CategoryIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	categories := models.Categories{
		models.Category{Name: "Clothes"},
		models.Category{Name: "Electronic"},
	}

	json.NewEncoder(w).Encode(categories)
}

func CategoryShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	fmt.Fprintln(w, "Category show:", categoryId)
}
