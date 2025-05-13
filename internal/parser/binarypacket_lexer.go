// Code generated from BinaryPacket.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type BinaryPacketLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var BinaryPacketLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func binarypacketlexerLexerInit() {
	staticData := &BinaryPacketLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'{'", "'}'", "'MetaData'", "':'", "'`'", "'Int64'", "'uInt16'",
		"'uInt32'", "'string'", "'Int32'", "'uInt8'", "'char'", "'char['", "']'",
		"'match'", "'\"'", "': '", "'ROOT'", "'packet'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "ROOT", "PACKET", "IDENTIFIER", "DIGIT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"ROOT", "PACKET", "IDENTIFIER", "DIGIT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 153, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 5, 19, 140, 8, 19, 10, 19,
		12, 19, 143, 9, 19, 1, 20, 1, 20, 1, 21, 4, 21, 148, 8, 21, 11, 21, 12,
		21, 149, 1, 21, 1, 21, 0, 0, 22, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 1, 0, 4, 3, 0, 65, 90,
		95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 3,
		0, 9, 10, 13, 13, 32, 32, 154, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5,
		1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0,
		0, 0, 1, 45, 1, 0, 0, 0, 3, 47, 1, 0, 0, 0, 5, 49, 1, 0, 0, 0, 7, 58, 1,
		0, 0, 0, 9, 60, 1, 0, 0, 0, 11, 62, 1, 0, 0, 0, 13, 68, 1, 0, 0, 0, 15,
		75, 1, 0, 0, 0, 17, 82, 1, 0, 0, 0, 19, 89, 1, 0, 0, 0, 21, 95, 1, 0, 0,
		0, 23, 101, 1, 0, 0, 0, 25, 106, 1, 0, 0, 0, 27, 112, 1, 0, 0, 0, 29, 114,
		1, 0, 0, 0, 31, 120, 1, 0, 0, 0, 33, 122, 1, 0, 0, 0, 35, 125, 1, 0, 0,
		0, 37, 130, 1, 0, 0, 0, 39, 137, 1, 0, 0, 0, 41, 144, 1, 0, 0, 0, 43, 147,
		1, 0, 0, 0, 45, 46, 5, 123, 0, 0, 46, 2, 1, 0, 0, 0, 47, 48, 5, 125, 0,
		0, 48, 4, 1, 0, 0, 0, 49, 50, 5, 77, 0, 0, 50, 51, 5, 101, 0, 0, 51, 52,
		5, 116, 0, 0, 52, 53, 5, 97, 0, 0, 53, 54, 5, 68, 0, 0, 54, 55, 5, 97,
		0, 0, 55, 56, 5, 116, 0, 0, 56, 57, 5, 97, 0, 0, 57, 6, 1, 0, 0, 0, 58,
		59, 5, 58, 0, 0, 59, 8, 1, 0, 0, 0, 60, 61, 5, 96, 0, 0, 61, 10, 1, 0,
		0, 0, 62, 63, 5, 73, 0, 0, 63, 64, 5, 110, 0, 0, 64, 65, 5, 116, 0, 0,
		65, 66, 5, 54, 0, 0, 66, 67, 5, 52, 0, 0, 67, 12, 1, 0, 0, 0, 68, 69, 5,
		117, 0, 0, 69, 70, 5, 73, 0, 0, 70, 71, 5, 110, 0, 0, 71, 72, 5, 116, 0,
		0, 72, 73, 5, 49, 0, 0, 73, 74, 5, 54, 0, 0, 74, 14, 1, 0, 0, 0, 75, 76,
		5, 117, 0, 0, 76, 77, 5, 73, 0, 0, 77, 78, 5, 110, 0, 0, 78, 79, 5, 116,
		0, 0, 79, 80, 5, 51, 0, 0, 80, 81, 5, 50, 0, 0, 81, 16, 1, 0, 0, 0, 82,
		83, 5, 115, 0, 0, 83, 84, 5, 116, 0, 0, 84, 85, 5, 114, 0, 0, 85, 86, 5,
		105, 0, 0, 86, 87, 5, 110, 0, 0, 87, 88, 5, 103, 0, 0, 88, 18, 1, 0, 0,
		0, 89, 90, 5, 73, 0, 0, 90, 91, 5, 110, 0, 0, 91, 92, 5, 116, 0, 0, 92,
		93, 5, 51, 0, 0, 93, 94, 5, 50, 0, 0, 94, 20, 1, 0, 0, 0, 95, 96, 5, 117,
		0, 0, 96, 97, 5, 73, 0, 0, 97, 98, 5, 110, 0, 0, 98, 99, 5, 116, 0, 0,
		99, 100, 5, 56, 0, 0, 100, 22, 1, 0, 0, 0, 101, 102, 5, 99, 0, 0, 102,
		103, 5, 104, 0, 0, 103, 104, 5, 97, 0, 0, 104, 105, 5, 114, 0, 0, 105,
		24, 1, 0, 0, 0, 106, 107, 5, 99, 0, 0, 107, 108, 5, 104, 0, 0, 108, 109,
		5, 97, 0, 0, 109, 110, 5, 114, 0, 0, 110, 111, 5, 91, 0, 0, 111, 26, 1,
		0, 0, 0, 112, 113, 5, 93, 0, 0, 113, 28, 1, 0, 0, 0, 114, 115, 5, 109,
		0, 0, 115, 116, 5, 97, 0, 0, 116, 117, 5, 116, 0, 0, 117, 118, 5, 99, 0,
		0, 118, 119, 5, 104, 0, 0, 119, 30, 1, 0, 0, 0, 120, 121, 5, 34, 0, 0,
		121, 32, 1, 0, 0, 0, 122, 123, 5, 58, 0, 0, 123, 124, 5, 32, 0, 0, 124,
		34, 1, 0, 0, 0, 125, 126, 5, 82, 0, 0, 126, 127, 5, 79, 0, 0, 127, 128,
		5, 79, 0, 0, 128, 129, 5, 84, 0, 0, 129, 36, 1, 0, 0, 0, 130, 131, 5, 112,
		0, 0, 131, 132, 5, 97, 0, 0, 132, 133, 5, 99, 0, 0, 133, 134, 5, 107, 0,
		0, 134, 135, 5, 101, 0, 0, 135, 136, 5, 116, 0, 0, 136, 38, 1, 0, 0, 0,
		137, 141, 7, 0, 0, 0, 138, 140, 7, 1, 0, 0, 139, 138, 1, 0, 0, 0, 140,
		143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 40, 1,
		0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 145, 7, 2, 0, 0, 145, 42, 1, 0, 0,
		0, 146, 148, 7, 3, 0, 0, 147, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149,
		147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152,
		6, 21, 0, 0, 152, 44, 1, 0, 0, 0, 3, 0, 141, 149, 1, 6, 0, 0,
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

