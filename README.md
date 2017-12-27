# remove_multi

TODO: rename this program as it also does multi replace instead of only multi remove.

## remove

This removes lines from the 2nd file from the 1st file.
It assumes that lines in 2nd file are regexps (_without_ `^` and `$` at the beginning and the end).

Usage:

```
remove_multi remove [source_file] [file_with_lines_to_remove]
```

## replace

This replaces lines from the 2nd file from the 1st file with lines from the 3rd file.
It assumes that lines in 2nd file are regexps (_without_ `^` and `$` at the beginning and the end).

Usage:

```
remove_multi replace [source_file] [file_with_lines_to_remove] [file_with_lines_to_replace]
```
