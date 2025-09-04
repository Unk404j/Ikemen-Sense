//go:build !windows
// +build !windows

package main

func InitGamepad() error                              { return nil }
func CloseGamepad()                                   {}
func Rumble(durationMs int)                           {}
func IsGamepadConnected() bool                        { return false }
func HasRumble() bool                                 { return false }
func ControllerName() string                          { return "" }
func AddSDLGamepadMappingsFromFile(path string) error { return nil }
