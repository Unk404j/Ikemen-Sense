//go:build windows && sdl
// +build windows,sdl

// sdl_input.go provides gamepad rumble support on Windows using SDL2.
// It initializes the SDL game controller and haptic subsystems and exposes
// a small set of helper functions for Lua scripts. All functions are
// defensive: if no controller is connected or the device lacks rumble
// capability, calls simply return without side effects.
package main

/*
#cgo windows CFLAGS: -I./external/SDL2/include
#cgo windows LDFLAGS: -L./external/SDL2/lib/x64 -lSDL2
#include "SDL.h"
#include "SDL_haptic.h"
#include "SDL_gamecontroller.h"
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"sync"
	"unsafe"
)

// Internal SDL handles protected by padLock so calls from multiple
// goroutines remain safe.
var (
	gc              *C.SDL_GameController // active controller instance
	hap             *C.SDL_Haptic         // haptic device used for fallback rumble
	rumbleSupported bool                  // cached capability flag
	padLock         sync.Mutex            // guards access to controller state
)

// setHint wraps SDL_SetHint while freeing intermediate C strings.
func setHint(name, value string) {
	cname := C.CString(name)
	cvalue := C.CString(value)
	C.SDL_SetHint(cname, cvalue)
	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(cvalue))
}

// InitGamepad prepares SDL's controller and haptic subsystems.
//
// Returns: error if SDL fails to initialize or a controller cannot be opened.
// Common issues: missing SDL2.dll beside the executable or unsupported
// controllers being connected.
func InitGamepad() error {
	padLock.Lock()
	defer padLock.Unlock()

	if C.SDL_Init(C.SDL_INIT_GAMECONTROLLER|C.SDL_INIT_HAPTIC) != 0 {
		return fmt.Errorf("SDL_Init: %s", C.GoString(C.SDL_GetError()))
	}

	// Force HIDAPI backends for broad device support.
	setHint("SDL_JOYSTICK_HIDAPI", "1")
	setHint("SDL_JOYSTICK_HIDAPI_PS4", "1")
	setHint("SDL_JOYSTICK_HIDAPI_PS5", "1")
	setHint("SDL_JOYSTICK_HIDAPI_XBOX", "1")

	if C.SDL_NumJoysticks() < 1 {
		// No controllers present. Not considered an error.
		return nil
	}

	gc = C.SDL_GameControllerOpen(0)
	if gc == nil {
		return fmt.Errorf("SDL_GameControllerOpen: %s", C.GoString(C.SDL_GetError()))
	}

	// Prefer rumble through the higher level controller API.
	if C.SDL_GameControllerHasRumble(gc) == C.SDL_TRUE {
		rumbleSupported = true
	} else {
		// Fallback to haptic API if available.
		joy := C.SDL_GameControllerGetJoystick(gc)
		hap = C.SDL_HapticOpenFromJoystick(joy)
		if hap != nil && C.SDL_HapticRumbleInit(hap) == 0 {
			rumbleSupported = true
		}
	}

	// Inform the user which controller was detected and whether it rumbles.
	fmt.Printf("Controller: %s\n", ControllerName())
	fmt.Printf("Rumble supported: %v\n", rumbleSupported)
	return nil
}

// CloseGamepad releases SDL resources. Safe to call multiple times.
func CloseGamepad() {
	padLock.Lock()
	defer padLock.Unlock()

	if hap != nil {
		C.SDL_HapticClose(hap)
		hap = nil
	}
	if gc != nil {
		C.SDL_GameControllerClose(gc)
		gc = nil
	}
	C.SDL_QuitSubSystem(C.SDL_INIT_GAMECONTROLLER | C.SDL_INIT_HAPTIC)
}

// IsGamepadConnected reports if a controller is currently open.
func IsGamepadConnected() bool {
	padLock.Lock()
	defer padLock.Unlock()
	return gc != nil
}

// HasRumble returns true when the connected controller provides rumble
// capabilities. Returns false if no controller is present or the device
// lacks vibration support.
func HasRumble() bool {
	padLock.Lock()
	defer padLock.Unlock()
	return rumbleSupported
}

// ControllerName returns the human-readable name reported by SDL for the
// connected controller. Returns an empty string when no controller is open.
func ControllerName() string {
	padLock.Lock()
	defer padLock.Unlock()
	if gc == nil {
		return ""
	}
	return C.GoString(C.SDL_GameControllerName(gc))
}

// Rumble activates the controller motors with the given intensity for the
// specified duration in milliseconds. If rumble is unsupported or no controller
// is connected, the call is ignored.
//
// Parameters:
//
//	intensity - strength in the range [0,1]
//	ms        - duration of the vibration in milliseconds.
//
// Common issues: some third-party drivers (e.g., HidHide/DS4Windows) may
// report controllers without rumble capability. In that case this function
// silently does nothing.
func Rumble(intensity float64, ms int) {
	padLock.Lock()
	defer padLock.Unlock()
	if gc == nil || !rumbleSupported {
		return
	}
	if intensity < 0 {
		intensity = 0
	} else if intensity > 1 {
		intensity = 1
	}
	duration := C.Uint32(ms)
	level := C.Uint16(intensity * 0xffff)
	// Try rumbling through the controller API first.
	if C.SDL_GameControllerRumble(gc, level, level, duration) == 0 {
		return
	}
	// If that fails, fall back to the generic haptic API.
	if hap != nil {
		C.SDL_HapticRumblePlay(hap, C.float(intensity), duration)
	}
}
