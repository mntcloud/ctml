package lexer

import (
	"errors"
	"io"

	base "github.com/mntcloud/ctml/internal"
)

const (
	BASIC base.State = iota
	IN_DOUBLE_QUOTES
)

type Lexer struct {
	reader io.Reader
	index  int
	state  base.State
	buf    []byte
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

	for ; l.index < len(l.buf); l.index++ {
		char := l.buf[l.index]

		if (char == ' ' || char == '\n') && token != nil && l.state == BASIC {
			return string(token), nil
		}

		if char == '"' {
			if l.state == BASIC {
				l.state = IN_DOUBLE_QUOTES
				continue
			} else {
				l.state = BASIC
				l.index++

				return string(token), nil
			}
		}

		if l.state == BASIC {
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
				return "<NL>", nil
			}

			if char != ' ' && char != '\n' && char != '\t' && char != ':' {
				token = append(token, char)
			}
		} else {
			token = append(token, char)
		}
	}

	return string(token), errors.New("token is probably chunked")
}
