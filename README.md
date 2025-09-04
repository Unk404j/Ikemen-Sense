# 🎮 Ikemen Sense

> A modern fork of the open-source fighting game engine **Ikemen GO**, reimagined with controller support, fresh UI/UX design, and an ambitious roadmap to become the ultimate evolution of MUGEN.
>Ikemen-Sense is not about ownership — it’s about pushing the engine forward.
>If someone more experienced wants to take the lead on technical direction, I’d be more than happy to support that vision and contribute in other ways. 
---

## 📌 What is Ikemen Sense?

**Ikemen Sense** is a modern fork of the Ikemen GO engine (a Go-based clone of M.U.G.E.N), with the goal of bringing this legendary 2D fighting engine to today's standards.

- 🎮 Native controller support (Xbox & PlayStation)
- 🕹️ Gamepad-ready menus and UI navigation
- ✨ Modern and responsive UX (feedback, vibration, rumble...)
- ⚙️ New gameplay modes and expanded features (coming soon)

---

# 🧭 Ikemen Sens - Roadmap & Todo

| Version | Goal | Status |
|---------|------|--------|
| `v1.00` | 🎮 Full modern gamepad support (SDL2, auto-detection, rumble) | ✅ IN PROGRESS |
| `v2.00` | 🧑‍🎨 Modern UI/UX (menu, buttons, animations) | ⏳ Upcoming |
| `v3.00` | 🎮 Rich game modes (VS CPU, tag, training, etc.) | ⏳ Upcoming |
| `v4.00` | ⚙️ Settings & accessibility (keybinds, UI scaling, filters) | ⏳ Upcoming |
| `v5.00` | 🚀 Future-ready (Android, online, mod/plugin support) | 🔮 Concept |

## ✅ Features Implemented

- ✅ SDL2 support with gamepad initialization
- ✅ Detection of Xbox / PlayStation controllers
- ✅ Rumble feedback via `SDL_GameControllerRumble`
- ✅ Code structure ready for Xbox/PS UI assets
- 🧪 Lua hook system for gameplay-triggered rumble (WIP)
- 🧪 Planned UI elements for contextual button display


---

## ✅ V0.99 - Fork Prototype

- [x] Go compilation working via MSYS2
- [x] Basic SDL2 integration
- [x] Vibration tested with `SDL_GameControllerRumble`
- [x] Clean fork created (`Ikemen Sens`)
- [x] Project structure organized (build, data, src...)

---

## 🚀 V1.00 – 🎮 Modern Gamepad System

Modern gamepad system: detection, UI and rumble.

- [ ] Detect controller type (Xbox / PlayStation / Other)
- [ ] Auto-enable vibration (with ON/OFF toggle later)
- [ ] Dynamic button mapping (menus: A/X enter, B/O back)
- [ ] Lua script `rumble.lua` for contextual vibration logic
- [ ] Contextual UI (Xbox / PS icons in menus)
- [ ] Unified and stable input API

---

## 🎨 V2.00 – UX / UI Modern Design

Redesign UI and UX with modern 2D fighting game style.

- [ ] Modern title screen (animated 2D illustration)
- [ ] Minimal and responsive UI
- [ ] Light / Dark theme switcher
- [ ] Gamepad navigation in menus
- [ ] Lifebar and victory screen visual overhaul
- [ ] Options menu inspired by recent Mortal Kombat games

---

## 🕹️ V3.00 – Rich Game Modes

More than arcade mode: variety and replay value.

- [ ] Add "Versus" mode (Player vs CPU)
- [ ] Dynamic stage selection
- [ ] Advanced training mode
- [ ] Add main hub / lobby
- [ ] Foundation for future Online play

---

## ⚙️ V4.00 – Advanced Settings / Accessibility

Customization, comfort and accessibility.

- [ ] Graphical settings menu
- [ ] Full button remapping (all input types)
- [ ] Vibration ON/OFF toggle
- [ ] Colorblind / contrast mode
- [ ] Crash recovery safe mode (logs, default settings)

---

## 🌐 V5.00 – Ready for the Future

Portability, compatibility, openness.

- [ ] Android port
- [ ] WebGL port (via Ebitengine)
- [ ] Fight replay & recording
- [ ] Cloud sync (profile, characters, saves)
- [ ] Web integration for voting / favorite characters


## 📦 Installation & Build

> You’ll need Go, SDL2, and a MinGW/MSYS2 environment under Windows.

```bash
# Clone the repo
git clone https://github.com/YourGitHubUsername/Ikemen-Sense.git
cd Ikemen-Sense

# Build
cd src
go build -o ../IkemenSense.exe
```

For detailed instructions, check the [original engine wiki](https://github.com/ikemen-engine/Ikemen-GO/wiki).

---

## 💡 Use Cases

Ikemen Sense is designed as a **modern and modular base** for:

- Custom fighting games
- Modern reinterpretations of classic MUGEN
- Game jams or rapid prototyping of 2D fighters
- Educational projects using Lua and Go

---

## 🔗 Useful Links

- [Ikemen Engine (original repo)](https://github.com/ikemen-engine/Ikemen-GO)
- [Ikemen Wiki & Docs](https://github.com/ikemen-engine/Ikemen-GO/wiki)
- [License (MIT)](./License.txt)

---

## 📜 License

This project is licensed under the **MIT License** (see [License.txt](./License.txt)).

Some screenpacks, fonts, or assets may fall under different licenses.  Certain non-code assets are licensed under CC-BY 3.0.

> ⚠️ **Disclaimer**: I'm still learning and exploring programming. If some assets or attributions are incorrect or incomplete, I sincerely apologize and will correct them upon request or discovery. 

---

*Project initiated by [Unk404j](https://github.com/Unk404j), driven by curiosity, passion, and guided by AI.*

See [License.txt](License.txt) for more details.

This project is an experimental fork by a beginner in game development. If any license has been miscredited, please accept our apologies. We’re working to improve and learn with every step.
