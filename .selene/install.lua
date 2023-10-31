local bin_name = require("build")
local gopath = os.getenv("GOPATH")

os.rename(bin_name, gopath .. "/bin/" .. bin_name)
