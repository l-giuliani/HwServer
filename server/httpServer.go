package server

import "net/http"
import "io"
import "it.etg/gpioServer/services"

func getGpioData(w http.ResponseWriter, r *http.Request) {
	gpioData := services.GpioGetData()
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(gpioData.GetFormattedData(gpioData))
	w.Write(services.FormatGpioData(gpioData))
}

func setGpioData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	w.Write(b)
}

func gpioHnd (w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getGpioData(w, r)
	} else if r.Method == "POST"{
		setGpioData(w, r)
	} else {
		//TODO return 401
	}
}

func Init() {
	http.HandleFunc("/gpio", gpioHnd)

	http.ListenAndServe(":3333", nil)
}