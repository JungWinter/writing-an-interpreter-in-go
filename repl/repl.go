package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"go-interpreter/evaluator"
	"go-interpreter/lexer"
	"go-interpreter/object"
	"go-interpreter/parser"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		// TODO: 멀티라인 입력 지원
		_, _ = fmt.Fprint(out, PROMPT)
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if p.Errs.ErrorOrNil() != nil {
			_, _ = fmt.Fprintf(out, "%s\n", strings.TrimSpace(p.Errs.Error()))
			continue
		}

		// TODO: 현재 환경을 디버깅 할 수 있는 구문 추가
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			_, _ = fmt.Fprintf(out, "%s\n", evaluated)
		}
	}
}
