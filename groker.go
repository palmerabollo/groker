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
    flag.Parse()

    g := grok.New()

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()

        values, _ := g.Parse(*pattern, line)
        delete(values, "")

        encoded, _ := json.Marshal(values)
        fmt.Println(string(encoded))
    }
}
