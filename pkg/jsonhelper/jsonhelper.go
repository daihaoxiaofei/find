package jsonhelper

import (
	jsonIterator "github.com/json-iterator/go"
)

var (
	IJson = jsonIterator.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by common package.
	Marshal = IJson.Marshal
	// Unmarshal is exported by common package.
	Unmarshal = IJson.Unmarshal
	// MarshalIndent is exported by common package.
	MarshalIndent = IJson.MarshalIndent
	// NewDecoder is exported by common package.
	NewDecoder = IJson.NewDecoder
	// NewEncoder is exported by common package.
	NewEncoder = IJson.NewEncoder
)
