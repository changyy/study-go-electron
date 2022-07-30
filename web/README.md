# Web development

## Setup

```
% source env_nvm.sh 
Now using node v16.15.1 (npm v8.11.0)
```

## Deploy

```
% npm run build

> study-go-electron@1.0.0 build
> webpack --mode production

asset bundle.js 10.4 KiB [emitted] [compared for emit] [minimized] (name: main)
asset index.html 161 bytes [compared for emit]
runtime modules 26.5 KiB 9 modules
./src/js/index.js 52 bytes [built] [code generated]
webpack 5.74.0 compiled successfully in 1070 ms
```

## Development

```
% npm run start

> study-go-electron@1.0.0 start
> webpack serve --mode development --open
```

## Coding Style

```
% npm init @eslint/config
✔ How would you like to use ESLint? · style
✔ What type of modules does your project use? · esm
✔ Which framework does your project use? · none
✔ Does your project use TypeScript? · No / Yes
✔ Where does your code run? · browser, node
✔ How would you like to define a style for your project? · guide
✔ Which style guide do you want to follow? · airbnb
✔ What format do you want your config file to be in? · JavaScript
Checking peerDependencies of eslint-config-airbnb-base@latest
The config that you've selected requires the following dependencies:

eslint-config-airbnb-base@latest eslint@^7.32.0 || ^8.2.0 eslint-plugin-import@^2.25.2
✔ Would you like to install them now? · No / Yes
✔ Which package manager do you want to use? · npm
Installing eslint-config-airbnb-base@latest, eslint@^7.32.0 || ^8.2.0, eslint-plugin-import@^2.25.2
```
