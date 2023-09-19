# README


## TODO

- [x] include webui via CID instead of via git submodule
- [ ] include webui via car file during pre-build hook
    * Blocked by:
        - [ ] https://github.com/ipfs/ipfs-webui/pull/1974
        - [ ] [There are no pre-build hooks with wails](https://github.com/wailsapp/wails/issues/1577)

## Getting started

### Install dependencies

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest # wails
go install github.com/filecoin-project/lassie/cmd/lassie@latest # lassie for downloading webui
go install github.com/ipld/go-car/cmd/car@latest # for extracting contents of car file downloaded from lassie
```

See https://wails.io/docs/gettingstarted/installation

### Run a local version

```bash
wails dev
```

The first run will take a long time since the ipfs/ipfs-webui package has so many dependencies to install.

## Building

To build a redistributable, production mode package, use `wails build`.

## Troubleshooting

### Blank page on app open

wails is weird.. I finally got this figured out by using `npx --yes serve -u build` as the `"frontend:dev:watcher"` setting.

### Live reload not working

You may want to follow https://gist.github.com/int128/e0cdec598c5b3db728ff35758abdbafd as mentioned by https://wails.io/docs/guides/application-development#external-dev-server

### TOO MANY OPEN FILES

Wails seems to leave processes and file handles open.. without a restart, the following worked for a quick fix:
```bash
ulimit -n 10240
```

### `wails dev` hangs on compiling frontend

**Example**
```bash
Building application for development...
  - No Install command. Skipping.
  - Compiling frontend:
```

You can try to do `wails dev -v 2` to see more info, but basically, if you set a non-exiting command (`npm run start`) for the `wails.json` file's `frontend:dev` command, wails will wait, the same as you would in your terminal, when you run that command. You should use a dev oriented build command in that slot, which is basically `npm run build` for the ipfs-webui project.
