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
		"", "'options'", "'{'", "'}'", "'='", "'MetaData'", "'string'", "'char'",
		"'char['", "']'", "'match'", "'['", "", "", "", "", "", "", "", "",
		"", "", "", "", "'root'", "'packet'", "", "','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "UINT8", "UINT16", "UINT32",
		"UINT64", "INT8", "INT16", "INT32", "INT64", "FLOAT32", "FLOAT64", "DIGITS",
		"STRING", "ROOT", "PACKET", "COLON", "COMMA", "SEMICOLON", "IDENTIFIER",
		"STRING_LITERAL", "LINE_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"packet", "optionDefinition", "optionDeclaration", "packetDefinition",
		"fieldDefinition", "metaDataDefinition", "metaDataDeclaration", "value",
		"type", "matchField", "matchPair", "list",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 145, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 5, 0, 28, 8, 0, 10, 0, 12, 0, 31, 9,
		0, 1, 1, 1, 1, 1, 1, 5, 1, 36, 8, 1, 10, 1, 12, 1, 39, 9, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 47, 8, 2, 1, 3, 3, 3, 50, 8, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 5, 3, 56, 8, 3, 10, 3, 12, 3, 59, 9, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 4, 3, 4, 66, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 72, 8, 5, 10,
		5, 12, 5, 75, 9, 5, 1, 5, 1, 5, 1, 6, 3, 6, 80, 8, 6, 1, 6, 1, 6, 3, 6,
		84, 8, 6, 1, 6, 3, 6, 87, 8, 6, 1, 7, 1, 7, 1, 7, 3, 7, 92, 8, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 109, 8, 8, 1, 8, 3, 8, 112, 8, 8, 1, 9, 1, 9, 1, 9, 1,
		9, 4, 9, 118, 8, 9, 11, 9, 12, 9, 119, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		3, 10, 127, 8, 10, 1, 10, 1, 10, 1, 10, 3, 10, 132, 8, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 5, 11, 138, 8, 11, 10, 11, 12, 11, 141, 9, 11, 1, 11, 1,
		11, 1, 11, 0, 0, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 1, 1,
		0, 22, 23, 166, 0, 29, 1, 0, 0, 0, 2, 32, 1, 0, 0, 0, 4, 42, 1, 0, 0, 0,
		6, 49, 1, 0, 0, 0, 8, 65, 1, 0, 0, 0, 10, 67, 1, 0, 0, 0, 12, 79, 1, 0,
		0, 0, 14, 91, 1, 0, 0, 0, 16, 111, 1, 0, 0, 0, 18, 113, 1, 0, 0, 0, 20,
		126, 1, 0, 0, 0, 22, 133, 1, 0, 0, 0, 24, 28, 3, 6, 3, 0, 25, 28, 3, 10,
		5, 0, 26, 28, 3, 2, 1, 0, 27, 24, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 26,
		1, 0, 0, 0, 28, 31, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0,
		30, 1, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 32, 33, 5, 1, 0, 0, 33, 37, 5, 2,
		0, 0, 34, 36, 3, 4, 2, 0, 35, 34, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 35,
		1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0,
		40, 41, 5, 3, 0, 0, 41, 3, 1, 0, 0, 0, 42, 43, 5, 29, 0, 0, 43, 44, 5,
		4, 0, 0, 44, 46, 3, 14, 7, 0, 45, 47, 5, 28, 0, 0, 46, 45, 1, 0, 0, 0,
		46, 47, 1, 0, 0, 0, 47, 5, 1, 0, 0, 0, 48, 50, 5, 24, 0, 0, 49, 48, 1,
		0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 5, 25, 0, 0, 52,
		53, 5, 29, 0, 0, 53, 57, 5, 2, 0, 0, 54, 56, 3, 8, 4, 0, 55, 54, 1, 0,
		0, 0, 56, 59, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60,
		1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 60, 61, 5, 3, 0, 0, 61, 7, 1, 0, 0, 0,
		62, 66, 5, 29, 0, 0, 63, 66, 3, 12, 6, 0, 64, 66, 3, 18, 9, 0, 65, 62,
		1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 64, 1, 0, 0, 0, 66, 9, 1, 0, 0, 0,
		67, 68, 5, 5, 0, 0, 68, 69, 5, 29, 0, 0, 69, 73, 5, 2, 0, 0, 70, 72, 3,
		12, 6, 0, 71, 70, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73,
		74, 1, 0, 0, 0, 74, 76, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 5, 3, 0,
		0, 77, 11, 1, 0, 0, 0, 78, 80, 3, 16, 8, 0, 79, 78, 1, 0, 0, 0, 79, 80,
		1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 83, 5, 29, 0, 0, 82, 84, 5, 30, 0,
		0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 86, 1, 0, 0, 0, 85, 87,
		5, 27, 0, 0, 86, 85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 13, 1, 0, 0, 0,
		88, 92, 3, 16, 8, 0, 89, 92, 5, 23, 0, 0, 90, 92, 5, 22, 0, 0, 91, 88,
		1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 90, 1, 0, 0, 0, 92, 15, 1, 0, 0, 0,
		93, 112, 5, 29, 0, 0, 94, 112, 5, 12, 0, 0, 95, 112, 5, 13, 0, 0, 96, 112,
		5, 14, 0, 0, 97, 112, 5, 15, 0, 0, 98, 112, 5, 16, 0, 0, 99, 112, 5, 17,
		0, 0, 100, 112, 5, 18, 0, 0, 101, 112, 5, 19, 0, 0, 102, 112, 5, 20, 0,
		0, 103, 112, 5, 21, 0, 0, 104, 112, 5, 6, 0, 0, 105, 112, 5, 7, 0, 0, 106,
		108, 5, 8, 0, 0, 107, 109, 5, 22, 0, 0, 108, 107, 1, 0, 0, 0, 108, 109,
		1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 5, 9, 0, 0, 111, 93, 1, 0,
		0, 0, 111, 94, 1, 0, 0, 0, 111, 95, 1, 0, 0, 0, 111, 96, 1, 0, 0, 0, 111,
		97, 1, 0, 0, 0, 111, 98, 1, 0, 0, 0, 111, 99, 1, 0, 0, 0, 111, 100, 1,
		0, 0, 0, 111, 101, 1, 0, 0, 0, 111, 102, 1, 0, 0, 0, 111, 103, 1, 0, 0,
		0, 111, 104, 1, 0, 0, 0, 111, 105, 1, 0, 0, 0, 111, 106, 1, 0, 0, 0, 112,
		17, 1, 0, 0, 0, 113, 114, 5, 10, 0, 0, 114, 115, 5, 29, 0, 0, 115, 117,
		5, 2, 0, 0, 116, 118, 3, 20, 10, 0, 117, 116, 1, 0, 0, 0, 118, 119, 1,
		0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 121, 1, 0, 0,
		0, 121, 122, 5, 3, 0, 0, 122, 19, 1, 0, 0, 0, 123, 127, 5, 22, 0, 0, 124,
		127, 5, 23, 0, 0, 125, 127, 3, 22, 11, 0, 126, 123, 1, 0, 0, 0, 126, 124,
		1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 5, 26,
		0, 0, 129, 131, 5, 29, 0, 0, 130, 132, 5, 27, 0, 0, 131, 130, 1, 0, 0,
		0, 131, 132, 1, 0, 0, 0, 132, 21, 1, 0, 0, 0, 133, 134, 5, 11, 0, 0, 134,
		139, 7, 0, 0, 0, 135, 136, 5, 27, 0, 0, 136, 138, 7, 0, 0, 0, 137, 135,
		1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 142, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 143, 5, 9, 0, 0,
		143, 23, 1, 0, 0, 0, 18, 27, 29, 37, 46, 49, 57, 65, 73, 79, 83, 86, 91,
		108, 111, 119, 126, 131, 139,
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
	PacketDslParserUINT8          = 12
	PacketDslParserUINT16         = 13
	PacketDslParserUINT32         = 14
	PacketDslParserUINT64         = 15
	PacketDslParserINT8           = 16
	PacketDslParserINT16          = 17
	PacketDslParserINT32          = 18
	PacketDslParserINT64          = 19
	PacketDslParserFLOAT32        = 20
	PacketDslParserFLOAT64        = 21
	PacketDslParserDIGITS         = 22
	PacketDslParserSTRING         = 23
	PacketDslParserROOT           = 24
	PacketDslParserPACKET         = 25
	PacketDslParserCOLON          = 26
	PacketDslParserCOMMA          = 27
	PacketDslParserSEMICOLON      = 28
	PacketDslParserIDENTIFIER     = 29
	PacketDslParserSTRING_LITERAL = 30
	PacketDslParserLINE_COMMENT   = 31
	PacketDslParserWS             = 32
)

