//go:build !windows && !linux && !darwin
// +build !windows,!linux,!darwin

package main

func InitGamepad() error           { return nil }
func CloseGamepad()                {}
func Rumble(durationMs int)        {}
func IsGamepadConnected() bool     { return false }
func HasRumble() bool              { return false }
func ControllerName(id int) string { return "" }
