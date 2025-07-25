# Lunar Calendar Backend API

## Tổng quan
Backend API cho ứng dụng lịch âm - dương, cung cấp các chức năng chuyển đổi lịch và tra cứu thông tin phong thủy theo truyền thống Việt Nam.

## Tính năng chính
✅ **Chuyển đổi lịch**
- Chuyển đổi từ dương lịch sang âm lịch
- Chuyển đổi từ âm lịch sang dương lịch
- Hỗ trợ tháng nhuận

✅ **Thông tin chi tiết**
- Can Chi (năm, tháng, ngày, giờ)
- Giờ hoàng đạo
- Tiết khí
- Thông tin ngày trong tuần, tuần trong năm

✅ **Lịch tháng**
- Hiển thị toàn bộ lịch tháng với thông tin âm/dương
- Đánh dấu ngày hôm nay và cuối tuần

✅ **Tìm ngày tốt/xấu**
- Phân tích ngày tốt/xấu theo phong thủy
- Hỗ trợ tìm kiếm theo mục đích (cưới hỏi, kinh doanh, du lịch)
- Tính điểm từ 1-10 cho mỗi ngày

## Cài đặt và chạy

### Yêu cầu
- Go 1.19+
- Git

### Cài đặt
```bash
# Clone repository
git clone <repository-url>
cd lunar/backend

# Build project
go build -o lunar-server main.go

# Chạy server
TELEGRAM_BOT_TOKEN=dummy WEBHOOK_URL=dummy ./lunar-server
```

Server sẽ chạy trên port 8080.

## API Endpoints

### 1. Health Check
```bash
curl -X GET http://localhost:8080/api/healthz
```

### 2. Thông tin ngày hôm nay
```bash
curl -X GET http://localhost:8080/api/calendar/today
```

### 3. Chuyển đổi dương lịch sang âm lịch
```bash
# GET method
curl -X GET "http://localhost:8080/api/calendar/solar-to-lunar?day=25&month=7&year=2025"

# POST method
curl -X POST http://localhost:8080/api/calendar/solar-to-lunar \
  -H "Content-Type: application/json" \
  -d '{"day": 25, "month": 7, "year": 2025, "timezone": 7.0}'
```

### 4. Chuyển đổi âm lịch sang dương lịch
```bash
# GET method
curl -X GET "http://localhost:8080/api/calendar/lunar-to-solar?lunar_day=1&lunar_month=6&lunar_year=2025"

# POST method
curl -X POST http://localhost:8080/api/calendar/lunar-to-solar \
  -H "Content-Type: application/json" \
  -d '{"lunar_day": 1, "lunar_month": 6, "lunar_year": 2025, "timezone": 7.0}'
```

### 5. Lịch tháng
```bash
# GET method
curl -X GET "http://localhost:8080/api/calendar/month?month=7&year=2025"

# POST method
curl -X POST http://localhost:8080/api/calendar/month \
  -H "Content-Type: application/json" \
  -d '{"month": 7, "year": 2025, "timezone": 7.0}'
```

### 6. Tìm ngày tốt/xấu
```bash
# GET method
curl -X GET "http://localhost:8080/api/calendar/good-days?start_month=7&start_year=2025&end_month=8&end_year=2025&purpose=wedding"

# POST method
curl -X POST http://localhost:8080/api/calendar/good-days \
  -H "Content-Type: application/json" \
  -d '{"start_month": 7, "start_year": 2025, "end_month": 8, "end_year": 2025, "purpose": "wedding", "timezone": 7.0}'
```

## Ví dụ Response

### Chuyển đổi lịch:
```json
{
  "solar": {
    "day": 25,
    "month": 7,
    "year": 2025,
    "leap_year": false,
    "day_of_year": 205,
    "week_of_year": 30,
    "weekday": "Thứ Sáu",
    "detail": "Dương Lịch: Thứ Sáu, 25/07/2025"
  },
  "lunar": {
    "lunar_day": 1,
    "lunar_month": 6,
    "lunar_year": 2025,
    "leap_month": 1,
    "month_name": "Sáu",
    "can_chi_year": "Ất Tỵ",
    "can_chi_month": "Quý Mùi",
    "can_chi_day": "Ất Mùi",
    "can_chi_hour": "Bính Tý",
    "gio_hoang_dao": "Dần (03-05), Mão (05-07), Tỵ (09-11), Thân (15-17), Tuất (19-21), Hợi (21-23)",
    "tiet_khi": "Tiểu mãn",
    "detail": "Âm Lịch: 00:00 - Giờ Bính Tý, Ngày Ất Mùi, Tháng Quý Mùi, Năm Ất Tỵ - Giờ tốt: Dần (03-05), Mão (05-07), Tỵ (09-11), Thân (15-17), Tuất (19-21), Hợi (21-23) - Tiết Khí: Tiểu mãn"
  }
}
```

### Ngày tốt/xấu:
```json
{
  "good_days": [
    {
      "date": "01/07/2025",
      "solar_day": 1,
      "solar_month": 7,
      "solar_year": 2025,
      "lunar_day": 7,
      "lunar_month": 6,
      "lunar_year": 2025,
      "can_chi_day": "Tân Mùi",
      "score": 7,
      "is_good_day": true,
      "reason": "Ngày lẻ tốt cho cưới hỏi",
      "gio_hoang_dao": "Dần (03-05), Mão (05-07), Tỵ (09-11), Thân (15-17), Tuất (19-21), Hợi (21-23)",
      "tiet_khi": "Lập hạ"
    }
  ],
  "bad_days": [...],
  "total": 31
}
```

## Tính năng nâng cao

### Hỗ trợ múi giờ
Tất cả API đều hỗ trợ parameter `timezone` (mặc định: +7 cho Việt Nam)

### CORS Support
API hỗ trợ CORS cho phép truy cập từ frontend web

### Error Handling
Tất cả lỗi đều trả về format chuẩn:
```json
{
  "error": "Error Type",
  "message": "Detailed error message"
}
```

## Cấu trúc dữ liệu

### Âm lịch
- **Can Chi**: Hệ thống 60 can chi truyền thống
- **Giờ hoàng đạo**: 12 giờ trong ngày với giờ tốt/xấu
- **Tiết khí**: 24 tiết khí trong năm
- **Tháng nhuận**: Hỗ trợ tháng nhuận trong năm âm lịch

### Phong thủy
- **Điểm số**: 1-10 (cao hơn = tốt hơn)
- **Phân loại**: Ngày tốt/xấu dựa trên nhiều yếu tố
- **Mục đích**: Tùy chỉnh theo mục đích sử dụng

## Development

### Cấu trúc project
```
backend/
├── main.go                    # Entry point
├── features/
│   ├── calendar_api.go        # API chuyển đổi lịch
│   ├── month_calendar_api.go  # API lịch tháng
│   ├── good_days_api.go       # API tìm ngày tốt
│   └── calendar/
│       ├── calendar.go        # Core calendar logic
│       ├── lunar.go           # Âm lịch calculations
│       └── solar.go           # Dương lịch calculations
├── shared/                    # Shared utilities
└── doc/                       # Documentation
```

### Testing
```bash
# Test all endpoints
curl -X GET http://localhost:8080/api/calendar/today | jq
curl -X GET "http://localhost:8080/api/calendar/solar-to-lunar?day=25&month=7&year=2025" | jq
curl -X GET "http://localhost:8080/api/calendar/good-days?start_month=7&start_year=2025&end_month=7&end_year=2025" | jq
```

## License
MIT License

## Contributing
1. Fork the project
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request
