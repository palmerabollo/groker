package main

import (
    "fmt"
    "flag"
    "bufio"
    "os"
    "encoding/json"
    "github.com/gemsi/grok"
)

func main() {
    pattern := flag.String("pattern", "%{GREEDYDATA:msg}", "a grok expression")
    namedcapture := flag.Bool("namedcapture", false, "parse only named captures (default is false)")
    flag.Parse()

    var g *grok.Grok

    if *namedcapture {
      g = grok.New(grok.NAMEDCAPTURE)
    } else {
      g = grok.New();
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
