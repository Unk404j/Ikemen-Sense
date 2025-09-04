-- Adds rumble feedback during fights for hits and damage.

if not hook or not hook.add then return end

local lastLife = {}

local function initLife()
    lastLife = {}
    for p = 1, 2 do
        if player(p) then
            lastLife[p] = life()
        end
    end
end

hook.add('launchFight', 'rumble_init_life', function()
    initLife()
end)

hook.add('loop', 'rumble_fight', function()
    if not Rumble or not Rumble.available or not Rumble.available() then return end
    for p = 1, 2 do
        if player(p) then
            if movehit() and Rumble.light then Rumble.light() end
            local l = life()
            if lastLife[p] and l < lastLife[p] and Rumble.heavy then
                Rumble.heavy()
            end
            lastLife[p] = l
        end
    end
end)

return true
