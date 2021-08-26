package main

import (
	"fmt"
	"log"
	"math/rand"
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

	var sources []*ndi.SourceT
	for len(sources) == 0 {
		fmt.Println("Looking for sources ...")
		ndi.FindWaitForSources(findInstance, 1000)
		sources = ndi.FindGetCurrentSources(findInstance)
	}

	recvInstance := ndi.RecvCreateV3(ndi.NewRecvCreateV3T())
	if recvInstance == nil {
		return fmt.Errorf("ndi recv_create_v3 failed.")
	}
	defer ndi.RecvDestory(recvInstance)

	ndi.RecvConnect(recvInstance, sources[rand.Intn(len(sources))])

	for start := time.Now(); time.Since(start) < time.Minute; {
		var (
			videoFrame ndi.VideoFrameV2T
			audioFrame ndi.AudioFrameV2T
		)
		switch ndi.RecvCaptureV2(recvInstance, &videoFrame, &audioFrame, nil, 5000) {
		case ndi.FrameTypeNone:
			fmt.Println("No data received.")
			break
		case ndi.FrameTypeVideo:
			fmt.Printf("Video data received (%dx%d).\n", videoFrame.Xres(), videoFrame.Yres())
			ndi.RecvFreeVideoV2(recvInstance, &videoFrame)
			break
		case ndi.FrameTypeAudio:
			fmt.Printf("Audio data received (%d samples).\n", audioFrame.NoSamples())
			ndi.RecvFreeAudioV2(recvInstance, &audioFrame)
			break
		}
	}

	return nil
}
