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
		"", "'{'", "'}'", "'MetaData'", "':'", "'`'", "'Int64'", "'uInt16'",
		"'uInt32'", "'string'", "'Int32'", "'uInt8'", "'char'", "'char['", "']'",
		"'match'", "'\"'", "': '", "'ROOT'", "'packet'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "ROOT", "PACKET", "IDENTIFIER", "DIGIT", "WS",
	}
	staticData.RuleNames = []string{
		"packet", "packetDefinition", "fieldDefinition", "metaDataDefinition",
		"metaDataDeclaration", "type", "description", "matchField",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 22, 96, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 5, 0, 19, 8, 0, 10,
		0, 12, 0, 22, 9, 0, 1, 1, 3, 1, 25, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1,
		31, 8, 1, 10, 1, 12, 1, 34, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 3, 2, 41,
		8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 4, 3, 47, 8, 3, 11, 3, 12, 3, 48, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 69, 8, 5, 10, 5, 12, 5, 72, 9, 5, 1, 5,
		3, 5, 75, 8, 5, 1, 6, 4, 6, 78, 8, 6, 11, 6, 12, 6, 79, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 4, 7, 88, 8, 7, 11, 7, 12, 7, 89, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 1, 1, 0, 5, 5, 104,
		0, 20, 1, 0, 0, 0, 2, 24, 1, 0, 0, 0, 4, 40, 1, 0, 0, 0, 6, 42, 1, 0, 0,
		0, 8, 52, 1, 0, 0, 0, 10, 74, 1, 0, 0, 0, 12, 77, 1, 0, 0, 0, 14, 81, 1,
		0, 0, 0, 16, 19, 3, 2, 1, 0, 17, 19, 3, 6, 3, 0, 18, 16, 1, 0, 0, 0, 18,
		17, 1, 0, 0, 0, 19, 22, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0, 20, 21, 1, 0, 0,
		0, 21, 1, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 23, 25, 5, 18, 0, 0, 24, 23,
		1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 27, 5, 19, 0, 0,
		27, 28, 5, 20, 0, 0, 28, 32, 5, 1, 0, 0, 29, 31, 3, 4, 2, 0, 30, 29, 1,
		0, 0, 0, 31, 34, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33,
		35, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 35, 36, 5, 2, 0, 0, 36, 3, 1, 0, 0,
		0, 37, 41, 5, 20, 0, 0, 38, 41, 3, 8, 4, 0, 39, 41, 3, 14, 7, 0, 40, 37,
		1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 39, 1, 0, 0, 0, 41, 5, 1, 0, 0, 0,
		42, 43, 5, 3, 0, 0, 43, 44, 5, 20, 0, 0, 44, 46, 5, 1, 0, 0, 45, 47, 3,
		8, 4, 0, 46, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48,
		49, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 5, 2, 0, 0, 51, 7, 1, 0, 0,
		0, 52, 53, 3, 10, 5, 0, 53, 54, 5, 20, 0, 0, 54, 55, 5, 4, 0, 0, 55, 56,
		5, 5, 0, 0, 56, 57, 3, 12, 6, 0, 57, 58, 5, 5, 0, 0, 58, 9, 1, 0, 0, 0,
		59, 75, 5, 6, 0, 0, 60, 75, 5, 7, 0, 0, 61, 75, 5, 8, 0, 0, 62, 75, 5,
		9, 0, 0, 63, 75, 5, 10, 0, 0, 64, 75, 5, 11, 0, 0, 65, 75, 5, 12, 0, 0,
		66, 70, 5, 13, 0, 0, 67, 69, 5, 21, 0, 0, 68, 67, 1, 0, 0, 0, 69, 72, 1,
		0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 73, 1, 0, 0, 0, 72,
		70, 1, 0, 0, 0, 73, 75, 5, 14, 0, 0, 74, 59, 1, 0, 0, 0, 74, 60, 1, 0,
		0, 0, 74, 61, 1, 0, 0, 0, 74, 62, 1, 0, 0, 0, 74, 63, 1, 0, 0, 0, 74, 64,
		1, 0, 0, 0, 74, 65, 1, 0, 0, 0, 74, 66, 1, 0, 0, 0, 75, 11, 1, 0, 0, 0,
		76, 78, 8, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 77, 1,
		0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 13, 1, 0, 0, 0, 81, 82, 5, 15, 0, 0, 82,
		83, 5, 20, 0, 0, 83, 87, 5, 1, 0, 0, 84, 85, 5, 16, 0, 0, 85, 86, 5, 20,
		0, 0, 86, 88, 5, 16, 0, 0, 87, 84, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89,
		87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 5, 17,
		0, 0, 92, 93, 5, 20, 0, 0, 93, 94, 5, 2, 0, 0, 94, 15, 1, 0, 0, 0, 10,
		18, 20, 24, 32, 40, 48, 70, 74, 79, 89,
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
	PacketDslParserEOF        = antlr.TokenEOF
	PacketDslParserT__0       = 1
	PacketDslParserT__1       = 2
	PacketDslParserT__2       = 3
	PacketDslParserT__3       = 4
	PacketDslParserT__4       = 5
	PacketDslParserT__5       = 6
	PacketDslParserT__6       = 7
	PacketDslParserT__7       = 8
	PacketDslParserT__8       = 9
	PacketDslParserT__9       = 10
	PacketDslParserT__10      = 11
	PacketDslParserT__11      = 12
	PacketDslParserT__12      = 13
	PacketDslParserT__13      = 14
	PacketDslParserT__14      = 15
	PacketDslParserT__15      = 16
	PacketDslParserT__16      = 17
	PacketDslParserROOT       = 18
	PacketDslParserPACKET     = 19
	PacketDslParserIDENTIFIER = 20
	PacketDslParserDIGIT      = 21
	PacketDslParserWS         = 22
)

