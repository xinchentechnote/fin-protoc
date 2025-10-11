// Code generated from grammar/PacketDsl.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // PacketDsl
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PacketDslParser struct {
	*antlr.BaseParser
}

var PacketDslParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func packetdslParserInit() {
	staticData := &PacketDslParserStaticData
	staticData.LiteralNames = []string{
		"", "'options'", "'{'", "'}'", "'='", "'@calculatedFrom('", "')'", "'@lengthOf('",
		"'('", "'string'", "'char'", "'char['", "']'", "'as'", "'['", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "'root'", "'packet'",
		"'repeat'", "'MetaData'", "'match'", "':'", "','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "UINT8",
		"UINT16", "UINT32", "UINT64", "INT8", "INT16", "INT32", "INT64", "FLOAT32",
		"FLOAT64", "DIGITS", "STRING", "PADDING_ATTR", "PADDING_CHAR", "ROOT",
		"PACKET", "REPEAT", "METADATA", "MATCH", "COLON", "COMMA", "SEMICOLON",
		"IDENTIFIER", "STRING_LITERAL", "LINE_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"packet", "optionDefinition", "optionDeclaration", "packetDefinition",
		"fieldDefinitionWithAttribute", "fieldDefinition", "metaDataDefinition",
		"lengthFieldDeclaration", "checkSumFieldDeclaration", "fieldAttribute",
		"calculatedFromAttribute", "lengthOfAttribute", "paddingAttribute",
		"metaDataDeclaration", "value", "type", "matchFieldDeclaration", "matchPair",
		"inerObjectDeclaration", "list",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 40, 234, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 0, 1,
		0, 5, 0, 44, 8, 0, 10, 0, 12, 0, 47, 9, 0, 1, 1, 1, 1, 1, 1, 5, 1, 52,
		8, 1, 10, 1, 12, 1, 55, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		63, 8, 2, 1, 3, 3, 3, 66, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 72, 8, 3,
		10, 3, 12, 3, 75, 9, 3, 1, 3, 1, 3, 1, 4, 5, 4, 80, 8, 4, 10, 4, 12, 4,
		83, 9, 4, 1, 4, 1, 4, 1, 5, 3, 5, 88, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 94, 8, 5, 1, 5, 1, 5, 3, 5, 98, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 107, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 113, 8, 6, 10,
		6, 12, 6, 116, 9, 6, 1, 6, 1, 6, 1, 7, 3, 7, 121, 8, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 126, 8, 7, 1, 7, 1, 7, 1, 8, 3, 8, 131, 8, 8, 1, 8, 1, 8, 1, 8,
		3, 8, 136, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 143, 8, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 3, 13, 159, 8, 13, 1, 13, 1, 13, 3, 13, 163, 8, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 3, 14, 170, 8, 14, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 3, 15, 187, 8, 15, 1, 15, 3, 15, 190, 8, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 4, 16, 198, 8, 16, 11, 16, 12, 16, 199, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 17, 3, 17, 207, 8, 17, 1, 17, 1, 17, 1, 17, 3,
		17, 212, 8, 17, 1, 18, 1, 18, 1, 18, 4, 18, 217, 8, 18, 11, 18, 12, 18,
		218, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 227, 8, 19, 10, 19,
		12, 19, 230, 9, 19, 1, 19, 1, 19, 1, 19, 0, 0, 20, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 0, 1, 1, 0, 25, 26,
		260, 0, 45, 1, 0, 0, 0, 2, 48, 1, 0, 0, 0, 4, 58, 1, 0, 0, 0, 6, 65, 1,
		0, 0, 0, 8, 81, 1, 0, 0, 0, 10, 106, 1, 0, 0, 0, 12, 108, 1, 0, 0, 0, 14,
		120, 1, 0, 0, 0, 16, 130, 1, 0, 0, 0, 18, 142, 1, 0, 0, 0, 20, 144, 1,
		0, 0, 0, 22, 148, 1, 0, 0, 0, 24, 152, 1, 0, 0, 0, 26, 158, 1, 0, 0, 0,
		28, 169, 1, 0, 0, 0, 30, 189, 1, 0, 0, 0, 32, 191, 1, 0, 0, 0, 34, 206,
		1, 0, 0, 0, 36, 213, 1, 0, 0, 0, 38, 222, 1, 0, 0, 0, 40, 44, 3, 6, 3,
		0, 41, 44, 3, 12, 6, 0, 42, 44, 3, 2, 1, 0, 43, 40, 1, 0, 0, 0, 43, 41,
		1, 0, 0, 0, 43, 42, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0,
		45, 46, 1, 0, 0, 0, 46, 1, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 49, 5, 1,
		0, 0, 49, 53, 5, 2, 0, 0, 50, 52, 3, 4, 2, 0, 51, 50, 1, 0, 0, 0, 52, 55,
		1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 56, 1, 0, 0, 0,
		55, 53, 1, 0, 0, 0, 56, 57, 5, 3, 0, 0, 57, 3, 1, 0, 0, 0, 58, 59, 5, 37,
		0, 0, 59, 60, 5, 4, 0, 0, 60, 62, 3, 28, 14, 0, 61, 63, 5, 36, 0, 0, 62,
		61, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 5, 1, 0, 0, 0, 64, 66, 5, 29, 0,
		0, 65, 64, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 68,
		5, 30, 0, 0, 68, 69, 5, 37, 0, 0, 69, 73, 5, 2, 0, 0, 70, 72, 3, 8, 4,
		0, 71, 70, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74,
		1, 0, 0, 0, 74, 76, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 5, 3, 0, 0,
		77, 7, 1, 0, 0, 0, 78, 80, 3, 18, 9, 0, 79, 78, 1, 0, 0, 0, 80, 83, 1,
		0, 0, 0, 81, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 84, 1, 0, 0, 0, 83,
		81, 1, 0, 0, 0, 84, 85, 3, 10, 5, 0, 85, 9, 1, 0, 0, 0, 86, 88, 5, 31,
		0, 0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90,
		3, 36, 18, 0, 90, 91, 5, 35, 0, 0, 91, 107, 1, 0, 0, 0, 92, 94, 5, 31,
		0, 0, 93, 92, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 107,
		3, 26, 13, 0, 96, 98, 5, 31, 0, 0, 97, 96, 1, 0, 0, 0, 97, 98, 1, 0, 0,
		0, 98, 99, 1, 0, 0, 0, 99, 100, 5, 37, 0, 0, 100, 107, 5, 35, 0, 0, 101,
		107, 3, 14, 7, 0, 102, 107, 3, 16, 8, 0, 103, 104, 3, 32, 16, 0, 104, 105,
		5, 35, 0, 0, 105, 107, 1, 0, 0, 0, 106, 87, 1, 0, 0, 0, 106, 93, 1, 0,
		0, 0, 106, 97, 1, 0, 0, 0, 106, 101, 1, 0, 0, 0, 106, 102, 1, 0, 0, 0,
		106, 103, 1, 0, 0, 0, 107, 11, 1, 0, 0, 0, 108, 109, 5, 32, 0, 0, 109,
		110, 5, 37, 0, 0, 110, 114, 5, 2, 0, 0, 111, 113, 3, 26, 13, 0, 112, 111,
		1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0,
		0, 0, 115, 117, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 118, 5, 3, 0, 0,
		118, 13, 1, 0, 0, 0, 119, 121, 3, 30, 15, 0, 120, 119, 1, 0, 0, 0, 120,
		121, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 5, 37, 0, 0, 123, 125,
		3, 22, 11, 0, 124, 126, 5, 38, 0, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1,
		0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 5, 35, 0, 0, 128, 15, 1, 0, 0,
		0, 129, 131, 3, 30, 15, 0, 130, 129, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0,
		131, 132, 1, 0, 0, 0, 132, 133, 5, 37, 0, 0, 133, 135, 3, 20, 10, 0, 134,
		136, 5, 38, 0, 0, 135, 134, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 137,
		1, 0, 0, 0, 137, 138, 5, 35, 0, 0, 138, 17, 1, 0, 0, 0, 139, 143, 3, 22,
		11, 0, 140, 143, 3, 20, 10, 0, 141, 143, 3, 24, 12, 0, 142, 139, 1, 0,
		0, 0, 142, 140, 1, 0, 0, 0, 142, 141, 1, 0, 0, 0, 143, 19, 1, 0, 0, 0,
		144, 145, 5, 5, 0, 0, 145, 146, 5, 26, 0, 0, 146, 147, 5, 6, 0, 0, 147,
		21, 1, 0, 0, 0, 148, 149, 5, 7, 0, 0, 149, 150, 5, 37, 0, 0, 150, 151,
		5, 6, 0, 0, 151, 23, 1, 0, 0, 0, 152, 153, 5, 27, 0, 0, 153, 154, 5, 8,
		0, 0, 154, 155, 5, 28, 0, 0, 155, 156, 5, 6, 0, 0, 156, 25, 1, 0, 0, 0,
		157, 159, 3, 30, 15, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159,
		160, 1, 0, 0, 0, 160, 162, 5, 37, 0, 0, 161, 163, 5, 38, 0, 0, 162, 161,
		1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165, 5, 35,
		0, 0, 165, 27, 1, 0, 0, 0, 166, 170, 3, 30, 15, 0, 167, 170, 5, 26, 0,
		0, 168, 170, 5, 25, 0, 0, 169, 166, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169,
		168, 1, 0, 0, 0, 170, 29, 1, 0, 0, 0, 171, 190, 5, 37, 0, 0, 172, 190,
		5, 15, 0, 0, 173, 190, 5, 16, 0, 0, 174, 190, 5, 17, 0, 0, 175, 190, 5,
		18, 0, 0, 176, 190, 5, 19, 0, 0, 177, 190, 5, 20, 0, 0, 178, 190, 5, 21,
		0, 0, 179, 190, 5, 22, 0, 0, 180, 190, 5, 23, 0, 0, 181, 190, 5, 24, 0,
		0, 182, 190, 5, 9, 0, 0, 183, 190, 5, 10, 0, 0, 184, 186, 5, 11, 0, 0,
		185, 187, 5, 25, 0, 0, 186, 185, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187,
		188, 1, 0, 0, 0, 188, 190, 5, 12, 0, 0, 189, 171, 1, 0, 0, 0, 189, 172,
		1, 0, 0, 0, 189, 173, 1, 0, 0, 0, 189, 174, 1, 0, 0, 0, 189, 175, 1, 0,
		0, 0, 189, 176, 1, 0, 0, 0, 189, 177, 1, 0, 0, 0, 189, 178, 1, 0, 0, 0,
		189, 179, 1, 0, 0, 0, 189, 180, 1, 0, 0, 0, 189, 181, 1, 0, 0, 0, 189,
		182, 1, 0, 0, 0, 189, 183, 1, 0, 0, 0, 189, 184, 1, 0, 0, 0, 190, 31, 1,
		0, 0, 0, 191, 192, 5, 33, 0, 0, 192, 193, 5, 37, 0, 0, 193, 194, 5, 13,
		0, 0, 194, 195, 5, 37, 0, 0, 195, 197, 5, 2, 0, 0, 196, 198, 3, 34, 17,
		0, 197, 196, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199,
		200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 5, 3, 0, 0, 202, 33, 1,
		0, 0, 0, 203, 207, 5, 25, 0, 0, 204, 207, 5, 26, 0, 0, 205, 207, 3, 38,
		19, 0, 206, 203, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 206, 205, 1, 0, 0, 0,
		207, 208, 1, 0, 0, 0, 208, 209, 5, 34, 0, 0, 209, 211, 5, 37, 0, 0, 210,
		212, 5, 35, 0, 0, 211, 210, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 35,
		1, 0, 0, 0, 213, 214, 5, 37, 0, 0, 214, 216, 5, 2, 0, 0, 215, 217, 3, 10,
		5, 0, 216, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0,
		218, 219, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 221, 5, 3, 0, 0, 221,
		37, 1, 0, 0, 0, 222, 223, 5, 14, 0, 0, 223, 228, 7, 0, 0, 0, 224, 225,
		5, 35, 0, 0, 225, 227, 7, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 230, 1, 0,
		0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 231, 1, 0, 0, 0,
		230, 228, 1, 0, 0, 0, 231, 232, 5, 12, 0, 0, 232, 39, 1, 0, 0, 0, 27, 43,
		45, 53, 62, 65, 73, 81, 87, 93, 97, 106, 114, 120, 125, 130, 135, 142,
		158, 162, 169, 186, 189, 199, 206, 211, 218, 228,
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

// PacketDslParserInit initializes any static state used to implement PacketDslParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPacketDslParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PacketDslParserInit() {
	staticData := &PacketDslParserStaticData
	staticData.once.Do(packetdslParserInit)
}

// NewPacketDslParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPacketDslParser(input antlr.TokenStream) *PacketDslParser {
	PacketDslParserInit()
	this := new(PacketDslParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PacketDslParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PacketDsl.g4"

	return this
}

// PacketDslParser tokens.
const (
	PacketDslParserEOF            = antlr.TokenEOF
	PacketDslParserT__0           = 1
	PacketDslParserT__1           = 2
	PacketDslParserT__2           = 3
	PacketDslParserT__3           = 4
	PacketDslParserT__4           = 5
	PacketDslParserT__5           = 6
	PacketDslParserT__6           = 7
	PacketDslParserT__7           = 8
	PacketDslParserT__8           = 9
	PacketDslParserT__9           = 10
	PacketDslParserT__10          = 11
	PacketDslParserT__11          = 12
	PacketDslParserT__12          = 13
	PacketDslParserT__13          = 14
	PacketDslParserUINT8          = 15
	PacketDslParserUINT16         = 16
	PacketDslParserUINT32         = 17
	PacketDslParserUINT64         = 18
	PacketDslParserINT8           = 19
	PacketDslParserINT16          = 20
	PacketDslParserINT32          = 21
	PacketDslParserINT64          = 22
	PacketDslParserFLOAT32        = 23
	PacketDslParserFLOAT64        = 24
	PacketDslParserDIGITS         = 25
	PacketDslParserSTRING         = 26
	PacketDslParserPADDING_ATTR   = 27
	PacketDslParserPADDING_CHAR   = 28
	PacketDslParserROOT           = 29
	PacketDslParserPACKET         = 30
	PacketDslParserREPEAT         = 31
	PacketDslParserMETADATA       = 32
	PacketDslParserMATCH          = 33
	PacketDslParserCOLON          = 34
	PacketDslParserCOMMA          = 35
	PacketDslParserSEMICOLON      = 36
	PacketDslParserIDENTIFIER     = 37
	PacketDslParserSTRING_LITERAL = 38
	PacketDslParserLINE_COMMENT   = 39
	PacketDslParserWS             = 40
)

// PacketDslParser rules.
const (
	PacketDslParserRULE_packet                       = 0
	PacketDslParserRULE_optionDefinition             = 1
	PacketDslParserRULE_optionDeclaration            = 2
	PacketDslParserRULE_packetDefinition             = 3
	PacketDslParserRULE_fieldDefinitionWithAttribute = 4
	PacketDslParserRULE_fieldDefinition              = 5
	PacketDslParserRULE_metaDataDefinition           = 6
	PacketDslParserRULE_lengthFieldDeclaration       = 7
	PacketDslParserRULE_checkSumFieldDeclaration     = 8
	PacketDslParserRULE_fieldAttribute               = 9
	PacketDslParserRULE_calculatedFromAttribute      = 10
	PacketDslParserRULE_lengthOfAttribute            = 11
	PacketDslParserRULE_paddingAttribute             = 12
	PacketDslParserRULE_metaDataDeclaration          = 13
	PacketDslParserRULE_value                        = 14
	PacketDslParserRULE_type                         = 15
	PacketDslParserRULE_matchFieldDeclaration        = 16
	PacketDslParserRULE_matchPair                    = 17
	PacketDslParserRULE_inerObjectDeclaration        = 18
	PacketDslParserRULE_list                         = 19
)

// IPacketContext is an interface to support dynamic dispatch.
type IPacketContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPacketDefinition() []IPacketDefinitionContext
	PacketDefinition(i int) IPacketDefinitionContext
	AllMetaDataDefinition() []IMetaDataDefinitionContext
	MetaDataDefinition(i int) IMetaDataDefinitionContext
	AllOptionDefinition() []IOptionDefinitionContext
	OptionDefinition(i int) IOptionDefinitionContext

	// IsPacketContext differentiates from other interfaces.
	IsPacketContext()
}

type PacketContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPacketContext() *PacketContext {
	var p = new(PacketContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_packet
	return p
}

func InitEmptyPacketContext(p *PacketContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_packet
}

func (*PacketContext) IsPacketContext() {}

func NewPacketContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PacketContext {
	var p = new(PacketContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_packet

	return p
}

func (s *PacketContext) GetParser() antlr.Parser { return s.parser }

func (s *PacketContext) AllPacketDefinition() []IPacketDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPacketDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IPacketDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPacketDefinitionContext); ok {
			tst[i] = t.(IPacketDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *PacketContext) PacketDefinition(i int) IPacketDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPacketDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPacketDefinitionContext)
}

func (s *PacketContext) AllMetaDataDefinition() []IMetaDataDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaDataDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IMetaDataDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaDataDefinitionContext); ok {
			tst[i] = t.(IMetaDataDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *PacketContext) MetaDataDefinition(i int) IMetaDataDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaDataDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaDataDefinitionContext)
}

func (s *PacketContext) AllOptionDefinition() []IOptionDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IOptionDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionDefinitionContext); ok {
			tst[i] = t.(IOptionDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *PacketContext) OptionDefinition(i int) IOptionDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionDefinitionContext)
}

func (s *PacketContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PacketContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PacketContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitPacket(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) Packet() (localctx IPacketContext) {
	localctx = NewPacketContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PacketDslParserRULE_packet)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5905580034) != 0 {
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PacketDslParserROOT, PacketDslParserPACKET:
			{
				p.SetState(40)
				p.PacketDefinition()
			}

		case PacketDslParserMETADATA:
			{
				p.SetState(41)
				p.MetaDataDefinition()
			}

		case PacketDslParserT__0:
			{
				p.SetState(42)
				p.OptionDefinition()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionDefinitionContext is an interface to support dynamic dispatch.
type IOptionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOptionDeclaration() []IOptionDeclarationContext
	OptionDeclaration(i int) IOptionDeclarationContext

	// IsOptionDefinitionContext differentiates from other interfaces.
	IsOptionDefinitionContext()
}

type OptionDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionDefinitionContext() *OptionDefinitionContext {
	var p = new(OptionDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_optionDefinition
	return p
}

func InitEmptyOptionDefinitionContext(p *OptionDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_optionDefinition
}

func (*OptionDefinitionContext) IsOptionDefinitionContext() {}

func NewOptionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionDefinitionContext {
	var p = new(OptionDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_optionDefinition

	return p
}

func (s *OptionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionDefinitionContext) AllOptionDeclaration() []IOptionDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IOptionDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionDeclarationContext); ok {
			tst[i] = t.(IOptionDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *OptionDefinitionContext) OptionDeclaration(i int) IOptionDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionDeclarationContext)
}

func (s *OptionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitOptionDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) OptionDefinition() (localctx IOptionDefinitionContext) {
	localctx = NewOptionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PacketDslParserRULE_optionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(PacketDslParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PacketDslParserIDENTIFIER {
		{
			p.SetState(50)
			p.OptionDeclaration()
		}

		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Match(PacketDslParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionDeclarationContext is an interface to support dynamic dispatch.
type IOptionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Value() IValueContext
	SEMICOLON() antlr.TerminalNode

	// IsOptionDeclarationContext differentiates from other interfaces.
	IsOptionDeclarationContext()
}

type OptionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionDeclarationContext() *OptionDeclarationContext {
	var p = new(OptionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_optionDeclaration
	return p
}

func InitEmptyOptionDeclarationContext(p *OptionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_optionDeclaration
}

func (*OptionDeclarationContext) IsOptionDeclarationContext() {}

func NewOptionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionDeclarationContext {
	var p = new(OptionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_optionDeclaration

	return p
}

func (s *OptionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *OptionDeclarationContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *OptionDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(PacketDslParserSEMICOLON, 0)
}

func (s *OptionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitOptionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) OptionDeclaration() (localctx IOptionDeclarationContext) {
	localctx = NewOptionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PacketDslParserRULE_optionDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(PacketDslParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Value()
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserSEMICOLON {
		{
			p.SetState(61)
			p.Match(PacketDslParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPacketDefinitionContext is an interface to support dynamic dispatch.
type IPacketDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKET() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ROOT() antlr.TerminalNode
	AllFieldDefinitionWithAttribute() []IFieldDefinitionWithAttributeContext
	FieldDefinitionWithAttribute(i int) IFieldDefinitionWithAttributeContext

	// IsPacketDefinitionContext differentiates from other interfaces.
	IsPacketDefinitionContext()
}

type PacketDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPacketDefinitionContext() *PacketDefinitionContext {
	var p = new(PacketDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_packetDefinition
	return p
}

func InitEmptyPacketDefinitionContext(p *PacketDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_packetDefinition
}

func (*PacketDefinitionContext) IsPacketDefinitionContext() {}

func NewPacketDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PacketDefinitionContext {
	var p = new(PacketDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_packetDefinition

	return p
}

func (s *PacketDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *PacketDefinitionContext) PACKET() antlr.TerminalNode {
	return s.GetToken(PacketDslParserPACKET, 0)
}

func (s *PacketDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *PacketDefinitionContext) ROOT() antlr.TerminalNode {
	return s.GetToken(PacketDslParserROOT, 0)
}

func (s *PacketDefinitionContext) AllFieldDefinitionWithAttribute() []IFieldDefinitionWithAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDefinitionWithAttributeContext); ok {
			len++
		}
	}

	tst := make([]IFieldDefinitionWithAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDefinitionWithAttributeContext); ok {
			tst[i] = t.(IFieldDefinitionWithAttributeContext)
			i++
		}
	}

	return tst
}

func (s *PacketDefinitionContext) FieldDefinitionWithAttribute(i int) IFieldDefinitionWithAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDefinitionWithAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDefinitionWithAttributeContext)
}

func (s *PacketDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PacketDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PacketDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitPacketDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) PacketDefinition() (localctx IPacketDefinitionContext) {
	localctx = NewPacketDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PacketDslParserRULE_packetDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserROOT {
		{
			p.SetState(64)
			p.Match(PacketDslParserROOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(67)
		p.Match(PacketDslParserPACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&148344114848) != 0 {
		{
			p.SetState(70)
			p.FieldDefinitionWithAttribute()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(76)
		p.Match(PacketDslParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldDefinitionWithAttributeContext is an interface to support dynamic dispatch.
type IFieldDefinitionWithAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldDefinition() IFieldDefinitionContext
	AllFieldAttribute() []IFieldAttributeContext
	FieldAttribute(i int) IFieldAttributeContext

	// IsFieldDefinitionWithAttributeContext differentiates from other interfaces.
	IsFieldDefinitionWithAttributeContext()
}

type FieldDefinitionWithAttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDefinitionWithAttributeContext() *FieldDefinitionWithAttributeContext {
	var p = new(FieldDefinitionWithAttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_fieldDefinitionWithAttribute
	return p
}

func InitEmptyFieldDefinitionWithAttributeContext(p *FieldDefinitionWithAttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_fieldDefinitionWithAttribute
}

func (*FieldDefinitionWithAttributeContext) IsFieldDefinitionWithAttributeContext() {}

func NewFieldDefinitionWithAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDefinitionWithAttributeContext {
	var p = new(FieldDefinitionWithAttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_fieldDefinitionWithAttribute

	return p
}

func (s *FieldDefinitionWithAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDefinitionWithAttributeContext) FieldDefinition() IFieldDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDefinitionContext)
}

func (s *FieldDefinitionWithAttributeContext) AllFieldAttribute() []IFieldAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldAttributeContext); ok {
			len++
		}
	}

	tst := make([]IFieldAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldAttributeContext); ok {
			tst[i] = t.(IFieldAttributeContext)
			i++
		}
	}

	return tst
}

func (s *FieldDefinitionWithAttributeContext) FieldAttribute(i int) IFieldAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldAttributeContext)
}

func (s *FieldDefinitionWithAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefinitionWithAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDefinitionWithAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitFieldDefinitionWithAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) FieldDefinitionWithAttribute() (localctx IFieldDefinitionWithAttributeContext) {
	localctx = NewFieldDefinitionWithAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PacketDslParserRULE_fieldDefinitionWithAttribute)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&134217888) != 0 {
		{
			p.SetState(78)
			p.FieldAttribute()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(84)
		p.FieldDefinition()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldDefinitionContext is an interface to support dynamic dispatch.
type IFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFieldDefinitionContext differentiates from other interfaces.
	IsFieldDefinitionContext()
}

type FieldDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDefinitionContext() *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_fieldDefinition
	return p
}

func InitEmptyFieldDefinitionContext(p *FieldDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_fieldDefinition
}

func (*FieldDefinitionContext) IsFieldDefinitionContext() {}

func NewFieldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_fieldDefinition

	return p
}

func (s *FieldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDefinitionContext) CopyAll(ctx *FieldDefinitionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LengthFieldContext struct {
	FieldDefinitionContext
}

func NewLengthFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LengthFieldContext {
	var p = new(LengthFieldContext)

	InitEmptyFieldDefinitionContext(&p.FieldDefinitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldDefinitionContext))

	return p
}

func (s *LengthFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthFieldContext) LengthFieldDeclaration() ILengthFieldDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthFieldDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthFieldDeclarationContext)
}

func (s *LengthFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitLengthField(s)

	default:
		return t.VisitChildren(s)
	}
}

type MetaFieldContext struct {
	FieldDefinitionContext
}

func NewMetaFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MetaFieldContext {
	var p = new(MetaFieldContext)

	InitEmptyFieldDefinitionContext(&p.FieldDefinitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldDefinitionContext))

	return p
}

