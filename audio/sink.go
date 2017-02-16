package audio

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gordonklaus/portaudio"
	"github.com/yeyus/dmr-audio-sink/handler"
)

// Initialize gets the default audio device and sets sample handler and cleanup
// routine for closing devices
func Initialize(instance *handler.AudioCodecSink) {
	portaudio.Initialize()
	stream, err := portaudio.OpenDefaultStream(0, 1, 8000, 800, func(out []float32) {
		//log.Printf("Requested to fill %d * float32 buffer, I have %d samples in buffer \n", len(out), instance.Buffer.Size())
		samples := instance.GetSamples(len(out))
		for i := range out {
			if i < len(samples) {
				out[i] = instance.Volume * samples[i]
			} else {
				// TODO: returning noise for easy debug
				//out[i] = instance.Volume * float32(rand.Float32())
				out[i] = 0
			}
		}
	})
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	go func() {
		for sig := range ch {
			// sig is a ^C, handle it
			log.Printf("Exiting with signal %v ...", sig)
			portaudio.Terminate()
		}
	}()

	stream.Start()

	instance.Stream = stream
}
