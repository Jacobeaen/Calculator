package rpn

// Проверка на правильность расставления скобок
func IsBracketsCorrect(source string) bool {
	count := 0
	for _, el := range source {
		if count < 0 {
			return false
		}

		if el == '(' {
			count++
		} else if el == ')' {
			count--
		}
	}

	return count == 0
}

// Является ли символ допустимым
func IsSymbolAllowed(symbol rune, array []rune) bool {
	for _, el := range array {
		if symbol == el {
			return true
		}
	}

	return false
}

// Проверка строки на правильный формат
func IsStringCorrect(source string) bool {
	array := []rune{
		'+', '0', '1',
		'-', '2', '3',
		'*', '4', '5',
		'/', '6', '7',
		'(', '8', '9',
		')', ' ', '.',
	}

	for _, el := range source {
		if !IsSymbolAllowed(el, array) {
			return false
		}
	}

	return IsBracketsCorrect(source)
}
