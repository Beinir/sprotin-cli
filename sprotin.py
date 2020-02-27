import webbrowser
import sys

try:
    searchString = sys.argv[1]
    langPara = sys.argv[2]

except IndexError:
    print("Not Enough Arguments.")
    sys.exit(1)

def search(searchString, langPara):
    if len(sys.argv) > 3:
        print("Too many arguments.")
        sys.exit(1)
    else:
        if langPara == 'en':
            webbrowser.open('https://sprotin.fo/dictionaries?_SearchFor=' + str(searchString) +
                        '&_SearchInflections=0&_SearchDescriptions=0&_DictionaryPage=1&_DictionaryId=3&_Group=')
        elif langPara == 'da':
            webbrowser.open('https://sprotin.fo/dictionaries?_SearchInflections=0&_SearchDescriptions=0&_SearchFor=' + str(
                searchString) + '&_DictionaryPage=1&_DictionaryId=5&_Group=')
        else:
            print("Error: You have typed something wrong." "\n" "Try: en or da as your last parameter.")


search(searchString, langPara)
