package phone

import (
	"strconv"
	"strings"
)
func GetPhone(s string) []string {
	var result []string

	return FindPhone(strings.Split(removeSpecialChar(s), " "), 0, result)
}

func removeSpecialChar (s string) string {
	rm := []string{",", ".", "+", "\n", "\r", "\t"}
	for _, v := range rm {
		s = strings.ReplaceAll(s, v, " ")
	}
	return s
}

func getPhoneVnValid(s string) bool {
	var check []string
	if len(s)==11 {
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

func FindPhone(data []string, from int, result []string) []string {
	//fmt.Println("Bat dau kiem tra phan tu: ", from)
	var num string // Number string tim duoc
	milestone := from
	//fmt.Println("(Milestone Checking: ", milestone, ")")

	for i:=from;i<len(data);i++ {
		//fmt.Println(i, "\t=>\t", data[i])
		if data[i] != "" {
			// check number
			if _, err := strconv.ParseInt(data[i],10,64); err == nil {
				//fmt.Println("Tim duoc day so: ----> ", data[i])
				num += data[i]
				//fmt.Println("Day so hien tai dang co: ", num)

				if len(num)==10 || len(num)==11{
					// check dau so valid 090, 091 092 .....
					if getPhoneVnValid(num) {
						//fmt.Println("---------------------------------!!! MATCH:", num)

						// cap nhat lai milestone hien tai
						milestone = i

						// luu vao mang chua ket qua
						result = append(result, num)
						continue
					} else {
						// dau so khong hop le
						// tiep tuc tim kiem lai bat dau tu Milestone + 1
						result = FindPhone(data, milestone+1, result)
						break
					}
				} else if len(num)>11 {
					//fmt.Println("Da lay qua 10 so, day khong phai la so dien thoai, reset de kiem lai lai bat dau tu Milestone: ", milestone+1)
					// tiep tuc tim kiem lai bat dau tu Milestone + 1
					result = FindPhone(data, milestone+1, result)
					break
				}
			} else {
				//fmt.Println("Phan tu ", i ," khong phai la day so, reset de kiem tra lai bat dau tu phan tu ", i+1)
				// thuc hien de qui de tiep tuc tim kiem
				result = FindPhone(data, i+1, result)
				break
			}
		}
	}

	return result
}