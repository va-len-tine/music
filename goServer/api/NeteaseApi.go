package api

import (
	"encoding/json"
	"goServer/global"
	"goServer/model"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NeteaseApi(kw string) []map[string]interface{}{
	uri := "http://localhost:3000/cloudsearch?keywords=" + url.QueryEscape(kw)
	resp,err := http.Get(uri)
	if err != nil{
		var result []map[string]interface{}
		m := map[string]interface{}{"author":"","name":global.CF.Local.Tips,"pic":"","src":global.CF.Local.Defaultmusic}
		result = append(result, m)
		return result
	}
	defer resp.Body.Close()
	bytes,_ := io.ReadAll(resp.Body)
	var v2 model.Netease2
	err = json.Unmarshal(bytes,&v2)
	ids := ""
	var res []map[string]interface{}
	for _,item := range v2.Result.Songs{
		m := make(map[string]interface{})
		m["name"] = item.Name
		m["author"] = item.Ar[0].Name
		m["pic"] = strings.Replace(item.Al.PicURL,"http://", "https://", -1)
		m["src"] = ""
		m["id"] = item.ID
		ids = ids + strconv.Itoa(item.ID) + ","
		res = append(res, m)
	}
	ids = strings.TrimRight(ids,",")

	resp,err = http.Get("http://localhost:3000/song/url?id=" + ids)
	if err != nil{
		var result []map[string]interface{}
		m := map[string]interface{}{"author":"","name":global.CF.Local.Tips,"pic":"","src":global.CF.Local.Defaultmusic}
		result = append(result, m)
		return result
	}
	defer resp.Body.Close()
	bytes,_ = io.ReadAll(resp.Body)
	var v3 model.Netease3
	err = json.Unmarshal(bytes,&v3)
	res1 := v3.Data

	var res2 []map[string]interface{}
	for _,i := range res{
		for _,j := range res1{
			if i["id"] == j.ID && j.URL != ""{
				i["src"] = strings.Replace(j.URL,"http://", "https://", -1)
				res2 = append(res2, i)
				break
			}
		}
	}

	if len(res2) == 0{
		var result []map[string]interface{}
		m := map[string]interface{}{"author":"","name":global.CF.Local.Tips,"pic":"","src":global.CF.Local.Defaultmusic}
		result = append(result, m)
		return result
	}
	return res2
}


func NeteaseApiTop(kw string) []map[string]interface{}{
	uri := "http://localhost:3000/playlist/detail?id=" + url.QueryEscape(kw)
	resp,err := http.Get(uri)
	if err != nil{
		var result []map[string]interface{}
		m := map[string]interface{}{"author":"","name":global.CF.Local.Tips,"pic":"","src":global.CF.Local.Defaultmusic}
		result = append(result, m)
		return result
	}
	defer resp.Body.Close()
	bytes,_ := io.ReadAll(resp.Body)
	var v1 model.Netease1
	err = json.Unmarshal(bytes,&v1)
	ids := ""
	var res []map[string]interface{}
	for _,item := range v1.Playlist.Tracks{
		m := make(map[string]interface{})
		m["name"] = item.Name
		m["author"] = item.Ar[0].Name
		m["pic"] = strings.Replace(item.Al.PicURL,"http://", "https://", -1)
		m["src"] = ""
		m["id"] = item.ID
		ids = ids + strconv.Itoa(item.ID) + ","
		res = append(res, m)
	}
	ids = strings.TrimRight(ids,",")

	resp,err = http.Get("http://localhost:3000/song/url?id=" + ids)
	if err != nil{
		var result []map[string]interface{}
		m := map[string]interface{}{"author":"","name":global.CF.Local.Tips,"pic":"","src":global.CF.Local.Defaultmusic}
		result = append(result, m)
		return result
	}
	defer resp.Body.Close()
	bytes,_ = io.ReadAll(resp.Body)
	var v3 model.Netease3
	err = json.Unmarshal(bytes,&v3)
	res1 := v3.Data

	var res2 []map[string]interface{}
	for _,i := range res{
		for _,j := range res1{
			if i["id"] == j.ID && j.URL != ""{
				i["src"] = strings.Replace(j.URL,"http://", "https://", -1)
				res2 = append(res2, i)
				break
			}
		}
		if len(res2)==20{
			break
		}
	}

	if len(res2) == 0{
		var result []map[string]interface{}
		m := map[string]interface{}{"author":"","name":global.CF.Local.Tips,"pic":"","src":global.CF.Local.Defaultmusic}
		result = append(result, m)
		return result
	}
	return res2
}
