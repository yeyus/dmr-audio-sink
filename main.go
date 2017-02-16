package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	audio "github.com/yeyus/dmr-audio-sink/audio"
	"github.com/yeyus/dmr-audio-sink/handler"
	audiosink "github.com/yeyus/dmr-audio-sink/proto"
	lane "gopkg.in/oleiade/lane.v1"
)

func main() {
	// setup command line usage
	cmd.Init()

	service := micro.NewService(
		micro.Name("com.ea7jmf.dmr.srv.audiocodecsink"),
	)

	service.Init()
	instance := handler.AudioCodecSink{
		Buffer: lane.NewDeque(),
		Volume: 1.0,
	}

	audio.Initialize(&instance)

	audiosink.RegisterAudioCodecSinkHandler(service.Server(), &instance)

	// TODO: is subscriber needed?

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
