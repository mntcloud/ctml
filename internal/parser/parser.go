package parser

import (
	"errors"
	"io"
	"strings"

	base "github.com/mntcloud/ctml/internal"
	"github.com/mntcloud/ctml/internal/lexer"
)

const (
	DEFAULT base.State = iota

	ELEMENT         // make focus that we're collecting information about an element and not about lines
	ELEMENT_ATTR    // in case of we have an attribute and we need a value
	ELEMENT_CONTENT // made especially for one line content
)

type Parser struct {
	lex lexer.Lexer

	AST     *Element
	current *Element

	state base.State

	lines    []string
	line     []string
	prevNest int
	nest     int
}

func New(r io.Reader) Parser {
	lex := lexer.New(r)

	return Parser{lex: lex}
}

// Created to be used as a finder for links to other .ctml pages
// on the local filesystem
func (p *Parser) GetLinks() (links []string, err error) {
	if p.AST == nil {
		return nil, errors.New("looks like we have empty AST, do a parsing first")
	}

	// we prefer to be sure, that root of
	// AST is always <body> tag, which doesn't have any links
	// by design and so what we need, it is only to check
	// the element's own children
	for _, ch := range p.AST.Children {
		el, ok := ch.(*Element)

		if !ok {
			continue
		}

		links = append(links, p.getLink(el)...)
	}

	return links, nil
}

func (p *Parser) getLink(node *Element) (links []string) {
	for k, v := range node.Attributes {
		if k == "href" || k == "src" {
			links = append(links, v)
		}
	}

	if len(node.Children) != 0 {
		for _, ch := range node.Children {
			el, ok := ch.(*Element)

			if !ok {
				continue
			}

			links = append(links, p.getLink(el)...)
		}
	}

	return links
}

func (p *Parser) Do() error {
	var tokPart string

	for {
		tok, err := p.lex.Next()

		if err == io.EOF {

			if p.lines != nil {
				c := Content{
					Parent: p.current,
					Value:  p.lines,
				}

				p.current.Children = append(p.current.Children, c)
			}

			break
		}

		if err != nil && err.Error() == "token is probably chunked" {
			tokPart = tok
			continue
		}

		if tokPart != "" {
			tok = tokPart + tok
			tokPart = ""
		}

		switch tok {
		case "<NL>":
			// this state controls an occasion below:
			// <elementName: this is text
			if p.state == ELEMENT_CONTENT {
				p.lines = append(p.lines, strings.Join(p.line, " "))

				c := Content{
					Parent: p.current,
					Value:  p.lines,
				}

				p.current.Children = append(p.current.Children, c)
				p.lines = nil
				p.line = nil

				p.current = p.current.Parent
			}

			if p.line != nil {
				p.lines = append(p.lines, strings.Join(p.line, " "))
				p.line = nil
			}

			p.prevNest = p.nest
			p.nest = 0
			p.state = DEFAULT
		case "<ELEMENT>":
			p.state = ELEMENT

			// there is described one case:
			// - exit out of identation block
			// 		<tag
			// 			content
			// 		<tag: another content
			if p.nest < p.prevNest && p.prevNest != 0 {
				if len(p.lines) > 0 {
					c := Content{
						Parent: p.current,
						Value:  p.lines,
					}

					p.current.Children = append(p.current.Children, c)
					p.lines = nil
				}

				index := 0
				levels := p.prevNest - p.nest

				// sometimes prevNest and nest difference
				// more than one and we need to do operation
				// swaping p.current with p.current.Parent
				// several times
				for {
					if index == levels {
						break
					}

					p.current = p.current.Parent
					index++
				}
			}

			/* allocate an empty element for now, we will populate it later */
			ch := Element{
				// required by go
				Attributes: make(map[string]string),
			}

			if p.AST == nil {
				p.AST = &ch
			} else {
				ch.Parent = p.current
				p.current.Children = append(p.current.Children, &ch)
			}

			p.current = &ch
		case "<DECL>":
			p.state = ELEMENT_CONTENT
		case "<IDENT>":
			p.nest++
		default:
			switch p.state {
			case ELEMENT:
				if p.current.Name == "" {
					p.current.Name = tok
					continue
				}

				if strings.HasPrefix(tok, "=") {
					res, _ := strings.CutPrefix(tok, "=")

					// attributes with true or false value can be simply mentioned without any value
					// this is by the default here, because we don't know what gonna happen next,
					// like the attribute has a value or it's another attribute next to it
					p.current.Attributes[res] = "true"
					p.state = ELEMENT_ATTR
				}

				if strings.HasPrefix(tok, ".") {
					res, _ := strings.CutPrefix(tok, ".")

					p.current.Classes = append(p.current.Classes, res)
				}
			case ELEMENT_ATTR:
				if strings.HasPrefix(tok, "=") {
					res, _ := strings.CutPrefix(tok, "=")

					p.current.Attributes[res] = "true"

					// keep the state
					continue
				}

				attrName := base.GetLastKey(p.current.Attributes)
				p.current.Attributes[attrName] = tok
				p.state = ELEMENT // get back to previous state

			case ELEMENT_CONTENT:
				fallthrough
			default:
				p.line = append(p.line, tok)
			}
		}
	}

	return nil
}
