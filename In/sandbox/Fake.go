package sandbox

import (
	"golang.org/x/sys/windows"
	"os"
)

//网上的沙盒等系统都是英文界面，我们可以利用这个进行判断
func IFlanguage(){
	a, _ := windows.GetUserPreferredUILanguages(windows.MUI_LANGUAGE_NAME)   //获取当前系统首选语言
	if a[0] != "zh-CN" {
		os.Exit(1)
	}
}
