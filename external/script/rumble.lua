-- rumble.lua augments the global Rumble table with helper presets.
-- Each helper triggers a short vibration with a built-in cooldown so
-- multiple rapid calls do not stack. Load with `pcall(require, 'rumble')`.

local M = Rumble or {}
local last = 0
local cooldown = 0.12 -- seconds

local function fire(ms)
    local now = os.clock()
    if now - last < cooldown then return end
    last = now
    if M and M.vibrate then
        M.vibrate(ms)
    end
end

--- Light rumble, useful for weak feedback cues.
function M.light()
    fire(40)
end

--- Menu selection confirmation rumble.
function M.menuSelect()
    fire(80)
end

--- Menu back/cancel rumble.
function M.menuBack()
    fire(100)
end

--- Short pulse used for cursor movement.
function M.move()
    fire(30)
end

return M

