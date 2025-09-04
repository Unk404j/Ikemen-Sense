-- rumble.lua
-- Interface Lua --> Go pour gérer les vibrations manette

-- Petite abstraction pour le moteur
Rumble = {}

-- Vérifie si le moteur Go expose la fonction (protection au cas où)
local hasRumble = (RumbleGamepad ~= nil)

-- Vibre fort pour les coups puissants
function Rumble.heavy()
    if hasRumble then
        RumbleGamepad(800) -- Durée en ms
    end
end

-- Vibration légère pour les petits coups ou impacts
function Rumble.light()
    if hasRumble then
        RumbleGamepad(300)
    end
end

-- Vibration lors d'une sélection dans le menu
function Rumble.menuSelect()
    if hasRumble then
        RumbleGamepad(150)
    end
end

-- Vibration lors d'un cancel / retour menu
function Rumble.menuBack()
    if hasRumble then
        RumbleGamepad(200)
    end
end

-- Fonction générique personnalisable
function Rumble.custom(ms)
    if hasRumble then
        RumbleGamepad(ms or 500)
    end
end

-- Pour usage global : ex.
-- Rumble.light() dans un hitdef
-- Rumble.menuSelect() lors d'un son de bouton pressé

return Rumble
