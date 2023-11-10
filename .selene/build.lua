local bin_name = Selene.os_type() == "windows" and "selene.exe" or "selene"

os.execute('go build -ldflags="-s -w" -o ' .. bin_name .. " main/main.go")

return bin_name
