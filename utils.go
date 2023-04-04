package phone

import (
	"strings"
)

func removeSpecialChar (s string) string {
	rm := []string{",", "-", ".", "+", "\n", "\r", "\t"}
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


func getPhoneVnValid(s string) bool {
	var check []string
	if len(s)==12 {
		check = []string{"840"}
	} else if len(s)==11 {
		check = []string{"843", "845", "847", "848", "849"}
	} else if len(s)==10 {
		check = []string{"03", "05", "07", "08", "09"}
	} else {
		return false
	}

	for _, v :=range check {
		if strings.Index(s, v)==0 {
			// hop le
			return true
		}
	}

	return false
}
