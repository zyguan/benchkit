// Code generated from Query.g4 by ANTLR 4.10.1. DO NOT EDIT.

package main

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type QueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var querylexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func querylexerLexerInit() {
	staticData := &querylexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'('", "')'", "','", "'<'", "'<='", "'>'", "'>='", "'='", "'!='",
		"'not'", "'and'", "'or'", "'is'", "'in'", "'like'", "'match'", "", "",
		"", "'null'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "LT", "LE", "GT", "GE", "EQ", "NE", "NOT", "AND", "OR",
		"IS", "IN", "LIKE", "MATCH", "STRING", "NUMBER", "BOOLEAN", "NULL",
		"IDENT", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "LT", "LE", "GT", "GE", "EQ", "NE", "NOT", "AND",
		"OR", "IS", "IN", "LIKE", "MATCH", "STRING", "NUMBER", "BOOLEAN", "NULL",
		"IDENT", "WS", "DQUOTA_STRING", "SQUOTA_STRING", "INT", "EXP",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 189, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 3, 16, 105, 8, 16, 1, 17, 3,
		17, 108, 8, 17, 1, 17, 1, 17, 1, 17, 4, 17, 113, 8, 17, 11, 17, 12, 17,
		114, 3, 17, 117, 8, 17, 1, 17, 3, 17, 120, 8, 17, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 131, 8, 18, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 5, 20, 140, 8, 20, 10, 20, 12, 20,
		143, 9, 20, 1, 21, 4, 21, 146, 8, 21, 11, 21, 12, 21, 147, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 156, 8, 22, 10, 22, 12, 22, 159, 9,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 167, 8, 23, 10, 23,
		12, 23, 170, 9, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 5, 24, 177, 8, 24,
		10, 24, 12, 24, 180, 9, 24, 3, 24, 182, 8, 24, 1, 25, 1, 25, 3, 25, 186,
		8, 25, 1, 25, 1, 25, 0, 0, 26, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 0, 47, 0, 49, 0, 51,
		0, 1, 0, 24, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 84,
		84, 116, 116, 2, 0, 65, 65, 97, 97, 2, 0, 68, 68, 100, 100, 2, 0, 82, 82,
		114, 114, 2, 0, 73, 73, 105, 105, 2, 0, 83, 83, 115, 115, 2, 0, 76, 76,
		108, 108, 2, 0, 75, 75, 107, 107, 2, 0, 69, 69, 101, 101, 2, 0, 77, 77,
		109, 109, 2, 0, 67, 67, 99, 99, 2, 0, 72, 72, 104, 104, 1, 0, 48, 57, 2,
		0, 85, 85, 117, 117, 2, 0, 70, 70, 102, 102, 3, 0, 65, 90, 95, 95, 97,
		122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32,
		2, 0, 34, 34, 92, 92, 2, 0, 39, 39, 92, 92, 1, 0, 49, 57, 2, 0, 43, 43,
		45, 45, 199, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 1, 53, 1,
		0, 0, 0, 3, 55, 1, 0, 0, 0, 5, 57, 1, 0, 0, 0, 7, 59, 1, 0, 0, 0, 9, 61,
		1, 0, 0, 0, 11, 64, 1, 0, 0, 0, 13, 66, 1, 0, 0, 0, 15, 69, 1, 0, 0, 0,
		17, 71, 1, 0, 0, 0, 19, 74, 1, 0, 0, 0, 21, 78, 1, 0, 0, 0, 23, 82, 1,
		0, 0, 0, 25, 85, 1, 0, 0, 0, 27, 88, 1, 0, 0, 0, 29, 91, 1, 0, 0, 0, 31,
		96, 1, 0, 0, 0, 33, 104, 1, 0, 0, 0, 35, 107, 1, 0, 0, 0, 37, 130, 1, 0,
		0, 0, 39, 132, 1, 0, 0, 0, 41, 137, 1, 0, 0, 0, 43, 145, 1, 0, 0, 0, 45,
		151, 1, 0, 0, 0, 47, 162, 1, 0, 0, 0, 49, 181, 1, 0, 0, 0, 51, 183, 1,
		0, 0, 0, 53, 54, 5, 40, 0, 0, 54, 2, 1, 0, 0, 0, 55, 56, 5, 41, 0, 0, 56,
		4, 1, 0, 0, 0, 57, 58, 5, 44, 0, 0, 58, 6, 1, 0, 0, 0, 59, 60, 5, 60, 0,
		0, 60, 8, 1, 0, 0, 0, 61, 62, 5, 60, 0, 0, 62, 63, 5, 61, 0, 0, 63, 10,
		1, 0, 0, 0, 64, 65, 5, 62, 0, 0, 65, 12, 1, 0, 0, 0, 66, 67, 5, 62, 0,
		0, 67, 68, 5, 61, 0, 0, 68, 14, 1, 0, 0, 0, 69, 70, 5, 61, 0, 0, 70, 16,
		1, 0, 0, 0, 71, 72, 5, 33, 0, 0, 72, 73, 5, 61, 0, 0, 73, 18, 1, 0, 0,
		0, 74, 75, 7, 0, 0, 0, 75, 76, 7, 1, 0, 0, 76, 77, 7, 2, 0, 0, 77, 20,
		1, 0, 0, 0, 78, 79, 7, 3, 0, 0, 79, 80, 7, 0, 0, 0, 80, 81, 7, 4, 0, 0,
		81, 22, 1, 0, 0, 0, 82, 83, 7, 1, 0, 0, 83, 84, 7, 5, 0, 0, 84, 24, 1,
		0, 0, 0, 85, 86, 7, 6, 0, 0, 86, 87, 7, 7, 0, 0, 87, 26, 1, 0, 0, 0, 88,
		89, 7, 6, 0, 0, 89, 90, 7, 0, 0, 0, 90, 28, 1, 0, 0, 0, 91, 92, 7, 8, 0,
		0, 92, 93, 7, 6, 0, 0, 93, 94, 7, 9, 0, 0, 94, 95, 7, 10, 0, 0, 95, 30,
		1, 0, 0, 0, 96, 97, 7, 11, 0, 0, 97, 98, 7, 3, 0, 0, 98, 99, 7, 2, 0, 0,
		99, 100, 7, 12, 0, 0, 100, 101, 7, 13, 0, 0, 101, 32, 1, 0, 0, 0, 102,
		105, 3, 45, 22, 0, 103, 105, 3, 47, 23, 0, 104, 102, 1, 0, 0, 0, 104, 103,
		1, 0, 0, 0, 105, 34, 1, 0, 0, 0, 106, 108, 5, 45, 0, 0, 107, 106, 1, 0,
		0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 116, 3, 49, 24,
		0, 110, 112, 5, 46, 0, 0, 111, 113, 7, 14, 0, 0, 112, 111, 1, 0, 0, 0,
		113, 114, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115,
		117, 1, 0, 0, 0, 116, 110, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 119,
		1, 0, 0, 0, 118, 120, 3, 51, 25, 0, 119, 118, 1, 0, 0, 0, 119, 120, 1,
		0, 0, 0, 120, 36, 1, 0, 0, 0, 121, 122, 7, 2, 0, 0, 122, 123, 7, 5, 0,
		0, 123, 124, 7, 15, 0, 0, 124, 131, 7, 10, 0, 0, 125, 126, 7, 16, 0, 0,
		126, 127, 7, 3, 0, 0, 127, 128, 7, 8, 0, 0, 128, 129, 7, 7, 0, 0, 129,
		131, 7, 10, 0, 0, 130, 121, 1, 0, 0, 0, 130, 125, 1, 0, 0, 0, 131, 38,
		1, 0, 0, 0, 132, 133, 7, 0, 0, 0, 133, 134, 7, 15, 0, 0, 134, 135, 7, 8,
		0, 0, 135, 136, 7, 8, 0, 0, 136, 40, 1, 0, 0, 0, 137, 141, 7, 17, 0, 0,
		138, 140, 7, 18, 0, 0, 139, 138, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141,
		139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 42, 1, 0, 0, 0, 143, 141, 1,
		0, 0, 0, 144, 146, 7, 19, 0, 0, 145, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0,
		0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149,
		150, 6, 21, 0, 0, 150, 44, 1, 0, 0, 0, 151, 157, 5, 34, 0, 0, 152, 153,
		5, 92, 0, 0, 153, 156, 9, 0, 0, 0, 154, 156, 8, 20, 0, 0, 155, 152, 1,
		0, 0, 0, 155, 154, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0,
		0, 157, 158, 1, 0, 0, 0, 158, 160, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160,
		161, 5, 34, 0, 0, 161, 46, 1, 0, 0, 0, 162, 168, 5, 39, 0, 0, 163, 164,
		5, 92, 0, 0, 164, 167, 9, 0, 0, 0, 165, 167, 8, 21, 0, 0, 166, 163, 1,
		0, 0, 0, 166, 165, 1, 0, 0, 0, 167, 170, 1, 0, 0, 0, 168, 166, 1, 0, 0,
		0, 168, 169, 1, 0, 0, 0, 169, 171, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171,
		172, 5, 39, 0, 0, 172, 48, 1, 0, 0, 0, 173, 182, 5, 48, 0, 0, 174, 178,
		7, 22, 0, 0, 175, 177, 7, 14, 0, 0, 176, 175, 1, 0, 0, 0, 177, 180, 1,
		0, 0, 0, 178, 176, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 182, 1, 0, 0,
		0, 180, 178, 1, 0, 0, 0, 181, 173, 1, 0, 0, 0, 181, 174, 1, 0, 0, 0, 182,
		50, 1, 0, 0, 0, 183, 185, 7, 10, 0, 0, 184, 186, 7, 23, 0, 0, 185, 184,
		1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 188, 3, 49,
		24, 0, 188, 52, 1, 0, 0, 0, 16, 0, 104, 107, 114, 116, 119, 130, 141, 147,
		155, 157, 166, 168, 178, 181, 185, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// QueryLexerInit initializes any static state used to implement QueryLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewQueryLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func QueryLexerInit() {
	staticData := &querylexerLexerStaticData
	staticData.once.Do(querylexerLexerInit)
}

// NewQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewQueryLexer(input antlr.CharStream) *QueryLexer {
	QueryLexerInit()
	l := new(QueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &querylexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Query.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// QueryLexer tokens.
const (
	QueryLexerT__0    = 1
	QueryLexerT__1    = 2
	QueryLexerT__2    = 3
	QueryLexerLT      = 4
	QueryLexerLE      = 5
	QueryLexerGT      = 6
	QueryLexerGE      = 7
	QueryLexerEQ      = 8
	QueryLexerNE      = 9
	QueryLexerNOT     = 10
	QueryLexerAND     = 11
	QueryLexerOR      = 12
	QueryLexerIS      = 13
	QueryLexerIN      = 14
	QueryLexerLIKE    = 15
	QueryLexerMATCH   = 16
	QueryLexerSTRING  = 17
	QueryLexerNUMBER  = 18
	QueryLexerBOOLEAN = 19
	QueryLexerNULL    = 20
	QueryLexerIDENT   = 21
	QueryLexerWS      = 22
)
