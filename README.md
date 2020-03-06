# sprotin-cli

### Is a small python program, which can run a fragment of the Sprotin.fo dictionary from Terminal. 

#### purpose
* #### Type the word, which you want to translate, and the preferred languages you want to translate from and to, into the terminal, and the following result will be displayed in the terminal.
---
#### Command examples

Ex 2.
```
sprotin --fo:fo bilur
```

Result:

```
motorakfar vanl. á fýra hjólum (til fólka- og vøruflutning), bil; sms. t.d. avlamis-, egin-, farma-, ferða-, her-, kassa-, last-, leigu-, lík-, mjólkar-, persón-, post-, rusk-, sjúkra-, sløkki-, tanga-, vøru-
```


Ex 2.
```
python sprotin.py --en:fo infant
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
* Faroese &#8594; Faroese
* English &#8594; Faroese
* Faroese &#8594; English
* Danish &#8594; Faroese
* Faroese &#8594; Danish

#### Hint: It would be smart to create an alias in your bash_profile.
* So a run could look something like this:
```
$ sprotin --fo:en eftirfylgjandi
```
