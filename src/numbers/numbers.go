package numbers

import (
	"math"
)

var numInv = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var maxTable = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func value(r string) int {
	if r == "I" {
		return 1
	}
	if r == "V" {
		return 5
	}
	if r == "X" {
		return 10
	}
	if r == "L" {
		return 50
	}
	if r == "C" {
		return 100
	}
	if r == "D" {
		return 500
	}
	if r == "M" {
		return 1000
	}

	return -1
}

func highestDecimal(n int) int {
	for _, v := range maxTable {
		if v <= n {
			return v
		}
	}
	return 1
}

func ConvertToRoman(n int) string {
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += numInv[v]
		n -= v
	}
	return out
}

func ConvertToDecimal(roman []string) int {
	dec := 0
	for i, j := range roman {
		l1 := value(j)

		if i+1 < len(roman) {
			l2 := value(roman[i+1])

			if l1 >= l2 {
				dec = dec + l1
			} else {
				dec = dec + l2 - l1
				i++
			}
		} else {
			dec = dec + l1
		}
	}

	return dec
}

func CheckPrimeNumber(nums []int) int {
	var primeNumber []int
	count := 0
	for i := 0; i < len(nums); i++ {
		sq_root := int(math.Sqrt(float64(nums[i])))
		for z := 2; z <= sq_root; z++ {
			if nums[i]%z == 0 {
				count++
			}
		}
		if count == 0 {
			primeNumber = append(primeNumber, nums[i])
		}
	}
	return primeNumber[len(primeNumber)-1]
}
