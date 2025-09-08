# 📦 Ikemen Sense — Gamepad Vibration Implementation & Usage (SDL2 Rumble)

## 🔧 Goal
Implement gamepad vibration support via SDL2 (`go-sdl2`) in Ikemen GO, exposed in Lua through a function `f_rumble(player, intensity, duration)`.

---

## To-Do
- [x] `go-sdl2` integrated into `go.mod` (in some builds)
- [x] `Rumble()` function properly coded in `input.go` or a dedicated file
- [x] Lua exposure in `script.go`: `f_rumble(p, strength, time)`
- [x] Clean integration into `Config.Input.Rumble`
- [x] Feature testable from `main.lua`, `menu.lua`, `fight.lua`, etc.
- [x] Can be used by M.U.G.E.N / Ikemen character creators

---

## 🧩 Patch: Add SDL2 + Extra Dependencies to go.mod

This patch introduces SDL2 support (`go-sdl2`), updates `golang.org/x/text` and adds `gopkg.in/ini.v1` as direct dependencies.

To apply manually:

```go
+	github.com/veandco/go-sdl2 v0.4.10
