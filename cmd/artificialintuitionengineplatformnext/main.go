// cmd/artificialintuitionengineplatformnext/main.go
package main

import (
"flag"
"log"
"os"

"artificialintuitionengineplatformnext/internal/artificialintuitionengineplatformnext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialintuitionengineplatformnext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
