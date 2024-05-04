package phone

import (
	"regexp"
	"strconv"
	"strings"
)

func HiddenPhone(s string) string {
	return ProcessPhone(strings.Split(removeSpecialChar(s), " "), 0, s)
}

// return phone match and string hidden phone
func ProcessPhone(data []string, from int, s string) string {
	//fmt.Println("Bat dau kiem tra phan tu: ", from)
	var num string // Number string tim duoc
	milestone := from
	//fmt.Println("(Milestone Checking: ", milestone, ")")
	var save []int
	for i:=from;i<len(data);i++ {
		//fmt.Println(i, "\t=>\t", data[i])
		if data[i] != "" {
			// Hàm kiểm tra ký tự có phải là ký tự đặc biệt không
			isSpecial := func(r rune) bool {
				return !unicode.IsLetter(r) && !unicode.IsNumber(r)
			}

			// Trim các ký tự đặc biệt ở đầu và cuối chuỗi
			data[i] = strings.TrimFunc(data[i], isSpecial)
			if _, err := strconv.ParseInt(data[i],10,64); err == nil {
				//fmt.Println("Tim duoc day so: ----> ", data[i])
				num += data[i]
				//fmt.Println("Day so hien tai dang co: ", num)

				save = append(save, i)
				if len(num)==10 || len(num)==11{
					// check dau so valid 090, 091 092 .....
					if getPhoneVnValid(num) {

						// danh dau offset de replace
						offsetStart := 0
						offsetEnd := 0

						for j:=0;j<save[0];j++ {
							offsetStart += len(data[j])+1
						}
						offsetEnd = offsetStart
						for k:=0;k<len(save);k++ {
							offsetEnd += len(data[save[k]])+1
							//fmt.Println("data[save[k]]", data[save[k]])
							//fmt.Println("offsetEnd", offsetEnd)
						}
						if offsetEnd > len(s) {
							offsetEnd = len(s)
						}

						//fmt.Println("save", save)
						//fmt.Println("1-----", s[:offsetStart])
						//fmt.Println("2-----", s[offsetStart:offsetEnd])
						//fmt.Println("3-----", s[offsetEnd:])
						reg, err := regexp.Compile("[0-9]")
						if err != nil {
							s = s[:offsetStart] + "***********" + s[offsetEnd:]
						} else {
							s = s[:offsetStart] + reg.ReplaceAllString(s[offsetStart:offsetEnd], "*") + s[offsetEnd:]
						}
						//fmt.Println("4-----", s)
						//fmt.Println("---------------------------------!!! MATCH:", num)

						// cap nhat lai milestone hien tai
						milestone = i

						// luu vao mang chua ket qua
						//result = append(result, num)
						continue
					} else {
						// dau so khong hop le
						// tiep tuc tim kiem lai bat dau tu Milestone + 1
						s = ProcessPhone(data, milestone+1, s)
						break
					}
				} else if len(num)>11 {
					//fmt.Println("Da lay qua 10 so, day khong phai la so dien thoai, reset de kiem lai lai bat dau tu Milestone: ", milestone+1)
					// tiep tuc tim kiem lai bat dau tu Milestone + 1
					s = ProcessPhone(data, milestone+1, s)
					save = save[:0]
					break
				}
			} else {
				//fmt.Println("Phan tu ", i ," khong phai la day so, reset de kiem tra lai bat dau tu phan tu ", i+1)
				// thuc hien de qui de tiep tuc tim kiem
				s = ProcessPhone(data, i+1, s)
				save = save[:0]
				break
			}
		}
	}

	return s
}
