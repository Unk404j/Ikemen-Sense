-- rumble.lua augments the global Rumble table with helper presets.
-- Each helper triggers a short vibration with a built-in cooldown so
-- multiple rapid calls do not stack. Load with `pcall(require, 'rumble')`.

local M = Rumble or {}
local last = 0
local cooldown = 0.12 -- seconds

local function fire(intensity, ms)
    local now = os.clock()
    if now - last < cooldown then return end
    last = now
    if M and M.vibrate then
        M.vibrate(intensity, ms)
    end
end

--- Light rumble, useful for weak feedback cues.
function M.light()
    fire(0.3, 40)
end

--- Menu selection confirmation rumble.
function M.menuSelect()
    fire(0.5, 80)
end

--- Menu back/cancel rumble.
function M.menuBack()
    fire(0.5, 100)
end

--- Short pulse used for cursor movement.
function M.move()
    fire(0.4, 30)
end

return M

