-- Hardened rumble wrapper for Go bindings ---------------------

local M = {}

local last = 0
local cooldown = 150 -- ms

local function connected()
    return GamepadIsConnected and GamepadIsConnected()
end

local function hasRumble()
    return GamepadHasRumble and GamepadHasRumble()
end

function M.available()
    return connected() and hasRumble()
end

function M.hasRumble()
    return hasRumble()
end

function M.controller_name()
    if GamepadControllerName then
        return GamepadControllerName() or ""
    end
    return ""
end

local function doRumble(ms)
    local now = os.clock() * 1000
    if now - last < cooldown then return end
    last = now
    if RumbleGamepad then RumbleGamepad(ms) end
end

function M.vibrate(ms)
    if M.available() then
        doRumble(ms or 500)
    end
end

function M.menuSelect() if M.available() then doRumble(150) end end
function M.menuBack()   if M.available() then doRumble(200) end end
function M.light()      if M.available() then doRumble(300) end end
function M.heavy()      if M.available() then doRumble(800) end end

return M
