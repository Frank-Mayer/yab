# Selene

<a href="https://en.wikipedia.org/wiki/Selene#/media/File:Clipeus_Selene_Terme.jpg">
<img src="https://upload.wikimedia.org/wikipedia/commons/0/04/Clipeus_Selene_Terme.jpg" alt="Clipeus Selene Terme" align="right" style="height: 6em; float: right;" />
</a>

Wouldn't it be great if you could use the same build tool for every project?
Regardless of operating system, programming language...

Selene is just that.

Use Lua scripts to define specific actions and execute them from the command line.

**Does that not already exist?**

The build tools Bazel and Gradle served as an example.
However, Gradle is only used in the Java ecosystem and Bazel is very complicated.
Both use a domain-specific language, which complicates familiarization and makes it difficult to find help.

Lua is a common and performant language that is used, for example, to configure [Neovim](https://github.com/neovim) or build World of Warcraft Mods.

Looking for an example configuration? Take a look at [this projects `.selene` folder](https://github.com/Frank-Mayer/selene/tree/main/.selene).

## Docs

Documentation is in the [DOCS.md](https://github.com/Frank-Mayer/selene/blob/main/DOCS.md) file.

## Badge

[![Selene Project](https://img.shields.io/badge/Selene_Project-2C2D72?logo=lua)](https://github.com/Frank-Mayer/selene)

```markdown
[![Selene Project](https://img.shields.io/badge/Selene_Project-2C2D72?logo=lua)](https://github.com/Frank-Mayer/selene)
```

## Usage

Run a specified sript:

```bash
selene build
```

Checks if one of the following files exist and runs it:

1. `./.selene/build.lua`
1. `$XDG_CONFIG_HOME/selene/build.lua`
1. `$APPDATA/selene/build.lua`

You can specify multiple arguments, which will be executed in the order in which they appear:

```bash
selene first second third
```

## Download

- [Windows amd64](https://frank-mayer.github.io/selene/selene-windows-amd64.exe)
- [Darwin amd64](https://frank-mayer.github.io/selene/selene-darwin-amd64)
- [Debian amd64](https://frank-mayer.github.io/selene/selene-debian-amd64)
