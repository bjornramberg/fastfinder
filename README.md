# FastFind 🔍
Just a a lightweight, blazing-fast Linux CLI tool written in Go for searching files and directories. 
It filters out the noise (like .git folders) and highlights directory matches in bold for better readability.

## Features
- Recursive Search: Deep-dives into subdirectories automatically.
- Skipping: Automatically ignores .git directories to save time and clutter.
- Visual Distinction: Clearly labels matches as [FILE] or [DIR], with directories highlighted in bold.
- Case-Insensitive: No need to worry about exact casing while searching.

## Installation
To install fastfind locally, ensure you have Go installed, then run:

```
go build -o fastfind main.go
```
 Move it to your path to use it anywhere (optional)
```
mv fastfind /usr/local/bin/
```
## Usage
The syntax is straightforward:


```
fastfind <search_term> [path]
```

Example 
```
fastfind index ./src/project-alpha
```

## Sample Output
```
Plaintext
[DIR]  ./src/config
[FILE] ./src/utils/config_helper.go
[FILE] ./docs/configuration.md
```
## Requirements
Go 1.16 or higher (as its uses os.DirEntry from path/filepath)
