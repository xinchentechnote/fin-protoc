### this is the simple roadmap for packet dsl

- packet dsl vscode extension
- [ ] implement format in go, use ffi and wasm or bin exec

## v0.1.x

- [x] support basic types: uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64, char[n],string
- [x] support match grammar,inner object field, repeat field(array), object field
- [x] support attributes @lengthOf, @calculatedFrom

## v0.2.x

- [x] support attributes @leftPad, @rightPad for fixed length string(char[n])

## v0.3.0

- [ ] support @tag attribute for fields
- [ ] support fix or step protocols
