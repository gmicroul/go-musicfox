package player

import (
	"io"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/minimp3"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"
	"github.com/go-musicfox/go-musicfox/pkg/configs"
	"github.com/go-musicfox/go-musicfox/pkg/constants"
	"github.com/pkg/errors"
	minimp3pkg "github.com/tosone/minimp3"
)

func DecodeSong(t SongType, r io.ReadSeekCloser) (streamer beep.StreamSeekCloser, format beep.Format, err error) {
	switch t {
	case Mp3:
		switch configs.ConfigRegistry.PlayerBeepMp3Decoder {
		case constants.BeepMiniMp3Decoder:
			minimp3pkg.BufferSize = 1024 * 50
			streamer, format, err = minimp3.Decode(r)
		default:
			streamer, format, err = mp3.Decode(r)
		}
	case Wav:
		streamer, format, err = wav.Decode(r)
	case Ogg:
		streamer, format, err = vorbis.Decode(r)
	case Flac:
		streamer, format, err = flac.Decode(r)
	default:
		err = errors.Errorf("Unknown song type(%d)", t)
	}
	return
}
