package handler

import (
	"errors"
	"fmt"
	"log"

	"github.com/gordonklaus/portaudio"
	"github.com/micro/go-micro/metadata"
	audiosink "github.com/yeyus/dmr-audio-sink/proto"
	lane "gopkg.in/oleiade/lane.v1"

	"golang.org/x/net/context"
)

// AudioCodecSink Data structure to keep service state variables
type AudioCodecSink struct {
	Buffer *lane.Deque
	Volume float32
	Stream *portaudio.Stream
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
	// TODO remove me
	log.Printf("Dequeuing: %v", a.Buffer.Shift())
	log.Printf("Buffer size: %v\n", a.Buffer.Size())
	// TODO end
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

	log.Printf("DecodedBits: %v\n", req.GetVoiceBits())
	a.Buffer.Append(req.GetVoiceBits())
	log.Printf("Buffer size: %v\n", a.Buffer.Size())
	rsp.DecodedBits = uint32(len(req.GetVoiceBits()))

	return nil
}
