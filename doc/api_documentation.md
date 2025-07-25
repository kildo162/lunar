# Lunar Calendar API Documentation

## Overview
API backend cho ứng dụng lịch âm - dương, cung cấp các chức năng chuyển đổi lịch và tra cứu thông tin phong thủy.

## Base URL
```
http://localhost:8080/api
```

## Endpoints

### 1. Health Check
```
GET /api/healthz
```
Kiểm tra trạng thái hoạt động của server.

**Response:**
```json
{
  "status": "OK",
  "message": "Server is running"
}
```

### 2. Lấy thông tin ngày hôm nay
```
GET /api/calendar/today?timezone=7.0
```

**Parameters:**
- `timezone` (optional): Múi giờ, mặc định là 7.0 (GMT+7)

**Response:**
```json
{
  "solar": {
    "day": 25,
    "month": 7,
    "year": 2025,
    "leap_year": false,
    "day_of_year": 206,
    "week_of_year": 30,
    "weekday": "Thứ Sáu",
    "detail": "Dương Lịch: Thứ Sáu, 25/07/2025"
  },
  "lunar": {
    "lunar_day": 1,
    "lunar_month": 6,
    "lunar_year": 2025,
    "leap_month": 0,
    "month_name": "Sáu",
    "can_chi_year": "Ất Tỵ",
    "can_chi_month": "Quý Mùi",
    "can_chi_day": "Bính Thìn",
    "can_chi_hour": "Mậu Tý",
    "gio_hoang_dao": "Tý (23-01), Mão (05-07), Thìn (07-09), Ngọ (11-13), Dậu (17-19), Hợi (21-23)",
    "tiet_khi": "Đại thử",
    "detail": "Âm Lịch: 00:00 - Giờ Mậu Tý, Ngày Bính Thìn, Tháng Quý Mùi, Năm Ất Tỵ - Giờ tốt: Tý (23-01), Mão (05-07), Thìn (07-09), Ngọ (11-13), Dậu (17-19), Hợi (21-23) - Tiết Khí: Đại thử"
  }
}
```

### 3. Chuyển đổi dương lịch sang âm lịch
```
GET /api/calendar/solar-to-lunar?day=25&month=7&year=2025&hour=10&minute=30&timezone=7.0
POST /api/calendar/solar-to-lunar
```

**GET Parameters:**
- `day`: Ngày (1-31)
- `month`: Tháng (1-12)  
- `year`: Năm
- `hour` (optional): Giờ
- `minute` (optional): Phút
- `second` (optional): Giây
- `timezone` (optional): Múi giờ, mặc định 7.0

**POST Body:**
```json
{
  "day": 25,
  "month": 7,
  "year": 2025,
  "hour": 10,
  "minute": 30,
  "second": 0,
  "timezone": 7.0
}
```

**Response:** Giống như endpoint `/today`

### 4. Chuyển đổi âm lịch sang dương lịch  
```
GET /api/calendar/lunar-to-solar?lunar_day=1&lunar_month=6&lunar_year=2025&leap_month=0&timezone=7.0
POST /api/calendar/lunar-to-solar
```

**GET Parameters:**
- `lunar_day`: Ngày âm lịch (1-30)
- `lunar_month`: Tháng âm lịch (1-12)
- `lunar_year`: Năm âm lịch
- `leap_month` (optional): Tháng nhuận (0 = không nhuận, 1 = nhuận)
- `timezone` (optional): Múi giờ, mặc định 7.0

**POST Body:**
```json
{
  "lunar_day": 1,
  "lunar_month": 6, 
  "lunar_year": 2025,
  "leap_month": 0,
  "timezone": 7.0
}
```

**Response:** Giống như endpoint `/today`

### 5. Lấy lịch tháng
```
GET /api/calendar/month?month=7&year=2025&timezone=7.0
POST /api/calendar/month
```

**GET Parameters:**
- `month` (optional): Tháng (1-12), mặc định tháng hiện tại
- `year` (optional): Năm, mặc định năm hiện tại
- `timezone` (optional): Múi giờ, mặc định 7.0

**POST Body:**
```json
{
  "month": 7,
  "year": 2025,
  "timezone": 7.0
}
```

**Response:**
```json
{
  "month": 7,
  "year": 2025,
  "month_name": "Tháng Bảy",
  "total_days": 31,
  "days": [
    {
      "solar_day": 1,
      "solar_month": 7,
      "solar_year": 2025,
      "lunar_day": 6,
      "lunar_month": 5,
      "lunar_year": 2025,
      "leap_month": 0,
      "weekday": "Thứ Ba",
      "can_chi_day": "Nhâm Tuất",
      "tiet_khi": "",
      "is_today": false,
      "is_weekend": false
    }
    // ... more days
  ]
}
```

### 6. Tìm ngày tốt/xấu
```
GET /api/calendar/good-days?start_month=7&start_year=2025&end_month=8&end_year=2025&purpose=wedding&timezone=7.0
POST /api/calendar/good-days
```

**GET Parameters:**
- `start_month`: Tháng bắt đầu (1-12)
- `start_year`: Năm bắt đầu
- `end_month`: Tháng kết thúc (1-12)
- `end_year`: Năm kết thúc
- `purpose` (optional): Mục đích (wedding, business, travel)
- `timezone` (optional): Múi giờ, mặc định 7.0

**POST Body:**
```json
{
  "start_month": 7,
  "start_year": 2025,
  "end_month": 8,
  "end_year": 2025,
  "purpose": "wedding",
  "timezone": 7.0
}
```

**Response:**
```json
{
  "good_days": [
    {
      "date": "25/07/2025",
      "solar_day": 25,
      "solar_month": 7,
      "solar_year": 2025,
      "lunar_day": 1,
      "lunar_month": 6,
      "lunar_year": 2025,
      "can_chi_day": "Bính Thìn",
      "score": 7,
      "is_good_day": true,
      "reason": "Ngày sóc/vọng (tốt)",
      "gio_hoang_dao": "Tý (23-01), Mão (05-07), Thìn (07-09), Ngọ (11-13), Dậu (17-19), Hợi (21-23)",
      "tiet_khi": "Đại thử"
    }
    // ... more good days
  ],
  "bad_days": [
    // ... bad days with similar structure but is_good_day: false
  ],
  "total": 62
}
```

## Error Response Format
Tất cả lỗi đều trả về theo format:
```json
{
  "error": "Error Type",
  "message": "Detailed error message"
}
```

## Status Codes
- `200 OK`: Thành công
- `400 Bad Request`: Dữ liệu đầu vào không hợp lệ
- `405 Method Not Allowed`: Phương thức HTTP không được hỗ trợ
- `500 Internal Server Error`: Lỗi server

## CORS Support
API hỗ trợ CORS với:
- `Access-Control-Allow-Origin: *`
- `Access-Control-Allow-Methods: GET, POST, OPTIONS`
- `Access-Control-Allow-Headers: Content-Type`

## Notes
- Tất cả endpoints đều hỗ trợ cả GET và POST methods (trừ `/healthz` và `/today` chỉ GET)
- Múi giờ mặc định là +7 (Việt Nam)
- Tìm kiếm ngày tốt được giới hạn tối đa 3 tháng để tránh quá tải
- Can chi, giờ hoàng đạo và tiết khí được tính theo lịch âm truyền thống Việt Nam
