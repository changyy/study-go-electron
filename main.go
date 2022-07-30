package main

import (
    "github.com/asticode/go-astilectron"
    "github.com/asticode/go-astikit"
    bootstrap "github.com/asticode/go-astilectron-bootstrap"
    "flag"
    "log"
    "fmt"
    "os"
)

var (
    AppName            string
    BuiltAt            string
    VersionAstilectron string
    VersionElectron    string
)

var (
    fs = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
    debug = fs.Bool("d", false, "enables the debug mode")
    developmentMode = fs.Bool("dev", false, "without bootstrap and bundler")
    w *astilectron.Window
)

func main() {
    l := log.New(log.Writer(), log.Prefix(), log.Flags())

    fs.Parse(os.Args[1:])

    if *developmentMode == true {
        a, _ := astilectron.New(log.New(os.Stderr, "", 0), astilectron.Options{
            AppName: "MyGoAstilectronProject",
        })
        defer a.Close()
    
        // Start astilectron
        a.Start()
    
        w, _ := a.NewWindow("./resources/app/index.html", &astilectron.WindowOptions{
            Center: astikit.BoolPtr(true),
            Height: astikit.IntPtr(600),
            Width:  astikit.IntPtr(600),
        })
        w.Create()
    
        // Blocking pattern
        a.Wait()
    } else {
        l.Printf("Running app built at %s\n", BuiltAt)
    
        if err := bootstrap.Run(bootstrap.Options{
            // https://github.com/asticode/go-astilectron-bootstrap/blob/master/run.go#L38
            Asset: Asset,
            AssetDir: AssetDir,
            // https://github.com/asticode/go-astilectron-bootstrap/blob/master/run.go#L49
            RestoreAssets: RestoreAssets,
            AstilectronOptions: astilectron.Options{
                AppName:            AppName,
                SingleInstance:     true,
                VersionAstilectron: VersionAstilectron,
                VersionElectron:    VersionElectron,
            },
            Debug:  *debug,
            Logger: l,
            OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
                w = ws[0]
                return nil
            },
            Windows: []*bootstrap.Window{{
                // https://github.com/asticode/go-astilectron-bootstrap/blob/master/run.go#L65
                Homepage:       "index.html",
                //MessageHandler: handleMessages,
                Options: &astilectron.WindowOptions{
                    //BackgroundColor: astikit.StrPtr("#333"),
                    Center:          astikit.BoolPtr(true),
                    Height:          astikit.IntPtr(700),
                    Width:           astikit.IntPtr(700),
                },
            }},
        }); err != nil {
            l.Fatal(fmt.Errorf("running bootstrap failed: %w", err))
        }
    }
}
