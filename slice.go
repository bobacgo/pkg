package main

func CrossMergeSlice[T any](arrs ...[]T) []T {
	var maxLen, totalLen int
	for _, arr := range arrs {
		arrLen := len(arr)
		if arrLen > maxLen {
			maxLen = arrLen
		}
		totalLen += arrLen
	}

	mArr := make([]T, 0, totalLen)
	for i := 0; i < maxLen; i++ {
		for _, arr := range arrs {
			if len(arr) > i {
				mArr = append(mArr, arr[i])
			}
		}
	}
	return mArr
}
