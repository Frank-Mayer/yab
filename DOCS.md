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

## Lua API Functions

**selene_os_type**

Returns the operating system type.

Parameters: None

Returns: "windows" or "unix" on the respective system.

**selene_args**

Returns the command line arguments passed to the program.

Parameters: None

Returns: A table containing the command line arguments.

