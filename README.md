# uni
replace tex symbols with unicode characters

Program uni translates tex-escaped text to unicode
The program prints a line of translated text for each input argument
to stdout.

If no arguments are given, it reads lines from stdin.
Line endings and missing final newline from stdin are preserved.
Uni translates only known patterns. Unrecognized patterns such as \n
are printed as is.

Examples: 
```
  $ uni "x = \alpha+\beta"
  x = α+β

  # show all patterns in a human readable form
  $ uni -h
	
  # list all patterns in the order of the replacements
  $ uni -l
```
  
Uni can be used in ktye/editor by selecting text and middle-clicking on `|uni` in the tag bar.
