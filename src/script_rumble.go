//go:build sdl
// +build sdl

package main

import lua "github.com/yuin/gopher-lua"

// registerRumble exposes controller vibration helpers to Lua scripts.
func registerRumble(l *lua.LState) {
	rumble := l.NewTable()
	l.SetGlobal("Rumble", rumble)

	// Rumble.vibrate(intensity, ms)
	l.SetField(rumble, "vibrate", l.NewFunction(func(l *lua.LState) int {
		intensity := float64(numArg(l, 1))
		ms := int(numArg(l, 2))
		if HasRumble() {
			Rumble(intensity, ms)
		}
		return 0
	}))

	// Rumble.available() -> bool
	l.SetField(rumble, "available", l.NewFunction(func(l *lua.LState) int {
		l.Push(lua.LBool(IsGamepadConnected()))
		return 1
	}))

	// Rumble.hasRumble() -> bool
	l.SetField(rumble, "hasRumble", l.NewFunction(func(l *lua.LState) int {
		l.Push(lua.LBool(HasRumble()))
		return 1
	}))

	// Rumble.controller_name() -> string
	l.SetField(rumble, "controller_name", l.NewFunction(func(l *lua.LState) int {
		l.Push(lua.LString(ControllerName()))
		return 1
	}))
}
