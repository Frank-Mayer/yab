# Yab

Yet another build tool

[![Deploy to Pages](https://github.com/Frank-Mayer/yab/actions/workflows/deploy.yml/badge.svg)](https://github.com/Frank-Mayer/yab/actions/workflows/deploy.yml)

Wouldn't it be great if you could use the same build tool for every project?
Regardless of operating system, programming language...

Yab is just that.

Use Lua scripts to define specific actions and execute them from the command line.

**Does that not already exist?**

No!

<table>
    <thead>
        <tr>
            <td></td>
            <td><sup>Heavily used</sup></td>
            <td><sup>Builtin support for many technologies</sup></td>
            <td><sup>Easy to setup and extend</sup></td>
            <td><sup>Basic syntax (loops, functions, ...)</sup></td>
            <td><sup>Parameters</sup></td>
            <td><sup>No domain specific language</sup></td>
            <td><sup>Cross-platform by default</sup></td>
            <td><sup>Does not make the codebase messy</sup></td>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Yab</td>
            <td>:x:</td>
            <td>:x:</td>
            <td>:white_check_mark:</td>
            <td>:white_check_mark:</td>
            <td>:white_check_mark:</td>
            <td>:white_check_mark:</td>
            <td>:white_check_mark:</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>Bazel</td>
            <td>:x:</td>
            <td>:white_check_mark:</td>
            <td>:x:</td>
            <td>:x:</td>
            <td>:x:</td>
            <td>:x:</td>
            <td>:white_check_mark:</td>
            <td>:x:</td>
        </tr>
        <tr>
            <td>Gradle</td>
            <td>:white_check_mark:</td>
            <td>:x:</td>
            <td>:x:</td>
            <td>:white_check_mark:</td>
            <td>:x:</td>
            <td>:x: / :white_check_mark:</td>
            <td>:white_check_mark:</td>
            <td>:white_check_mark:</td>
        </tr>
        <tr>
            <td>Make</td>
            <td>:white_check_mark:</td>
            <td>:x:</td>
            <td>:white_check_mark:</td>
            <td>:x:</td>
            <td>:white_check_mark:</td>
            <td>:x:</td>
            <td>:x:</td>
            <td>:white_check_mark:</td>
        </tr>
    </tbody>
</table>

Lua is a common and performant language.
Yab offers some useful functions in addition to the Lua standard library that might be useful when building configurations.

Looking for an example configuration?
Take a look at [this projects `.yab` folder](https://github.com/Frank-Mayer/yab/tree/main/.yab).

## Installation

### Download prebuild

https://frank-mayer.github.io/yab/

### Install using Go

```bash
go install github.com/Frank-Mayer/yab/cmd/yab@latest
```

## Docs

Documentation is in the [DOCS.md](https://github.com/Frank-Mayer/yab/blob/main/DOCS.md) file.

## Usage

Run one or more configs:

```bash
yab [configs ...]
```

Pass arguments to the scripts:

```bash
yab [configs ...] -- [args ...]
```

A config is a lua file inside the config directory.

The following directories are used as configs (first found wins)

1. `./.yab/`
1. `$XDG_CONFIG_HOME/yab/`
1. `$APPDATA/yab/`
1. `$HOME/.config/yab/`

## Lua definitions

Run `yab --def` to create a definitions file in your global config directory.
Use this to configure your Lua language server.

Global config is one of those directories:

1. `$XDG_CONFIG_HOME/yab/`
1. `$APPDATA/yab/`
1. `$HOME/.config/yab/`

## GitHub Actions

```yaml
- name: Setup Yab
  uses: Frank-Mayer/yab-setup@v1.0.0
```

## Badge

[![Yab Project](https://img.shields.io/badge/Yab_Project-2C2D72?logo=lua)](https://github.com/Frank-Mayer/yab)

```markdown
[![Yab Project](https://img.shields.io/badge/Yab_Project-2C2D72?logo=lua)](https://github.com/Frank-Mayer/yab)
```
