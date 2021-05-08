package main

import (
	"strings"
)

type Header struct {
	cookie  []string
	content *strings.Reader
}
