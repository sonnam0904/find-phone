package phone

import (
	"strconv"
	"strings"
)
func GetPhone(s string) []string {
	var result []string
	return FindPhone(strings.Split(strings.ReplaceAll(s, ".", " "), " "), 0, result)
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

				if len(num) == 10 {
					// stop vi tim dc 10 so de check valid so dien thoại
					//fmt.Println("---------------------------------!!! MATCH:", num)

					// cap nhat lai milestone hien tai
					milestone = i
					// check dau so valid 090, 091 092 .....

					// luu vao mang chua ket qua
					result = append(result, num)
					continue
				} else if len(num)>10 {
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