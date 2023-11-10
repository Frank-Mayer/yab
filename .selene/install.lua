local bin_name = require("build")
local gopath = os.getenv("GOPATH")
if gopath == nil then
    print("GOPATH is not set")
    os.exit(1)
end

os.rename(bin_name, gopath .. "/bin/" .. bin_name)
