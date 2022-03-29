package fjson

import (
	jsoniter "github.com/json-iterator/go"
)

var ConfigCompatibleWithStandardLibrary = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	UseNumber:              true,
}.Froze()
