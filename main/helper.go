package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("%s\n", elapsed)
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	i, j := 0, len(s)-1
	for {
		if i > j {
			break
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	if n == 2 {
// 		return true
// 	}
// 	if n%2 == 0 {
// 		return false
// 	}
// 	for i := 3; i <= int(math.Ceil(math.Sqrt(float64(n)))); i += 2 {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i = i + 6
	}
	return true
}

func add(a, b string) string {
	diff := len(b) - len(a)
	abs := int(math.Abs(float64(diff)))
	for i := 1; i <= abs; i++ {
		if diff > 0 {
			a = fmt.Sprintf("%d%s", 0, a)
		} else {
			b = fmt.Sprintf("%d%s", 0, b)
		}
	}
	x, y, temp := 0, 0, 0
	first, second := 0, 0
	result := ""
	for i := len(b) - 1; i >= 0; i-- {
		x, _ = strconv.Atoi(string(b[i]))
		y, _ = strconv.Atoi(string(a[i]))
		temp = x + y
		if first > 0 {
			temp++
		}
		if i == 0 {
			second = temp
		} else {
			if temp > 9 {
				first = temp / 10
				second = temp - 10
			} else {
				first = 0
				second = temp
			}
		}
		s := strconv.Itoa(second)
		result = s + result
	}
	return result
}

func multiply(a, b string) string {
	prod, s := "", ""
	x, y, temp := 0, 0, 0
	first, second := 0, 0
	zeros := 0
	arr := []string{}
	for i := len(b) - 1; i >= 0; i-- {
		prod = ""
		first = 0
		second = 0
		for j := len(a) - 1; j >= 0; j-- {
			x, _ = strconv.Atoi(string(b[i]))
			y, _ = strconv.Atoi(string(a[j]))
			temp = x * y
			if first > 0 {
				temp += first
			}
			if j == 0 {
				second = temp
			} else {
				if temp > 9 {
					first = temp / 10
					second = temp % 10
				} else {
					first = 0
					second = temp
				}
			}
			s = strconv.Itoa(second)
			prod = s + prod
		}
		zeros = len(b) - i - 1
		for k := 1; k <= zeros; k++ {
			prod = fmt.Sprintf("%s%d", prod, 0)
		}
		arr = append(arr, prod)
	}
	result := ""
	for _, v := range arr {
		result = add(result, v)
	}
	return result
}

func pow(x, y string) string {
	n, _ := strconv.Atoi(y)
	result := "1"
	for i := 1; i <= n; i++ {
		result = multiply(result, x)
	}
	return result
}
