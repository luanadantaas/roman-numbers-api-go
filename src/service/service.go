package service

import (
	"net/http"
	"sort"
	"strings"

	"github.com/luanadantaas/roman-numbers-api-go/src/entity"
	"github.com/luanadantaas/roman-numbers-api-go/src/numbers"

	"github.com/gin-gonic/gin"
)

func FindPrimeNum(c *gin.Context) {
	var response entity.ApiResponse
	var number entity.Number

	c.Bind(&response)

	prime := splitRomanNumbers(response.Text)

	sort.Ints(prime)
	maxPrime := numbers.CheckPrimeNumber(prime)
	romanMaxPrime := numbers.ConvertToRoman(maxPrime)

	number.Number = romanMaxPrime
	number.Value = maxPrime
	c.JSON(http.StatusOK, number)
}

func splitRomanNumbers(text string) []int {
	var prime []int
	letters := strings.Split(text, "")
	//"text": "AXIBIV"
	for i := 0; i < len(letters); i++ {
		if (i > 0 && ((letters[i-1] != "I") && (letters[i-1] != "V") && (letters[i-1] != "X") && (letters[i-1] != "L") && (letters[i-1] != "C") && (letters[i-1] != "D") && (letters[i-1] != "M"))) || i == 0 {
			if (letters[i] == "I") || (letters[i] == "V") || (letters[i] == "X") || (letters[i] == "L") || (letters[i] == "C") || (letters[i] == "D") || (letters[i] == "M") {
				var num []string
				num = append(num, letters[i])
				for j := i + 1; j < len(letters); j++ {
					if (letters[j] == "I") || (letters[j] == "V") || (letters[j] == "X") || (letters[j] == "L") || (letters[j] == "C") || (letters[j] == "D") || (letters[j] == "M") {
						num = append(num, letters[j])
					} else {
						break
					}
				}
				number := numbers.ConvertToDecimal(num)
				if number != 0 {
					prime = append(prime, number)
				}
			}
		}
	}
	return prime
}
