package api

import (
	"encoding/json"
	"goServer/global"
	"goServer/model"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func QQApi(kw string) []map[string]string{
	uri := "http://localhost:3300/search?key=" + url.QueryEscape(kw)
	resp,err := http.Get(uri)
	if err != nil{
		var result []map[string]string
		m := map[string]string{"author":"","name":"出了点小意外...","pic":"","src":global.CF.Local.Defaultmusic}
		result = append(result, m)
		return result
	}
	defer resp.Body.Close()
	bytes,_ := io.ReadAll(resp.Body)
	var v2 model.QQ2
	err = json.Unmarshal(bytes,&v2)
	ids := ""
	var res []map[string]string
	for _,item := range v2.Data.List{
		m := make(map[string]string)
		m["name"] = item.Songname
		m["author"] = item.Singer[0].Name
		m["pic"] = ""
		m["src"] = ""
		m["id"] = item.Songmid
		ids = ids + item.Songmid + ","
		res = append(res, m)
	}
	ids = strings.TrimRight(ids,",")
	resp,err = http.Get("http://localhost:3300/song/urls?id=" + ids)
	if err != nil{
		var result []map[string]string
		m := map[string]string{"author":"","name":"出了点小意外...","pic":"","src":global.CF.Local.Defaultmusic}
		result = append(result, m)
		return result
	}
	defer resp.Body.Close()
	bytes,_ = io.ReadAll(resp.Body)
	var v3 map[string]map[string]string
	err = json.Unmarshal(bytes,&v3)
	for _,i := range res{
		i["src"] = v3["data"][i["id"]]
	}

	var res1 []map[string]string
	for _,i := range res{
		if i["src"] != ""{
			res1 = append(res1, i)
		}
	}
	return res1
}
