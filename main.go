package main

import (
	"blog/global"
	"blog/internal/model"
	"blog/internal/routes"
	"blog/pkg/setting"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 程式的執行順序是全域變數初始化 -> init 方法 -> main 方法
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err: %v", err)
	}
	err = setupDbEngine()
	if err != nil {
		log.Fatalf("init setupDbEngine err: %v", err)
	}
}

func main() {
	gin.SetMode(global.Server.Mode)
	router := routes.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.Server.Port,
		Handler:        router,
		ReadTimeout:    global.Server.ReadTimeout,
		WriteTimeout:   global.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 請求表頭最大位元組數，數字 1 向左位移 20 位，即 1 乘以 (2^{20})。結果是 1,048,576。
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSetting("Server", &global.Server)
	if err != nil {
		return err
	}
	err = setting.ReadSetting("App", &global.App)
	if err != nil {
		return err
	}
	err = setting.ReadSetting("Database", &global.Database)
	if err != nil {
		return err
	}
	global.Server.ReadTimeout *= time.Second
	global.Server.WriteTimeout *= time.Second
	return nil
}

func setupDbEngine() error {
	var err error
	global.DbEngine, err = model.NewDbEngine(global.Database)
	if err != nil {
		return err
	}
	return nil
}
