local bin_name = require("build")

os.rename(bin_name, os.getenv("GOPATH") .. "/bin/" .. bin_name)
