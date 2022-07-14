package util

// DuplicateUintSlice uint Array deduplicate
func DeDuplicateUintSlice(input []uint) (output []uint) {
	output = make([]uint, 0, 10)

	if len(input) == 0 {
		return
	}

	var duplicate = make(map[uint]struct{})

	for _, v := range input {
		if _, ok := duplicate[v]; !ok {
			duplicate[v] = struct{}{}
			output = append(output, v)
		}
	}

	return
}
