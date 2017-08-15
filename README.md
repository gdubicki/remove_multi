# remove_multi

This program takes paths of 2 files and removes all lines from the 2nd file from the 1st file.
It assumes that lines in 2nd file are regexps (_without_ `^` and `$` at the beginning and the end).
Only all lines in order are removed from the first file.

Usage:

```
remove_multi [source_file] [file_with_lines_to_remove]
```
