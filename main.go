//go:generate esc -o generator/static.go -pkg generator tmpl types api

package main

import (
	"github.com/llonchj/godoo/cmd"
)

func main() {
	cmd.Execute()
}
