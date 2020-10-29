package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adrianosela/guardduty/cache"
	"github.com/adrianosela/guardduty/generator"
	"github.com/adrianosela/guardduty/store"
)

func main() {
	db := store.NewFileSystemStore("./categorization.json")
	cache := cache.NewMemoryCache(db, time.Second*5)
	stream := generator.Stream(100, time.Millisecond*500)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case finding, open := <-stream:
			if !open {
				return
			}
			severity, ok := cache.GetSeverity(finding)
			if !ok {
				severity = "[UNKNOWN]"
			}
			log.Printf("Processed %s severity finding %s", severity, finding)
		case <-sig:
			return
		}
	}

}
