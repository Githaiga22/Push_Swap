package sortMinimum

func SortOperations(operations []string) []string {
	var res []string
	for i := 0; i < len(operations); i++ {
		if operations[i] == "sb" && operations[i+1] == "sa" {
			res = append(res, "ss")
			i++
			continue
		} else if operations[i] == "rb" && operations[i + 1] == "ra" {
			res = append(res, "rr")
			i++
			continue
		} else if operations[i] == "rrb" && operations[i + 1] == "rra" {
			res = append(res, "rrr")
			i++
			continue
		}
		res = append(res, operations[i])
	}

	return res
}
