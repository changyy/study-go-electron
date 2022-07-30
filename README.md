# Using go-astilectron-bundler

https://github.com/asticode/go-astilectron-bundler

```
% go get -u github.com/asticode/go-astilectron-bundler/astilectron-bundler
% go install github.com/asticode/go-astilectron-bundler/astilectron-bundler
% rm -rf ./output 
% ~/go/bin/astilectron-bundler cc
% ~/go/bin/astilectron-bundler
% tree output 
output
├── darwin-amd64
│   └── MyGoElectronProject.app
│       └── Contents
│           ├── Info.plist
│           └── MacOS
│               └── MyGoElectronProject
├── linux-amd64
│   └── MyGoElectronProject
└── windows-amd64
    └── MyGoElectronProject.exe

6 directories, 4 files
```