// PacketDslParser rules.
const (
	PacketDslParserRULE_packet              = 0
	PacketDslParserRULE_optionDefinition    = 1
	PacketDslParserRULE_optionDeclaration   = 2
	PacketDslParserRULE_packetDefinition    = 3
	PacketDslParserRULE_fieldDefinition     = 4
	PacketDslParserRULE_metaDataDefinition  = 5
	PacketDslParserRULE_metaDataDeclaration = 6
	PacketDslParserRULE_value               = 7
	PacketDslParserRULE_type                = 8
	PacketDslParserRULE_matchField          = 9
	PacketDslParserRULE_matchPair           = 10
	PacketDslParserRULE_list                = 11
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
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&50331682) != 0 {
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PacketDslParserROOT, PacketDslParserPACKET:
			{
				p.SetState(24)
				p.PacketDefinition()
			}

		case PacketDslParserT__4:
			{
				p.SetState(25)
				p.MetaDataDefinition()
			}

		case PacketDslParserT__0:
			{
				p.SetState(26)
				p.OptionDefinition()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(31)
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
		p.SetState(32)
		p.Match(PacketDslParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PacketDslParserIDENTIFIER {
		{
			p.SetState(34)
			p.OptionDeclaration()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
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
		p.SetState(42)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(PacketDslParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Value()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserSEMICOLON {
		{
			p.SetState(45)
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
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserROOT {
		{
			p.SetState(48)
			p.Match(PacketDslParserROOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(51)
		p.Match(PacketDslParserPACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&541062592) != 0 {
		{
			p.SetState(54)
			p.FieldDefinition()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
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

// IFieldDefinitionContext is an interface to support dynamic dispatch.
type IFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	MetaDataDeclaration() IMetaDataDeclarationContext
	MatchField() IMatchFieldContext

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

func (s *FieldDefinitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *FieldDefinitionContext) MetaDataDeclaration() IMetaDataDeclarationContext {
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

func (s *FieldDefinitionContext) MatchField() IMatchFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMatchFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMatchFieldContext)
}

func (s *FieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitFieldDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) FieldDefinition() (localctx IFieldDefinitionContext) {
	localctx = NewFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PacketDslParserRULE_fieldDefinition)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(PacketDslParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.MetaDataDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.MatchField()
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
	p.EnterRule(localctx, 10, PacketDslParserRULE_metaDataDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(PacketDslParserT__4)
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

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&541061568) != 0 {
		{
			p.SetState(70)
			p.MetaDataDeclaration()
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

// IMetaDataDeclarationContext is an interface to support dynamic dispatch.
type IMetaDataDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Type_() ITypeContext
	STRING_LITERAL() antlr.TerminalNode
	COMMA() antlr.TerminalNode

	// IsMetaDataDeclarationContext differentiates from other interfaces.
	IsMetaDataDeclarationContext()
}

type MetaDataDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *MetaDataDeclarationContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PacketDslParserCOMMA, 0)
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
	p.EnterRule(localctx, 12, PacketDslParserRULE_metaDataDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(78)
			p.Type_()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(81)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserSTRING_LITERAL {
		{
			p.SetState(82)
			p.Match(PacketDslParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserCOMMA {
		{
			p.SetState(85)
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
	p.EnterRule(localctx, 14, PacketDslParserRULE_value)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserT__5, PacketDslParserT__6, PacketDslParserT__7, PacketDslParserUINT8, PacketDslParserUINT16, PacketDslParserUINT32, PacketDslParserUINT64, PacketDslParserINT8, PacketDslParserINT16, PacketDslParserINT32, PacketDslParserINT64, PacketDslParserFLOAT32, PacketDslParserFLOAT64, PacketDslParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Type_()
		}

	case PacketDslParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Match(PacketDslParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserDIGITS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
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
	p.EnterRule(localctx, 16, PacketDslParserRULE_type)
	var _la int

	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(PacketDslParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserUINT8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(PacketDslParserUINT8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserUINT16:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
			p.Match(PacketDslParserUINT16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserUINT32:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(96)
			p.Match(PacketDslParserUINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserUINT64:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(97)
			p.Match(PacketDslParserUINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserINT8:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(98)
			p.Match(PacketDslParserINT8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserINT16:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(99)
			p.Match(PacketDslParserINT16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserINT32:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(100)
			p.Match(PacketDslParserINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserINT64:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(101)
			p.Match(PacketDslParserINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserFLOAT32:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(102)
			p.Match(PacketDslParserFLOAT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserFLOAT64:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(103)
			p.Match(PacketDslParserFLOAT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__5:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(104)
			p.Match(PacketDslParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__6:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(105)
			p.Match(PacketDslParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__7:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(106)
			p.Match(PacketDslParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == PacketDslParserDIGITS {
			{
				p.SetState(107)
				p.Match(PacketDslParserDIGITS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(110)
			p.Match(PacketDslParserT__8)
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

// IMatchFieldContext is an interface to support dynamic dispatch.
type IMatchFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	AllMatchPair() []IMatchPairContext
	MatchPair(i int) IMatchPairContext

	// IsMatchFieldContext differentiates from other interfaces.
	IsMatchFieldContext()
}

type MatchFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMatchFieldContext() *MatchFieldContext {
	var p = new(MatchFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_matchField
	return p
}

func InitEmptyMatchFieldContext(p *MatchFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_matchField
}

func (*MatchFieldContext) IsMatchFieldContext() {}

func NewMatchFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchFieldContext {
	var p = new(MatchFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_matchField

	return p
}

func (s *MatchFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *MatchFieldContext) AllMatchPair() []IMatchPairContext {
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

func (s *MatchFieldContext) MatchPair(i int) IMatchPairContext {
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

func (s *MatchFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitMatchField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) MatchField() (localctx IMatchFieldContext) {
	localctx = NewMatchFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PacketDslParserRULE_matchField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(PacketDslParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Match(PacketDslParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&12584960) != 0) {
		{
			p.SetState(116)
			p.MatchPair()
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
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
	p.EnterRule(localctx, 20, PacketDslParserRULE_matchPair)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserDIGITS:
		{
			p.SetState(123)
			p.Match(PacketDslParserDIGITS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserSTRING:
		{
			p.SetState(124)
			p.Match(PacketDslParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__10:
		{
			p.SetState(125)
			p.List()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(128)
		p.Match(PacketDslParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserCOMMA {
		{
			p.SetState(130)
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
	p.EnterRule(localctx, 22, PacketDslParserRULE_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(PacketDslParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PacketDslParserDIGITS || _la == PacketDslParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PacketDslParserCOMMA {
		{
			p.SetState(135)
			p.Match(PacketDslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PacketDslParserDIGITS || _la == PacketDslParserSTRING) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(142)
		p.Match(PacketDslParserT__8)
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
