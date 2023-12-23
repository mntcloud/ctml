package lexer

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	template := `body:
	header:
		button: content
		button: content
	main:
		h1: content
		h2: content 
		p: 
			content
	`

	shouldBe := []string{
		"<ELEMENT>", "body", "<NL>",
		"<IDENT>", "<ELEMENT>", "header", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "button", "<DECL>", "content", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "button", "<DECL>", "content", "<NL>",
		"<IDENT>", "<ELEMENT>", "main", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "h1", "<DECL>", "content", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "h2", "<DECL>", "content", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "p", "<NL>",
		"<IDENT>", "<IDENT>", "<IDENT>", "content", "<NL>",
	}

	if err := procedure(template, shouldBe); err != nil {
		t.Fatal(err.Error())
	}
}

func TestLexerWithAttr(t *testing.T) {
	template := `body:
	header:
		button =active false: content
		button =active true: content
	main:
		h1 .class: content
		h2 .class: content 
		p: 
			content
	`
	shouldBe := []string{
		"<ELEMENT>", "body", "<NL>",
		"<IDENT>", "<ELEMENT>", "header", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "button", "=active", "false", "<DECL>", "content", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "button", "=active", "true", "<DECL>", "content", "<NL>",
		"<IDENT>", "<ELEMENT>", "main", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "h1", ".class", "<DECL>", "content", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "h2", ".class", "<DECL>", "content", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "p", "<NL>",
		"<IDENT>", "<IDENT>", "<IDENT>", "content", "<NL>",
	}

	if err := procedure(template, shouldBe); err != nil {
		t.Fatal(err.Error())
	}
}

func TestLexerWithAttrValueInDoubleQuotes(t *testing.T) {
	template := `<body
	<header
		<button =active "falsy false": content
		<button =active true: content
	`
	shouldBe := []string{
		"<ELEMENT>", "body", "<NL>",
		"<IDENT>", "<ELEMENT>", "header", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "button", "=active", "falsy false", "<DECL>", "content", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "button", "=active", "true", "<DECL>", "content", "<NL>",
	}

	if err := procedure(template, shouldBe); err != nil {
		t.Fatal(err.Error())
	}
}

func TestLexerWithUnicode(t *testing.T) {
	template := `body:
	header:
		button =active false: контент
		button =active true: контент
	main:
		h1 .class: контент
		h2 .class: контент 
		p: 
			контент
	`
	shouldBe := []string{
		"<ELEMENT>", "body", "<NL>",
		"<IDENT>", "<ELEMENT>", "header", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "button", "=active", "false", "<DECL>", "контент", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "button", "=active", "true", "<DECL>", "контент", "<NL>",
		"<IDENT>", "<ELEMENT>", "main", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "h1", ".class", "<DECL>", "контент", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "h2", ".class", "<DECL>", "контент", "<NL>",
		"<IDENT>", "<IDENT>", "<ELEMENT>", "p", "<NL>",
		"<IDENT>", "<IDENT>", "<IDENT>", "контент", "<NL>",
	}

	if err := procedure(template, shouldBe); err != nil {
		t.Fatal(err.Error())
	}
}

func TestLexerWithEscapes(t *testing.T) {
	template := `<body
	\:example
	`
	shouldBe := []string{
		"<ELEMENT>", "body", "<NL>",
		"<IDENT>", ":example",
	}

	if err := procedure(template, shouldBe); err != nil {
		t.Fatalf("got error: %s", err)
	}
}
func TestLexerTwoWordsInRow(t *testing.T) {
	template := `<body
	<someElement =attr "hello, world": lorem ipsum
	`
	shouldBe := []string{
		"<ELEMENT>", "body", "<NL>",
		"<IDENT>", "<ELEMENT>", "someElement", "=attr", "hello, world", "<DECL>", "lorem", "ipsum", "<NL>",
	}

	if err := procedure(template, shouldBe); err != nil {
		t.Fatalf("got error: %s", err)
	}
}

func procedure(doc string, shouldBe []string) error {
	rd := strings.NewReader(doc)
	lex := New(rd)

	var temp string

	res := []string{}
	for {
		tok, err := lex.Next()

		if err == io.EOF {
			break
		}

		if err != nil && err.Error() == "token is probably chunked" {
			temp = tok
			continue
		}

		if temp != "" {
			res = append(res, temp+tok)
			temp = ""
		} else {
			res = append(res, tok)
		}
	}

	for i, v := range shouldBe {
		if v != res[i] {
			return fmt.Errorf("got: %s, should be: %s", res[i], v)
		}
	}

	return nil
}
