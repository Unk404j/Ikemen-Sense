# 🎮 Ikemen Sense

> **Ikemen-Sense is not about ownership — it’s about pushing the engine forward.**  
> If someone more experienced wants to take the lead on technical direction, I’d be more than happy to support that vision and contribute in other ways.  
>  
> **The dream:** make Ikemen-Sense the ultimate open-source fighting game engine of the next decade — modern, extensible, and community-driven.

---

## 📌 What is Ikemen Sense?

**Ikemen Sense** is a modern fork of **Ikemen GO** (a Go-based clone of M.U.G.E.N), aiming to bring this legendary 2D fighting engine to today's standards.

- 🎮 **Native controller support** (Xbox & PlayStation)
- 🕹️ **Gamepad-ready menus** and UI navigation
- ✨ **Modern UX** (feedback, vibration, rumble…)
- ⚙️ **New gameplay modes** and expanded features (coming soon)

---

## 🧽 Roadmap

| Version | Goal | Status |
|--------|------|--------|
| **v1.00** | 🎮 Modern gamepad system (SDL2, rumble) | ✅ In progress |
| **v2.00** | 🧑‍🎨 Modern UI/UX (menus, animations) | ⏳ Upcoming |
| **v3.00** | 🎮 Rich game modes (VS CPU, tag, training) | ⏳ Upcoming |
| **v4.00** | ⚙️ Accessibility (keybinds, UI scaling, filters) | ⏳ Upcoming |
| **v5.00** | 🚀 Future-ready (Android, online, mod/plugin) | 🔮 Concept |

---

## ✅ Current Features

- ✅ SDL2 support & controller detection  
- ✅ Xbox / PlayStation gamepad compatibility  
- ✅ Rumble feedback via `SDL_GameControllerRumble`  
- ✅ Project structure reorganized for modularity  
- 🧲 Lua hook system for gameplay-triggered rumble (WIP)  
- 🧲 Contextual UI elements (Xbox/PS icons in menus – planned)

---

## 💻 Installation & Build

> Requires Go. SDL2 headers are included under `external/SDL2/include`.
> Due to size and licensing reasons the Windows binaries (`SDL2.dll`, `SDL2.lib`)
> are **not** stored in the repository. Download them from the
> [official SDL2 releases](https://github.com/libsdl-org/SDL/releases) and place
> them in `external/SDL2/lib/x64/`. Keep a copy of `SDL2.dll` next to the built
> `IkemenSense.exe`.

```bash
git clone https://github.com/YourGitHubUsername/Ikemen-Sense.git
cd Ikemen-Sense
set CGO_ENABLED=1
go build -tags "windows sdl" -o IkemenSense.exe ./src
```

For detailed steps, check the [original engine wiki](https://github.com/ikemen-engine/Ikemen-GO/wiki).

### Windows Gamepad Rumble

With the SDL2 binaries in place, controller vibration can be triggered with the
following Lua snippet:

```lua
if Rumble.available() and Rumble.hasRumble() then
    Rumble.vibrate(0.5, 200)
end
```

Optional presets are available via `pcall(require, 'rumble')`.

---

## 💡 Use Cases

Ikemen Sense is a **modern and modular base** for:

* Custom fighting games
* Modern reinterpretations of classic M.U.G.E.N
* Game jams & rapid prototyping
* Educational projects using Lua and Go

---

## 🌟 Long-Term Vision

The goal is to turn **Ikemen-Sense** into **the ultimate open-source fighting game engine of the next decade** — a platform that is:

* 💻 **Modern & Extensible** – ready for HD assets, cross-platform builds, and plugins
* 🕹️ **Player-Friendly** – smooth controller support, accessibility options, online play
* 🧑‍💻 **Dev-Friendly** – clean architecture, full documentation, easy onboarding
* 🎮 **Community-Driven** – powered by contributors, modders, and fighting game fans

Ikemen-Sense is not just a fork — it’s a step towards a future where creating a professional-quality 2D fighting game is **as accessible as making a mod**, and where the community drives innovation together.

---

## 🤖 AI-Assisted Development

Ikemen-Sense is partly developed with **AI-assisted coding** (OpenAI Codex / ChatGPT).
AI tools helped prototype features, refactor code, and draft documentation — but every feature is tested, reviewed, and guided by a clear human vision.

As a fun nod to this collaboration, there are plans to add **GPT.exe**, an AI-inspired playable character, as a reference to the project's AI-driven origins.

---

## 🤝 Contributing & Credits

This project is **community-focused**:

* Contributions are welcome — bug fixes, features, documentation, and refactors.
* If a more experienced developer wants to take technical leadership, I will gladly support and contribute in other ways.
* Feedback, issues, and suggestions are encouraged to help shape the future of Ikemen-Sense.

**Credits:**

* The [Ikemen GO contributors](https://github.com/ikemen-engine/Ikemen-GO)
* The M.U.G.E.N community
* Everyone helping to build this fork and push it forward

---

## 📜 License

This project is under the **MIT License** (see [License.txt](./License.txt)).
Some screenpacks, fonts, or assets may fall under other licenses (e.g., CC-BY 3.0).

> ⚠️ **Disclaimer:** I’m still learning programming and open-source collaboration.
> If some assets or attributions are incorrect, I sincerely apologize and will fix them as soon as they are reported.

---

*Project initiated by [Unk404j](https://github.com/Unk404j), driven by curiosity, passion, and guided by AI.*

