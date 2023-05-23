package main

import (
	"net/http"
	"redis_learning/service"
	"redis_learning/middleware"

)

func main() {
	middleware.RedisInit()


	http.HandleFunc("/test" , func (w http.ResponseWriter , r *http.Request)  {
		w.Write([]byte("hello world"))
	})
	http.HandleFunc("/like" , service.LikeService)

	http.ListenAndServe(":8099" , nil )

}