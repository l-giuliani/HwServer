package server

import "net/http"
import "io"
import "strconv"
import "strings"
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

func wdHndStart(w http.ResponseWriter, r *http.Request) {
	time := strings.TrimPrefix(r.URL.Path, "/wd/start/")
	t, err := strconv.Atoi(time)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	res := services.StartWd(uint8(t))
	if res {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(500)
	}
}

func wdHndReset(w http.ResponseWriter, r *http.Request) {
	time := strings.TrimPrefix(r.URL.Path, "/wd/reset/")
	t, err := strconv.Atoi(time)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	res := services.ResetWd(uint8(t))
	if res {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(500)
	}
}

func wdHndStop(w http.ResponseWriter, r *http.Request) {
	res := services.StopWd()
	if res {
		w.WriteHeader(204)
	} else {
		w.WriteHeader(500)
	}
}

func wdHnd (w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { 
		action := strings.TrimPrefix(r.URL.Path, "/wd/")
		if strings.Contains(action, "reset") {
			wdHndReset(w, r)
		} else if strings.Contains(action, "start") {
			wdHndStart(w, r)
		} else if action == "stop" {
			wdHndStop(w, r)
		} else {
			w.WriteHeader(500)
		}
	} else {
		w.WriteHeader(405)
	}
}

func Init() {
	http.HandleFunc("/gpio", gpioHnd)
	http.HandleFunc("/wd/", wdHnd)
	
	http.ListenAndServe(":3333", nil)
}