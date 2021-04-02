package youtube

import "github.com/barakmich/gochromecast/primitives"

//GetScreenIDRequest for getting a screen ID for an existing youtube application.
type GetScreenIDRequest struct {
	primitives.PayloadHeaders
	ScreenID int `json:"screen_ids"`
}
