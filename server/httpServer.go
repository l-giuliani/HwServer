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
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	
	res := services.GpioWrite(data)
	if !res {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(204)
	}	
}

func gpioHnd (w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getGpioData(w, r)
	} else if r.Method == "POST"{
		setGpioData(w, r)
	} else {
		w.WriteHeader(405)
	}
}

func Init() {
	http.HandleFunc("/gpio", gpioHnd)

	http.ListenAndServe(":3333", nil)
}