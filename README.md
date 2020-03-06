# sprotin-cli

### Is a small python program, which can run a fragment of the Sprotin.fo dictionary from Terminal. 

#### The idea is that you type a word, which you want to translate, and the preferred languages you want to translate from and to, into the terminal, and the following result will be displayed in the terminal.
---
Command examples

```
python sprotin.py --fo:en eftirfylgjandi
```
Resuls
```
following (t.d. the following lines); ensuing (t.d. the ensuing battle); succeeding (t.d. the succeeding events)

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
