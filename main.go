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
	PrependWord string
	SearchWord string
	DisplayWord string
	WordList string
	InflexCats string
	ShortInflectedForm string
	InflectedForm []string
	Explanation string
	Origin string
	OriginSource string
	GrammarComment string
	WordNr string
	Index int
	Phonetic string
	Date string
	Groups []string
	ShortInflection string
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

	url_format := "https://sprotin.fo/dictionary_search_json.php?DictionaryId=1&DictionaryPage=%d&SearchFor=%s&SearchInflections=0&SearchDescriptions=0&Group=&SkipOtherDictionariesResults=0&SkipSimilarWords=0"
	var word string
	var target = "fo:fo"

	flag.Parse()

	if len(flag.Args()) == 2 {
		target = flag.Args()[0]
		word = flag.Args()[1]
	} else if len(flag.Args()) == 1 {
		word = flag.Args()[0]
	} else {
		os.Exit(2)
	}

	dictionaryId, ok := targets[target]
	if !ok {
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
		log.Fatal(err)
	}

	for _, word := range result.Words {
		print_word(word)
	}
}

func print_word(word Word) {
	re := regexp.MustCompile(`<.*?>`)
	explReg := re.ReplaceAllString(word.Explanation, "")
	searchReg := re.ReplaceAllString(word.SearchWord, "")
	fmt.Println(searchReg)
	fmt.Printf("%s\n", Red(explReg))
}