package api

import (
	"encoding/json"
	"goServer/error"
	"goServer/model"
	"io"
	"net/http"
	"net/url"
)

func MiguApi(kw string) []map[string]interface{}{
	result := make([]map[string]interface{},0)
	uri := "http://localhost:3400/search?keyword=" + url.QueryEscape(kw)
	resp,err := http.Get(uri)
	error.HandleErr(err,"请求出错")
	defer resp.Body.Close()
	bytes,_ := io.ReadAll(resp.Body)
	var v model.Migu
	err = json.Unmarshal(bytes, &v)
	for _,item := range v.Data.List{
		m := make(map[string]interface{})
		m["name"] = item.Name
		m["author"] = item.Artists[0].Name
		m["pic"] = item.Album.PicURL
		m["src"] = item.URL
		if m["src"] != nil{
			result = append(result, m)
		}
	}
	return result
}
