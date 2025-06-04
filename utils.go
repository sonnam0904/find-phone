package phone

import (
	"strings"
)

func removeSpecialChar(s string) string {
	rm := []string{":", ",", "-", ".", "+", "\n", "\r", "\t"}
	for _, v := range rm {
		s = strings.ReplaceAll(s, v, " ")
	}
	//reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	//if err != nil {
	//	return s
	//}
	//s = reg.ReplaceAllString(s, " ")
	return s
}

// Kiểm tra số điện thoại quốc tế có hợp lệ không
func getPhoneIntlValid(s string) bool {
	// Kiểm tra nếu là đầu số Việt Nam (84), trả về false để xử lý riêng theo rule cũ
	if strings.HasPrefix(s, "84") {
		return false
	}

	// Danh sách mã quốc gia từ list_tel_code.md
	intlCodes := map[string]bool{
		"1":  true, // United States, Canada
		"7":  true, // Russia
		"20": true, // Egypt
		"27": true, // South Africa
		"30": true, // Greece
		"31": true, // Netherlands
		"32": true, // Belgium
		"33": true, // France
		"34": true, // Spain
		"39": true, // Italy
		"41": true, // Switzerland
		"43": true, // Austria
		"44": true, // United Kingdom
		"45": true, // Denmark
		"46": true, // Sweden
		"47": true, // Norway
		"49": true, // Germany
		"52": true, // Mexico
		"54": true, // Argentina
		"55": true, // Brazil
		"60": true, // Malaysia
		"61": true, // Australia
		"62": true, // Indonesia
		"63": true, // Philippines
		"64": true, // New Zealand
		"65": true, // Singapore
		"66": true, // Thailand
		"81": true, // Japan
		"82": true, // South Korea
		// "84": true,   // Vietnam - Đã loại bỏ để xử lý riêng theo rule cũ
		"86":  true, // China
		"90":  true, // Turkey
		"91":  true, // India
		"92":  true, // Pakistan
		"234": true, // Nigeria
		"254": true, // Kenya
		"351": true, // Portugal
		"353": true, // Ireland
		"358": true, // Finland
		"852": true, // Hong Kong
		"880": true, // Bangladesh
		"886": true, // Taiwan
		"966": true, // Saudi Arabia
		"971": true, // UAE
		"972": true, // Israel
		"974": true, // Qatar
	}

	// Kiểm tra độ dài số điện thoại quốc tế (tối đa 15 số)
	if len(s) > 15 {
		return false
	}

	// Kiểm tra mã quốc gia
	for code := range intlCodes {
		if strings.Index(s, code) == 0 {
			return true
		}
	}

	return false
}

func getPhoneVnValid(s string) bool {
	var check []string
	if len(s) == 12 {
		check = []string{"840"}
	} else if len(s) == 11 {
		check = []string{"843", "845", "847", "848", "849"}
	} else if len(s) == 10 {
		check = []string{"03", "05", "07", "08", "09"}
	} else {
		// Kiểm tra số điện thoại quốc tế
		return getPhoneIntlValid(s)
	}

	for _, v := range check {
		if strings.Index(s, v) == 0 {
			// hop le
			return true
		}
	}

	// Nếu không phải số Việt Nam, kiểm tra số quốc tế
	return getPhoneIntlValid(s)
}
