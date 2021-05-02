# zouch

Create a new file from a template.

Inspired by [touch-alt](https://github.com/akameco/touch-alt) and [touch_erb](https://github.com/himanoa/touch_erb).

## Usage

```sh
NAME:
   zouch - Create a new file from a template

USAGE:
   zouch [files...]
   zouch --list
   zouch --preview [files...]
   zouch --add     [files...]

GLOBAL OPTIONS:
   --list, -l     list template files (default: false)
   --preview, -p  show template preview (default: false)
   --add, -A      add [files...] as new templates (default: false)
   -r             create directories as required (default: false)
   --force, -f    force update existing files (default: false)
   --verbose, -V  display verbose output (default: false)
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Examples

```sh
$ cat ~/.config/zouch/templates/today.txt
Today is {{ Now.Format "2006-01-02" }}!

$ zouch today.txt

$ cat today.txt
Today is 2021-05-02!
```

## Installation

### From source

```sh
$ go get -u github.com/Ryooooooga/zouch
```

### From precompiled binary

https://github.com/Ryooooooga/zouch/releases/

### Using [zinit](https://github.com/zdharma/zinit)

Add the following to your `.zshrc`.

```sh
zinit ice lucid wait"0" as"program" from"gh-r" \
    pick"zouch*/zouch"
zinit light 'Ryooooooga/zouch'
```
