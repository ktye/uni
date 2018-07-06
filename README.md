# uni
replace tex symbols with unicode characters

Uni translates tex-escaped text to unicode.
The program prints a line of translated text for each input argument
to stdout or translates a stream from stdin.

Line endings and missing final newline from stdin are preserved.

Uni translates only known patterns. Unrecognized patterns such as \n
are printed unchanged.

Examples: 
```
  $ uni "x = \alpha+\beta"
  x = α+β

  # show all patterns in a human readable form
  $ uni -h
	
  # list all patterns in the order of the replacements
  $ uni -l

  # show each rune line by line with it's code point
  uni -d '\alpha+\beta'
  'α' '\u03b1' 0x3b1 01661
  '+' '+' 0x2b 053
  'β' '\u03b2' 0x3b2 01662
```
  
Uni can be used in ktye/editor by selecting text and middle-clicking on `|uni` in the tag bar.
