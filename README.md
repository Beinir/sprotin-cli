# sprotin-cli

### Is a small python program, which can access a fragment of the Sprotin.fo dictionary from Terminal. 

#### Usage
* #### Type the word, which you want to translate, and the preferred languages you want to translate from and to, into the terminal, and the result will be displayed in the terminal.
---
#### Command examples

Ex 1.
```
$ sprotin --fo:fo bilur
```

Result:

```
motorakfar vanl. á fýra hjólum (til fólka- og vøruflutning), bil; sms. t.d. avlamis-, egin-, farma-, ferða-, her-, kassa-, last-, leigu-, lík-, mjólkar-, persón-, post-, rusk-, sjúkra-, sløkki-, tanga-, vøru-
```


Ex 2.
```
$ python sprotin.py --en:fo infant
```
Resuls: 
```
infant:
1 lítið barn; infant (in arms) pinkubarn, vøggubarn, havingarbarn (t.d. the enormous death-rate among infants) 
2 (løgfr.) ómyndingur, ikki komin til lógaldurs (t.d. contracts made by an infant) 
3 l barna- (t.d. years), barnsligur; sum er um at koma fyri seg (t.d. civilization; industry)

infanta:
infantinna (sponsk el. portugisisk prinsessa)

infante:
infant (spanskur el. portugisiskur prinsur)

infanticide:
barnamorð; barnamordari

infanticipate:
(am.) S ganga við barni
```
#### At this moment in time, the program can only translate: 
* Faroese &#8594; Faroese ```sprotin.py --fo:fo bilur```
* English &#8594; Faroese ```sprotin.py --en:fo car```
* Faroese &#8594; English ```sprotin.py --fo:en bilur```
* Danish &#8594; Faroese ```sprotin.py --da:fo bil```
* Faroese &#8594; Danish ```sprotin.py --fo:da bilur```

#### Hint: It would be smart to create an alias in your bash_profile.
* So commands could look something like this:
```
$ sprotin --fo:fo bilur
$ sprotin --en:fo infant
```