func (s *MetaFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaFieldContext) MetaDataDeclaration() IMetaDataDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaDataDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaDataDeclarationContext)
}

func (s *MetaFieldContext) REPEAT() antlr.TerminalNode {
	return s.GetToken(PacketDslParserREPEAT, 0)
}

func (s *MetaFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitMetaField(s)

	default:
		return t.VisitChildren(s)
	}
}

type InerObjectFieldContext struct {
	FieldDefinitionContext
}

func NewInerObjectFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InerObjectFieldContext {
	var p = new(InerObjectFieldContext)

	InitEmptyFieldDefinitionContext(&p.FieldDefinitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldDefinitionContext))

	return p
}

func (s *InerObjectFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InerObjectFieldContext) InerObjectDeclaration() IInerObjectDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInerObjectDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInerObjectDeclarationContext)
}

func (s *InerObjectFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PacketDslParserCOMMA, 0)
}

func (s *InerObjectFieldContext) REPEAT() antlr.TerminalNode {
	return s.GetToken(PacketDslParserREPEAT, 0)
}

func (s *InerObjectFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitInerObjectField(s)

	default:
		return t.VisitChildren(s)
	}
}

type ObjectFieldContext struct {
	FieldDefinitionContext
}

func NewObjectFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectFieldContext {
	var p = new(ObjectFieldContext)

	InitEmptyFieldDefinitionContext(&p.FieldDefinitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldDefinitionContext))

	return p
}

