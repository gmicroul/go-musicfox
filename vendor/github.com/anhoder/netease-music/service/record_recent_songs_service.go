package service

import (
	"github.com/anhoder/netease-music/util"
)

type RecordRecentSongsService struct {
	Limit string `json:"limit" form:"limit"`
}

func (service *RecordRecentSongsService) RecordRecentSongs() (float64, []byte) {
	options := &util.Options{
		Crypto: "weapi",
	}
	data := make(map[string]string)
	if service.Limit == "" {
		data["limit"] = "100"
	} else {
		data["limit"] = service.Limit
	}

	code, reBody, _ := util.CreateRequest("POST", `https://music.163.com/api/play-record/song/list`, data, options)

	return code, reBody
}
