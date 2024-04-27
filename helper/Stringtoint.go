package helper

import "strconv"

func ParseInt(element string) (int, error) {
	if len(element) == 0 {
		return -1, nil
	}
	result, err := strconv.Atoi(element)
	if err != nil {
		return -1, err
	}

	return result, nil
}
