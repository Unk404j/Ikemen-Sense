# 🎮 Ikemen Sense

**Ikemen Sense** is not about ownership — it's about pushing the engine forward.  
This fork is a **living laboratory** for exploring new features, documenting internals, and imagining what a modern open-source fighting game engine could be.

If someone more experienced wants to take the lead on technical direction, I’d be happy to support that vision and contribute in other ways.  
The dream: **make Ikemen Sense the most approachable, extensible, and community-driven evolution of Ikemen GO.**

---

## ⚠️ Experimental Playground

This repository is **first and foremost an experimentation sandbox**.  
Features, UI ideas, and code comments here are **prototypes** — they may change, break, or never be merged upstream.

**Nothing here is final.**  
Think of this fork as a place where:
- 🧪 **UI/UX experiments** can be tried safely  
- 🛠️ New ideas can be prototyped without risk to upstream  
- 📚 Documentation can be refined to onboard new contributors  

If you see something interesting, feel free to reuse it or give feedback.

---

## 📌 What is Ikemen Sense?

Ikemen Sense is a **documentation-first fork** of [Ikemen GO](https://github.com/ikemen-engine/Ikemen-GO) (a Go-based clone of M.U.G.E.N).  
The current focus is **commenting, documenting, and clarifying** the engine internals — to make it easier for new contributors and future development.

- ✅ **No functional changes on `main` branch yet**  
- ✅ **All code runs identically to upstream** — only comments and doc improvements are present  

---

## 🧽 Roadmap

| Version | Goal | Status |
|--------|------|--------|
| **v0.10** | 📝 Comment and document core engine | ✅ In progress |
| **v1.00** | 🎮 Add modern gamepad system (SDL2, rumble) | ⏳ Planned |
| **v2.00** | 🧑‍🎨 Modern UI/UX (menus, animations, options) | ⏳ Planned |
| **v3.00** | 🎮 Rich game modes (VS CPU, tag, training) | ⏳ Planned |
| **v4.00** | ⚙️ Accessibility (keybinds, scaling, filters) | ⏳ Planned |
| **v5.00** | 🚀 Future-ready (Android, online, mod/plugin) | 🔮 Concept |

---

## 🎨 UI/UX Experiments

Ikemen Sense will gradually explore **visual and usability upgrades**.  
Below are *concept mockups* for how the engine could evolve visually (not implemented yet).

| Current (vanilla) | Concept (Ikemen Sense) |
|-------------------|-----------------------|
| *(screenshot placeholder)* | *(concept art placeholder)* |

> **Note:** These are experiments — they may be iterated, replaced, or dropped entirely.

---

## ✅ Current State

- 📚 **Commented source code** for better readability  
- 🧠 **Explanations for key structs, functions, and modules**  
- ❌ **No new gameplay/UI features yet** — behavior identical to upstream  
- 🎓 Acts as a **learning resource** for Go, Lua, and Ikemen GO internals  

---

## 💻 Installation & Build

The build process is identical to upstream **Ikemen GO**.

```bash
git clone https://github.com/Unk404j/Ikemen-Sense.git
cd Ikemen-Sense/src
go build -o ../IkemenSense.exe
