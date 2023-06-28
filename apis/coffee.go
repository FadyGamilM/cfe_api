package apis

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FadyGamilM/cfe_api/core/services"
)

func GetCoffees(w http.ResponseWriter, r *http.Request) {
	var coffee_repo services.Coffee

	coffees_domain, err := coffee_repo.GetAllCoffees()
	if err != nil {
		log.Println(err)
		return
	}

	response, err := json.Marshal(coffees_domain)

	if err != nil {
		log.Println(err)
	}
	w.Write(response)
}
