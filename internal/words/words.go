package words

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Word struct {
	Word string
	Freq int
}

type WordGen struct {
	list []Word
}

func NewWordGen() *WordGen {
	list := generateList("assets/englishwords.txt")
	return &WordGen{
		list,
	}
}

func generateList(source_file string) []Word {
	list, err := getWordList(source_file)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	list = filterInvalid(list)

	return list
}

func filterInvalid(wordlist []Word) []Word {
	newWordlist := []Word{}

	for _, e := range wordlist {
		if !IsLetter(e.Word) {
			continue
		}

		if wordLength(e.Word) != 5 {
			continue
		}

		newWordlist = append(newWordlist, e)
	}

	return newWordlist
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func wordLength(s string) int {
	return utf8.RuneCountInString(s)
}

func (wl *WordGen) RandomWord() Word {
	return wl.list[rand.Intn(min(len(wl.list), 1000))]
}

func getWordList(filename string) ([]Word, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const maxCapacity = 10 * 1024 * 1024

	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	list := []Word{}

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		word := tokens[0]
		freq, _ := strconv.Atoi(tokens[1])

		list = append(list, Word{
			word,
			freq,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
