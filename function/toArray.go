package push

import (
	"strconv"
	"strings"
)

func ToArray(str string) ([]int, string) {
	arr := strings.Split(str, " ")
	var res []int
	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			return nil, "not a number"
		}
		res = append(res, num)
	}
	return res, ""
}
