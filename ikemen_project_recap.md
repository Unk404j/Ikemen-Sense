# Projet IkemenSens - Récapitulatif & To-Do

## 📅 Contexte du projet

Tu travailles sur une version forkée d'Ikemen GO pour créer un jeu de combat 2D personnalisé. Le but est de :

- Créer une expérience cohérente et old school
- Intégrer une gestion native des vibrations via SDL2
- Ajouter une UI avec boutons PlayStation / Xbox
- Avoir un menu personnalisé, des stages et persos customisés
- Distribuer une version prête à jouer multiplateforme (PC d'abord, Android peut-être)

---

## 🔢 Étapes réalisées

- [x] Fork Ikemen GO ✅
- [x] Compilation réussie avec MSYS2 + Go + Mingw64 ✅
- [x] Création d'une arborescence de projet (scripts, data, etc) ✅
- [x] Extraction de personnages (Sol Badguy, Scorpion, Kung Fu Man) ✅
- [x] Analyse des fichiers .def, .cns, .cmd, .air, .sff, .snd ✅
- [x] Ajout de stages, palettes, UI, scripts Lua ✅
- [x] Test d'intégration de vibrations via SDL2 (en cours de planification) ✅
- [x] Compilation propre (reclone du repo sans les build écrasés) ✅

---

## 📈 To-Do List (next steps)

### 🔹 SDL2 - Support des vibrations
- [ ] Fork SDL2 pour ajouter support XInput/DInput
- [ ] Hook dans le moteur Go pour déclencher vibrations
- [ ] Liaison avec Lua : créer une fonction `vibrate(strength, duration)`
- [ ] Test sur un hit : appeler vibration lors d'une collision

### 🔹 UI PlayStation / Xbox
- [ ] Créer assets (sprites ou shaders) pour boutons PS/Xbox
- [ ] Détection de type de manette (via SDL)
- [ ] Affichage dynamique des touches (menu + HUD in-game)

### 🔹 Fonction VS CPU
- [ ] Ajouter un mode "Player VS CPU" hors arcade
- [ ] Créer interface de sélection custom (menu)
- [ ] Ajout script Lua pour gérer les entrées et lancer un match

### 🔹 Outils / scripts
- [ ] Script pour ajouter automatiquement les stages dans `select.def`
- [ ] Script .bat de génération d'arborescence de projet
- [ ] Auto-ajout de nouveaux persos au roster avec prévisualisation

### 🔹 Divers / idées
- [ ] Port Android (long terme)
- [ ] Mode "Spectateur" ou "Replay"
- [ ] Gestion tactile pour Android (overlay de boutons)
- [ ] Patch de compatibilité manettes retro (PS2/USB, etc.)

---

## 🔍 Suggestions de noms de fork

- IkemenSens ✨ (DualSense + Sens = Sensation, cohérent avec la vibe vibration/manette)
- IkemenCore
- IkemenReloaded
- FightEngineX
- MUGENSense

---

## 💡 Conseils

- ✅ Prends le temps de documenter ton code, surtout les scripts Lua
- ✅ Versionne tout avec Git (même le .def si tu modifies beaucoup)
- ✅ Ne cherche pas à tout faire en même temps. Priorise : SDL2 > UI > Game Modes

Quand tu veux attaquer une section, je peux t'accompagner ligne par ligne.

