package lexer

import (
	"errors"
	"io"

	base "github.com/mntcloud/ctml/internal"
)

const (
	DEFAULT base.State = iota
	IN_DOUBLE_QUOTES
	NEWLINE
)

type IdentType int

const (
	NOTHING IdentType = iota
	SPACES
	TABS
)

type Lexer struct {
	reader io.Reader
	index  int
	state  base.State
	buf    []byte

	spaceCount int
	// identType  IdentType
}

func New(r io.Reader) Lexer {
	return Lexer{reader: r}
}

func (l *Lexer) Next() (string, error) {
	if len(l.buf) == l.index || l.buf == nil {
		l.buf = make([]byte, 128)
		l.index = 0

		n, err := l.reader.Read(l.buf)

		if err == io.EOF {
			return "", err
		}

		l.buf = l.buf[:n]
	}

	var token []byte

	// TODO: refactoring
	for ; l.index < len(l.buf); l.index++ {
		char := l.buf[l.index]

		if (char == ' ' || char == '\n' || char == ':') && token != nil && l.state == DEFAULT {
			return string(token), nil
		}

		if char == '"' {
			if l.state == DEFAULT {
				l.state = IN_DOUBLE_QUOTES
				continue
			} else {
				l.state = DEFAULT
				l.index++

				return string(token), nil
			}
		}

		// TODO: either tabs or spaces for identation
		if l.state == NEWLINE {
			if l.spaceCount == 4 {
				l.spaceCount = 0
				return "<IDENT>", nil
			}

			if char == ' ' {
				l.spaceCount++
			}

			if char == '\t' {
				l.spaceCount = 0
				l.index++
				return "<IDENT>", nil
			}

			if char != '\t' && char != ' ' {
				l.state = DEFAULT
			}
		}

		if l.state == DEFAULT {
			// escape a character
			if char == '\\' {
				l.index++
				token = append(token, l.buf[l.index])
				continue
			}

			if char == '<' {
				l.index++
				return "<ELEMENT>", nil
			}

			if char == ':' {
				l.index++
				return "<DECL>", nil
			}

			if char == '\t' {
				l.index++
				return "<IDENT>", nil
			}

			if char == '\n' {
				l.index++
				l.state = NEWLINE
				return "<NL>", nil
			}

			if char != ' ' && char != '\n' && char != '\t' && char != ':' {
				token = append(token, char)
			}
		}

		if l.state == IN_DOUBLE_QUOTES {
			token = append(token, char)
		}

	}

	return string(token), errors.New("token is probably chunked")
}
