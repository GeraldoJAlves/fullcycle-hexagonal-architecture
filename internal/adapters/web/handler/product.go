package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/adapters/dto"
	"github.com/geraldojalves/fullcycle-hexagonal-architecture/internal/application"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, s application.ProductServiceInterface) {
	r.Handle("/products/{id}", n.With(
		negroni.Wrap(getProduct(s)),
	)).Methods("GET", "OPTIONS")
	r.Handle("/products", n.With(
		negroni.Wrap(createProduct(s)),
	)).Methods("POST", "OPTIONS")
}

func getProduct(s application.ProductServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		vars := mux.Vars(r)
		productId := vars["id"]

		product, err := s.Get(productId)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(product)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func createProduct(s application.ProductServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		var productDto = dto.NewProduct()

		if err := json.NewDecoder(r.Body).Decode(&productDto); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError(err.Error()))
			return
		}
		// var product = application.Product{}
		// if _, err := productDto.Bind(&product); err != nil {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	jsonError(err.Error())
		// 	return
		// }

		product, err := s.Create(productDto.Name, productDto.Price)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
		}

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
		}
	}
}
