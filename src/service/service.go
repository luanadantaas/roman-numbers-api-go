package service

import (
	"desafio/src/entity"
	"desafio/src/numbers"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

func FindPrimeNum(c *gin.Context) {
	var response entity.ApiResponse
	var number entity.Number
	var prime []int
	c.Bind(&response)

	letters := strings.Split(response.Text, "")

	for i, _ := range letters {
		if (i > 0 && ((letters[i-1] != "I") && (letters[i-1] != "V") && (letters[i-1] != "X") && (letters[i-1] != "L") && (letters[i-1] != "C") && (letters[i-1] != "D") && (letters[i-1] != "M"))) || i == 0 {
			if (letters[i] == "I") || (letters[i] == "V") || (letters[i] == "X") || (letters[i] == "L") || (letters[i] == "C") || (letters[i] == "D") || (letters[i] == "M") {
				var num []string
				num = append(num, letters[i])
				j := i + 1
				for j, _ = range letters {
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

	sort.Ints(prime)
	maxPrime := numbers.CheckPrimeNumber(prime)
	romanMaxPrime := number.ConvertToRoman(maxPrime)

	number.Number = romanMaxPrime
	number.Value = maxPrime
	c.JSON(http.StatusOK)
}
