# 🎮 Ikemen Sense

> **Ikemen-Sense is not about ownership — it’s about pushing the engine forward.**  
> If someone more experienced wants to take the lead on technical direction, I’d be more than happy to support that vision and contribute in other ways.  
>  
> **The dream:** make Ikemen-Sense the ultimate open-source fighting game engine of the next decade — modern, extensible, and community-driven.

---

## 📌 What is Ikemen Sense?

**Ikemen Sense** is a **documentation-first fork** of **Ikemen GO** (a Go-based clone of M.U.G.E.N).  
The current focus is **commenting, documenting, and clarifying the engine internals** — to make it easier for new contributors and future development.

No functional changes are present on the `main` branch yet.  
All code runs identically to upstream — only comments have been added.

---

## 🧽 Roadmap

| Version | Goal | Status |
|--------|------|--------|
| **v0.10** | 📝 Comment and document core engine | ✅ In progress |
| **v1.00** | 🎮 Add modern gamepad system (SDL2, rumble) | ⏳ Planned |
| **v2.00** | 🧑‍🎨 Modern UI/UX (menus, animations) | ⏳ Planned |
| **v3.00** | 🎮 Rich game modes (VS CPU, tag, training) | ⏳ Planned |
| **v4.00** | ⚙️ Accessibility (keybinds, UI scaling, filters) | ⏳ Planned |
| **v5.00** | 🚀 Future-ready (Android, online, mod/plugin) | 🔮 Concept |

---

## ✅ Current State

- ✅ **Commented source code** for better readability  
- ✅ Added explanations for key structs, functions, and modules  
- ❌ No features enabled yet (engine behavior unchanged)  
- 📚 Acts as a **learning resource** for Go, Lua, and Ikemen GO internals

---

## 💻 Installation & Build

> The build process is identical to upstream Ikemen GO.  
> This fork adds no functional changes (yet).

```bash
git clone https://github.com/YourGitHubUsername/Ikemen-Sense.git
cd Ikemen-Sense/src
go build -o ../IkemenSense.exe
