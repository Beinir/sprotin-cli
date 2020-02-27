import webbrowser
import sys

searchString1 = sys.argv[1]
searchString2 = sys.argv[2]


def search(searchString1, searchString2):
    if searchString2 == 'en':
        webbrowser.open('https://sprotin.fo/dictionaries?_SearchFor=' + str(searchString1) +
                        '&_SearchInflections=0&_SearchDescriptions=0&_DictionaryPage=1&_DictionaryId=3&_Group=')
    if searchString2 == 'da':
        webbrowser.open('https://sprotin.fo/dictionaries?_SearchInflections=0&_SearchDescriptions=0&_SearchFor=' + str(
            searchString1) +
                        '&_DictionaryPage=1&_DictionaryId=5&_Group=')
    else:
        print("Error: You have typed something wrong." "\n" "Try: en or da as your last parameter.")


search(searchString1, searchString2)
