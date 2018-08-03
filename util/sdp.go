package util

import (
	"net"
	"time"

	sdp "github.com/gortc/sdp"
	log "github.com/sirupsen/logrus"
)

func SDP_DecodeMessage(tData []byte) *sdp.Message {
	m := new(sdp.Message)
	session, err := sdp.DecodeSession(tData, nil)
	if err != nil {
		log.Fatal(err)
	}
	decoder := sdp.NewDecoder(session)
	if err := decoder.Decode(m); err != nil {
		log.Error(err)
	}
	if m.Version != 0 {
		log.Error("wat", m.Version)
	}
	return m
}

func SDP_EncodeMessage(m *sdp.Message) []byte {
	s := make(sdp.Session, 0, 100)
	s = m.Append(s)
	buf := make([]byte, 0, 1024)
	buf = s.AppendTo(buf)
	return buf
}

func SDP_GenerateOffer(name string, address string, val ...interface{}) (*sdp.Message, error) {
	audio := sdp.Media{
		Title: "audiotitle",
		Description: sdp.MediaDescription{
			Type:     "audio",
			Port:     49170,
			Format:   "0",
			Protocol: "RTP/AVP",
		},
		Connection: sdp.ConnectionData{
			NetworkType: "IN",
			AddressType: "IP4",
			IP:          net.ParseIP("224.2.1.1"),
			TTL:         127,
		},
	}
	video := sdp.Media{
		Title: "videotitle",
		Description: sdp.MediaDescription{
			Type:     "video",
			Port:     51372,
			Format:   "99",
			Protocol: "RTP/AVP",
		},
		Bandwidths: sdp.Bandwidths{
			sdp.BandwidthApplicationSpecific: 66781,
		},
		Encryption: sdp.Encryption{
			Method: "prompt",
		},
	}
	video.AddAttribute("rtpmap", "99", "h263-1998/90000")

	m := &sdp.Message{
		Origin: sdp.Origin{
			Username:       "jdoe",
			SessionID:      2890844526,
			SessionVersion: 2890842807,
			Address:        "10.47.16.5",
		},
		Name:  "SDP Seminar",
		Info:  "A Seminar on the session description protocol",
		URI:   "http://www.example.com/seminars/sdp.pdf",
		Email: "j.doe@example.com (Jane Doe)",
		Phone: "12345",
		Connection: sdp.ConnectionData{
			IP:  net.ParseIP("224.2.17.12"),
			TTL: 127,
		},
		Bandwidths: sdp.Bandwidths{
			sdp.BandwidthConferenceTotal: 154798,
		},
		Timing: []sdp.Timing{
			{
				Start:  sdp.NTPToTime(2873397496),
				End:    sdp.NTPToTime(2873404696),
				Repeat: 7 * time.Hour * 24,
				Active: 3600 * time.Second,
				Offsets: []time.Duration{
					0,
					25 * time.Hour,
				},
			},
		},
		TZAdjustments: []sdp.TimeZone{
			{
				sdp.NTPToTime(2882844526),
				time.Hour * -1,
			},
			{
				sdp.NTPToTime(2898848070),
				time.Hour * 0,
			},
		},
		Encryption: sdp.Encryption{
			Method: "clear",
			Key:    "ab8c4df8b8f4as8v8iuy8re",
		},
		Medias: []sdp.Media{audio, video},
	}
	m.AddFlag("recvonly")

	return m, nil
}

func SDP_GenerateAnswer(offer *sdp.Message) (*sdp.Message, error) {
	return nil, nil
}
