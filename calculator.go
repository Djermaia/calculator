package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Читаем ввод
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')

	// Вычисляем результат выражения
	result, err := calculate(expression)
	if err != nil {
		// Если произошла ошибка, выводим ее
		fmt.Println("Ошибка:", err)
		return
	}

	// Выводим результат
	fmt.Println("Результат:", result)
}

func calculate(expression string) (string, error) {
	// Разделяем выражение на отдельные токены
	tokens := strings.Fields(expression)
	if len(tokens) != 3 {
		return "", fmt.Errorf("формат математической операции не удовлетворяет заданию — два операнда и один " +
			"оператор (+, -, /, *)")
	}

	// Парсим первое число
	a, err := parseNumber(tokens[0])
	if err != nil {
		return "", err
	}

	// Получаем оператор
	operator := tokens[1]

	// Парсим второе число
	b, err := parseNumber(tokens[2])
	if err != nil {
		return "", err
	}

	// Проверка на соответствие типов чисел (арабские или римские)
	isRomanA := isRomanNumeral(tokens[0])
	isRomanB := isRomanNumeral(tokens[2])
	if isRomanA != isRomanB {
		return "", fmt.Errorf("используются одновременно разные системы счисления")
	}

	// Выполняем операцию в зависимости от оператора
	result := 0
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		return "", fmt.Errorf("не соответствует одному из операторов: +, -, /, *")
	}

	// Если числа были в римской системе счисления, преобразуем результат в римские цифры
	if isRomanA {
		if result <= 0 {
			return "", fmt.Errorf("результатом работы калькулятора с римскими числами могут быть только " +
				"положительные числа")
		}
		return toRoman(result), nil
	}
	// Конвертируем результат в строку и возвращаем его
	return strconv.Itoa(result), nil
}

func parseNumber(str string) (int, error) {
	// Проверяем, является ли строка числом в арабской системе счисления
	if val, err := strconv.Atoi(str); err == nil {
		if val < 1 || val > 10 {
			return 0, fmt.Errorf("число выходит за диапазон от 1 до 10")
		}
		return val, nil
	}

	// Проверяем, является ли строка числом в римской системе счисления
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	val, ok := romanNumerals[str]
	if !ok {
		return 0, fmt.Errorf("неверное число")
	}

	return val, nil
}

func isRomanNumeral(str string) bool {
	// Проверка, является ли строка римским числом
	romanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	for _, numeral := range romanNumerals {
		if str == numeral {
			return true
		}
	}

	return false
}

func toRoman(n int) string {
	// Преобразование арабского числа в римское
	romanValues := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNumerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var result strings.Builder
	// Итерация по значениям и добавление соответствующего римского числа в результирующую строку
	for i := 0; i < len(romanValues); i++ {
		for n >= romanValues[i] {
			result.WriteString(romanNumerals[i])
			n -= romanValues[i]
		}
	}

	return result.String()
}
