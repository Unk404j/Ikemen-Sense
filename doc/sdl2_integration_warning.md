
# ⚠️ Pourquoi intégrer go-sdl2 directement n’est pas idéal pour Ikemen Sense

> Excellente question, connard curieux — tu fais bien de creuser.

---

## 🔴 Pourquoi ce n’est pas idéal pour Ikemen Sense (moteur clean et modulaire)

### 1. 🧱 `go-sdl2` est un gros monolithe
- Il t’embarque tout : son, vidéo, events, haptique, joysticks…
- Même si tu n’utilises que le rumble (`sdl.Haptic`), il va linker tout le reste.
- Ikemen utilise déjà GLFW (ou Kinc), donc SDL devient redondant → conflit potentiel.

⚠️ Deux libs qui veulent gérer les manettes, les events, les fenêtres = conflit possible.

---

### 2. 🤯 Incompatibilités potentielles / conflits runtime
- SDL peut initialiser ses propres devices HID (gamepads, haptics), qui entrent en conflit avec ceux de GLFW.
- Problèmes possibles :
  - SDL capture la manette en exclu,
  - GLFW détecte 0 device,
  - SDL ne relâche pas les ressources à la fermeture du jeu.

🎮 Résultat : bugs d’input ou vibration en boucle pour certains périphériques (PS5/Xbox…).

---

### 3. 🧩 Ça casse la logique modulaire des build tags Go
- Ikemen est conçu pour des backends swappables :
  - `+build glfw`
  - `+build kinc`
  - `+build sdl` (potentiellement)
- Si tu mets `go-sdl2` en dur dans `input.go`, tu brises cette modularité.

---

## ✅ Comment l’utiliser proprement malgré tout

### 🔨 En cloisonnant

1. Crée un fichier `input_rumble_sdl.go` → `+build sdl`
2. Contient : `InitHaptic`, `Rumble`, etc.
3. Crée un `input_rumble_stub.go` → `+build !sdl`
4. Implémente une interface `Rumbler` (ou fonctions vides par défaut)

---

### 🧠 Exemple minimal

```go
// input_rumble_sdl.go
//go:build sdl

package input

import "github.com/veandco/go-sdl2/sdl"

func InitRumble(device int) {
    // init SDL haptic here
}
```

```go
// input_rumble_stub.go
//go:build !sdl

package input

func InitRumble(device int) {
    // do nothing
}
```

---

## 🧠 TL;DR — Comparatif

| Critère                    | Couplage dur à SDL | Module SDL isolé |
|----------------------------|--------------------|-------------------|
| Simplicité                 | ✅ Facile           | ❌ Faut structurer |
| Compatibilité moteur       | ❌ Fragile          | ✅ Safe            |
| Cross-platform             | ❌ Risqué           | ✅ Géré via tags   |
| Code propre / maintenable | ❌ Nope             | ✅ Yep             |
| Style Ikemen Sense         | ❌ Triche           | ✅ Classe          |

---

Tu veux que je te prépare ce setup modulaire avec les deux fichiers `input_rumble_sdl.go` et `input_rumble_stub.go` tout prêts à brancher dans ton fork ?
