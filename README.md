## Summary
- integrate SDL2-based controller rumble for Windows builds
- expose rumble controls to Lua and add helper module with cooldown presets
- bundle SDL2 headers and document manual placement of Windows binaries

## Testing
- `GOOS=windows GOARCH=amd64 go build -tags=windows -o IkemenSense.exe ./src` *(fails: build constraints exclude all Go files in go-gl)*

------
