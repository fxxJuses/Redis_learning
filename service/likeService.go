package service

import (
	"fmt"
	"net/http"
	"redis_learning/middleware"
	"time"
)

type Like struct{
	VideoId  int64
	UserId int64
	action int 
	delete int
}


func LikeService(w http.ResponseWriter , r *http.Request){
	if r.Method != "POST"{
		w.Write([]byte("Request Method must be POST , but get " + r.Method))
		return 
	}

	r.ParseForm()
	videoId := r.Form.Get("videoId")
	userId := r.Form.Get("userId")
	action := r.Form.Get("action")


	if action == "1"{
		err := middleware.Rdb.Set(videoId , userId , time.Second * 10).Err()
		if err != nil {
			panic(err)
		}
		w.Write([]byte("redis set videoId successfully"))
		w.WriteHeader(http.StatusOK)
		fmt.Println("redis set videoId successfully")
	}else if action == "2" {
		value , err := middleware.Rdb.Get(videoId).Result()
		w.Write([]byte(value))
		w.WriteHeader(http.StatusOK)
		if err != nil {
			panic(err)
		}else{
			middleware.Rdb.Del(videoId)
		}
	}
}