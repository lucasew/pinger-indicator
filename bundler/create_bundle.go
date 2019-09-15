package main

import (
    "os"
    "io/ioutil"
    "fmt"
)
func main() {
    f, err := os.Create("bundle.go")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    fmt.Fprintf(f, `
// gerado por create_bundle.go
package main

var (
    icoGoodNet = %s
    icoBadNet = %s
    icoNoNet = %s
)
    `,
        imgRepresentation("./icones/goodnet.png"),
        imgRepresentation("./icones/badnet.png"),
        imgRepresentation("./icones/nointernet.png"),
    )
}

func imgRepresentation(img string) string {
    return fmt.Sprintf("%#v", loadImg(img))
}

func loadImg(img string) []byte {
    fmt.Printf("Carregando imagem %s...\n", img)
    f, err := ioutil.ReadFile(img)
    if err != nil {
        panic(err)
    }
    return f
}
