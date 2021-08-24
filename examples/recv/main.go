package main

import (
	"fmt"
	"log"

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

	var sources []ndi.SourceT
	for len(sources) == 0 {
		fmt.Println("Looking for sources ...")
		ndi.FindWaitForSources(findInstance, 1000)
		sources = ndi.FindGetCurrentSources(findInstance)
	}

	ndi.RecvCreateV3(ndi.NewRecvCreateV3T())
	return nil
}
