import requests
from typing import Optional, List
from fastapi import FastAPI
from pydantic import BaseModel
import pymysql
import urllib.request
import os
import urllib.parse

#默认播放链接,有bug,会将服务器上的歌曲大小变成0
#dft_url = 'http://39.101.203.25/music/c45ae8acfea229a31cd7bc85ce136669.mp3'
dft_url = 'http://39.101.203.25/music/14b91afdc1e2d83ec06ff7c3aea1286a.mp3'

class Item(BaseModel):
    name: Optional[str] = None
    author: Optional[str] = None
    src: Optional[str] = None
    pic: Optional[str] = None


app = FastAPI()


def netease_music(kw):
    if kw == '19723756' or kw == '3779629' or kw == '2884035' or kw == '3778678':
        url = 'http://localhost:3000/playlist/detail?id=' + kw
        res = requests.get(url).json()['playlist']['tracks'][:20]
    else:
        url = 'http://localhost:3000/cloudsearch?keywords=' + kw
        res = requests.get(url).json()['result']['songs']
    result = []
    ids = ''
    for item in res:
        name = item['name']
        id = item['id']
        ids = ids + str(id) + ','
        author = item['ar'][0]['name']
        pic = item['al']['picUrl']
        src = None
        result.append({"id": id, "name": name, "author": author, "src": src, "pic": pic})
    res1 = requests.get('http://localhost:3000/song/url?id=' + ids[:-1]).json()['data']
    R = []
    for i in result:
        for j in res1:
            if j['id'] == i['id'] and j['url'] is not None:
                i['src'] = j['url']
                R.append(i)
    return R


def qq_music(kw):
    url = 'http://localhost:3300/search?key=' + kw
    try:
        res = requests.get(url).json()['data']['list']
    except Exception as e:
        return [{"src":dft_url,"name":"搜了个寂寞...啥也没有","author":"","pic":""}]
    result = []
    ids = ''
    for item in res:
        name = item['songname']
        id = item['songmid']
        ids = ids + str(id) + ','
        author = item['singer'][0]['name']
        pic = ''
        src = None
        result.append({"id": id, "name": name, "author": author, "src": src, "pic": pic})
    res1 = requests.get('http://localhost:3300/song/urls?id=' + ids[:-1]).json()['data']
    R = []
    for i in result:
        for j in res1.keys():
            if j == i['id'] and res1[j] is not None:
                i['src'] = res1[j]
                R.append(i)
    return R


def migu_music_bak(kw):
    url = 'http://localhost:3400/search?keyword=' + kw
    try:
        res = requests.get(url, timeout=2).json()
    except Exception as e:
        try:
            res = requests.get(url, timeout=2).json()
        except Exception as e:
            return [{"src":dft_url,"name":"搜了个寂寞...啥也没有","author":"","pic":""}]
        result = []
        for item in res['data']['list']:
            name = item['name']
            author = item['artists'][0]['name']
            pic = item['album']['picUrl']
            src = item['url']
            if src is not None:
                result.append({"name": name, "author": author, "src": src, "pic": pic})
        return result


def migu_music(kw):
    url = 'http://localhost:3400/search?keyword=' + kw
    try:
        res = requests.get(url, timeout=2).json()
    except Exception as e:
        return [{"src":dft_url,"name":"搜了个寂寞...啥也没有","author":"","pic":""}]
    result = []
    for item in res['data']['list']:
        name = item['name']
        author = item['artists'][0]['name']
        pic = item['album']['picUrl']
        src = item['url']
        if src is not None:
            result.append({"name": name, "author": author, "src": src, "pic": pic})
    return result


def download(item):
    nm = urllib.parse.unquote(item['src']).split('/')[-1].split('?')[0]
    file_path = '/usr/share/nginx/html/music/' + nm
    try:
        with open(file_path, "wb") as f:
            f.write(urllib.request.urlopen(item['src']).read())
    except Exception as e:
        print(f"下载错误！{e}")
    else:
        try:
            url = 'http://39.101.203.25/music/' + nm
            db = pymysql.connect(host='localhost',user = "godman",password = "godman",db = "music")
            cursor = db.cursor()
            cursor.execute("insert into music_info values(%s, %s, %s, %s)" %(db.escape(url), db.escape(item['name']), db.escape(item['author']), db.escape(item['pic'])))
            db.commit()
            cursor.close()
            db.close()
            with open(file_path, "wb") as f:
                f.write(urllib.request.urlopen(item['src']).read())
        except Exception as e:
            print(f"下载错误！{e}")


def magic(kw):
    result = []
    db = pymysql.connect(host='localhost',user = "godman",password = "godman",db = "music")
    cursor = db.cursor()
    sql = "select * from music_info where name like %s or author like %s;"
    params = ['%' + kw + '%', '%' + kw + '%']
    cursor.execute(sql, params)
    fd = ('src', 'name', 'author', 'pic')
    for item in cursor.fetchmany(20):
        result.append(dict(zip(fd, item)))
    if not result:
        result.append({"src":dft_url,"name":"搜了个寂寞...啥也没有","author":"","pic":""})
        sql = "select * from music_info limit 20;"
        cursor.execute(sql)
        for item in cursor.fetchmany(20):
            result.append(dict(zip(fd, item)))
    cursor.close()
    db.close()
    return result


@app.get("/magic/api")
async def deal_item(keyword: str = '林俊杰'):
    return magic(keyword)

@app.get("/netease/api")
async def deal_item(keyword: str = '林俊杰'):
    return netease_music(keyword)

@app.get("/qq/api")
async def deal_item(keyword: str = '林俊杰'):
    return qq_music(keyword)

@app.get("/migu/api")
async def deal_item(keyword: str = '林俊杰'):
    return migu_music(keyword)

@app.post("/download/api")
async def deal_item(item: Item):
    download(dict(item))
