package control

import (
	"fmt"
	"net/http"
)

func getAllAppliances(w http.ResponseWriter, r *http.Request) {
	// get appliances by room
	// get from runtime
}

func getAppliancesByRoom(room string, w http.ResponseWriter, r *http.Request) {
	// get appliances by room
	// get from runtime
}

func controlApplianceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Control Appliance Handler`")
	if r.Method == http.MethodGet {
		params := r.URL.Query()

		room := params.Get("room")
		if room != "" {
			getAppliancesByRoom(room, w, r)
		}

		//id := r.URL.Query().Get("id")
		//t := r.URL.Query().Get("type")
	} else {
		getAllAppliances(w, r)
	}
}
