# Documentation v1.1.0

## Usage

yab [configs ...]

yab [configs ...] -- [args ...]

Configs are Lua files in your local `.yab` folder or in the global config folder.

### Flags

**--debug**

Enables debug logging.

**--silent**

Disables logging.

## Command Line Arguments

**yab [--version, -v]**

Prints the version of the program.

**yab [--help, -h]**

Prints this help.

**yab --def**

Creates definitions file in global config.

**yab [--update, --upgrade, -u]**

Updates the Yab binary to the latest version.

## Lua API Functions (in the `Yab` global table)

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
* executable `string`

**Returns:** true if the executable is available, false otherwise.

### 𝑓 stdall

*Call a shell command and return the full output (stdout + stderr) in one string.*

**Parameters:** 
* command `string`

**Returns:** The output of the command.

### 𝑓 stdout

*Call a shell command and return the output (stdout) in one string.*

**Parameters:** 
* command `string`

**Returns:** The output of the command.

### 𝑓 stderr

*Call a shell command and return the error output (stderr) in one string.*

**Parameters:** 
* command `string`

**Returns:** The output of the command.

### 𝑓 js_run

*Run a script from the `package.json` file using the first javascript package manager found. Trying pnpm, bun and npm in that order.*

**Parameters:** 
* script `string`

**Returns:** None

### 𝑓 js_install

*Install dependencies from `package.json` using the first javascript package manager found. Trying pnpm, bun and npm in that order.*

**Parameters:** None

**Returns:** None

### 𝑓 git_clone_or_pull

*Clones a git repository to a specified destination. If the repository already exists, it will pull the latest changes instead.*

**Parameters:** 
* url `string`
* destination `string`

**Returns:** None

### 𝑓 zip

*Create a zip file containing the given files.*

**Parameters:** 
* files `table`
* output `string`

**Returns:** None

**Example:**

```lua
Yab.zip(
	{'foo.txt', 'bar.txt', 'baz/'},
	'archive.zip'
)
```

### 𝑓 watch

*Watch file or directory paths for changes and call a function when a change occurs. The callback function will be called with the file path and the event type as arguments. The event type can be one of 'create', 'write', 'remove', 'rename' or 'chmod'.*

**Parameters:** 
* paths `table`
* callback `function(string, string)`

**Returns:** None

**Example:**

```lua
Yab.watch('foo.txt', function(file, event)
	print('foo.txt changed!')
end)
```

### 𝑓 block

*Block the current thread and wait for an interrupt signal.*

**Parameters:** None

**Returns:** None

**Example:**

```lua
Yab.block()
```

### 𝑓 find

*Find files matching a pattern in a directory.*

**Parameters:** 
* pattern `string`

**Returns:** A table containing the matching file paths.

**Example:**

```lua
Yab.find('*.txt')
```

### 𝑓 find

*Find files matching a pattern in a directory.*

**Parameters:** 
* root `string`
* pattern `string`

**Returns:** A table containing the matching file paths.

**Example:**

```lua
Yab.find('foo', '*.txt')
```

### 𝑓 fileinfo

*Get information about a file.*

**Parameters:** 
* path `string`

**Returns:** A table containing the file information (name, size, mode, modtime, isdir, sys). See https://pkg.go.dev/io/fs#FileInfo for details.

**Example:**

```lua
local foo_info = Yab.fileinfo('foo.txt')
print(foo_info.size)
```

