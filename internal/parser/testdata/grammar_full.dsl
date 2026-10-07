// Grammar full-coverage fixture.
// Exercises every PacketDsl grammar rule; the CI job `codegen-verify`
// compiles and runs the generated code for all six target languages,
// so every construct here must stay supported by every generator.
// Line comments are part of the grammar.

// All eight options, including the optional trailing semicolon variant.
options {
	StringPrefixLenType = u16;
	ArrayPrefixLenType = u16;
	LittleEndian = false;
	JavaPackage = "com.grammar.full.messages";
	GoPackage = "messages";
	GoModule = "fin-protoc-codegen-verify";
	FixedStringPadFromLeft = false;
	FixedStringPadChar = ' ';
}

// MetaData block: basic type, char[n], zchar[n], dynamic string and a
// reference to an earlier metadata entry.
MetaData CommonDefs {
	u64 Price `价格`,
	char[12] SecurityID `证券代码`,
	zchar[8] ZCode `零终止字符串`,
	string Note `动态字符串`,
	Price RefPrice `引用元数据`,
}

// Empty packet body (fieldDefinitionWithAttribute* may be empty).
packet EmptyMsg {
}

packet Logon {
	// @leftPad with explicit pad char
	@leftPad('0')
	char[10] UserName `用户名`,
	zchar[8] Token `零终止令牌`,
	string Password `密码`,
	uint64 ClientId `客户端ID`,
	u16 HeartbeatInterval `心跳间隔`,
}

packet StrA {
	char[6] Marker `标记`,
	f32 Ratio `比率`,
}

packet StrB {
	char[6] Marker `标记`,
	f64 Ratio `比率`,
}

packet Item {
	char[4] Code `代码`,
	u32 Qty `数量`,
}

packet InlineItem {
	char[4] NCode `编码`,
	u32 NQty `数量`,
}

root packet GrammarFull {
	// Basic types, short alias forms.
	char Side `买卖方向`,
	u8 F_u8,
	i8 F_i8,
	u16 MsgType `消息类型`,
	i16 F_i16,
	u32 F_u32,
	i32 F_i32,
	u64 F_u64,
	i64 F_i64,
	f32 F_f32,
	f64 F_f64,
	// Basic types, long alias forms.
	uint8 A_u8,
	int8 A_i8,
	uint16 A_u16,
	int16 A_i16,
	uint32 A_u32,
	int32 A_i32,
	uint64 A_u64,
	int64 A_i64,
	float32 A_f32,
	float64 A_f64,
	// Dynamic strings: char[] alias and string.
	char[] CharArr `char[]动态串`,
	string Note `动态字符串`,
	// Metadata-typed fields: metadata name + field name forms.
	Note MetaNote `元数据动态串`,
	Price RefPriceField `元数据引用`,
	// @tag attribute.
	@tag(9)
	u32 Tagged `带tag字段`,
	// Padding attribute variants: explicit space right, NUL right, default left.
	@rightPad(' ')
	char[8] PadSpace `空格右填充`,
	@rightPad('\x00')
	char[6] PadNul `NUL右填充`,
	@leftPad()
	char[5] PadDefaultLeft `默认左填充`,
	// Length field (@lengthOf) targeting a following object field.
	@lengthOf(ItemBlock)
	u16 ItemLen `块长度`,
	Item ItemBlock `对象块`,
	// Repeat: basic type, dynamic string, object reference.
	repeat u32 RepeatU32 `重复整数`,
	repeat string RepeatStr `重复字符串`,
	repeat Item RepeatItems `重复对象`,
	// Repeat inline (nested) object declaration.
	repeat InnerBlk {
		char[4] ICode `内联代码`,
		u32 IQty `内联数量`,
	},
	// Non-repeat inline (nested) object declaration.
	InlineObj {
		char[4] OCode `内联代码`,
		u32 OQty `内联数量`,
	},
	// Object references: with explicit field name and bare form.
	InlineItem NamedObj `命名对象引用`,
	EmptyMsg BareEmpty `空对象引用`,
	// Match field: digits key, list-of-digits key.
	match MsgType as Body {
		1: Logon,
		[2, 3]: EmptyMsg,
	},
	// Match field: string key field with string and list-of-strings keys.
	string SessKind `会话类型`,
	match SessKind as Sess {
		"fast": StrA,
		["slow", "idle"]: StrB,
	},
	// Checksum field with doc string.
	@calculatedFrom("CRC32")
	u32 Checksum `校验和`,
}
