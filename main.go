package main

import (
	"LiveProxySpeedTest/global"
	"LiveProxySpeedTest/internal/task"
	"LiveProxySpeedTest/internal/utils"
	"LiveProxySpeedTest/pkg/logging"
	"LiveProxySpeedTest/pkg/setting"
	"flag"
	"fmt"
	"log"
	"runtime"
)

var table string

func init() {
	flag.StringVar(&table, "t", "aishang", "指定测速类型")
	flag.Parse()
	// 配置初始化
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	// zap初始化
	logging.Init()
}

func main() {
	task.InitRandSeed() // 置随机数种子
	// 开始延迟测速
	pingData := task.NewPing(table).Run().FilterDelay()
	// 开始下载测速
	speedData := utils.DownloadSpeedSet(pingData)
	utils.ExportCsv(speedData) // 输出文件
	speedData.Print()          // 打印结果

	endPrint()
}

func setupSetting() error {
	newSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = newSetting.ReadSection("Log", &global.LogSetting)
	if err != nil {
		return err
	}

	return nil
}

func endPrint() {
	if utils.NoPrintResult() {
		return
	}
	if runtime.GOOS == "windows" { // 如果是 Windows 系统，则需要按下 回车键 或 Ctrl+C 退出（避免通过双击运行时，测速完毕后直接关闭）
		fmt.Printf("按下 回车键 或 Ctrl+C 退出。")
		fmt.Scanln()
	}
}
