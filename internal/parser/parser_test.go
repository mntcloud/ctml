package parser

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/mntcloud/ctml/internal/lexer"
)

func TestParserGeneral(t *testing.T) {
	template := `<body
	<header
		<button: content
		<button: content
	<main
		<h1: content
		<h2: content 
		<p 
			hello, it's me! 
			you didn't expect me?
	`

	shouldBe := Element{
		Name: "body",
		Children: []Child{
			Element{
				Name: "header",
				Children: []Child{
					Element{
						Name: "button",
						Children: []Child{
							Content{
								Value: []string{"content"},
							},
						},
					},
					Element{
						Name: "button",
						Children: []Child{
							Content{
								Value: []string{"content"},
							},
						},
					},
				},
			},
			Element{
				Name: "main",
				Children: []Child{
					Element{
						Name: "h1",
						Children: []Child{
							Content{
								Value: []string{"content"},
							},
						},
					},
					Element{
						Name: "h2",
						Children: []Child{
							Content{
								Value: []string{"content"},
							},
						},
					},
					Element{
						Name: "p",
						Children: []Child{
							Content{
								Value: []string{
									"hello, it's me!",
									"you didn't expect me?",
								},
							},
						},
					},
				},
			},
		},
	}

	if err := procedure(template, shouldBe); err != nil {
		t.Fatalf("got error: %s", err)
	}
}

func TestParserWithAttr(t *testing.T) {
	template := `<body
	<someElement =attr "hello, world": lorem ipsum
	`

	shouldBe := Element{
		Name: "body",
		Children: []Child{
			Element{
				Name: "someElement",
				Attributes: map[string]string{
					"attr": "hello, world",
				},
				Children: []Child{
					Content{
						Value: []string{"lorem ipsum"},
					},
				},
			},
		},
	}

	if err := procedure(template, shouldBe); err != nil {
		t.Fatalf("got error: %s", err)
	}
}

func procedure(template string, shouldBe Element) error {
	strr := strings.NewReader(template)
	lx := lexer.New(strr)

	pr := New(lx)
	pr.Do()

	if err := elementCheck(shouldBe, *pr.AST, 0); err != nil {
		return err
	}

	return nil
}

func elementCheck(selSb Element, selGot Element, nest int) error {
	if selSb.Name == selGot.Name && len(selSb.Children) == len(selGot.Children) {

		// TODO: add check for attributes and classes

		for i := range selSb.Children {
			switch shouldbe := selSb.Children[i].(type) {
			case Element:
				got, ok := selGot.Children[i].(*Element)

				if !ok {
					return errors.New("should be: Element; got: Content")
				}

				if err := elementCheck(shouldbe, *got, nest+1); err != nil {
					return err
				}
			case Content:
				got, ok := selGot.Children[i].(Content)

				if !ok {
					return errors.New("should be: Content; got: Element")
				}

				if err := contentCheck(shouldbe, got, nest+1); err != nil {
					return err
				}
			default:
				return nil
			}
		}

		return nil
	}

	return fmt.Errorf(
		"something went wrong: shouldbe: %s, len %d; got: %s, len %d;",
		selSb.Name, len(selSb.Children), selGot.Name, len(selGot.Children))
}

func contentCheck(selSb Content, selGot Content, nest int) error {
	if len(selGot.Value) == len(selSb.Value) {
		for i := range selSb.Value {
			if selSb.Value[i] != selGot.Value[i] {
				return fmt.Errorf("got: %s; should be: %s", selGot.Value[i], selSb.Value[i])
			}
		}
	} else {
		return fmt.Errorf(
			"%s, length doesn't match: got: %d; shouldbe: %d",
			selGot.Value, len(selGot.Value), len(selSb.Value))
	}

	return nil
}
