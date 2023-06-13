package utils

import (
	"crypto/rand"
	"log"
	"math"
	"math/big"
	"regexp"
	"sort"
	"strconv"
)

func TypeUrlStringProto() string {
	value := "type.googleapis.com/google.protobuf.StringValue"
	return value
}

func GenerateRandomNumber(len int) (int, error) {
	maxLimit := int64(int(math.Pow10(len)) - 1)
	lowLimit := int(math.Pow10(len - 1))

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return 0, err
	}
	randomNumberInt := int(randomNumber.Int64())

	// Handling integers between 0, 10^(n-1) .. for n=4, handling cases between (0, 999)
	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}

	// Never likely to occur, kust for safe side.
	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}
	return randomNumberInt, nil
}
func IntegerToRoman(numberStr string) string {
	number, _ := strconv.Atoi(numberStr)
	romanMap := map[int]string{
		1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L",
		90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M",
	}
	// create a slice of slices
	rows := len(romanMap)
	matrix := make([][]string, rows)
	var key_slice []int
	for k, _ := range romanMap {
		key_slice = append(key_slice, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(key_slice)))
	row := 0
	for _, key := range key_slice {
		// convert int key to string key
		skey := strconv.Itoa(key)
		matrix[row] = []string{skey, romanMap[key]}
		row++

	}
	result := ""
	for _, item := range matrix {
		// convert string to int
		den, err := strconv.Atoi(item[0])
		if err != nil {
			panic(err)
		}
		sym := item[1]
		for number >= den {
			result += sym
			number -= den
		}
	}
	return result
}
func SanitizeString(word string) string {

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	afterSenitize := re.ReplaceAllString(word, "")
	return afterSenitize
}
