# Selene

[![Deploy to Pages](https://github.com/Frank-Mayer/selene/actions/workflows/deploy.yml/badge.svg)](https://github.com/Frank-Mayer/selene/actions/workflows/deploy.yml)

Wouldn't it be great if you could use the same build tool for every project?
Regardless of operating system, programming language...

Selene is just that.

Use Lua scripts to define specific actions and execute them from the command line.

**Does that not already exist?**

The build tools Bazel and Gradle served as an example.
However, Gradle is mainly used in the Java ecosystem and Bazel is very complicated.
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

## Quickstart

[Download](https://frank-mayer.github.io/selene/) the latest binary for your os and architecture.

Place the downloaded binary wherever you want.
To be able to use the self-update functionality ensure that access rights aren’t a problem.

Rename the binary to `selene.exe`/`selene` depending on your OS (if you don’t want to alwys type your OS name and architecture).

Ensure that the location of the binary is available in the `PATH` environment variable to make it globally available.

Remember to restart your shell session after changing any environment variables.

Try `selene --version` to check if the istallation was successful.

Use `selene --help` to see the documentation of your installed version.

## Usage

Run one or more sripts:

```bash
selene [configs ...]
```

Pass arguments to the scripts:

```bash
selene [configs ...] -- [args ...]
```

The following folders are searched for configs:

1. `./.selene/`
1. `$XDG_CONFIG_HOME/selene/`
1. `$APPDATA/selene/`
1. `$HOME/.config/selene/`

## Lua definitions

When you initialize a project with `selene --init` a definitions file is created in one of those directories:

1. `$XDG_CONFIG_HOME/selene/`
1. `$APPDATA/selene/`
1. `$HOME/.config/selene/`

## GitHub Actions

```yaml
- name: Setup Selene
  uses: Frank-Mayer/selene-setup@v1.0.0
```

## Download

- Darwin&nbsp;&nbsp; [amd64](https://frank-mayer.github.io/selene/selene-darwin-amd64)&nbsp;&nbsp;&nbsp;[arm64](https://frank-mayer.github.io/selene/selene-darwin-arm64)
- Debian&nbsp;&nbsp; [amd64](https://frank-mayer.github.io/selene/selene-linux-amd64)&nbsp;&nbsp;&nbsp;[arm64](https://frank-mayer.github.io/selene/selene-linux-arm64)
- Windows&nbsp;&nbsp; [amd64](https://frank-mayer.github.io/selene/selene-windows-amd64.exe)&nbsp;&nbsp;&nbsp;[arm64](https://frank-mayer.github.io/selene/selene-windows-arm64.exe)

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
