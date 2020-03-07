package helpers

func CheckValueInt64(slice []int64, item int64) bool {
	set := make(map[int64]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
