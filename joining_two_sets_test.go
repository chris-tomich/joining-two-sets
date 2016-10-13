package joining_two_sets

import (
	"testing"
	"math/rand"
	"time"
	"strconv"
)

func MakeWord(maxSize int) string {
	var letters = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	wordSize := r.Intn(maxSize);

	isAlpha := r.Intn(10) >= 3

	var word string

	for i := 0; i < wordSize; i++ {
		if isAlpha {
			word = word + letters[r.Intn(26)]
		} else {
			word = word + strconv.Itoa(r.Intn(10))
		}
	}

	return word
}

func GetTwoDiffSizedSets(size int) ([]string, []string) {
	bagOfWords := make([]string, 0, size)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		bagOfWords = append(bagOfWords, MakeWord(10))
	}

	wordSubSetLen := (len(bagOfWords) / 10) * 2

	wordSubSet := make([]string, wordSubSetLen)

	for i := 0; i < wordSubSetLen; i++ {
		if (i % 2) > 0 {
			wordSubSet[i] = bagOfWords[r.Intn(len(bagOfWords))]
		} else {
			wordSubSet[i] = MakeWord(10)
		}
	}

	return bagOfWords, wordSubSet
}

func GetTwoMatchingSizeSets(size int) ([]string, []string) {
	bagOfWords := make([]string, 0, size)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		bagOfWords = append(bagOfWords, MakeWord(10))
	}

	wordSubSetLen := len(bagOfWords)

	wordSubSet := make([]string, wordSubSetLen)

	for i := 0; i < wordSubSetLen; i++ {
		if (i % 5) == 0 {
			wordSubSet[i] = bagOfWords[r.Intn(len(bagOfWords))]
		} else {
			wordSubSet[i] = MakeWord(10)
		}
	}

	return bagOfWords, wordSubSet
}

func BenchmarkJoinTwoSets(b *testing.B) {
	largeSet, smallSet := GetTwoMatchingSizeSets(1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		largeSetMap := make(map[string]int)

		for i, word := range largeSet {
			largeSetMap[word] = i
		}

		newSet := make([]string, 0, len(smallSet))

		for _, word := range smallSet {
			if _, ok := largeSetMap[word]; ok {
				newSet = append(newSet, word)
			}
		}
	}
}
