//go:build !sdl
// +build !sdl

// sdl_input_stub.go provides no-op implementations of the gamepad rumble
// helpers when the sdl build tag is not enabled. The real implementation
// lives in sdl_input.go and uses SDL2. These stubs allow the rest of the
// engine to compile without conditional checks.
package main

// InitGamepad performs no initialization when SDL support is disabled.
// Return value: always nil since no work is done.
func InitGamepad() error { return nil }

// CloseGamepad is a no-op placeholder when SDL support is disabled.
func CloseGamepad() {}

// IsGamepadConnected always reports false when SDL support is disabled.
func IsGamepadConnected() bool { return false }

// HasRumble always reports false when SDL support is disabled.
func HasRumble() bool { return false }

// ControllerName returns an empty string when SDL support is disabled.
func ControllerName() string { return "" }

// Rumble performs no action when SDL support is disabled.
func Rumble(intensity float64, ms int) {}
