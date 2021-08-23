package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mix3/ndi-go"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if !ndi.Initialize() {
		return fmt.Errorf("ndi initialize failed.")
	}
	defer ndi.Destroy()

	findInstance := ndi.FindCreateV2(ndi.NewFindCreateT(
		ndi.NewFindCreateTOptionShowLocalSources(true),
	))
	if findInstance == nil {
		return fmt.Errorf("ndi find_create_v2 failed.")
	}
	defer ndi.FindDestroy(findInstance)

	for start := time.Now(); time.Now().Sub(start) < time.Minute; {
		if !ndi.FindWaitForSources(findInstance, 5000) {
			fmt.Println("No change to the sources found.")
			continue
		}

		sources := ndi.FindGetCurrentSources(findInstance)
		fmt.Printf("Network sources (%d found).\n", len(sources))
		for i, v := range sources {
			fmt.Printf("%d. %s\n", i+1, v.Name())
		}
	}
	return nil
}
