package main

/*
#include <stdlib.h>  // 需要这个来声明 free()
*/
import "C"
import "github.com/xinchentechnote/fin-protoc/internal/parser"

// FormatPacketDslExport formats the packet DSL string from C char pointer and returns a C char pointer.
// The function takes a C-style string (char*) as input and returns a newly allocated C-style string.
// Caller is responsible for freeing the returned string using C.free().
//
// Parameters:
//
//	dsl - C string containing the packet DSL to format
//
// Returns:
//
//	Formatted DSL as C string, or error message if formatting fails
//
//export FormatPacketDslExport
func FormatPacketDslExport(dsl *C.char) *C.char {
	goDsl := C.GoString(dsl)
	formatted, err := parser.FormatPacketDsl(goDsl)
	if err != nil {
		// rerurn error msg
		return C.CString("Error:" + err.Error())
	}
	return C.CString(formatted)
}

// CompilePacketDslExport compiles the packet DSL string from C char pointer and writes the compiled output to C char pointer.
//
//export CompilePacketDslExport
func CompilePacketDslExport(dsl *C.char, output *C.char, lang C.uint) {
	goDsl := C.GoString(dsl)
	goOutput := C.GoString(output)
	goLang := int(lang)

	println("dsl:", goDsl)
	println("output:", goOutput)
	println("lang:", goLang)
}
