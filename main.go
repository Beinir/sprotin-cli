package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	. "github.com/logrusorgru/aurora"
)

type Word struct {
	PrependWord        string
	SearchWord         string
	DisplayWord        string
	WordList           string
	InflexCats         string
	ShortInflectedForm string
	InflectedForm      []string
	Explanation        string
	Origin             string
	OriginSource       string
	GrammarComment     string
	WordNr             string
	Index              int
	Phonetic           string
	Date               string
	Groups             []string
	ShortInflection    string
}

type Search struct {
	Words []Word `json:"words"`
}

func main() {

	targets := map[string]int{
		"fo:fo": 1,
		"fo:en": 2,
		"en:fo": 3,
		"fo:da": 4,
		"da:fo": 5,
	}

	var url_format string
	var target string
	var word string
	target = "fo:fo"

	flag.Parse()

	if len(flag.Args()) == 2 {
		target = flag.Args()[0]
		word = flag.Args()[1]
	} else if len(flag.Args()) == 1 {
		word = flag.Args()[0]
	} else {
		os.Exit(2)
	}

	// Ohhhh This is a hack!!!!!
	if target == "fo:en" {
		targets["fo:en"] = targets["fo:fo"]
		url_format = "https://sprotin.fo/dictionary_search_json.php?DictionaryId=2&DictionaryPage=%d&SearchFor=%s&SearchInflections=0&SearchDescriptions=0&Group=&SkipOtherDictionariesResults=0&SkipSimilarWords=0"
	} else if target == "fo:fo" {
		url_format = "https://sprotin.fo/dictionary_search_json.php?DictionaryId=1&DictionaryPage=%d&SearchFor=%s&SearchInflections=0&SearchDescriptions=0&Group=&SkipOtherDictionariesResults=0&SkipSimilarWords=0"
	} else if target == "en:fo" {
		targets["en:fo"] = targets["fo:fo"]
		url_format = "https://sprotin.fo/dictionary_search_json.php?DictionaryId=3&DictionaryPage=%d&SearchFor=%s&SearchInflections=0&SearchDescriptions=0&Group=&SkipOtherDictionariesResults=0&SkipSimilarWords=0"
	} else if target == "fo:da" {
		targets["fo:da"] = targets["fo:fo"]
		url_format = "https://sprotin.fo/dictionary_search_json.php?DictionaryId=4&DictionaryPage=%d&SearchFor=%s&SearchInflections=0&SearchDescriptions=0&Group=&SkipOtherDictionariesResults=0&SkipSimilarWords=0"
	} else if target == "da:fo" {
		targets["da:fo"] = targets["fo:fo"]
		url_format = "https://sprotin.fo/dictionary_search_json.php?DictionaryId=5&DictionaryPage=%d&SearchFor=%s&SearchInflections=0&SearchDescriptions=0&Group=&SkipOtherDictionariesResults=0&SkipSimilarWords=0"
	} else {
		fmt.Println(Red("Error in the Language specification Parameter"))
		os.Exit(2)
	}

	dictionaryId, ok := targets[target]
	if !ok {
		fmt.Println(Red("argument error"))
		os.Exit(2)
	}

	url := fmt.Sprintf(url_format, dictionaryId, word)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var result Search
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println(Green("Hey"))
		log.Fatal(err)
	}

	if len(result.Words) < 1 {
		fmt.Println(Green("Einki Úrslit/No result"))
	}

	i := 1

	for _, word := range result.Words {

		print_word(word, i)
		i++
	}
	fmt.Println(Blue("Done"))
}

func print_word(word Word, i int) {
	re := regexp.MustCompile(`<.*?>`)
	if i <= 6 {
		explReg := re.ReplaceAllString(word.Explanation, "")
		searchReg := re.ReplaceAllString(word.SearchWord, "")
		fmt.Println(i, searchReg)
		fmt.Printf("%s\n", Green(explReg))
	}
}
