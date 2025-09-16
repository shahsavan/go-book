//go:build generate
// +build generate

package main

// This file is not built into the app.
// It is used with `go generate ./schemas/avro`.

import (
	_ "github.com/actgardner/gogen-avro/v10/cmd" // Avro generator
)

//go:generate go run github.com/actgardner/gogen-avro/v10/cmd@latest -package avro assignment_created.avsc
