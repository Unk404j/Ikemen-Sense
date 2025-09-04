//go:build !windows
// +build !windows

package main

func InitGamepad() error       { return nil }
func CloseGamepad()            {}
func Rumble(durationMs int)    {}
func IsGamepadConnected() bool { return false }
