local args = Yab.args()
local join = function(t)
    local s = ""
    for _, v in ipairs(t) do
        s = s .. " " .. v
    end
    return s
end
os.execute("go run cmd/yab/yab.go " .. join(args))
