package youtube

import (
	"net/url"
	"strconv"

	"github.com/barakmich/gochromecast/generic"
	"github.com/imroc/req"
)

const actionSetPlaylist = "setPlaylist"

//components for the initializeQueueRequest parameters and body
const listIDKey = "_listId"
const actionKey = "__sc"
const currentTimeKey = "_currentTime"
const currentIndexKey = "_currentIndex"
const audioOnlyKey = "_audioOnly"
const countKey = "count"
const videoIDKey = "_videoId"

const defaultTime = "0"
const defaultIndex = -1
const defaultAudioOnlySetting = "false"
const defaultCount = 1

type initializeQueueRequestParams struct {
	VideoID             string
	ListID              string
	LoungeID            string
	RequestCount        int
	SessionRequestCount int
	SessionID           string
	GSessionID          string
}

func createInitializeQueueRequest(params initializeQueueRequestParams) generic.RequestComponents {
	requestCount := params.RequestCount
	header := req.Header{
		loungeIDHeader: params.LoungeID,
	}

	for k, v := range defaultHeaders {
		header[k] = v
	}

	reqParams := req.Param{
		sessionIDKey:  params.SessionID,
		gSessionIDKey: params.GSessionID,
		requestIDKey:  requestCount,
		versionKey:    bindVersion,
		cVersionKey:   bindCVersion,
	}
	index := strconv.Itoa(defaultIndex)
	count := strconv.Itoa(defaultCount)
	body := map[string][]string{
		listIDKey:       {params.ListID},
		actionKey:       {actionSetPlaylist},
		currentTimeKey:  {defaultTime},
		currentIndexKey: {index},
		audioOnlyKey:    {defaultAudioOnlySetting},
		videoIDKey:      {params.VideoID},
		countKey:        {count},
	}
	formattedBody := FormatSessionParameters(body, params.SessionRequestCount)
	return generic.RequestComponents{
		URL:    bindURL,
		Header: header,
		Params: reqParams,
		Body:   url.Values(formattedBody).Encode(),
	}

}
