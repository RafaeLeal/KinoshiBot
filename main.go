package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "bytes"
    "math/rand"
    "github.com/rogpeppe/go-charset/charset"
    _ "github.com/rogpeppe/go-charset/data"
    "time"
)
func toISO88591(utf8 string) string {
    buf := new(bytes.Buffer)

    w, err := charset.NewWriter("latin1", buf)
    if err != nil {
        panic(err)
    }
    defer w.Close()

    fmt.Fprintf(w, utf8)
    return buf.String()
}

func kinoshitaSays(phrase string) {
    fmt.Printf("%s  - KinsohiBot\n", toISO88591(phrase))
}

func kinoshitaAsksBool(ynquestion string, s string, n string) {
    var resp rune
    fmt.Printf("%s (y/n) - KinoshiBot > ", toISO88591(ynquestion))
    fmt.Scanf("%c", &resp)
    if(resp == 'y') {
        kinoshitaSays(s)
    } else {
        kinoshitaSays(n)
    }
}

func kinoshitaSaysOneOf(quotes []string) {
    rand.Seed(time.Now().Unix())
    kinoshitaSays(quotes[rand.Intn(len(quotes))])
}
func listen(scanner *bufio.Scanner) bool {
    fmt.Print(" - ")
    return scanner.Scan()
}
func main() {
    var input string
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("KinoshiBot started!")
    kinoshitaSays(" - E aí galera?")

    for listen(scanner) {
        fmt.Print(" - ")
        input = scanner.Text()
        isQuestion := strings.Contains(input, "?")
        if (isQuestion) {
            kinoshitaAsksBool(
                "Você viu no site?",
                "Ta muito errado. -1 ponto!",
                "Estão vai lá ler o site e não me incomode.")
        } else {
            quotes := []string {
                "Vocês estão entendendo?",
                "Você ta estudando?",
                "Eu falo para vocês copiarem, mas vocês não copiam..",
                "Vocês vão mal na prova" }
            kinoshitaSaysOneOf(quotes)
        }
    }
    if err := scanner.Err(); err != nil {
        os.Exit(1)
    }

}