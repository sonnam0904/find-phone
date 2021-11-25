package phone

import (
	"testing"
	"fmt"
)

func TestGetPhone(t *testing.T) {
	testData := []struct{
		input, expected string
	}{
		{"10/ 08- 2020. 09872  33333 111 22 day ne", ""},
		{"8888888888888 097.999 99.     09 xxxxxx xxx", ""},
		{"8888888888888 097.999.9999. 09 xxxxxx xxx", ""},
		{"0987 6567 xxxx xxx 87 xxxx 098 8987 765 xxxx", ""},
		{"0987 6567 xxxx xxx 87 xxxx 098 8987 765 x. xxx 0988888 888", ""},
		{"0987 6567 xxxx xxx 87 xxxx 888876765. 4543234 xxxx. xxx .0988888 888", ""},
		{"097.999.9998 097.999.9999. 09 xxxxxx xxx", ""},
		{"Zalo dùng số điện thoại 0386500999", ""},
		{"10/ 08- 2020. 09872  33333 111 22 day ne", ""},
	}

	for _,i := range testData {
		phone := GetPhone(i.input)
		fmt.Println(phone)
	}
}

func BenchmarkGetPhone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testData := []struct{
			input, expected string
		}{
			{"10/ 08- 2020. 09872  33333 111 22 day ne", ""},
			{"8888888888888 097.999 99.     09 xxxxxx xxx", ""},
			{"8888888888888 097.999.9999. 09 xxxxxx xxx", ""},
			{"0987 6567 xxxx xxx 87 xxxx 098 8987 765 xxxx", ""},
			{"0987 6567 xxxx xxx 87 xxxx 098 8987 765 x. xxx 0988888 888", ""},
			{"0987 6567 xxxx xxx 87 xxxx 888876765. 4543234 xxxx. xxx .0988888 888", ""},
			{"097.999.9998 097.999.9999. 09 xxxxxx xxx", ""},
			{"Zalo dùng số điện thoại 0386500999", ""},
			{"10/ 08- 2020. 09872  33333 111 22 day ne", ""},
		}
		for _,i := range testData {
			_ = GetPhone(i.input)
			//fmt.Println(phone)
		}
	}
}