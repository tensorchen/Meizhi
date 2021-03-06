package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/Meizhi/today", Today)
	http.HandleFunc("/Meizhi/api/today", ApiToday)
	http.HandleFunc("/Meizhi/api/swiper", ApiSwiper)
	http.HandleFunc("/health-check", HealthCheck)
	http.ListenAndServe(":8000", nil)
}
