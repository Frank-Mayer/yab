# selene

<a href="https://en.wikipedia.org/wiki/Selene#/media/File:Clipeus_Selene_Terme.jpg">
<img src="https://upload.wikimedia.org/wikipedia/commons/0/04/Clipeus_Selene_Terme.jpg" alt="Clipeus Selene Terme" align="right" style="height: 6em; float: right;" />
</a>

Selene is a versatile and user-friendly build tool configured in Lua, designed to streamline your development workflow. Whether you're compiling, packaging, or automating tasks, Selene empowers you to define and manage your build process with ease. Say goodbye to complex configuration files and hello to the simplicity and flexibility of Lua. Get started today and unlock the full potential of your projects!

Looking for an example? Take a look at [this projects `.selene` folder](https://github.com/Frank-Mayer/selene/tree/main/.selene).

## Docs

Documentation is in the [DOCS.md](https://github.com/Frank-Mayer/selene/blob/main/DOCS.md) file.

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