// PacketDslParser rules.
const (
	PacketDslParserRULE_packet              = 0
	PacketDslParserRULE_packetDefinition    = 1
	PacketDslParserRULE_fieldDefinition     = 2
	PacketDslParserRULE_metaDataDefinition  = 3
	PacketDslParserRULE_metaDataDeclaration = 4
	PacketDslParserRULE_type                = 5
	PacketDslParserRULE_description         = 6
	PacketDslParserRULE_matchField          = 7
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
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&786440) != 0 {
		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case PacketDslParserROOT, PacketDslParserPACKET:
			{
				p.SetState(16)
				p.PacketDefinition()
			}

		case PacketDslParserT__2:
			{
				p.SetState(17)
				p.MetaDataDefinition()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(22)
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
	p.EnterRule(localctx, 2, PacketDslParserRULE_packetDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PacketDslParserROOT {
		{
			p.SetState(23)
			p.Match(PacketDslParserROOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(26)
		p.Match(PacketDslParserPACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.Match(PacketDslParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1097664) != 0 {
		{
			p.SetState(29)
			p.FieldDefinition()
		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(35)
		p.Match(PacketDslParserT__1)
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
	p.EnterRule(localctx, 4, PacketDslParserRULE_fieldDefinition)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.Match(PacketDslParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__5, PacketDslParserT__6, PacketDslParserT__7, PacketDslParserT__8, PacketDslParserT__9, PacketDslParserT__10, PacketDslParserT__11, PacketDslParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.MetaDataDeclaration()
		}

	case PacketDslParserT__14:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.MatchField()
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
	p.EnterRule(localctx, 6, PacketDslParserRULE_metaDataDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(PacketDslParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(PacketDslParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16320) != 0) {
		{
			p.SetState(45)
			p.MetaDataDeclaration()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(PacketDslParserT__1)
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
	Type_() ITypeContext
	IDENTIFIER() antlr.TerminalNode
	Description() IDescriptionContext

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

func (s *MetaDataDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, 0)
}

func (s *MetaDataDeclarationContext) Description() IDescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
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
	p.EnterRule(localctx, 8, PacketDslParserRULE_metaDataDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Type_()
	}
	{
		p.SetState(53)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Match(PacketDslParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Match(PacketDslParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(56)
		p.Description()
	}
	{
		p.SetState(57)
		p.Match(PacketDslParserT__4)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDIGIT() []antlr.TerminalNode
	DIGIT(i int) antlr.TerminalNode

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

func (s *TypeContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(PacketDslParserDIGIT)
}

func (s *TypeContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(PacketDslParserDIGIT, i)
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
	p.EnterRule(localctx, 10, PacketDslParserRULE_type)
	var _la int

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PacketDslParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Match(PacketDslParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Match(PacketDslParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.Match(PacketDslParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__8:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(62)
			p.Match(PacketDslParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__9:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(63)
			p.Match(PacketDslParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__10:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(64)
			p.Match(PacketDslParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__11:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(65)
			p.Match(PacketDslParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PacketDslParserT__12:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(66)
			p.Match(PacketDslParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PacketDslParserDIGIT {
			{
				p.SetState(67)
				p.Match(PacketDslParserDIGIT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(73)
			p.Match(PacketDslParserT__13)
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

// IDescriptionContext is an interface to support dynamic dispatch.
type IDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDescriptionContext differentiates from other interfaces.
	IsDescriptionContext()
}

type DescriptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionContext() *DescriptionContext {
	var p = new(DescriptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_description
	return p
}

func InitEmptyDescriptionContext(p *DescriptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PacketDslParserRULE_description
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PacketDslParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }
func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PacketDslVisitor:
		return t.VisitDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PacketDslParser) Description() (localctx IDescriptionContext) {
	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PacketDslParserRULE_description)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8388574) != 0) {
		{
			p.SetState(76)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == PacketDslParserT__4 {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

		p.SetState(79)
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

// IMatchFieldContext is an interface to support dynamic dispatch.
type IMatchFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode

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

func (s *MatchFieldContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PacketDslParserIDENTIFIER)
}

func (s *MatchFieldContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PacketDslParserIDENTIFIER, i)
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
	p.EnterRule(localctx, 14, PacketDslParserRULE_matchField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(PacketDslParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(PacketDslParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PacketDslParserT__15 {
		{
			p.SetState(84)
			p.Match(PacketDslParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Match(PacketDslParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.Match(PacketDslParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
		p.Match(PacketDslParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(PacketDslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(PacketDslParserT__1)
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
