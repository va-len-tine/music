package model

type QQ2 struct {
	Result int `json:"result"`
	Data struct {
		List []struct {
			Albumid int `json:"albumid"`
			Albummid string `json:"albummid"`
			Albumname string `json:"albumname"`
			AlbumnameHilight string `json:"albumname_hilight"`
			Alertid int `json:"alertid"`
			BelongCD int `json:"belongCD"`
			CdIdx int `json:"cdIdx"`
			Chinesesinger int `json:"chinesesinger"`
			Docid string `json:"docid"`
			Grp []interface{} `json:"grp"`
			Interval int `json:"interval"`
			Isonly int `json:"isonly"`
			Lyric string `json:"lyric"`
			LyricHilight string `json:"lyric_hilight"`
			MediaMid string `json:"media_mid"`
			Msgid int `json:"msgid"`
			NewStatus int `json:"newStatus"`
			Nt int64 `json:"nt"`
			Pay struct {
				Payalbum int `json:"payalbum"`
				Payalbumprice int `json:"payalbumprice"`
				Paydownload int `json:"paydownload"`
				Payinfo int `json:"payinfo"`
				Payplay int `json:"payplay"`
				Paytrackmouth int `json:"paytrackmouth"`
				Paytrackprice int `json:"paytrackprice"`
			} `json:"pay"`
			Preview struct {
				Trybegin int `json:"trybegin"`
				Tryend int `json:"tryend"`
				Trysize int `json:"trysize"`
			} `json:"preview"`
			Pubtime int `json:"pubtime"`
			Pure int `json:"pure"`
			Singer []struct {
				ID int `json:"id"`
				Mid string `json:"mid"`
				Name string `json:"name"`
				NameHilight string `json:"name_hilight"`
			} `json:"singer"`
			Size128 int `json:"size128"`
			Size320 int `json:"size320"`
			Sizeape int `json:"sizeape"`
			Sizeflac int `json:"sizeflac"`
			Sizeogg int `json:"sizeogg"`
			Songid int `json:"songid"`
			Songmid string `json:"songmid"`
			Songname string `json:"songname"`
			SongnameHilight string `json:"songname_hilight"`
			StrMediaMid string `json:"strMediaMid"`
			Stream int `json:"stream"`
			Switch int `json:"switch"`
			T int `json:"t"`
			Tag int `json:"tag"`
			Type int `json:"type"`
			Ver int `json:"ver"`
			Vid string `json:"vid"`
			Format string `json:"format,omitempty"`
			Songurl string `json:"songurl,omitempty"`
		} `json:"list"`
		PageNo int `json:"pageNo"`
		PageSize int `json:"pageSize"`
		Total int `json:"total"`
		Key string `json:"key"`
		T int `json:"t"`
		Type string `json:"type"`
	} `json:"data"`
}
