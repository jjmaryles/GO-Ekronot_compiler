package parser

import (
	p "GO-Ekronot_compiler/common/parser"
	"strconv"
	"strings"
)

// Line types
const (
	A_Type = "a-instruction"
	C_Type = "c-instruction"
	Label  = "label"
)

// New returns new Parser
func New(sourceFile string) *p.Parser {
	return p.New(sourceFile)
}

// CommandType returns type of line constant
func CommandType(c string) string {
	switch {
	case strings.HasPrefix(c, "(") && strings.HasSuffix(c, ")"):
		return Label // label lines
	case strings.HasPrefix(c, "@"):
		return A_Type // a-instruction
	default:
		return C_Type // c-instruction
	}
}

// CommandArgs returns command args
// (dest, comp, jump) in case of c-instruction
func CommandArgs(s string) (a string, b string, c string) {
	a, b, c = "", "", ""
	switch CommandType(s) {
	case Label:
		a = s[1 : len(s)-1]
	case A_Type:
		a = s[1:]
	case C_Type:
		compInd := strings.Index(s, "=")
		jumpInd := strings.Index(s, ";")
		if jumpInd != -1 {
			c = s[jumpInd+1:]
		} else {
			jumpInd = len(s)
		}
		if compInd == -1 {
			b = s[:jumpInd]
		} else {
			a = s[:compInd]
			b = s[compInd+1 : jumpInd]
		}
	}
	return
}

// returns true if variable
func IsVar(s string) bool {
	_, err := strconv.ParseInt(s, 10, 16)
	return err != nil
}
