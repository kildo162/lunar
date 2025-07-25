package api

import (
	"backend/features/service"
	"net/http"
)

// Handler chuyển đổi dương lịch sang âm lịch
func HandleConvertToLunar(w http.ResponseWriter, r *http.Request) {
	service.ConvertToLunarService(w, r)
}

// Handler chuyển đổi âm lịch sang dương lịch
func HandleConvertToSolar(w http.ResponseWriter, r *http.Request) {
	service.ConvertToSolarService(w, r)
}

// Handler lấy thông tin ngày hôm nay
func HandleGetToday(w http.ResponseWriter, r *http.Request) {
	service.GetTodayService(w, r)
}

// Handler lấy lịch tháng
func HandleGetMonthCalendar(w http.ResponseWriter, r *http.Request) {
	service.GetMonthCalendarService(w, r)
}

// Handler tìm ngày tốt/xấu
func HandleSearchGoodDays(w http.ResponseWriter, r *http.Request) {
	service.SearchGoodDaysService(w, r)
}

// Handler API root
func HandleAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Lunar Calendar API is running"}`))
}

// Handler healthz
func HandleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "OK"}`))
}
