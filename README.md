# find-phone

Thư viện Go để tìm kiếm và xử lý số điện thoại Việt Nam trong văn bản.

## Cài đặt

```bash
go get github.com/sonnam0904/find-phone
```

## Tính năng

- Tìm kiếm số điện thoại Việt Nam và quốc tế trong chuỗi văn bản
- Hỗ trợ nhiều định dạng số điện thoại (từ 10 đến 15 chữ số)
- Hỗ trợ các đầu số phổ biến của Việt Nam (03, 05, 07, 08, 09, 843, 845, 847, 848, 849, 840)
- Hỗ trợ mã quốc gia của hơn 40 quốc gia trên thế giới
- Ẩn số điện thoại trong văn bản (thay thế bằng dấu *)
- Tự động xử lý các ký tự đặc biệt (dấu chấm, dấu phẩy, dấu gạch ngang, v.v.)

## Cách sử dụng

### Import thư viện

```go
import (
	"github.com/sonnam0904/find-phone"
)
```

### Tìm kiếm số điện thoại trong chuỗi

```go
// Lấy tất cả số điện thoại từ chuỗi văn bản
text := "Liên hệ với tôi qua số 0987654321 hoặc 0912345678"
phones := phone.GetPhone(text)
// Kết quả: []string{"0987654321", "0912345678"}
```

### Ẩn số điện thoại trong chuỗi

```go
// Thay thế tất cả số điện thoại bằng dấu *
text := "Liên hệ với tôi qua số 0987654321"
hiddenText := phone.HiddenPhone(text)
// Kết quả: "Liên hệ với tôi qua số **********"
```

## Định dạng số điện thoại được hỗ trợ

### Số điện thoại Việt Nam
- Số điện thoại 10 chữ số: 03xxxxxxxx, 05xxxxxxxx, 07xxxxxxxx, 08xxxxxxxx, 09xxxxxxxx
- Số điện thoại 11 chữ số: 843xxxxxxxx, 845xxxxxxxx, 847xxxxxxxx, 848xxxxxxxx, 849xxxxxxxx
- Số điện thoại 12 chữ số: 840xxxxxxxxx

### Số điện thoại quốc tế
Hỗ trợ mã quốc gia của hơn 40 quốc gia, bao gồm:
- +1 (Mỹ, Canada)
- +44 (Anh)
- +49 (Đức)
- +33 (Pháp)
- +39 (Ý)
- +34 (Tây Ban Nha)
- +61 (Úc)
- +91 (Ấn Độ)
- +81 (Nhật Bản)
- +82 (Hàn Quốc)
- +86 (Trung Quốc)
- +66 (Thái Lan)
- +60 (Malaysia)
- +65 (Singapore)
- +62 (Indonesia)
- Và nhiều quốc gia khác

Xem danh sách đầy đủ trong file `list_tel_code.md`

## Yêu cầu

- Go 1.17 trở lên

## Giấy phép

MIT
