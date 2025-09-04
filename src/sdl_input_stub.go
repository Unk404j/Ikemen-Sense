//go:build !windows

// sdl_input_stub.go provides no-op implementations of the gamepad rumble
// helpers on non-Windows platforms. The real implementation lives in
// sdl_input.go and uses SDL2. These stubs allow the rest of the engine to
// compile on other platforms without conditional checks.
package main

// InitGamepad performs no initialization on non-Windows builds.
// Return value: always nil since no work is done.
func InitGamepad() error { return nil }

// CloseGamepad is a no-op placeholder for non-Windows builds.
func CloseGamepad() {}

// IsGamepadConnected always reports false outside Windows.
func IsGamepadConnected() bool { return false }

// HasRumble always reports false outside Windows.
func HasRumble() bool { return false }

// ControllerName returns an empty string outside Windows.
func ControllerName() string { return "" }

// Rumble performs no action on non-Windows builds.
func Rumble(ms int) {}
