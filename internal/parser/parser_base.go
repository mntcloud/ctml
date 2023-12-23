package parser

import (
	"fmt"
	"strings"
)

type Element struct {
	Name       string
	Parent     *Element
	Attributes map[string]string
	Classes    []string
	Children   []Child
}

func (el Element) print() string {
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
			tree = tree + v.print() + " "
		}

		return fmt.Sprintf("<%s %s>%s</%s>", el.Name, strings.Join(attr, " "), tree, el.Name)
	}

	return fmt.Sprintf("<%s>", el.Name)
}

type Content struct {
	Parent *Element
	Value  []string
}

func (c Content) print() string {
	return strings.Join(c.Value, "\n")
}
