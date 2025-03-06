package main

import (
    "fmt"
    "log"
    "os"

    "github.com/a-catgirl-dev/gotta/pkg/cli"
    "github.com/a-catgirl-dev/gotta/pkg/storage"
)

func getConfigDir() string {
    return os.ExpandEnv("$HOME/.local/state/gotta")
}

func ensureConfigDir(configDir string) error {
    return os.MkdirAll(configDir, 0755)
}

func main() {
    configDir := getConfigDir()
    cfgPath := ensureConfigDir(configDir)
    if cfgPath != nil {
        log.Fatal("failed to create ~/.local/state/gotta")
    }
    cfg := fmt.Sprintf("%s/gotta.json", configDir)
    store := storage.NewStorage(cfg)
    cli := cli.NewCLI(*store)
    if err := cli.Run(os.Args[1:]); err != nil {
        log.Fatal(err)
    }
}

