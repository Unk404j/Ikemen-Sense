# 🎮 Ikemen Sense

> A modern fork of the open-source fighting game engine **Ikemen GO**, reimagined with controller support, fresh UI/UX design, and an ambitious roadmap to become the ultimate evolution of MUGEN.

---

## 📌 What is Ikemen Sense?

**Ikemen Sense** is a modern fork of the Ikemen GO engine (a Go-based clone of M.U.G.E.N), with the goal of bringing this legendary 2D fighting engine to today's standards.

- 🎮 Native controller support (Xbox & PlayStation)
- 🕹️ Gamepad-ready menus and UI navigation
- ✨ Modern and responsive UX (feedback, vibration, rumble...)
- ⚙️ New gameplay modes and expanded features (coming soon)
- 💻 Stable base for new games like **Crossworld**

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


## ✅ V0.99 - Prototype de Fork

- [x] Compilation Go opérationnelle via MSYS2
- [x] Intégration de base SDL2
- [x] Vibration testée avec `SDL_GameControllerRumble`
- [x] Création d'un fork propre (`Ikemen Sens`)
- [x] Structure du projet organisée (build, data, src...)

---

## 🚀 V1.00 – 🎮 Modern Gamepad System

Système de manettes moderne, auto-détection, UI et rumble.

- [ ] Détection du type de manette (Xbox / PlayStation / Autre)
- [ ] Activation auto des vibrations (avec option ON/OFF plus tard)
- [ ] Mapping dynamique des boutons (menus : A/X entrer, B/O retour)
- [ ] Script Lua `rumble.lua` pour gestion contextuelle des vibrations
- [ ] Affichage contextuel de l’interface (icônes Xbox / PS dans menus)
- [ ] API d’entrée unifiée et stable

---

## 🎨 V2.00 – UX / UI Design Moderne

Interface utilisateur repensée dans l’esprit des jeux actuels.

- [ ] Écran titre modernisé (illustration 2D dynamique / animé)
- [ ] UI minimaliste et responsive
- [ ] Thèmes clairs/sombres
- [ ] Menus accessibles à la manette (navigation fluide)
- [ ] Refonte visuelle des lifebars et écrans de victoire
- [ ] Menus d’option lisibles et inspirés des Mortal Kombat récents

---

## 🕹️ V3.00 – Modes de Jeu Riches

Plus que l’arcade : du choix, de la variété, de la rejouabilité.

- [ ] Ajout d’un mode *Versus* (Joueur vs CPU)
- [ ] Sélection de stage dynamique
- [ ] Mode *Training* avancé
- [ ] Ajout d’un hub/menu central type « lobby »
- [ ] Base pour *Online play* futur

---

## ⚙️ V4.00 – Paramétrage Avancé / Accessibilité

Personnalisation, accessibilité et options modernes.

- [ ] Menu des options moderne et graphique
- [ ] Remapping complet des touches (tous supports)
- [ ] Activation/désactivation vibration
- [ ] Mode daltonien / contraste
- [ ] Mode sans échec au démarrage si crash (log + safe mode)

---

## 🌐 V5.00 – Ready for the Future

Ouverture, compatibilité, extensibilité.

- [ ] Port Android
- [ ] Port WebGL via Ebitengine
- [ ] Enregistrement/replay de combats
- [ ] Support cloud : profil, persos, saves
- [ ] Intégration possible avec site web pour vote/persos populaires

---

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

- Custom fighting games like **Crossworld**
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
