# Line modifier

Unix command program for modifying lines from various inputs e.g. STDIN, files

## Prerequisites

-   [Go 1.22+](https://go.dev/doc/install)

## Install

```bash
$ go install github.com/mluksic/lm@latest
```

## Usage

Prefix the lines in `example_file` with `"600"`

```bash
$ lm -f example_file -p 600
```

Pipe to STDIN

```bash
$ cat example_line | lm -p 600 -s ,
```

### Parameters
- -f Modify lines in the target file
- -p Add prefix to each line
- -s Add suffix to each line

## Test

```bash
$ go test -v ./..
```

## Authors

👤 **Miha Luksic**
