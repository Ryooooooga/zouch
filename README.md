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

### Variables

| name                | type      | e.g.                                  |
|:--------------------|:----------|:--------------------------------------|
| `.Filename`         | `string`  | `today.txt`                           |
| `.Filepath`         | `string`  | full path of `.Filename`              |
| `.TemplateFilename` | `string`  | `~/.config/zouch/templates/today.txt` |
| `.TemplateFilepath` | `string`  | same as `.TemplateFilename`           |


### Functions

| name               | signature                                                          | implementation                                           |
|:-------------------|:-------------------------------------------------------------------|:---------------------------------------------------------|
| `Now`              | `func() time.Time`                                                 | `time.Now`                                               |
| `Base`             | `func(string) string`                                              | `path.Base`                                              |
| `Ext`              | `func(string) string`                                              | `path.Ext`                                               |
| `Dir`              | `func(string) string`                                              | `path.Dir`                                               |
| `Abs`              | `func(string) (string, error)`                                     | `filepath.Abs`                                           |
| `Getwd`            | `func(string) (string, error)`                                     | `os.Getwd`                                               |
| `Getenv`           | `func(string) string`                                              | `os.Getenv`                                              |
| `HasPrefix`        | `func(string) bool`                                                | `strings.HasPrefix`                                      |
| `HasSuffix`        | `func(string) bool`                                                | `strings.HasSuffix`                                      |
| `TrimPrefix`       | `func(string, string) string`                                      | `strings.TrimPrefix`                                     |
| `TrimSuffix`       | `func(string, string) string`                                      | `strings.TrimSuffix`                                     |
| `LowerCamelCase`   | `func(string) string`                                              | `strcase.LowerCamelCase`                                 |
| `UpperCamelCase`   | `func(string) string`                                              | `strcase.UpperCamelCase`                                 |
| `SnakeCase`        | `func(string) string`                                              | `strcase.SnakeCase`                                      |
| `UpperSnakeCase`   | `func(string) string`                                              | `strcase.UpperSnakeCase`                                 |
| `KebabCase`        | `func(string) string`                                              | `strcase.KebabCase`                                      |
| `UpperKebabCase`   | `func(string) string`                                              | `strcase.UpperKebabCase`                                 |
| `Replace`          | `func(string, string, string, int) string`                         | `strings.Replace`                                        |
| `ReplaceAll`       | `func(string, string, string) string`                              | `strings.ReplaceAll`                                     |
| `Shell`            | `func(command string) (string, error)`                             | `exec.Command("/bin/sh", "-c", command).Output()`        |
| `RegexReplaceAll`  | `func(src string, pattern string, replace string) (string, error)` | `regexp.Compile(pattern).ReplaceAllString(src, replace)` |

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

### Using Homebrew

```sh
$ brew install ryooooooga/tap/zouch
```
