local bin_name = Yab.os_type() == "windows" and "yab.exe" or "yab"

os.execute('go build -ldflags="-s -w" -o ' .. bin_name .. " cmd/yab/yab.go")

return bin_name