func (s *ObjectFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *ObjectFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PacketDslParserCOMMA, 0)
}

func (s *ObjectFieldContext) REPEAT() antlr.TerminalNode {
	return s.GetToken(PacketDslParserREPEAT, 0)
}

func (s *ObjectFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitObjectField(s)

	default:
		return t.VisitChildren(s)
	}
}

type CheckSumFieldContext struct {
	FieldDefinitionContext
}

func NewCheckSumFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CheckSumFieldContext {
	var p = new(CheckSumFieldContext)

	InitEmptyFieldDefinitionContext(&p.FieldDefinitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldDefinitionContext))

	return p
}

func (s *CheckSumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CheckSumFieldContext) CheckSumFieldDeclaration() ICheckSumFieldDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICheckSumFieldDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICheckSumFieldDeclarationContext)
}

func (s *CheckSumFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitCheckSumField(s)

	default:
		return t.VisitChildren(s)
	}
}

type MatchFieldContext struct {
	FieldDefinitionContext
}

func NewMatchFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MatchFieldContext {
	var p = new(MatchFieldContext)

	InitEmptyFieldDefinitionContext(&p.FieldDefinitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*FieldDefinitionContext))

	return p
}

func (s *MatchFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchFieldContext) MatchFieldDeclaration() IMatchFieldDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchFieldDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchFieldDeclarationContext)
}

