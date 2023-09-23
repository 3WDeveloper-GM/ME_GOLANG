package method

func CountElementArray(nums []int) (map[int]int, int) {
	dict := make(map[int]int)
	counter := 0
	for _, val := range nums {
		if _, ok := dict[val]; ok {
			dict[val]++
		} else {
			dict[val] = 1
		}
		if dict[val] > counter {
			counter++
		}
	}
	return dict, counter
}

func DictToarray(dict map[int]int) [][]int {
	b := make([][]int, 0)
	for len(dict) > 0 {
		a := make([]int, 0)
		for key, value := range dict {
			if value != 0 {
				dict[key] = value - 1
				a = append([]int{key}, a...)
				// fmt.Println(a)
			} else if value == 0 {
				delete(dict, key)
			}
		}
		if len(dict) == 0 {
			break
		}
		b = append(b, [][]int{a}...)
		// fmt.Println(len(dict))
		// fmt.Println(a)
	}
	return b
}
