package words

import (
	"bufio"
	"errors"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/mileusna/srb"
)

type Word struct {
	Word string
	Freq int
}

type WordGen struct {
	list []Word
}

func NewWordGen(lang string) *WordGen {
	list := generateList(lang)

	log.Println(lang, len(list))

	return &WordGen{
		list,
	}
}

func generateList(lang string) []Word {
	source_file := "assets/" + lang + "words.txt"
	list, err := getWordList(source_file)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	list = filterInvalid(list, lang)

	return list
}

func filterInvalid(wordlist []Word, lang string) []Word {
	newWordlist := []Word{}

	var blacklist []string
	var dict []string
	if lang == "serbian" {
		dict, _ = readDict()
		blacklist, _ = readBlacklist()
	}

	for _, e := range wordlist {
		if lang == "serbian" {
			e.Word = srb.ToCyr(e.Word)
		}

		if !IsLetter(e.Word) {
			continue
		}

		if wordLength(e.Word) != 5 {
			continue
		}

		if lang == "serbian" {
			if existsIn(e.Word, blacklist) || !existsIn(e.Word, dict) {
				continue
			}
		}

		newWordlist = append(newWordlist, e)
	}

	return newWordlist
}

func existsIn(word string, list []string) bool {
	for _, str := range list {
		if str == word {
			return true
		}
	}
	return false
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

func readDict() ([]string, error) {
	file, err := os.Open("assets/serbiandict.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const maxCapacity = 10 * 1024 * 1024

	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	list := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, "/")
		word := tokens[0]

		list = append(list, word)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func readBlacklist() ([]string, error) {
	file, err := os.Open("assets/serbianblacklist.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const maxCapacity = 10 * 1024 * 1024

	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	list := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		list = append(list, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (wg *WordGen) GetN(count int) ([]string, error) {
	size := min(1000, len(wg.list))

	if count < 0 || count > size {
		return nil, errors.New("invalid count")
	}

	options := make([]int, size)

	for i := range options {
		options[i] = i
	}

	rand.Shuffle(size, func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	result := make([]string, count)

	for i := range count {
		result[i] = wg.list[options[i]].Word
	}

	return result, nil
}
