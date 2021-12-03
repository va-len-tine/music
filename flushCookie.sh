data="RK=iJD4y34bl2; ptcz=e96106ead89270d88df901e0d064f6429a6fc1bb11ed4c6be1ccdd64c4bcb7e1; fqm_pvqid=d2b124a3-37db-4d5f-b5d8-e09f7f3114ca; pgv_pvid=8182091358; ts_refer=ADTAGmyqq; ts_uid=5407687754; tmeLoginType=2; ptui_loginuin=2292752593; euin=ow-qowSkow4qoz**; fqm_sessionid=cf13f8f5-b1bf-45b6-9298-70f7e3d4f61e; pgv_info=ssid=s6113047572; _qpsvr_localtk=0.14099130248337843; login_type=1; psrf_access_token_expiresAt=1646182495; psrf_qqrefresh_token=8FE9E71200A92673E2A4AA7F3A9E0019; wxunionid=; psrf_qqaccess_token=A7DA9148CC406830BB9EBE7DD771FED2; wxrefresh_token=; psrf_qqopenid=88206CD98E2706E3F9B9E4CB6045E5F5; wxopenid=; psrf_qqunionid=5836864AEE9E76BAF901B2AA780879FD; uin=o2292752593; skey=@0IYQWTsLR"

curl --location --request GET 'http://localhost:3300/user/setCookie' \
--header 'content-type: application/json' \
--header 'Cookie: NMTID=00OwCrdl-XUu-wEYED_sF6QNOYSoYcAAAF9Dj3rKA' \
--data "{\"data\":\"$data\"}"
