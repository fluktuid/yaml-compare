# Yaml Comparator

Master: ![CircleCI](https://circleci.com/gh/fluktuid/yaml-compare/tree/master.svg?style=svg&circle-token=0c4a4b1a381b3523aefc795a09552596cb1c7eea)\
Develop: ![CircleCI](https://circleci.com/gh/fluktuid/yaml-compare/tree/master.svg?style=svg&circle-token=0c4a4b1a381b3523aefc795a09552596cb1c7eea)

## Installation

yml-compare is available using the standard go get command.

Install by running:\
`go get github.com/fluktuid/yaml-compare`

Run tests by running:\
`There are currently no tests`

### As executable (compile yourself)
``` bash
$ cd $GOPATH/src/yaml-compare/main
$ go build main.go
$ cp main /usr/local/bin/yml
```

## Usage

If you want to compare two files you can use
``` bash
$ yml [flags] [file0] [file1]
```

There is also a 'cat mode' where one file will be printed out:
``` bash
$ yml [flags] [file0]
```
This Mode will resolve the yml anchors (_&anchor_) by default.

### Examples

``` bash
$ yml [flags] [file0]
```

### Flags

| flag |  | Description | Default | Notes |
|---|---|---|---|---|
| -h | --help | print help | - |
| -w | --white | print without ANSI Colors | false |
| -p | --print | print files after anchor resolving | false | true if only one file is specified |
| -c | --print-cpomplete | print the complete diff file after comparing | false |
| -L | --print-line-types | print the Line Types, e.g. 'ListItem' | false |
| -R | --resolve-anchors | Resolve Yml Anchors| true |
| -A | --beware-anchors | Beware anchor names | false | only valid if `-R` is used |
| -P | --beware-pointer | Beware Pointer names | false | only valid if `-R` is used |
| -f | --full-qualifier-name  | false | use full-qualifier names, e.g. 'step[0].instrument' | **alpha**, buggy |

## Used Libraries
| Library | License | Source | Notes |
|---|---|---|---|
| aurora | WTFPL | [Link](https://github.com/logrusorgru/aurora) |   |
| bufio |  |  | go sdk |
| bytes |  |  | go sdk |
| errors |  |  | go sdk |
| fmt |  |  | go sdk |
| io |  |  | go sdk |
| os |  |  | go sdk |
| pflag | BSD-3-Clause | [Link](https://github.com/spf13/pflag/blob/master/LICENSE) |(c) 2012 Alex Ogier (c) 2012 The Go Authors |
| reflect |  |  | go sdk |
| regexp |  |  | go sdk |
| strconv |  |  | go sdk |
| strings |  |  | go sdk |

## Todo
- [ ] Single-Line Lists `[ e0 e1 e2 ... ]`
- [ ] Single-Line Maps `{ k0:v0, k1:v1, k2:v2, ...}`
- [ ] Tests
- [ ] Write Manpage
- [ ] Add to Package System
    - Homebrew
    - Arch User Repository
- [ ] Write Manpage
- [ ] Analyze directories
    - [x] basic
    - [ ] compare by file names
- [ ] More Flags
    - [ ] 'No-Change'\
          Prints changes as _add_ and _remove_
    - [ ] 'out'\
          Print diff to file
    - [ ] 'no-diff-chars'\
          do not print `+`,`-` and `~`
