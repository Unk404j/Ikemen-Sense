//go:build !sdl
// +build !sdl

package main

import lua "github.com/yuin/gopher-lua"

// registerRumble is a no-op when SDL support is disabled.
func registerRumble(l *lua.LState) {}
