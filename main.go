package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/kilianp07/AthleteIQBox/transmitter"
	"github.com/kilianp07/AthleteIQBox/transmitter/data"
	"github.com/kilianp07/AthleteIQBox/utils"
)

func main() {
	var (
		conf        *transmitter.Conf
		err         error
		server      *transmitter.Transmitter
		wg          = &sync.WaitGroup{}
		errChan     = make(chan error)
		successChan = make(chan bool)
		ctx, cancel = context.WithCancel(context.Background())
		osSignal    = make(chan os.Signal)
	)

	// Read configuration
	if err = utils.ReadJSONFile("box.json", &conf); err != nil {
		fmt.Println("Failed to read configuration:", err)
		os.Exit(1)
	}

	// Initialize transmitter
	server, err = transmitter.New(conf)
	if err != nil {
		fmt.Println("Failed to initialize transmitter: ", err)
		os.Exit(1)
	}

	// Watch for a CTRL+C signal
	signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for sig := range osSignal {
			fmt.Println("Received signal: ", sig)
			fmt.Println("Shutting down...")
			cancel()
		}
	}()

	// Start bluetooth server
	wg.Add(1)
	go server.Start(wg, errChan, successChan, ctx)
	select {
	case err := <-errChan:
		fmt.Println("Error starting server: ", err)
		os.Exit(1)
	case <-successChan:
		fmt.Println("Server started")
	}

	// Update data
	wg.Add(1)
	go update(ctx, server, wg)

	wg.Wait()

}

func update(ctx context.Context, server *transmitter.Transmitter, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context done")
			return
		case <-ticker.C:
			values := data.Position{
				Latitude:  rand.Float64(),
				Longitude: rand.Float64(),
			}

			if err := server.Update("position", values); err != nil {
				fmt.Println("Error updating position with values", values, ":", err)
				return
			}
		}
	}
}
