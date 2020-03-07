import json
import re
import sys
import requests
from termcolor import colored

try:
    langPara = sys.argv[1]
    searchString = sys.argv[2]
except IndexError:
    print("Not Enough Arguments.")
    sys.exit(1)

def search(langPara, searchString):
    if len(sys.argv) > 3:
        print("Too many arguments.")
        sys.exit(1)
    else:
        if langPara == '--fo:fo':
            data = requests.get('https://sprotin.fo/dictionary_search_json.php?DictionaryId=1&DictionaryPage=1&SearchFor=' + str(searchString) +
                                '&SearchInflections=0&SearchDescriptions=0&Group=&SkipOtherDictionariesResults=0&SkipSimilarWords=0')
            parse_json(data)
        elif langPara == '--en:fo':
            data = requests.get('https://sprotin.fo/dictionary_search_json.php?DictionaryId=3&DictionaryPage=1&SearchFor=' + str(searchString) +
                                '&SearchInflections=0&SearchDescriptions=0&Group=&SkipOtherDictionariesResults=0&SkipSimilarWords=0')
            parse_json(data)

        elif langPara == '--fo:en':
            data = requests.get('https://sprotin.fo/dictionary_search_json.php?DictionaryId=2&DictionaryPage=1&SearchFor=' + str(searchString) +
                                '&SearchInflections=0&SearchDescriptions=0&Group=&SkipOtherDictionariesResults=0&SkipSimilarWords=0')
            parse_json(data)

        elif langPara == '--da:fo':
            data = requests.get('https://sprotin.fo/dictionary_search_json.php?DictionaryId=5&DictionaryPage=1&SearchFor=' + str(searchString) +
                                '&SearchInflections=0&SearchDescriptions=0&Group=&SkipOtherDictionariesResults=0&SkipSimilarWords=0')
            parse_json(data)
        elif langPara == '--fo:da':
            data = requests.get('https://sprotin.fo/dictionary_search_json.php?DictionaryId=1&DictionaryPage=1&SearchFor=' + str(searchString) +
                                '&SearchInflections=0&SearchDescriptions=0&Group=&SkipOtherDictionariesResults=0&SkipSimilarWords=0')
            parseJson(data)
        else:
            print("Error: You have typed something wrong." "\n" "Try: en or da as your last parameter.")

def parse_json(data):
    res = json.loads(data.text)
    i = 0
    for word in res["words"]:
        if i < 5:
            searchWord = (word["SearchWord"])
            result = (word["Explanation"])
            reg = re.sub('<.*?>', '', result)
            searchReg = re.sub('<.*?>', '', searchWord)
            print('\033[1m' + searchReg + ':')
            print(colored(reg, 'green'))
            print("")
            i = i + 1
        else:
            break
    no_result(data, res)

def no_result(data, res):
    error = res["message"]
    if error != None:
        print(error)
    else: 
        return None
    print("")

search(langPara, searchString)