// BinaryPacketLexerInit initializes any static state used to implement BinaryPacketLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBinaryPacketLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BinaryPacketLexerInit() {
	staticData := &BinaryPacketLexerLexerStaticData
	staticData.once.Do(binarypacketlexerLexerInit)
}

// NewBinaryPacketLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBinaryPacketLexer(input antlr.CharStream) *BinaryPacketLexer {
	BinaryPacketLexerInit()
	l := new(BinaryPacketLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &BinaryPacketLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "BinaryPacket.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BinaryPacketLexer tokens.
const (
	BinaryPacketLexerT__0       = 1
	BinaryPacketLexerT__1       = 2
	BinaryPacketLexerT__2       = 3
	BinaryPacketLexerT__3       = 4
	BinaryPacketLexerT__4       = 5
	BinaryPacketLexerT__5       = 6
	BinaryPacketLexerT__6       = 7
	BinaryPacketLexerT__7       = 8
	BinaryPacketLexerT__8       = 9
	BinaryPacketLexerT__9       = 10
	BinaryPacketLexerT__10      = 11
	BinaryPacketLexerT__11      = 12
	BinaryPacketLexerT__12      = 13
	BinaryPacketLexerT__13      = 14
	BinaryPacketLexerT__14      = 15
	BinaryPacketLexerT__15      = 16
	BinaryPacketLexerT__16      = 17
	BinaryPacketLexerROOT       = 18
	BinaryPacketLexerPACKET     = 19
	BinaryPacketLexerIDENTIFIER = 20
	BinaryPacketLexerDIGIT      = 21
	BinaryPacketLexerWS         = 22
)
