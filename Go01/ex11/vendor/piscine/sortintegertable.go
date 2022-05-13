package piscine

func tableLen(table []int) int {
	length := 0
	for range table {
		length++
	}
	return length
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func quickSort(table []int, left int, right int) {
	if left >= right {
		return
	}
	i := left
	j := right
	pivot := table[left]
	for {
		for table[i] < pivot {
			i++
		}
		for table[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		swap(&table[i], &table[j])
		i++
		j--
	}
	quickSort(table, left, i-1)
	quickSort(table, j+1, right)
}

func SortIntegerTable(table []int) {
	length := tableLen(table)
	quickSort(table, 0, length-1)
}
