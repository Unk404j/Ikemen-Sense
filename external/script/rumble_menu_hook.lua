-- external/script/rumble_menu_hook.lua
-- Adds menu & options rumble via Ikemen's hook system, without editing menu.lua/options.lua directly.

local function safe_call(fn)
    if not fn then return false end
    local ok, res = pcall(fn)
    return ok and res
end

-- Simple edge guard to avoid multi-triggering the same input many times per frame.
local last = { up=false, down=false, left=false, right=false, ok=false, back=false }

local function detect_and_rumble(section)
    -- Ensure Rumble is available
    if not Rumble or not Rumble.available or not Rumble.available() then return end

    -- Cancel / Back
    if safe_call(esc) or main.f_input(main.t_players, {'m'}) then
        if not last.back then
            if Rumble.menuBack then Rumble.menuBack() end
            last = { up=false, down=false, left=false, right=false, ok=false, back=true }
        end
        return
    end

    -- OK / Validate
    if main.f_input(main.t_players, {'pal', 's'}) then
        if not last.ok then
            if Rumble.menuSelect then Rumble.menuSelect() end
            last = { up=false, down=false, left=false, right=false, ok=true, back=false }
        end
        return
    end

    -- Horizontal move (value change)
    if main.f_input(main.t_players, {'$F'}) then
        if not last.right then
            if Rumble.light then Rumble.light() end
            last = { up=false, down=false, left=false, right=true, ok=false, back=false }
        end
        return
    elseif main.f_input(main.t_players, {'$B'}) then
        if not last.left then
            if Rumble.light then Rumble.light() end
            last = { up=false, down=false, left=true, right=false, ok=false, back=false }
        end
        return
    end

    -- Vertical move (cursor up/down). Some menus only use U/D.
    if main.f_input(main.t_players, {'$U'}) then
        if not last.up then
            if Rumble.light then Rumble.light() end
            last = { up=true, down=false, left=false, right=false, ok=false, back=false }
        end
        return
    elseif main.f_input(main.t_players, {'$D'}) then
        if not last.down then
            if Rumble.light then Rumble.light() end
            last = { up=false, down=true, left=false, right=false, ok=false, back=false }
        end
        return
    end

    -- No relevant input this frame: reset
    last = { up=false, down=false, left=false, right=false, ok=false, back=false }
end

-- Hook into both menu and options loop iterations
if hook and hook.add then
    hook.add('menu.menu.loop', 'rumble_menu_nav', function(section, tbl)
        detect_and_rumble(section)
    end)
    hook.add('options.menu.loop', 'rumble_options_nav', function(section, tbl)
        detect_and_rumble(section)
    end)
end

return true
