# Documentation

## Usage:

selene [configs ...]

selene [configs ...] -- [args ...]

## Command Line Arguments

**selene [--version, -v]**

Prints the version of the program.

**selene [--help, -h]**

Prints this help.

**selene --init**

Initializes a new Selene project.

**selene [--update, --upgrade, -u]**

Updates the Selene binary to the latest version.

## Lua API Functions (in the `Selene` global table)

### 𝑓 os_type

*Returns the operating system type.*

**Parameters:** None

**Returns:** "windows", "linux" or "darwin" on the respective system.

### 𝑓 os_arch

*Returns the operating system architecture.*

**Parameters:** None

**Returns:** "amd64" or "arm64" on the respective system.

### 𝑓 args

*Returns the command line arguments passed to the program.*

**Parameters:** None

**Returns:** A table containing the command line arguments.

### 𝑓 check_exec

*Checks if an executable is available in the system's PATH.*

**Parameters:** 
* executable: *string*

**Returns:** true if the executable is available, false otherwise.

### 𝑓 stdall

*Call a shell command and return the full output (stdout + stderr) in one string.*

**Parameters:** 
* command: *string*

**Returns:** The output of the command.

### 𝑓 stdout

*Call a shell command and return the output (stdout) in one string.*

**Parameters:** 
* command: *string*

**Returns:** The output of the command.

### 𝑓 stderr

*Call a shell command and return the error output (stderr) in one string.*

**Parameters:** 
* command: *string*

**Returns:** The output of the command.

### 𝑓 js_run

*Run a script from the `package.json` file using the first javascript package manager found. Trying pnpm, yarn, bun and npm in that order.*

**Parameters:** 
* script: *string*

**Returns:** true if a javascript package manager was found, false otherwise.

### 𝑓 js_install

*Install dependencies from `package.json` using the first javascript package manager found. Trying pnpm, yarn, bun and npm in that order.*

**Parameters:** None

**Returns:** true if a javascript package manager was found, false otherwise.

