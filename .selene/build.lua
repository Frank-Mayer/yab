local bin_name = selene_os_type() == "windows" and "selene.exe" or "selene"

os.execute("go build -o " .. bin_name .. " main/main.go")

return bin_name
