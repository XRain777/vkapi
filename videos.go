package vkapi

// ===============================
//  VideoSave
// ===============================

// VideoSaveParams параметры метода VideoSave.
type VideoSaveParams struct {
	Name           string
	Description    string
	IsPrivate      bool
	WallPost       bool
	Link           string
	GroupID        int
	AlbumID        int
	PrivacyView    string
	PrivacyComment string
	NoComments     bool
	Repeat         bool
	Compression    bool
}

// VideoSaveResp структура, возвращаемая методом VideoSave.
type VideoSaveResp struct {
	AccessKey   string `json:"access_key"`
	Description string `json:"description"`
	OwnerID     int    `json:"owner_id"`
	Title       string `json:"title"`
	UploadURL   string `json:"upload_url"`
	VideoID     int    `json:"video_id"`
}

// VideoSave Возвращает адрес сервера, необходимый для загрузки, и данные видеозаписи.
func (api *API) VideoSave(p VideoSaveParams) (*VideoSaveResp, error) {
	resp, err := api.Request("video.save", p, new(VideoSaveResp))
	if err != nil {
		return nil, err
	}
	return resp.(*VideoSaveResp), nil
}
