package service

import (
	"math/rand"
)

var primeNumber = []int{47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109}
var referenceCharMap = []int{2, 3, 2, 2, 2, 1, 2, 1, 3, 3, 1, 2, 2, 0, 1, 0, 2, 0, 1, 3, 3, 0, 0, 3, 3, 2, 0, 3, 2, 1, 1, 2, 0, 2, 3, 2, 0, 2, 0, 0, 2, 0, 3, 3, 3, 0, 1, 2, 0, 1, 2, 0, 0, 1, 0, 0, 1, 3, 3, 1, 0, 1, 2, 2, 2, 3, 2, 3, 3, 1, 1, 1, 0, 3, 3, 0, 2, 3, 1, 2, 1, 3, 3, 3, 0, 0, 3, 0, 3, 2, 1, 2, 3, 3, 3, 1, 2, 2, 0, 2, 0, 2, 0, 2, 3, 3, 1, 0, 2}
var referenceAlter1 = []int{3, 0, 0, 1, 1, 0, 3, 3, 1, 2, 0, 3, 3, 1, 0, 1, 0, 3, 0, 0, 1, 3, 3, 1, 1, 1, 1, 1, 0, 3, 2, 3, 1, 3, 1, 0, 1, 0, 2, 1, 1, 2, 1, 2, 2, 1, 0, 3, 1, 2, 3, 2, 1, 3, 2, 1, 0, 1, 2, 0, 1, 0, 3, 3, 3, 0, 0, 2, 1, 3, 3, 0, 2, 1, 2, 2, 3, 0, 3, 1, 3, 2, 1, 1, 2, 3, 0, 1, 1, 0, 0, 1, 0, 1, 0, 2, 0, 3, 2, 1, 2, 0, 3, 0, 2, 1, 2, 3, 0}
var referenceAlter2 = []int{1, 1, 3, 0, 0, 2, 1, 2, 0, 1, 2, 0, 1, 2, 3, 3, 3, 1, 2, 2, 0, 1, 1, 2, 0, 0, 3, 0, 3, 0, 3, 0, 2, 1, 2, 1, 2, 1, 3, 3, 0, 1, 2, 0, 1, 3, 2, 0, 3, 3, 0, 3, 3, 0, 3, 3, 3, 0, 0, 2, 2, 2, 1, 1, 1, 1, 3, 1, 2, 0, 0, 3, 3, 0, 1, 1, 0, 1, 0, 3, 0, 1, 2, 0, 3, 1, 2, 3, 2, 1, 2, 0, 2, 2, 2, 3, 1, 0, 1, 0, 1, 3, 1, 3, 0, 2, 3, 2, 3}
var referenceAlter3 = []int{0, 2, 1, 3, 3, 3, 0, 0, 2, 0, 3, 1, 0, 3, 2, 2, 1, 2, 3, 1, 2, 2, 2, 0, 2, 3, 2, 2, 1, 2, 0, 1, 3, 0, 0, 3, 3, 3, 1, 2, 3, 3, 0, 1, 0, 2, 3, 1, 2, 0, 1, 1, 2, 2, 1, 2, 2, 2, 1, 3, 3, 3, 0, 0, 0, 2, 1, 0, 0, 2, 2, 2, 1, 2, 0, 3, 1, 2, 2, 0, 2, 0, 0, 2, 1, 2, 1, 2, 0, 3, 3, 3, 1, 0, 1, 0, 3, 1, 3, 3, 3, 1, 2, 1, 1, 0, 0, 1, 1}

func GenerateRandomPrime() int {
	i := rand.Intn(17) - 1
	return primeNumber[i]
}

func GenerateMapChar(index int) string {
	k := referenceCharMap[index]
	return returnCharOnInt(k)
}

func returnCharOnInt(index int) string {
	switch index {
	case 0:
		return "a"
	case 1:
		return "b"
	case 2:
		return "c"
	case 3:
		return "d"
	default:
		return "0"
	}
}

func GenerateMapCharAlter(index int) []string {
	k1 := referenceAlter1[index]
	k2 := referenceAlter2[index]
	k3 := referenceAlter3[index]

	return []string{returnCharOnInt(k1), returnCharOnInt(k2), returnCharOnInt(k3)}
}
