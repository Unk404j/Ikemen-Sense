## Summary
- integrate SDL2-based controller rumble for Windows builds
- expose rumble controls to Lua and add helper module with cooldown presets
- bundle SDL2 headers and document manual placement of Windows binaries

## Testing
- `GOOS=windows GOARCH=amd64 go build -tags=windows -o IkemenSense.exe ./src` *(fails: build constraints exclude all Go files in go-gl)*

------
## Usefull doc for the project :

https://pkg.go.dev/github.com/veandco/go-sdl2/sdl

### Overview :
Package sdl is SDL2 wrapped for Go users. It enables interoperability between Go and the SDL2 library which is written in C. That means the original SDL2 installation is required for this to work. SDL2 is a cross-platform development library designed to provide low level access to audio, keyboard, mouse, joystick, and graphics hardware via OpenGL and Direct3D.
