// cmd/ercnftmetadataindexernetwork/main.go
package main

import (
"flag"
"log"
"os"

"ercnftmetadataindexernetwork/internal/ercnftmetadataindexernetwork"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := ercnftmetadataindexernetwork.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
