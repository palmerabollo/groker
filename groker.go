package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/gemsi/grok"
)

func main() {
	pattern := flag.String("pattern", "%{GREEDYDATA:msg}", "a grok expression")
	namedcapture := flag.Bool("namedcapture", false, "parse only named captures (default is false)")
	flag.Parse()

	var g *grok.Grok

	if *namedcapture {
		g = grok.NewWithConfig(&grok.Config{NamedCapturesOnly: true})
	} else {
		g = grok.New()
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		values, _ := g.Parse(*pattern, line)
		delete(values, "")

		encoded, _ := json.Marshal(values)
		fmt.Println(string(encoded))
	}
}
