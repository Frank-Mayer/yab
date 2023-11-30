# Selene

[![Deploy to Pages](https://github.com/Frank-Mayer/selene/actions/workflows/deploy.yml/badge.svg)](https://github.com/Frank-Mayer/selene/actions/workflows/deploy.yml)

Wouldn't it be great if you could use the same build tool for every project?
Regardless of operating system, programming language...

Selene is just that.

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
            <td>Selene</td>
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
            <td>:white_check_mark:</td>
            <td>:white_check_mark:</td>
            <td>:x:</td>
            <td>:x:</td>
            <td>:white_check_mark:</td>
        </tr>
    </tbody>
</table>

Lua is a common and performant language.
Selene offers some useful functions in addition to the Lua standard library that might be useful when building configurations.

Looking for an example configuration?
Take a look at [this projects `.selene` folder](https://github.com/Frank-Mayer/selene/tree/main/.selene).

## Installation

### Download prebuild

https://frank-mayer.github.io/selene/

### Install using Go

```bash
go install github.com/Frank-Mayer/selene/cmd/selene@latest
```

## Docs

Documentation is in the [DOCS.md](https://github.com/Frank-Mayer/selene/blob/main/DOCS.md) file.

## Usage

Run one or more configs:

```bash
selene [configs ...]
```

Pass arguments to the scripts:

```bash
selene [configs ...] -- [args ...]
```

A config is a lua file inside the config directory.

The following directories are used as configs (first found wins)

1. `./.selene/`
1. `$XDG_CONFIG_HOME/selene/`
1. `$APPDATA/selene/`
1. `$HOME/.config/selene/`

## Lua definitions

Run `selene --def` to create a definitions file in your global config directory.
Use this to configure your Lua language server.

Global config is one of those directories:

1. `$XDG_CONFIG_HOME/selene/`
1. `$APPDATA/selene/`
1. `$HOME/.config/selene/`

## GitHub Actions

```yaml
- name: Setup Selene
  uses: Frank-Mayer/selene-setup@v1.0.0
```

## Badge

[![Selene Project](https://img.shields.io/badge/Selene_Project-2C2D72?logo=lua)](https://github.com/Frank-Mayer/selene)

```markdown
[![Selene Project](https://img.shields.io/badge/Selene_Project-2C2D72?logo=lua)](https://github.com/Frank-Mayer/selene)
```

## Etymology

<a href="https://en.wikipedia.org/wiki/Selene#/media/File:Clipeus_Selene_Terme.jpg">
<img src="https://upload.wikimedia.org/wikipedia/commons/0/04/Clipeus_Selene_Terme.jpg" alt="Clipeus Selene Terme" align="right" style="height: 6em; float: right;" />
</a>

Lua is Portuguese for moon. Selene is the Greek goddess of the moon.

> In ancient Greek mythology and religion, Selene (/sɪˈliːniː/; Greek: Σελήνη pronounced [selɛ̌ːnɛː] seh-LEH-neh, meaning "Moon") is the goddess and personification of the Moon.
> Also known as Mene, she is traditionally the daughter of the Titans Hyperion and Theia, and sister of the sun god Helios and the dawn goddess Eos.
> She drives her moon chariot across the heavens.
>
> https://en.wikipedia.org/wiki/Selene
