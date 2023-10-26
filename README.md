# selene

<img src="https://upload.wikimedia.org/wikipedia/commons/0/04/Clipeus_Selene_Terme.jpg" alt="Clipeus Selene Terme" align="right" style="height: 6em; float: right;" />

Selene is a versatile and user-friendly build tool configured in Lua, designed to streamline your development workflow. Whether you're compiling, packaging, or automating tasks, Selene empowers you to define and manage your build process with ease. Say goodbye to complex configuration files and hello to the simplicity and flexibility of Lua. Get started today and unlock the full potential of your projects!

## Usage

```bash
selene build
```

Checks of one of the following files exist and runs it:
1. `./selene/build.lua`
1. `$XDG_CONFIG_HOME/selene/build.lua`
1. `$APPDATA/selene/build.lua`

```bash
selene first second third
```

You can specify multiple arguments, which will be executed in the order in which they appear.

```bash
selene
```

If you dont provide any argument `init` is the default.
If there is no init file, it will be created.

## Download

- [Windows amd64](https://frank-mayer.github.io/selene/selene-windows-amd64.exe)
- [Darwin amd64](https://frank-mayer.github.io/selene/selene-darwin-amd64)
