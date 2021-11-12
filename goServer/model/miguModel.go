package model

type Migu struct {
	Result int `json:"result"`
	Data struct {
		List []struct {
			Name string `json:"name"`
			ID string `json:"id"`
			Cid string `json:"cid"`
			MvID string `json:"mvId"`
			URL interface{} `json:"url"`
			Album struct {
				PicURL string `json:"picUrl"`
				Name string `json:"name"`
				ID string `json:"id"`
			} `json:"album"`
			Artists []struct {
				ID string `json:"id"`
				Name string `json:"name"`
			} `json:"artists"`
		} `json:"list"`
		Total int `json:"total"`
	} `json:"data"`
}
