package error

import "log"

func HandleErr(err error, info string){
	if err != nil{
		log.Fatal(info, ":", err)
	}
}
