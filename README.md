# groker

Simple library to parse grok patterns with go.

This small tool reads from standard input, applies a grok pattern and dumps the captured fields as JSON.

## Examples

```bash
cat apache.log | groker -pattern=%{COMMONAPACHELOG}
```

```bash
echo "Hello 123" | groker -pattern="%{WORD:word} %{WORD:number}"
// output: {"number":"123","word":"Hello"}
```

```bash
echo "2015-06-16T01:47:07.665Z" | ./groker -pattern="%{TIMESTAMP_ISO8601:time}"
{"HOUR":"","ISO8601_TIMEZONE":"Z","MINUTE":"","MONTHDAY":"16","MONTHNUM":"06","SECOND":"07.665","YEAR":"2015","time":"2015-06-16T01:47:07.665Z"}
```

```bash
echo "2015-06-16T01:47:07.665Z" | ./groker -pattern="%{TIMESTAMP_ISO8601:time}" -namedcapture
// output: {"time":"2015-06-16T01:47:07.665Z"}
```

## Installation

It is based on [gemsi/grok](https://github.com/gemsi/grok).

```bash
go get github.com/gemsi/grok
go build
```

## TODO

- Support more output formats.
- Differentiate numbers, booleans and strings.
