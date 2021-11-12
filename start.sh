pm2 delete all
cd ./QQMusicApi/
pm2 start --name qq_music npm -- start
cd ../MiguMusicApi/
pm2 start --name migu_music npm -- start
cd ../NeteaseCloudMusicApi/
pm2 start --name netease_music app.js
