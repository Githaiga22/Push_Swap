package push

import "fmt"

func PushMinimum(input []int) []string {
	var res []string
	arr := append([]int{}, input...)
	var arr2 []int
	max := 0
	j := 0
	hld := 0

	for !IsSorted(arr) {

		if arr[0] > arr[1] {
			arr[1], arr[0] = arr[0], arr[1]
			res = append(res, "sa")
			continue
		}

		for i := 0; i < len(arr); i++ {
			if arr[i] > max && arr[i] != hld {
				max = arr[i]
				j = i
			}
		}

		if j == len(arr)-1 {
			j = 0
			hld = max
			max = 0
			continue
		}

		for i := 0; i <= j-1; i++ {
			fmt.Println(arr[i])
			if i < len(arr)-1 {
				res = append(res, "pb")
				arr2, res = SortArr2(arr[i], arr2, res)
			}
		}

		for i := 0; i <= j - 1; i++ {
			if i < len(arr)-1 {
				arr = PushToB(arr)
			}
		}

	}
	for _, v := range arr2 {
		arr = append([]int{v}, arr...)
		res = append(res, "pa")
	}
	fmt.Println(res)
	fmt.Println(arr)
	fmt.Println(arr2)
	return res
}

func SortArr2(num int, arr2 []int, res []string) ([]int, []string) {
	if arr2 == nil {
		arr2 = append(arr2, num)
		return arr2, res
	}
	if num > arr2[0] {
		arr2 = append([]int{num}, arr2...)
	} else if num < arr2[len(arr2)-1] {
		arr2 = append(arr2, num)
		res = append(res, "rb")
	} else if num < arr2[0] && num > arr2[len(arr2)-1] {
		arr2 = append(arr2, num)
		arr2[0], arr2[1] = arr2[1], arr2[0]
	}
	return arr2, res
}

func PushToB(arr []int) (res []int) {
	for i := 1; i < len(arr); i++ {
		res = append(res, arr[i])
	}
	fmt.Println(res)
	return res
}
