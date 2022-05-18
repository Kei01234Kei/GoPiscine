package piscine

func strLen(s string) int {
	length := 0
	for range s {
		length++
	}
	return length
}

func Index(s string, toFind string) int {
	toFindLength := strLen(toFind)
	if toFindLength == 0 {
		return -1
	}
	sameLength := 0
	index := 0
	for i, _ := range s {
		for j, _ := range toFind {
			if s[i+j] == toFind[j] {
				sameLength++
			} else {
				break
			}
		}
		if sameLength == toFindLength {
			return index
		}
		sameLength = 0
		index++
	}
	return -1
}
