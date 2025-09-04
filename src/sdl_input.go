//go:build windows
// +build windows

package main

/*
#cgo LDFLAGS: -lSDL2
#include <SDL2/SDL.h>

SDL_GameController* gController = NULL;

int InitGamepad() {
    if (SDL_Init(SDL_INIT_GAMECONTROLLER) < 0) {
        return -1;
    }
    if (SDL_NumJoysticks() < 1) {
        return -2;
    }
    gController = SDL_GameControllerOpen(0);
    if (gController == NULL) {
        return -3;
    }
    return 0;
}

void CloseGamepad() {
    if (gController != NULL) {
        SDL_GameControllerClose(gController);
        gController = NULL;
    }
    SDL_QuitSubSystem(SDL_INIT_GAMECONTROLLER);
}

int IsGamepadConnected() {
    return gController != NULL;
}

void RumbleGamepad(Uint16 lowFreq, Uint16 highFreq, Uint32 durationMs) {
    if (gController != NULL) {
        SDL_GameControllerRumble(gController, lowFreq, highFreq, durationMs);
    }
}
*/
import "C"

import (
	"fmt"
)

// InitGamepad initializes the SDL2 game controller.
func InitGamepad() error {
	result := C.InitGamepad()
	switch result {
	case 0:
		fmt.Println("SDL2 Gamepad initialized.")
		return nil
	case -1:
		return fmt.Errorf("SDL_Init(SDL_INIT_GAMECONTROLLER) failed")
	case -2:
		return fmt.Errorf("No joysticks detected")
	case -3:
		return fmt.Errorf("Failed to open game controller")
	default:
		return fmt.Errorf("Unknown error: %d", result)
	}
}

// CloseGamepad properly closes the SDL2 game controller.
func CloseGamepad() {
	C.CloseGamepad()
}

// Rumble triggers a vibration (in milliseconds).
func Rumble(durationMs int) {
	C.RumbleGamepad(0xFFFF, 0xFFFF, C.Uint32(durationMs))
}

// IsGamepadConnected returns true if an SDL game controller is open.
func IsGamepadConnected() bool {
	return C.IsGamepadConnected() != 0
}