func (s *MatchFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PacketDslParserCOMMA, 0)
}

func (s *MatchFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitMatchField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) FieldDefinition() (localctx IFieldDefinitionContext) {
	localctx = NewFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PacketDslParserRULE_fieldDefinition)
	var _la int

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInerObjectFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PacketDslParserREPEAT {
			{
				p.SetState(86)
				p.Match(PacketDslParserREPEAT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(89)
			p.InerObjectDeclaration()
		}
		{
			p.SetState(90)
			p.Match(PacketDslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewMetaFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PacketDslParserREPEAT {
			{
				p.SetState(92)
				p.Match(PacketDslParserREPEAT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(95)
			p.MetaDataDeclaration()
		}

	case 3:
		localctx = NewObjectFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PacketDslParserREPEAT {
			{
				p.SetState(96)
				p.Match(PacketDslParserREPEAT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(99)
			p.Match(PacketDslParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(PacketDslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewLengthFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(101)
			p.LengthFieldDeclaration()
		}

	case 5:
		localctx = NewCheckSumFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(102)
			p.CheckSumFieldDeclaration()
		}

	case 6:
		localctx = NewMatchFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(103)
			p.MatchFieldDeclaration()
		}
		{
			p.SetState(104)
			p.Match(PacketDslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMetaDataDefinitionContext is an interface to support dynamic dispatch.
type IMetaDataDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	METADATA() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	AllMetaDataDeclaration() []IMetaDataDeclarationContext
	MetaDataDeclaration(i int) IMetaDataDeclarationContext

	// IsMetaDataDefinitionContext differentiates from other interfaces.
	IsMetaDataDefinitionContext()
}

type MetaDataDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetaDataDefinitionContext() *MetaDataDefinitionContext {
	var p = new(MetaDataDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_metaDataDefinition
	return p
}

func InitEmptyMetaDataDefinitionContext(p *MetaDataDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_metaDataDefinition
}

func (*MetaDataDefinitionContext) IsMetaDataDefinitionContext() {}

func NewMetaDataDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaDataDefinitionContext {
	var p = new(MetaDataDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_metaDataDefinition

	return p
}

func (s *MetaDataDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MetaDataDefinitionContext) METADATA() antlr.TerminalNode {
	return s.GetToken(PacketDslParserMETADATA, 0)
}

func (s *MetaDataDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *MetaDataDefinitionContext) AllMetaDataDeclaration() []IMetaDataDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMetaDataDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IMetaDataDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMetaDataDeclarationContext); ok {
			tst[i] = t.(IMetaDataDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *MetaDataDefinitionContext) MetaDataDeclaration(i int) IMetaDataDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaDataDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaDataDeclarationContext)
}

func (s *MetaDataDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaDataDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetaDataDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitMetaDataDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) MetaDataDefinition() (localctx IMetaDataDefinitionContext) {
	localctx = NewMetaDataDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PacketDslParserRULE_metaDataDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(PacketDslParserMETADATA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&137472478720) != 0 {
		{
			p.SetState(111)
			p.MetaDataDeclaration()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(PacketDslParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILengthFieldDeclarationContext is an interface to support dynamic dispatch.
type ILengthFieldDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	LengthOfAttribute() ILengthOfAttributeContext
	COMMA() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Type_() ITypeContext
	STRING_LITERAL() antlr.TerminalNode

	// IsLengthFieldDeclarationContext differentiates from other interfaces.
	IsLengthFieldDeclarationContext()
}

type LengthFieldDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyLengthFieldDeclarationContext() *LengthFieldDeclarationContext {
	var p = new(LengthFieldDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_lengthFieldDeclaration
	return p
}

func InitEmptyLengthFieldDeclarationContext(p *LengthFieldDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_lengthFieldDeclaration
}

func (*LengthFieldDeclarationContext) IsLengthFieldDeclarationContext() {}

func NewLengthFieldDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthFieldDeclarationContext {
	var p = new(LengthFieldDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_lengthFieldDeclaration

	return p
}

func (s *LengthFieldDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthFieldDeclarationContext) GetName() antlr.Token { return s.name }

func (s *LengthFieldDeclarationContext) SetName(v antlr.Token) { s.name = v }

func (s *LengthFieldDeclarationContext) LengthOfAttribute() ILengthOfAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthOfAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthOfAttributeContext)
}

func (s *LengthFieldDeclarationContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PacketDslParserCOMMA, 0)
}

func (s *LengthFieldDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *LengthFieldDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *LengthFieldDeclarationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(PacketDslParserSTRING_LITERAL, 0)
}

func (s *LengthFieldDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthFieldDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthFieldDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitLengthFieldDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) LengthFieldDeclaration() (localctx ILengthFieldDeclarationContext) {
	localctx = NewLengthFieldDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PacketDslParserRULE_lengthFieldDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(119)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(122)

		var _m = p.Match(PacketDslParserIDENTIFIER)

		localctx.(*LengthFieldDeclarationContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.LengthOfAttribute()
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserSTRING_LITERAL {
		{
			p.SetState(124)
			p.Match(PacketDslParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(127)
		p.Match(PacketDslParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICheckSumFieldDeclarationContext is an interface to support dynamic dispatch.
type ICheckSumFieldDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	CalculatedFromAttribute() ICalculatedFromAttributeContext
	COMMA() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Type_() ITypeContext
	STRING_LITERAL() antlr.TerminalNode

	// IsCheckSumFieldDeclarationContext differentiates from other interfaces.
	IsCheckSumFieldDeclarationContext()
}

type CheckSumFieldDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyCheckSumFieldDeclarationContext() *CheckSumFieldDeclarationContext {
	var p = new(CheckSumFieldDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_checkSumFieldDeclaration
	return p
}

func InitEmptyCheckSumFieldDeclarationContext(p *CheckSumFieldDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_checkSumFieldDeclaration
}

func (*CheckSumFieldDeclarationContext) IsCheckSumFieldDeclarationContext() {}

func NewCheckSumFieldDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CheckSumFieldDeclarationContext {
	var p = new(CheckSumFieldDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_checkSumFieldDeclaration

	return p
}

func (s *CheckSumFieldDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *CheckSumFieldDeclarationContext) GetName() antlr.Token { return s.name }

func (s *CheckSumFieldDeclarationContext) SetName(v antlr.Token) { s.name = v }

func (s *CheckSumFieldDeclarationContext) CalculatedFromAttribute() ICalculatedFromAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculatedFromAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculatedFromAttributeContext)
}

func (s *CheckSumFieldDeclarationContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PacketDslParserCOMMA, 0)
}

func (s *CheckSumFieldDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *CheckSumFieldDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *CheckSumFieldDeclarationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(PacketDslParserSTRING_LITERAL, 0)
}

func (s *CheckSumFieldDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CheckSumFieldDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CheckSumFieldDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitCheckSumFieldDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) CheckSumFieldDeclaration() (localctx ICheckSumFieldDeclarationContext) {
	localctx = NewCheckSumFieldDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PacketDslParserRULE_checkSumFieldDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(129)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(132)

		var _m = p.Match(PacketDslParserIDENTIFIER)

		localctx.(*CheckSumFieldDeclarationContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.CalculatedFromAttribute()
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserSTRING_LITERAL {
		{
			p.SetState(134)
			p.Match(PacketDslParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(137)
		p.Match(PacketDslParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldAttributeContext is an interface to support dynamic dispatch.
type IFieldAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LengthOfAttribute() ILengthOfAttributeContext
	CalculatedFromAttribute() ICalculatedFromAttributeContext
	PaddingAttribute() IPaddingAttributeContext

	// IsFieldAttributeContext differentiates from other interfaces.
	IsFieldAttributeContext()
}

type FieldAttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldAttributeContext() *FieldAttributeContext {
	var p = new(FieldAttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_fieldAttribute
	return p
}

func InitEmptyFieldAttributeContext(p *FieldAttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_fieldAttribute
}

func (*FieldAttributeContext) IsFieldAttributeContext() {}

func NewFieldAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldAttributeContext {
	var p = new(FieldAttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_fieldAttribute

	return p
}

func (s *FieldAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldAttributeContext) LengthOfAttribute() ILengthOfAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthOfAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthOfAttributeContext)
}

func (s *FieldAttributeContext) CalculatedFromAttribute() ICalculatedFromAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculatedFromAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculatedFromAttributeContext)
}

func (s *FieldAttributeContext) PaddingAttribute() IPaddingAttributeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPaddingAttributeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPaddingAttributeContext)
}

func (s *FieldAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitFieldAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) FieldAttribute() (localctx IFieldAttributeContext) {
	localctx = NewFieldAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PacketDslParserRULE_fieldAttribute)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.LengthOfAttribute()
		}

	case PacketDslParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.CalculatedFromAttribute()
		}

	case PacketDslParserPADDING_ATTR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(141)
			p.PaddingAttribute()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICalculatedFromAttributeContext is an interface to support dynamic dispatch.
type ICalculatedFromAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFrom returns the from token.
	GetFrom() antlr.Token

	// SetFrom sets the from token.
	SetFrom(antlr.Token)

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsCalculatedFromAttributeContext differentiates from other interfaces.
	IsCalculatedFromAttributeContext()
}

type CalculatedFromAttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	from   antlr.Token
}

func NewEmptyCalculatedFromAttributeContext() *CalculatedFromAttributeContext {
	var p = new(CalculatedFromAttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_calculatedFromAttribute
	return p
}

func InitEmptyCalculatedFromAttributeContext(p *CalculatedFromAttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_calculatedFromAttribute
}

func (*CalculatedFromAttributeContext) IsCalculatedFromAttributeContext() {}

func NewCalculatedFromAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalculatedFromAttributeContext {
	var p = new(CalculatedFromAttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_calculatedFromAttribute

	return p
}

func (s *CalculatedFromAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *CalculatedFromAttributeContext) GetFrom() antlr.Token { return s.from }

func (s *CalculatedFromAttributeContext) SetFrom(v antlr.Token) { s.from = v }

func (s *CalculatedFromAttributeContext) STRING() antlr.TerminalNode {
	return s.GetToken(PacketDslParserSTRING, 0)
}

func (s *CalculatedFromAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculatedFromAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalculatedFromAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitCalculatedFromAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) CalculatedFromAttribute() (localctx ICalculatedFromAttributeContext) {
	localctx = NewCalculatedFromAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PacketDslParserRULE_calculatedFromAttribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(PacketDslParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)

		var _m = p.Match(PacketDslParserSTRING)

		localctx.(*CalculatedFromAttributeContext).from = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Match(PacketDslParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILengthOfAttributeContext is an interface to support dynamic dispatch.
type ILengthOfAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFrom returns the from token.
	GetFrom() antlr.Token

	// SetFrom sets the from token.
	SetFrom(antlr.Token)

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsLengthOfAttributeContext differentiates from other interfaces.
	IsLengthOfAttributeContext()
}

type LengthOfAttributeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	from   antlr.Token
}

func NewEmptyLengthOfAttributeContext() *LengthOfAttributeContext {
	var p = new(LengthOfAttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_lengthOfAttribute
	return p
}

func InitEmptyLengthOfAttributeContext(p *LengthOfAttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_lengthOfAttribute
}

func (*LengthOfAttributeContext) IsLengthOfAttributeContext() {}

func NewLengthOfAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthOfAttributeContext {
	var p = new(LengthOfAttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_lengthOfAttribute

	return p
}

func (s *LengthOfAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthOfAttributeContext) GetFrom() antlr.Token { return s.from }

func (s *LengthOfAttributeContext) SetFrom(v antlr.Token) { s.from = v }

func (s *LengthOfAttributeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *LengthOfAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthOfAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthOfAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitLengthOfAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) LengthOfAttribute() (localctx ILengthOfAttributeContext) {
	localctx = NewLengthOfAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PacketDslParserRULE_lengthOfAttribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(PacketDslParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)

		var _m = p.Match(PacketDslParserIDENTIFIER)

		localctx.(*LengthOfAttributeContext).from = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(PacketDslParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPaddingAttributeContext is an interface to support dynamic dispatch.
type IPaddingAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPadding returns the padding token.
	GetPadding() antlr.Token

	// SetPadding sets the padding token.
	SetPadding(antlr.Token)

	// Getter signatures
	PADDING_ATTR() antlr.TerminalNode
	PADDING_CHAR() antlr.TerminalNode

	// IsPaddingAttributeContext differentiates from other interfaces.
	IsPaddingAttributeContext()
}

type PaddingAttributeContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	padding antlr.Token
}

func NewEmptyPaddingAttributeContext() *PaddingAttributeContext {
	var p = new(PaddingAttributeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_paddingAttribute
	return p
}

func InitEmptyPaddingAttributeContext(p *PaddingAttributeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_paddingAttribute
}

func (*PaddingAttributeContext) IsPaddingAttributeContext() {}

func NewPaddingAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PaddingAttributeContext {
	var p = new(PaddingAttributeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_paddingAttribute

	return p
}

func (s *PaddingAttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *PaddingAttributeContext) GetPadding() antlr.Token { return s.padding }

func (s *PaddingAttributeContext) SetPadding(v antlr.Token) { s.padding = v }

func (s *PaddingAttributeContext) PADDING_ATTR() antlr.TerminalNode {
	return s.GetToken(PacketDslParserPADDING_ATTR, 0)
}

func (s *PaddingAttributeContext) PADDING_CHAR() antlr.TerminalNode {
	return s.GetToken(PacketDslParserPADDING_CHAR, 0)
}

func (s *PaddingAttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PaddingAttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PaddingAttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitPaddingAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) PaddingAttribute() (localctx IPaddingAttributeContext) {
	localctx = NewPaddingAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PacketDslParserRULE_paddingAttribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(PacketDslParserPADDING_ATTR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(PacketDslParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)

		var _m = p.Match(PacketDslParserPADDING_CHAR)

		localctx.(*PaddingAttributeContext).padding = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(PacketDslParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMetaDataDeclarationContext is an interface to support dynamic dispatch.
type IMetaDataDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Getter signatures
	COMMA() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Type_() ITypeContext
	STRING_LITERAL() antlr.TerminalNode

	// IsMetaDataDeclarationContext differentiates from other interfaces.
	IsMetaDataDeclarationContext()
}

type MetaDataDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyMetaDataDeclarationContext() *MetaDataDeclarationContext {
	var p = new(MetaDataDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_metaDataDeclaration
	return p
}

func InitEmptyMetaDataDeclarationContext(p *MetaDataDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_metaDataDeclaration
}

func (*MetaDataDeclarationContext) IsMetaDataDeclarationContext() {}

func NewMetaDataDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaDataDeclarationContext {
	var p = new(MetaDataDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_metaDataDeclaration

	return p
}

func (s *MetaDataDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MetaDataDeclarationContext) GetName() antlr.Token { return s.name }

func (s *MetaDataDeclarationContext) SetName(v antlr.Token) { s.name = v }

func (s *MetaDataDeclarationContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PacketDslParserCOMMA, 0)
}

func (s *MetaDataDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *MetaDataDeclarationContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *MetaDataDeclarationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(PacketDslParserSTRING_LITERAL, 0)
}

func (s *MetaDataDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaDataDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetaDataDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitMetaDataDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) MetaDataDeclaration() (localctx IMetaDataDeclarationContext) {
	localctx = NewMetaDataDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PacketDslParserRULE_metaDataDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(158)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(157)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(160)

		var _m = p.Match(PacketDslParserIDENTIFIER)

		localctx.(*MetaDataDeclarationContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserSTRING_LITERAL {
		{
			p.SetState(161)
			p.Match(PacketDslParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(164)
		p.Match(PacketDslParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() ITypeContext
	STRING() antlr.TerminalNode
	DIGITS() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(PacketDslParserSTRING, 0)
}

func (s *ValueContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(PacketDslParserDIGITS, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PacketDslParserRULE_value)
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserT__8, PacketDslParserT__9, PacketDslParserT__10, PacketDslParserUINT8, PacketDslParserUINT16, PacketDslParserUINT32, PacketDslParserUINT64, PacketDslParserINT8, PacketDslParserINT16, PacketDslParserINT32, PacketDslParserINT64, PacketDslParserFLOAT32, PacketDslParserFLOAT64, PacketDslParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(166)
			p.Type_()
		}

	case PacketDslParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Match(PacketDslParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserDIGITS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(168)
			p.Match(PacketDslParserDIGITS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	UINT8() antlr.TerminalNode
	UINT16() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	INT8() antlr.TerminalNode
	INT16() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	FLOAT32() antlr.TerminalNode
	FLOAT64() antlr.TerminalNode
	DIGITS() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *TypeContext) UINT8() antlr.TerminalNode {
	return s.GetToken(PacketDslParserUINT8, 0)
}

func (s *TypeContext) UINT16() antlr.TerminalNode {
	return s.GetToken(PacketDslParserUINT16, 0)
}

func (s *TypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(PacketDslParserUINT32, 0)
}

func (s *TypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(PacketDslParserUINT64, 0)
}

func (s *TypeContext) INT8() antlr.TerminalNode {
	return s.GetToken(PacketDslParserINT8, 0)
}

func (s *TypeContext) INT16() antlr.TerminalNode {
	return s.GetToken(PacketDslParserINT16, 0)
}

func (s *TypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(PacketDslParserINT32, 0)
}

func (s *TypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(PacketDslParserINT64, 0)
}

func (s *TypeContext) FLOAT32() antlr.TerminalNode {
	return s.GetToken(PacketDslParserFLOAT32, 0)
}

func (s *TypeContext) FLOAT64() antlr.TerminalNode {
	return s.GetToken(PacketDslParserFLOAT64, 0)
}

func (s *TypeContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(PacketDslParserDIGITS, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PacketDslParserRULE_type)
	var _la int

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Match(PacketDslParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserUINT8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Match(PacketDslParserUINT8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserUINT16:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(173)
			p.Match(PacketDslParserUINT16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserUINT32:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(174)
			p.Match(PacketDslParserUINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserUINT64:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(175)
			p.Match(PacketDslParserUINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserINT8:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(176)
			p.Match(PacketDslParserINT8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserINT16:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(177)
			p.Match(PacketDslParserINT16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserINT32:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(178)
			p.Match(PacketDslParserINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserINT64:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(179)
			p.Match(PacketDslParserINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserFLOAT32:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(180)
			p.Match(PacketDslParserFLOAT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserFLOAT64:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(181)
			p.Match(PacketDslParserFLOAT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__8:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(182)
			p.Match(PacketDslParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__9:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(183)
			p.Match(PacketDslParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__10:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(184)
			p.Match(PacketDslParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PacketDslParserDIGITS {
			{
				p.SetState(185)
				p.Match(PacketDslParserDIGITS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(188)
			p.Match(PacketDslParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatchFieldDeclarationContext is an interface to support dynamic dispatch.
type IMatchFieldDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMatchKey returns the matchKey token.
	GetMatchKey() antlr.Token

	// GetMatchName returns the matchName token.
	GetMatchName() antlr.Token

	// SetMatchKey sets the matchKey token.
	SetMatchKey(antlr.Token)

	// SetMatchName sets the matchName token.
	SetMatchName(antlr.Token)

	// Getter signatures
	MATCH() antlr.TerminalNode
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllMatchPair() []IMatchPairContext
	MatchPair(i int) IMatchPairContext

	// IsMatchFieldDeclarationContext differentiates from other interfaces.
	IsMatchFieldDeclarationContext()
}

type MatchFieldDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser    antlr.Parser
	matchKey  antlr.Token
	matchName antlr.Token
}

func NewEmptyMatchFieldDeclarationContext() *MatchFieldDeclarationContext {
	var p = new(MatchFieldDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_matchFieldDeclaration
	return p
}

func InitEmptyMatchFieldDeclarationContext(p *MatchFieldDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_matchFieldDeclaration
}

func (*MatchFieldDeclarationContext) IsMatchFieldDeclarationContext() {}

func NewMatchFieldDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchFieldDeclarationContext {
	var p = new(MatchFieldDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_matchFieldDeclaration

	return p
}

func (s *MatchFieldDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchFieldDeclarationContext) GetMatchKey() antlr.Token { return s.matchKey }

func (s *MatchFieldDeclarationContext) GetMatchName() antlr.Token { return s.matchName }

func (s *MatchFieldDeclarationContext) SetMatchKey(v antlr.Token) { s.matchKey = v }

func (s *MatchFieldDeclarationContext) SetMatchName(v antlr.Token) { s.matchName = v }

func (s *MatchFieldDeclarationContext) MATCH() antlr.TerminalNode {
	return s.GetToken(PacketDslParserMATCH, 0)
}

func (s *MatchFieldDeclarationContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PacketDslParserIDENTIFIER)
}

func (s *MatchFieldDeclarationContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, i)
}

func (s *MatchFieldDeclarationContext) AllMatchPair() []IMatchPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMatchPairContext); ok {
			len++
		}
	}

	tst := make([]IMatchPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMatchPairContext); ok {
			tst[i] = t.(IMatchPairContext)
			i++
		}
	}

	return tst
}

func (s *MatchFieldDeclarationContext) MatchPair(i int) IMatchPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchPairContext)
}

func (s *MatchFieldDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchFieldDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchFieldDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitMatchFieldDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) MatchFieldDeclaration() (localctx IMatchFieldDeclarationContext) {
	localctx = NewMatchFieldDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PacketDslParserRULE_matchFieldDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Match(PacketDslParserMATCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)

		var _m = p.Match(PacketDslParserIDENTIFIER)

		localctx.(*MatchFieldDeclarationContext).matchKey = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Match(PacketDslParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)

		var _m = p.Match(PacketDslParserIDENTIFIER)

		localctx.(*MatchFieldDeclarationContext).matchName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&100679680) != 0) {
		{
			p.SetState(196)
			p.MatchPair()
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(201)
		p.Match(PacketDslParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMatchPairContext is an interface to support dynamic dispatch.
type IMatchPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	DIGITS() antlr.TerminalNode
	STRING() antlr.TerminalNode
	List() IListContext
	COMMA() antlr.TerminalNode

	// IsMatchPairContext differentiates from other interfaces.
	IsMatchPairContext()
}

type MatchPairContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchPairContext() *MatchPairContext {
	var p = new(MatchPairContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_matchPair
	return p
}

func InitEmptyMatchPairContext(p *MatchPairContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_matchPair
}

func (*MatchPairContext) IsMatchPairContext() {}

func NewMatchPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchPairContext {
	var p = new(MatchPairContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_matchPair

	return p
}

func (s *MatchPairContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchPairContext) COLON() antlr.TerminalNode {
	return s.GetToken(PacketDslParserCOLON, 0)
}

func (s *MatchPairContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *MatchPairContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(PacketDslParserDIGITS, 0)
}

func (s *MatchPairContext) STRING() antlr.TerminalNode {
	return s.GetToken(PacketDslParserSTRING, 0)
}

func (s *MatchPairContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *MatchPairContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PacketDslParserCOMMA, 0)
}

func (s *MatchPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitMatchPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) MatchPair() (localctx IMatchPairContext) {
	localctx = NewMatchPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PacketDslParserRULE_matchPair)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserDIGITS:
		{
			p.SetState(203)
			p.Match(PacketDslParserDIGITS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserSTRING:
		{
			p.SetState(204)
			p.Match(PacketDslParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__13:
		{
			p.SetState(205)
			p.List()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(208)
		p.Match(PacketDslParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserCOMMA {
		{
			p.SetState(210)
			p.Match(PacketDslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInerObjectDeclarationContext is an interface to support dynamic dispatch.
type IInerObjectDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllFieldDefinition() []IFieldDefinitionContext
	FieldDefinition(i int) IFieldDefinitionContext

	// IsInerObjectDeclarationContext differentiates from other interfaces.
	IsInerObjectDeclarationContext()
}

type InerObjectDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInerObjectDeclarationContext() *InerObjectDeclarationContext {
	var p = new(InerObjectDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_inerObjectDeclaration
	return p
}

func InitEmptyInerObjectDeclarationContext(p *InerObjectDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_inerObjectDeclaration
}

func (*InerObjectDeclarationContext) IsInerObjectDeclarationContext() {}

func NewInerObjectDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InerObjectDeclarationContext {
	var p = new(InerObjectDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_inerObjectDeclaration

	return p
}

func (s *InerObjectDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *InerObjectDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *InerObjectDeclarationContext) AllFieldDefinition() []IFieldDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IFieldDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldDefinitionContext); ok {
			tst[i] = t.(IFieldDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *InerObjectDeclarationContext) FieldDefinition(i int) IFieldDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldDefinitionContext)
}

func (s *InerObjectDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InerObjectDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InerObjectDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitInerObjectDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) InerObjectDeclaration() (localctx IInerObjectDeclarationContext) {
	localctx = NewInerObjectDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PacketDslParserRULE_inerObjectDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(214)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&148209896960) != 0) {
		{
			p.SetState(215)
			p.FieldDefinition()
		}

		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(220)
		p.Match(PacketDslParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDIGITS() []antlr.TerminalNode
	DIGITS(i int) antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllDIGITS() []antlr.TerminalNode {
	return s.GetTokens(PacketDslParserDIGITS)
}

func (s *ListContext) DIGITS(i int) antlr.TerminalNode {
	return s.GetToken(PacketDslParserDIGITS, i)
}

func (s *ListContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(PacketDslParserSTRING)
}

func (s *ListContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(PacketDslParserSTRING, i)
}

func (s *ListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PacketDslParserCOMMA)
}

func (s *ListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PacketDslParserCOMMA, i)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PacketDslParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(PacketDslParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PacketDslParserDIGITS || _la == PacketDslParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PacketDslParserCOMMA {
		{
			p.SetState(224)
			p.Match(PacketDslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PacketDslParserDIGITS || _la == PacketDslParserSTRING) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(231)
		p.Match(PacketDslParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
