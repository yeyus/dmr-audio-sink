package handler

import (
	"errors"
	"fmt"
	"log"

	"github.com/gordonklaus/portaudio"
	"github.com/micro/go-micro/metadata"
	audiosink "github.com/yeyus/dmr-audio-sink/proto"
	lane "gopkg.in/oleiade/lane.v1"

	"github.com/pd0mz/go-dsd"
	"golang.org/x/net/context"
)

// AudioCodecSink Data structure to keep service state variables
type AudioCodecSink struct {
	Buffer      *lane.Deque
	Volume      float32
	Stream      *portaudio.Stream
	DecoderAMBE *dsd.AMBE
}

// SetVolume handler for setting output volume
func (a *AudioCodecSink) SetVolume(ctx context.Context, req *audiosink.SetVolumeRequest, rsp *audiosink.VolumeResponse) error {
	md, _ := metadata.FromContext(ctx)
	log.Printf("Received AudioCodecSink.SetVolume request with metadata: %v", md)
	a.Volume = req.Volume
	rsp.Volume = req.Volume
	return nil
}

// GetVolume handler for getting current output volume
func (a *AudioCodecSink) GetVolume(ctx context.Context, req *audiosink.GetVolumeRequest, rsp *audiosink.VolumeResponse) error {
	log.Printf("Received AudioCodecSink.GetVolume")
	rsp.Volume = a.Volume
	return nil
}

// Decode handler for received encoded audio frames and enqueuing to the buffer
func (a *AudioCodecSink) Decode(ctx context.Context, req *audiosink.DecodeRequest, rsp *audiosink.DecodeResponse) error {
	log.Printf("Receiver AudioCodecSink.Decode request")

	if req.GetCodec() != audiosink.DecodeRequest_AMBE3600X2450 {
		rsp.Error = "Unsupported codec mode"
		return errors.New(rsp.Error)
	}

	if req.GetClientId() == 0 {
		rsp.Error = "Invalid client id"
		return errors.New(rsp.Error)
	}

	if len(req.GetVoiceBits())%216 != 0 {
		rsp.Error = fmt.Sprintf("Unsupported frame size: %d", len(req.GetVoiceBits()))
		return errors.New(rsp.Error)
	}

	newSamples, err := a.decodeSamples(req.GetCodec(), req.GetVoiceBits())
	if err != nil {
		rsp.Error = err.Error()
	}

	log.Printf("new samples: %d, buffer size: %v\n", newSamples, a.Buffer.Size())
	rsp.DecodedBits = uint32(newSamples)

	return nil
}

// GetSamples will get you either maxSamples or maximum sample count available
// in the samples buffer
func (a *AudioCodecSink) GetSamples(maxSamples int) []float32 {
	sampleCount := min(maxSamples, a.Buffer.Size())

	out := make([]float32, sampleCount)
	for i := 0; i < sampleCount; i++ {
		out[i] = a.Buffer.Shift().(float32)
	}

	return out
}

func (a *AudioCodecSink) decodeSamples(codec audiosink.DecodeRequest_Codec, data []byte) (int, error) {
	sampleCount := 0

	if codec != audiosink.DecodeRequest_AMBE3600X2400 {
		for i := 0; i < len(data); i = i + 216 {
			decoded, err := a.DecoderAMBE.Decode(data[i : i+216])
			if err != nil {
				return sampleCount, err
			}

			for j := 0; j < len(decoded); j++ {
				a.Buffer.Append(decoded[j])
				sampleCount++
			}
		}
	} else {
		return 0, errors.New("Unsupported codec mode")
	}

	log.Printf("Inserted %d new samples", sampleCount)
	return sampleCount, nil
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
