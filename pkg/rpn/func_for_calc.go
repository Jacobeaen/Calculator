package rpn

// Кол-во элементов в срезе меньше указаного
func MinLength(element, start int, open []int) int {
	count := 0
	for _, open_br := range open[start:] {
		if element > open_br {
			count += 1
		} else {
			break
		}
	}

	return count
}

// Есть ли пара [i, j] в массиве [[], [], ... , []]
func IsPairInArray(pair []int, array [][]int) bool {
	for _, brackets := range array {
		if brackets[0] == pair[0] && brackets[1] == pair[1] {
			return true
		}
	}

	return false
}

// Есть ли символ в строке
func IsSymbolInString(el rune, source []rune) bool {
	for _, x := range source {
		if x == el {
			return true
		}
	}

	return false
}

// Индексы открывающей и закрывающей ее скобок
func PairsBracketsIndexes(source string, open []int) ([][]int, error) {
	result := make([][]int, 0)

	if len(open) == 0 {
		return nil, ErrBrackets
	}

	for _, open_br := range open {
		close_br := open_br + 1
		count := 1

		for count != 0 {
			if source[close_br] == '(' {
				count++
			} else if source[close_br] == ')' {
				count--
			}
			close_br++
		}
		close_br--

		pair := []int{open_br, close_br}
		result = append(result, pair)
	}

	return result, nil
}

// Есть в строке скобки
func IsSubstringHaveBrackets(source string, pair []int) bool {
	start_i := pair[0] + 1
	end_i := pair[1]

	for _, symbol := range source[start_i:end_i] {
		if symbol == '(' || symbol == ')' {
			return false
		}
	}

	return true
}

// Список скобок без внутрених скобок
func OnlySimpleBreakets(source string, array [][]int) (result [][]int) {
	for _, pair := range array {
		if IsSubstringHaveBrackets(source, pair) {
			if !IsPairInArray(pair, result) {
				result = append(result, pair)
			}
		}
	}

	return result
}

// Слайс из индексов закрывающих скобок
func OpenBracketsIndexes(source string) (result []int, err error) {
	for index, el := range source {
		if el == '(' {
			result = append(result, index)
		}
	}

	if len(result) != 0 {
		return result, nil
	}

	return result, ErrBrackets
}

// Последний элемент массива
func GetLastElement(array []string) string {
	return array[len(array)-1]
}

func IsSymbolDigit(symbol rune) bool {
	return symbol >= '0' && symbol <= '9'
}
