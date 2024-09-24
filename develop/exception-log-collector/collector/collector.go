package collector

import (
    "bufio"
    "log"
    "os"

    "github.com/fsnotify/fsnotify"
)

type LogCollector struct {
    paths  []string
    output chan string
}

func NewLogCollector(paths []string, output chan string) *LogCollector {
    return &LogCollector{
        paths:  paths,
        output: output,
    }
}

func (lc *LogCollector) Start() {
    for _, path := range lc.paths {
        go lc.watchFile(path)
    }
}

func (lc *LogCollector) watchFile(path string) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Printf("Failed to create watcher: %v", err)
        return
    }
    defer watcher.Close()

    file, err := os.Open(path)
    if err != nil {
        log.Printf("Failed to open file %s: %v", path, err)
        return
    }
    defer file.Close()

    // Move to the end of the file
    file.Seek(0, os.SEEK_END)
    reader := bufio.NewReader(file)

    err = watcher.Add(path)
    if err != nil {
        log.Printf("Failed to add watcher for file %s: %v", path, err)
        return
    }

    for {
        select {
        case event := <-watcher.Events:
            if event.Op&fsnotify.Write == fsnotify.Write {
                for {
                    line, err := reader.ReadString('\n')
                    if err != nil {
                        break
                    }
                    lc.output <- line
                }
            }
        case err := <-watcher.Errors:
            log.Printf("Watcher error: %v", err)
            return
        }
    }
}
