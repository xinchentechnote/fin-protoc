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
		"", "'options'", "'{'", "'}'", "'='", "'@lengthOf('", "')'", "'@calculatedFrom('",
		"'string'", "'char'", "'char['", "']'", "'as'", "'['", "", "", "", "",
		"", "", "", "", "", "", "", "", "'root'", "'packet'", "'repeat'", "'MetaData'",
		"'match'", "':'", "','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "UINT8", "UINT16",
		"UINT32", "UINT64", "INT8", "INT16", "INT32", "INT64", "FLOAT32", "FLOAT64",
		"DIGITS", "STRING", "ROOT", "PACKET", "REPEAT", "METADATA", "MATCH",
		"COLON", "COMMA", "SEMICOLON", "IDENTIFIER", "STRING_LITERAL", "LINE_COMMENT",
		"WS",
	}
	staticData.RuleNames = []string{
		"packet", "optionDefinition", "optionDeclaration", "packetDefinition",
		"fieldDefinitionWithAttribute", "fieldDefinition", "metaDataDefinition",
		"lengthFieldDeclaration", "checkSumFieldDeclaration", "fieldAttribute",
		"lengthOfAttribute", "calculatedFromAttribute", "metaDataDeclaration",
		"value", "type", "matchFieldDeclaration", "matchPair", "inerObjectDeclaration",
		"list",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 37, 226, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 5, 0, 42, 8,
		0, 10, 0, 12, 0, 45, 9, 0, 1, 1, 1, 1, 1, 1, 5, 1, 50, 8, 1, 10, 1, 12,
		1, 53, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 61, 8, 2, 1, 3,
		3, 3, 64, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 70, 8, 3, 10, 3, 12, 3, 73,
		9, 3, 1, 3, 1, 3, 1, 4, 5, 4, 78, 8, 4, 10, 4, 12, 4, 81, 9, 4, 1, 4, 1,
		4, 1, 5, 3, 5, 86, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 92, 8, 5, 1, 5,
		1, 5, 3, 5, 96, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 105,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 111, 8, 6, 10, 6, 12, 6, 114, 9, 6,
		1, 6, 1, 6, 1, 7, 3, 7, 119, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 124, 8, 7, 1,
		7, 1, 7, 1, 8, 3, 8, 129, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 134, 8, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 3, 9, 140, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 3, 12, 151, 8, 12, 1, 12, 1, 12, 3, 12, 155, 8,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 3, 13, 162, 8, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 3, 14, 179, 8, 14, 1, 14, 3, 14, 182, 8, 14, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 4, 15, 190, 8, 15, 11, 15, 12, 15, 191,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 3, 16, 199, 8, 16, 1, 16, 1, 16, 1,
		16, 3, 16, 204, 8, 16, 1, 17, 1, 17, 1, 17, 4, 17, 209, 8, 17, 11, 17,
		12, 17, 210, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 219, 8, 18,
		10, 18, 12, 18, 222, 9, 18, 1, 18, 1, 18, 1, 18, 0, 0, 19, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 0, 1, 1, 0,
		24, 25, 252, 0, 43, 1, 0, 0, 0, 2, 46, 1, 0, 0, 0, 4, 56, 1, 0, 0, 0, 6,
		63, 1, 0, 0, 0, 8, 79, 1, 0, 0, 0, 10, 104, 1, 0, 0, 0, 12, 106, 1, 0,
		0, 0, 14, 118, 1, 0, 0, 0, 16, 128, 1, 0, 0, 0, 18, 139, 1, 0, 0, 0, 20,
		141, 1, 0, 0, 0, 22, 145, 1, 0, 0, 0, 24, 150, 1, 0, 0, 0, 26, 161, 1,
		0, 0, 0, 28, 181, 1, 0, 0, 0, 30, 183, 1, 0, 0, 0, 32, 198, 1, 0, 0, 0,
		34, 205, 1, 0, 0, 0, 36, 214, 1, 0, 0, 0, 38, 42, 3, 6, 3, 0, 39, 42, 3,
		12, 6, 0, 40, 42, 3, 2, 1, 0, 41, 38, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41,
		40, 1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0,
		0, 44, 1, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 47, 5, 1, 0, 0, 47, 51, 5,
		2, 0, 0, 48, 50, 3, 4, 2, 0, 49, 48, 1, 0, 0, 0, 50, 53, 1, 0, 0, 0, 51,
		49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 1, 0, 0, 0, 53, 51, 1, 0, 0,
		0, 54, 55, 5, 3, 0, 0, 55, 3, 1, 0, 0, 0, 56, 57, 5, 34, 0, 0, 57, 58,
		5, 4, 0, 0, 58, 60, 3, 26, 13, 0, 59, 61, 5, 33, 0, 0, 60, 59, 1, 0, 0,
		0, 60, 61, 1, 0, 0, 0, 61, 5, 1, 0, 0, 0, 62, 64, 5, 26, 0, 0, 63, 62,
		1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 5, 27, 0, 0,
		66, 67, 5, 34, 0, 0, 67, 71, 5, 2, 0, 0, 68, 70, 3, 10, 5, 0, 69, 68, 1,
		0, 0, 0, 70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72,
		74, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 74, 75, 5, 3, 0, 0, 75, 7, 1, 0, 0,
		0, 76, 78, 3, 18, 9, 0, 77, 76, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 77,
		1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 1, 0, 0, 0, 81, 79, 1, 0, 0, 0,
		82, 83, 3, 10, 5, 0, 83, 9, 1, 0, 0, 0, 84, 86, 5, 28, 0, 0, 85, 84, 1,
		0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 3, 34, 17, 0,
		88, 89, 5, 32, 0, 0, 89, 105, 1, 0, 0, 0, 90, 92, 5, 28, 0, 0, 91, 90,
		1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 105, 3, 24, 12,
		0, 94, 96, 5, 28, 0, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97,
		1, 0, 0, 0, 97, 98, 5, 34, 0, 0, 98, 105, 5, 32, 0, 0, 99, 105, 3, 14,
		7, 0, 100, 105, 3, 16, 8, 0, 101, 102, 3, 30, 15, 0, 102, 103, 5, 32, 0,
		0, 103, 105, 1, 0, 0, 0, 104, 85, 1, 0, 0, 0, 104, 91, 1, 0, 0, 0, 104,
		95, 1, 0, 0, 0, 104, 99, 1, 0, 0, 0, 104, 100, 1, 0, 0, 0, 104, 101, 1,
		0, 0, 0, 105, 11, 1, 0, 0, 0, 106, 107, 5, 29, 0, 0, 107, 108, 5, 34, 0,
		0, 108, 112, 5, 2, 0, 0, 109, 111, 3, 24, 12, 0, 110, 109, 1, 0, 0, 0,
		111, 114, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113,
		115, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 115, 116, 5, 3, 0, 0, 116, 13, 1,
		0, 0, 0, 117, 119, 3, 28, 14, 0, 118, 117, 1, 0, 0, 0, 118, 119, 1, 0,
		0, 0, 119, 120, 1, 0, 0, 0, 120, 121, 5, 34, 0, 0, 121, 123, 3, 20, 10,
		0, 122, 124, 5, 35, 0, 0, 123, 122, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124,
		125, 1, 0, 0, 0, 125, 126, 5, 32, 0, 0, 126, 15, 1, 0, 0, 0, 127, 129,
		3, 28, 14, 0, 128, 127, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 1,
		0, 0, 0, 130, 131, 5, 34, 0, 0, 131, 133, 3, 22, 11, 0, 132, 134, 5, 35,
		0, 0, 133, 132, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0,
		135, 136, 5, 32, 0, 0, 136, 17, 1, 0, 0, 0, 137, 140, 3, 20, 10, 0, 138,
		140, 3, 22, 11, 0, 139, 137, 1, 0, 0, 0, 139, 138, 1, 0, 0, 0, 140, 19,
		1, 0, 0, 0, 141, 142, 5, 5, 0, 0, 142, 143, 5, 34, 0, 0, 143, 144, 5, 6,
		0, 0, 144, 21, 1, 0, 0, 0, 145, 146, 5, 7, 0, 0, 146, 147, 5, 25, 0, 0,
		147, 148, 5, 6, 0, 0, 148, 23, 1, 0, 0, 0, 149, 151, 3, 28, 14, 0, 150,
		149, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 154,
		5, 34, 0, 0, 153, 155, 5, 35, 0, 0, 154, 153, 1, 0, 0, 0, 154, 155, 1,
		0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 5, 32, 0, 0, 157, 25, 1, 0, 0,
		0, 158, 162, 3, 28, 14, 0, 159, 162, 5, 25, 0, 0, 160, 162, 5, 24, 0, 0,
		161, 158, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 162,
		27, 1, 0, 0, 0, 163, 182, 5, 34, 0, 0, 164, 182, 5, 14, 0, 0, 165, 182,
		5, 15, 0, 0, 166, 182, 5, 16, 0, 0, 167, 182, 5, 17, 0, 0, 168, 182, 5,
		18, 0, 0, 169, 182, 5, 19, 0, 0, 170, 182, 5, 20, 0, 0, 171, 182, 5, 21,
		0, 0, 172, 182, 5, 22, 0, 0, 173, 182, 5, 23, 0, 0, 174, 182, 5, 8, 0,
		0, 175, 182, 5, 9, 0, 0, 176, 178, 5, 10, 0, 0, 177, 179, 5, 24, 0, 0,
		178, 177, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180,
		182, 5, 11, 0, 0, 181, 163, 1, 0, 0, 0, 181, 164, 1, 0, 0, 0, 181, 165,
		1, 0, 0, 0, 181, 166, 1, 0, 0, 0, 181, 167, 1, 0, 0, 0, 181, 168, 1, 0,
		0, 0, 181, 169, 1, 0, 0, 0, 181, 170, 1, 0, 0, 0, 181, 171, 1, 0, 0, 0,
		181, 172, 1, 0, 0, 0, 181, 173, 1, 0, 0, 0, 181, 174, 1, 0, 0, 0, 181,
		175, 1, 0, 0, 0, 181, 176, 1, 0, 0, 0, 182, 29, 1, 0, 0, 0, 183, 184, 5,
		30, 0, 0, 184, 185, 5, 34, 0, 0, 185, 186, 5, 12, 0, 0, 186, 187, 5, 34,
		0, 0, 187, 189, 5, 2, 0, 0, 188, 190, 3, 32, 16, 0, 189, 188, 1, 0, 0,
		0, 190, 191, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192,
		193, 1, 0, 0, 0, 193, 194, 5, 3, 0, 0, 194, 31, 1, 0, 0, 0, 195, 199, 5,
		24, 0, 0, 196, 199, 5, 25, 0, 0, 197, 199, 3, 36, 18, 0, 198, 195, 1, 0,
		0, 0, 198, 196, 1, 0, 0, 0, 198, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0,
		200, 201, 5, 31, 0, 0, 201, 203, 5, 34, 0, 0, 202, 204, 5, 32, 0, 0, 203,
		202, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 33, 1, 0, 0, 0, 205, 206, 5,
		34, 0, 0, 206, 208, 5, 2, 0, 0, 207, 209, 3, 10, 5, 0, 208, 207, 1, 0,
		0, 0, 209, 210, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0,
		211, 212, 1, 0, 0, 0, 212, 213, 5, 3, 0, 0, 213, 35, 1, 0, 0, 0, 214, 215,
		5, 13, 0, 0, 215, 220, 7, 0, 0, 0, 216, 217, 5, 32, 0, 0, 217, 219, 7,
		0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0, 220, 218, 1, 0, 0,
		0, 220, 221, 1, 0, 0, 0, 221, 223, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 223,
		224, 5, 11, 0, 0, 224, 37, 1, 0, 0, 0, 27, 41, 43, 51, 60, 63, 71, 79,
		85, 91, 95, 104, 112, 118, 123, 128, 133, 139, 150, 154, 161, 178, 181,
		191, 198, 203, 210, 220,
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
	PacketDslParserUINT8          = 14
	PacketDslParserUINT16         = 15
	PacketDslParserUINT32         = 16
	PacketDslParserUINT64         = 17
	PacketDslParserINT8           = 18
	PacketDslParserINT16          = 19
	PacketDslParserINT32          = 20
	PacketDslParserINT64          = 21
	PacketDslParserFLOAT32        = 22
	PacketDslParserFLOAT64        = 23
	PacketDslParserDIGITS         = 24
	PacketDslParserSTRING         = 25
	PacketDslParserROOT           = 26
	PacketDslParserPACKET         = 27
	PacketDslParserREPEAT         = 28
	PacketDslParserMETADATA       = 29
	PacketDslParserMATCH          = 30
	PacketDslParserCOLON          = 31
	PacketDslParserCOMMA          = 32
	PacketDslParserSEMICOLON      = 33
	PacketDslParserIDENTIFIER     = 34
	PacketDslParserSTRING_LITERAL = 35
	PacketDslParserLINE_COMMENT   = 36
	PacketDslParserWS             = 37
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
	PacketDslParserRULE_lengthOfAttribute            = 10
	PacketDslParserRULE_calculatedFromAttribute      = 11
	PacketDslParserRULE_metaDataDeclaration          = 12
	PacketDslParserRULE_value                        = 13
	PacketDslParserRULE_type                         = 14
	PacketDslParserRULE_matchFieldDeclaration        = 15
	PacketDslParserRULE_matchPair                    = 16
	PacketDslParserRULE_inerObjectDeclaration        = 17
	PacketDslParserRULE_list                         = 18
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
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&738197506) != 0 {
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PacketDslParserROOT, PacketDslParserPACKET:
			{
				p.SetState(38)
				p.PacketDefinition()
			}

		case PacketDslParserMETADATA:
			{
				p.SetState(39)
				p.MetaDataDefinition()
			}

		case PacketDslParserT__0:
			{
				p.SetState(40)
				p.OptionDefinition()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(45)
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
		p.SetState(46)
		p.Match(PacketDslParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PacketDslParserIDENTIFIER {
		{
			p.SetState(48)
			p.OptionDeclaration()
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
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
		p.SetState(56)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(PacketDslParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Value()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserSEMICOLON {
		{
			p.SetState(59)
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
	AllFieldDefinition() []IFieldDefinitionContext
	FieldDefinition(i int) IFieldDefinitionContext

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

func (s *PacketDefinitionContext) AllFieldDefinition() []IFieldDefinitionContext {
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

func (s *PacketDefinitionContext) FieldDefinition(i int) IFieldDefinitionContext {
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
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserROOT {
		{
			p.SetState(62)
			p.Match(PacketDslParserROOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(65)
		p.Match(PacketDslParserPACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(66)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(67)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18538809088) != 0 {
		{
			p.SetState(68)
			p.FieldDefinition()
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(74)
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
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PacketDslParserT__4 || _la == PacketDslParserT__6 {
		{
			p.SetState(76)
			p.FieldAttribute()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
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

	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInerObjectFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PacketDslParserREPEAT {
			{
				p.SetState(84)
				p.Match(PacketDslParserREPEAT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(87)
			p.InerObjectDeclaration()
		}
		{
			p.SetState(88)
			p.Match(PacketDslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewMetaFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PacketDslParserREPEAT {
			{
				p.SetState(90)
				p.Match(PacketDslParserREPEAT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(93)
			p.MetaDataDeclaration()
		}

	case 3:
		localctx = NewObjectFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PacketDslParserREPEAT {
			{
				p.SetState(94)
				p.Match(PacketDslParserREPEAT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(97)
			p.Match(PacketDslParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
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
			p.SetState(99)
			p.LengthFieldDeclaration()
		}

	case 5:
		localctx = NewCheckSumFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(100)
			p.CheckSumFieldDeclaration()
		}

	case 6:
		localctx = NewMatchFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(101)
			p.MatchFieldDeclaration()
		}
		{
			p.SetState(102)
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
		p.SetState(106)
		p.Match(PacketDslParserMETADATA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17196631808) != 0 {
		{
			p.SetState(109)
			p.MetaDataDeclaration()
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(115)
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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(117)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(120)

		var _m = p.Match(PacketDslParserIDENTIFIER)

		localctx.(*LengthFieldDeclarationContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.LengthOfAttribute()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserSTRING_LITERAL {
		{
			p.SetState(122)
			p.Match(PacketDslParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(125)
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
	COMMA() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	CalculatedFromAttribute() ICalculatedFromAttributeContext
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

func (s *CheckSumFieldDeclarationContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PacketDslParserCOMMA, 0)
}

func (s *CheckSumFieldDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

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
	p.SetState(128)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(127)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(130)

		var _m = p.Match(PacketDslParserIDENTIFIER)

		localctx.(*CheckSumFieldDeclarationContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(131)
		p.CalculatedFromAttribute()
	}

	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserSTRING_LITERAL {
		{
			p.SetState(132)
			p.Match(PacketDslParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(135)
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
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.LengthOfAttribute()
		}

	case PacketDslParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.CalculatedFromAttribute()
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
	p.EnterRule(localctx, 20, PacketDslParserRULE_lengthOfAttribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(PacketDslParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)

		var _m = p.Match(PacketDslParserIDENTIFIER)

		localctx.(*LengthOfAttributeContext).from = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
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
	p.EnterRule(localctx, 22, PacketDslParserRULE_calculatedFromAttribute)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Match(PacketDslParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)

		var _m = p.Match(PacketDslParserSTRING)

		localctx.(*CalculatedFromAttributeContext).from = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
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
	p.EnterRule(localctx, 24, PacketDslParserRULE_metaDataDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(149)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(152)

		var _m = p.Match(PacketDslParserIDENTIFIER)

		localctx.(*MetaDataDeclarationContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserSTRING_LITERAL {
		{
			p.SetState(153)
			p.Match(PacketDslParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(156)
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
	p.EnterRule(localctx, 26, PacketDslParserRULE_value)
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserT__7, PacketDslParserT__8, PacketDslParserT__9, PacketDslParserUINT8, PacketDslParserUINT16, PacketDslParserUINT32, PacketDslParserUINT64, PacketDslParserINT8, PacketDslParserINT16, PacketDslParserINT32, PacketDslParserINT64, PacketDslParserFLOAT32, PacketDslParserFLOAT64, PacketDslParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Type_()
		}

	case PacketDslParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Match(PacketDslParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserDIGITS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
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
	p.EnterRule(localctx, 28, PacketDslParserRULE_type)
	var _la int

	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)
			p.Match(PacketDslParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserUINT8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Match(PacketDslParserUINT8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserUINT16:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(165)
			p.Match(PacketDslParserUINT16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserUINT32:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(166)
			p.Match(PacketDslParserUINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserUINT64:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(167)
			p.Match(PacketDslParserUINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserINT8:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(168)
			p.Match(PacketDslParserINT8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserINT16:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(169)
			p.Match(PacketDslParserINT16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserINT32:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(170)
			p.Match(PacketDslParserINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserINT64:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(171)
			p.Match(PacketDslParserINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserFLOAT32:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(172)
			p.Match(PacketDslParserFLOAT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserFLOAT64:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(173)
			p.Match(PacketDslParserFLOAT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__7:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(174)
			p.Match(PacketDslParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__8:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(175)
			p.Match(PacketDslParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__9:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(176)
			p.Match(PacketDslParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PacketDslParserDIGITS {
			{
				p.SetState(177)
				p.Match(PacketDslParserDIGITS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(180)
			p.Match(PacketDslParserT__10)
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
	p.EnterRule(localctx, 30, PacketDslParserRULE_matchFieldDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(PacketDslParserMATCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)

		var _m = p.Match(PacketDslParserIDENTIFIER)

		localctx.(*MatchFieldDeclarationContext).matchKey = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(PacketDslParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)

		var _m = p.Match(PacketDslParserIDENTIFIER)

		localctx.(*MatchFieldDeclarationContext).matchName = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&50339840) != 0) {
		{
			p.SetState(188)
			p.MatchPair()
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(193)
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
	p.EnterRule(localctx, 32, PacketDslParserRULE_matchPair)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserDIGITS:
		{
			p.SetState(195)
			p.Match(PacketDslParserDIGITS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserSTRING:
		{
			p.SetState(196)
			p.Match(PacketDslParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__12:
		{
			p.SetState(197)
			p.List()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(200)
		p.Match(PacketDslParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserCOMMA {
		{
			p.SetState(202)
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
	p.EnterRule(localctx, 34, PacketDslParserRULE_inerObjectDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(206)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18538809088) != 0) {
		{
			p.SetState(207)
			p.FieldDefinition()
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(212)
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
	p.EnterRule(localctx, 36, PacketDslParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(PacketDslParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PacketDslParserDIGITS || _la == PacketDslParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PacketDslParserCOMMA {
		{
			p.SetState(216)
			p.Match(PacketDslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PacketDslParserDIGITS || _la == PacketDslParserSTRING) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(222)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(223)
		p.Match(PacketDslParserT__10)
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
