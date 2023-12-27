package parser

import (
	"fmt"
	"strings"
)

type Child interface {
	Print(ident int) string
}

type Element struct {
	Name       string
	Parent     *Element
	Attributes map[string]string
	Classes    []string
	Children   []Child
}

func (el Element) Print(ident int) string {
	attr := []string{}

	if len(attr) > 0 {
		for k, v := range el.Attributes {
			attr = append(attr, fmt.Sprintf("%s=\"%s\"", k, v))
		}
	}

	if len(el.Classes) > 0 {
		attr = append(attr, fmt.Sprintf("class=\"%s\"", strings.Join(el.Classes, " ")))
	}

	if len(el.Children) > 0 {
		tree := ""

		for _, v := range el.Children {
			tree = tree + v.Print(ident+1) + " "
		}

		return fmt.Sprintf("<%s %s>%s</%s>", el.Name, strings.Join(attr, " "), tree, el.Name)
	}

	return fmt.Sprintf("<%s>", el.Name)
}

type Content struct {
	Parent *Element
	Value  []string
}

func (c Content) Print(ident int) string {
	return strings.Join(c.Value, "\n")
}
