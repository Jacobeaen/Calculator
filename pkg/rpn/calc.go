package rpn

import (
	"strconv"
	"strings"
)

// Операция над двумя числами
func Operation(array []float64, operation rune) (float64, error) {
	if len(array) < 2 {
		return 0, ErrOperationWithoutNumbers
	}

	number1 := array[len(array)-2]
	number2 := array[len(array)-1]

	if operation == '+' {
		return number1 + number2, nil
	} else if operation == '-' {
		return number1 - number2, nil
	} else if operation == '*' {
		return number1 * number2, nil
	} else if operation == '/' {
		if number2 != 0 {
			return number1 / number2, nil
		}
		return 0, ErrZeroDivison
	}

	return 0, ErrUnknownOperation
}

func CalculateExpression(source string) (float64, error) {
	numbers := make([]float64, 0)
	operations := make([]rune, 0)
	signs := []rune{
		'+',
		'-',
		'*',
		'/',
	}
	rating := map[rune]int{
		'+': 1,
		'-': 1,
		'/': 2,
		'*': 2,
	}

	var last_digit string = "_"
	var last_sign rune = '_'
	var last_symbol rune = '_'

	count_numbers := 0
	count_operations := 0

	for _, symbol := range source {
		if IsSymbolDigit(symbol) || symbol == '.' {

			// Если еще не было символов
			if last_digit == "_" {
				numbers = append(numbers, float64(symbol-'0'))

				count_numbers++
				last_digit = string(symbol)

				// Если два или больше цифр подряд
			} else if IsSymbolDigit(last_symbol) || last_symbol == '.' {
				x := last_digit + string(symbol)

				n, _ := strconv.ParseFloat(x, 64)

				count_numbers--
				numbers[count_numbers] = float64(n)
				count_numbers++

				last_digit = x

				// Если просто одиночная цифра
			} else if symbol == '.' {
				last_digit = last_digit + string(symbol)

			} else {
				numbers = append(numbers, float64(symbol-'0'))

				last_digit = string(symbol)
				count_numbers++
			}

		} else if IsSymbolInString(symbol, signs) {
			if len(numbers) < 1 {
				return 0, ErrOperationWithoutNumbers

				// Если первый знак - просто добавляем
			} else if len(operations) == 0 {
				operations = append(operations, symbol)

				count_operations++
				last_sign = symbol
			} else {
				if rating[last_sign] >= rating[symbol] {
					if len(numbers) < 2 {
						return 0, ErrOperationWithoutNumbers
					}
					result, err := Operation(numbers, last_sign)
					if err != nil {
						return 0, err
					}

					count_numbers -= 2
					numbers = numbers[:count_numbers+1]
					numbers[count_numbers] = result
					count_numbers++

					count_operations--
					operations = operations[:count_operations]
					operations = append(operations, symbol)

					if len(numbers) > 1 && len(operations) > 1 {
						if rating[last_sign] >= rating[symbol] {
							result2, _ := Operation(numbers, operations[len(operations)-2])

							count_numbers -= 2
							numbers = numbers[:count_numbers+1]
							numbers[count_numbers] = result2
							count_numbers++

							operations = operations[count_operations:]
							count_operations--
						}
					}
					last_sign = symbol
					count_operations++

				} else {
					operations = append(operations, symbol)

					count_operations++
					last_sign = symbol
				}
			}
		}
		last_symbol = symbol
	}

	j := len(numbers) - 1
	for i := len(operations) - 1; i >= 0; i-- {
		result, err := Operation(numbers, operations[i])

		if err != nil {
			return 0, err
		}
		numbers = numbers[:j]
		j--
		numbers[j] = result

		operations = operations[:i]
	}

	return numbers[0], nil

}

func Calc(source string) (float64, error) {
	source = strings.ReplaceAll(source, " ", "")
	if len(source) == 0 {
		return 0, ErrEmptyExpression
	}

	_, err := IsStringCorrect(source)

	if err != nil {
		return 0, err
	}

	for {
		x, _ := OpenBracketsIndexes(source)
		z, _ := PairsBracketsIndexes(source, x)
		arr := OnlySimpleBreakets(source, z)

		if len(arr) != 0 {
			for i := len(arr) - 1; i >= 0; i-- {
				start_br := arr[i][0]
				end_br := arr[i][1]

				substring := source[start_br+1 : end_br]
				result, _ := CalculateExpression(substring)

				substitution := strconv.FormatFloat(result, 'f', 2, 64)
				source = source[:start_br] + substitution + source[end_br+1:]
			}

		} else {
			result, err := CalculateExpression(source)

			if err != nil {
				return 0, err
			}
			return result, nil
		}
	}
}
