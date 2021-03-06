pm2 delete all
cd ./QQMusicApi/
pm2 start --name qq_music npm -- start
cd ../MiguMusicApi/
pm2 start --name migu_music npm -- start
cd ../NeteaseCloudMusicApi/
pm2 start --name netease_music app.js

kill -9 `ps -ef|grep main|grep -v grep|awk '{print $2}'`
cd ../goServer/
go build main.go
nohup ./main >>./log/access.log 2>&1 &

sleep 3
bash ../flushCookie.sh
netstat -nltp
systemctl restart nginx
