// 以 62 進位的方式進行數字與文字間的轉換，如： FF 代表 945。
//
// 關於表示 62 進位的用字符如下：
//   0-9, A-Z, a-z
package carry62

import "strings"

var symbol62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var symbol62List = strings.Split(symbol62, "")

// 10 進位轉 62 進位。
func ToString(num10 int) string {
	var remainder int
	var num62 string

	if num10 == 0 {
		return "0"
	}

	for num10 != 0 {
		remainder = num10 % 62
		num62 = symbol62List[remainder] + num62
		num10 = num10 / 62
	}

	return num62
}

// 62 進位轉 10 進位。
// 如果出現不允許字符就回傳 -1。
func ToNumber(num62 string) int {
	var num10 int
	listTxt := []byte(num62)
	multiple := 1

	for idx, leng := 0, len(listTxt); idx < leng; idx++ {
		val := byte(listTxt[leng-1-idx])
		num := strings.IndexByte(symbol62, val)
		if num != -1 {
			num10 = num10 + num*multiple

			multiple = multiple * 62
		} else {
			return -1
		}
	}

	return int(num10)
}
