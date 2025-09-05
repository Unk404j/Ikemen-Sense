
# 📦 Ikemen Sense — Implémentation & Usage de la Vibration Manette (SDL2 Rumble)

## 🔧 Objectif
Implémenter le support de vibration manette via SDL2 (`go-sdl2`) dans Ikemen GO, exposé en Lua via une fonction `f_rumble(player, intensity, duration)`.

---

## à faire
- [x] `go-sdl2` intégré au `go.mod` (dans certains builds)
- [x] Fonction `Rumble()` codée proprement dans `input.go` ou fichier dédié
- [x] Exposition Lua dans `script.go` : `f_rumble(p, strength, time)`
- [x] Intégration clean dans `Config.Input.Rumble`
- [x] Feature testable depuis `main.lua`, `menu.lua`, `fight.lua`, etc.
- [x] Peut être utilisée par les créateurs de personnages M.U.G.E.N / Ikemen

---

## ⚙️ Build Tags (Go)
### 🎯 À quoi ça sert ?
Permet de compiler certaines fonctionnalités uniquement si souhaité.

### 🧱 Exemple :
```go
//go:build sdl
```

### ✅ Commande de compilation :
```bash
go build -tags sdl -o ikemen.exe
```

---

## 🧠 Utilisation en Lua :
```lua
f_rumble(1, 100, 300) -- player 1, intensité 100, durée 300ms
```

---

## 🎮 Exemple en `.cns` (perso) :
```cns
[State 210, Rumble]
type = LuaFunc
trigger1 = movehit = 1
name = "f_rumble"
params = 1, 100, 300
```

---

## 🧩 Possibilités d’utilisation
- Hit = rumble (via `movehit`)
- KO = rumble fort
- Super = rumble long
- GetHit = rumble léger
- Test dans les menus (`menu.lua`, `main.lua`)
- Hook dans `common1.cns` possible

---

## 🧱 Architecture recommandée
- `input_rumble_sdl.go` → implémentation réelle (avec `//go:build sdl`)
- `input_rumble_stub.go` → fallback vide (avec `//go:build !sdl`)
- Utiliser `Config.Input.Rumble` comme flag runtime
- Exposition Lua dans `script.go`

---

## 🛡️ Avantages
- Code propre et modulaire
- Pas de conflit avec GLFW (si bien isolé)
- Compatible avec tous les builds
- Repos propre, maintenable, mod-friendly

---

## 📘 À documenter pour les créateurs :
- Signature Lua
- Restrictions SDL
- Exemples CNS
- Conditions safe d’appel (`movehit`, `triggeronce`, etc.)

---

## 🧪 Tests recommandés :
- `menu.lua` → bouton test vibration
- `fight.lua` → vibration sur hit
- `main.lua` → test au démarrage
- Différentes manettes (PS/Xbox/8bitdo)
