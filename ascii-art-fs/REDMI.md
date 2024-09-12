## <span style="; font-size: 45px">Ascii-Art-Fs</span>
- [Objectives](#objectives)

- [Instalation](#Instalation)

- [Usage](#Usage)

- [Contributions](#Contributions)

<span style=" font-size: 30px">   Objectives </span>

Ascii-art is a program that takes a string as an argument and outputs the string in a graphical representation using ASCII. What we mean by graphical representation using ASCII is writing the received string using ASCII characters.

<span style=" font-size: 30px">  Instalation </span>

-  <span style="color:red; font-size: 15px">Clone this project</span>    [clone](https://learn.zone01oujda.ma/git/yjaouhar/ascii-art-fs.git/).
-  <span style="color:red; font-size: 15px"> Rune the is program withe : </span>


```console
Usage: go run . [STRING] [BANNER]

EX: go run . something standard
```

- **String** : It is the text that will bo convert to ascii-art
- **Banner** : It is the name of the template. I know some templates may be hard to read, just do not obsess about it.


- **banners** [here](https://learn.zone01oujda.ma/git/root/public/src/branch/master/subjects/ascii-art/).

<span style="font-size: 30px">  Usage </span>

```console
ascii-art-fs$ go run . "hello" standard | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
ascii-art-fs$ go run . "Hello There!" shadow | cat -e
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
ascii-art-fs$ go run . "Hello There!" thinkertoy | cat -e
                                                $
o  o     o o           o-O-o o                o $
|  |     | |             |   |                | $
O--O o-o | | o-o         |   O--o o-o o-o o-o o $
|  | |-' | | | |         |   |  | |-' |   |-'   $
o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                $
                                                $
```
- This project should handle an input with numbers, letters, spaces, special characters and \n.

<span style=" font-size: 30px">   Contributions</span>

- JAOUHARY YASSINE  [Git](https://learn.zone01oujda.ma/git/yjaouhar)
- BAJADY JAMAL   [Git](https://learn.zone01oujda.ma/git/jbajady)

