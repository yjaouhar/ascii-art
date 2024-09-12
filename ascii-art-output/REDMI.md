<h1 style="; font-size: 45px">Ascii-Art-output</h1>

- [Objectives](#objectives)

- [Instalation](#Instalation)

- [Usage](#Usage)

- [Contributions](#Contributions)
<h2> objectives</h2>

<p>Ascii-art is a program that takes a string as an argument and outputs the string in a graphical representation using ASCII. What we mean by graphical representation using ASCII is writing the received string using ASCII characters. </p>


<h2> instalation</h2>

- <span> clone this project:  </span> [git clone](https://learn.zone01oujda.ma/git/yjaouhar/ascii-art-output).
- <span>Rune the is program withe :</span>

```console
Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard
```
- **Option** : It is the name of the file you want to write to.
- **String** : It is the text that will bo convert to ascii-art
- **Banner** : It is the name of the template. I know some templates may be hard to read, just do not obsess about it.


- **banners** [here](https://learn.zone01oujda.ma/git/root/public/src/branch/master/subjects/ascii-art/).

<h2> ussge</h2>

```console
student$ go run . --output=banner.txt "hello" standard
student$ cat -e banner.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
student$ go run . --output=banner.txt 'Hello There!' shadow
student$ cat -e banner.txt
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
student$
```
- <span>This project should handle an input with numbers, letters, spaces, special characters and \n.</span>

<h2> contributions</h2>

- <span>JAOUHARY YASSINE</span> [Git](https://learn.zone01oujda.ma/git/yjaouhar).
- <span>BAJADY JAMAL</span> [Git](https://learn.zone01oujda.ma/git/jbajady).
