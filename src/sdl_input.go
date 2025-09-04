//go:build windows
// +build windows

package main

/*
#cgo LDFLAGS: -lSDL2
#include <SDL2/SDL.h>

static SDL_GameController* gController = NULL;
static SDL_Haptic*        gHaptic    = NULL;

int InitGamepad() {
    SDL_SetHint(SDL_HINT_JOYSTICK_HIDAPI,            "1");
    SDL_SetHint(SDL_HINT_JOYSTICK_HIDAPI_PS4,        "1");
    SDL_SetHint(SDL_HINT_JOYSTICK_HIDAPI_PS4_RUMBLE, "1");
    SDL_SetHint(SDL_HINT_JOYSTICK_HIDAPI_PS5,        "1");
    SDL_SetHint(SDL_HINT_JOYSTICK_HIDAPI_PS5_RUMBLE, "1");
    SDL_SetHint(SDL_HINT_JOYSTICK_HIDAPI_XBOX,       "1");

    if (SDL_Init(SDL_INIT_GAMECONTROLLER | SDL_INIT_HAPTIC) < 0) return -1;
    if (SDL_NumJoysticks() < 1)                              return -2;

    gController = SDL_GameControllerOpen(0);
    if (!gController) return -3;

    if (SDL_GameControllerHasRumble(gController) == SDL_TRUE) return 0;

    SDL_Joystick* js = SDL_GameControllerGetJoystick(gController);
    if (!js) return -4;
    gHaptic = SDL_HapticOpenFromJoystick(js);
    if (!gHaptic) return -5;
    if (SDL_HapticRumbleInit(gHaptic) != 0) return -6;
    return 1; // using haptic fallback
}

void CloseGamepad() {
    if (gHaptic)    { SDL_HapticClose(gHaptic); gHaptic = NULL; }
    if (gController){ SDL_GameControllerClose(gController); gController = NULL; }
    SDL_QuitSubSystem(SDL_INIT_HAPTIC | SDL_INIT_GAMECONTROLLER);
}

int IsGamepadConnected() { return gController != NULL; }

const char* ControllerName() {
    if (!gController) return "";
    const char* n = SDL_GameControllerName(gController);
    return n ? n : "";
}

int HasRumble() {
    if (gController && SDL_GameControllerHasRumble(gController) == SDL_TRUE) return 1;
    if (gHaptic) return 1;
    return 0;
}

void RumbleGamepad(Uint16 low, Uint16 high, Uint32 ms) {
    if (gController && SDL_GameControllerHasRumble(gController) == SDL_TRUE) {
        SDL_GameControllerRumble(gController, low, high, ms);
    } else if (gHaptic) {
        SDL_HapticRumblePlay(gHaptic, 1.0f, ms);
    }
}
*/
import "C"

import "fmt"

// InitGamepad initializes SDL2 controller + haptic.
func InitGamepad() error {
	switch r := C.InitGamepad(); r {
	case 0:
		fmt.Println("SDL2 gamepad initialized (controller rumble).")
		return nil
	case 1:
		fmt.Println("SDL2 gamepad initialized (haptic fallback).")
		return nil
	case -1:
		return fmt.Errorf("SDL_Init failed")
	case -2:
		return fmt.Errorf("no joysticks detected")
	case -3:
		return fmt.Errorf("failed to open game controller")
	case -4:
		return fmt.Errorf("failed to get joystick from controller")
	case -5:
		return fmt.Errorf("failed to open haptic from joystick")
	case -6:
		return fmt.Errorf("failed to init haptic rumble")
	default:
		return fmt.Errorf("unknown error %d", r)
	}
}

// CloseGamepad shuts down controller/haptic subsystems.
func CloseGamepad() { C.CloseGamepad() }

// Rumble triggers vibration (full strength, ms duration).
func Rumble(ms int) { C.RumbleGamepad(0xFFFF, 0xFFFF, C.Uint32(ms)) }

// IsGamepadConnected reports if a controller is opened.
func IsGamepadConnected() bool { return C.IsGamepadConnected() != 0 }

// HasRumble reports rumble capability.
func HasRumble() bool { return C.HasRumble() != 0 }

// ControllerName returns SDL name.
func ControllerName() string { return C.GoString(C.ControllerName()) }
