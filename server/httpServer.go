package server

import "net/http"
import "it.etg/gpioServer/services"

func getGpioData(w http.ResponseWriter, r *http.Request) {
	gpioData := services.GpioGetData()
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(gpioData.GetFormattedData(gpioData))
	w.Write(gpioData.GetFormattedData())
}

func gpioHnd (w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getGpioData(w, r)
	} else if r.Method == "POST"{

	}
}

func Init() {
	http.HandleFunc("/gpio", gpioHnd)

	http.ListenAndServe(":3333", nil)
}