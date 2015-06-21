# grok

Simple library to parse grok patterns with go.

This small tool reads an input, applies a grok pattern and dumps the captured fields as JSON.

```bash
cat apache.log | grok -pattern=%{COMMONAPACHELOG}
```

```bash
echo "Hello 123" | ./grok -pattern="%{WORD:word} %{WORD:number}"
// output: {"number":"123","word":"Hello"}
```

It is based on [gemsi/grok](https://github.com/gemsi/grok).

## Installation

TODO

## TODO

- Support more output formats
