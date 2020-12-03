# timeprefix

Timeprefix reads stdin and lines in the following output:

```
time-passed-since-previous-line current-timestamp | contents-of-line-1
time-passed-since-previous-line current-timestamp | contents-of-line-2
```

## Usage

```
$ sh -c 'set -x; sleep 1; echo hey' 2>&1 | timeprefix
846.859µs       2020-12-03T23:05:46.383126224+03:00 | + sleep 1
1.000926902s    2020-12-03T23:05:47.384052996+03:00 | + echo hey
48.232µs        2020-12-03T23:05:47.384101293+03:00 | hey
```

The first two columns show:
* how much time passed since the previous line
* current date time in RFC3339 with nanoseconds

# Installation

```
go get -v github.com/kovetskiy/timeprefix
```

# License

MIT
