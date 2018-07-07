package main

import (
	"fmt"
	"os"

	parser "cql-grammar/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

/* type TreeShapeListener struct {
	*parser.BaseCQL3Listener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
} */

func initParser(filePath string) *parser.CQL3Parser {
	input, _ := antlr.NewFileStream(filePath)
	lexer := parser.NewCQL3Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCQL3Parser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	return p
}

func main() {
	p := initParser(os.Args[1])
	tree := p.Statements()
	//antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)

	fmt.Println(tree.ToStringTree([]string{}, p))
}
