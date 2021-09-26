package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/lexer"
	"github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/token"
)

const PROMT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, "%v", PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "exit" {
			return
		}
		l := lexer.New(line)

		for toke := l.NextToken(); toke.Type != token.EOF; toke = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", toke)
		}
	}

}
