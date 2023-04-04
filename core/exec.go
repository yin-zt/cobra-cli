package core

import (
	"context"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/henrylee2cn/mahonia"
	"github.com/yin-zt/cobra-cli/config"
	"github.com/yin-zt/cobra-cli/utils"
	"io/ioutil"
	random "math/rand"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// Shell 使用方式：gmd shell  -d path -f file -t 12 -u -a -x
// 如果本地存在fpath文件(path+file); 且没有位置参数-u；则执行fpath + 其他位置参数组成的命令
func (this *Cli) Shell() {
	//var err error
	//var includeReg *regexp.Regexp
	src := ""
	argv := utils.GetArgsMap()
	file := ""
	dir := ""
	update := "0"
	debug := "0"
	timeout := -1
	ok := true
	// -u 则是
	if v, ok := argv["u"]; ok {
		update = v
	}

	// 如果位置参数带 -x 则是使用bash执行
	if v, ok := argv["x"]; ok {
		debug = v
	}

	// -t 指定执行时间
	if v, ok := argv["t"]; ok {
		timeout, _ = strconv.Atoi(v)
	}

	// -f 指定文件
	if file, ok = argv["f"]; !ok {
		fmt.Println("-f(filename) is required")
		return
	}

	// -d 则是指定目录；如果没有-d参数，则使用 shell目录
	if dir, ok = argv["d"]; !ok {
		dir = "shell"
	}

	// ScriptPath -> /tmp/script/shell
	path := config.StoreCmd + dir
	if !utils.IsExist(path) {
		log.Debug(os.MkdirAll(path, 0777))
	}
	os.Chmod(path, 0777)

	//includeRegStr := `#include\s+-f\s+(?P<filename>[a-zA-Z.0-9_-]+)?\s+-d\s+(?P<dir>[a-zA-Z.0-9_-]+)?|#include\s+-d\s+(?P<dir2>[a-zA-Z.0-9_-]+)?\s+-f\s+(?P<filename2>[a-zA-Z.0-9_-]+)?`
	// 函数传入目录和文件名两个参数，请求cli server的 http://cli server/cli/shell接口
	//DownloadShell := func(dir, file string) (string, error) {
	//	req := httplib.Post(this.Conf.EnterURL + "/" + this.Conf.DefaultModule + "/shell")
	//	req.Param("dir", dir)
	//	req.Param("file", file)
	//	return req.String()
	//}

	// includestr字符串中包含如下格式："-d path -f filename"
	// 函数最后会去调用DownloadShell
	//DownloadIncludue := func(includeStr string) string {
	//	type DF struct {
	//		Dir  string
	//		File string
	//	}
	//	df := DF{}
	//	parts := strings.Split(includeStr, " ")
	//	for i, v := range parts {
	//		if v == "-d" {
	//			df.Dir = parts[i+1]
	//		}
	//		if v == "-f" {
	//			df.File = parts[i+1]
	//		}
	//	}
	//
	//	if s, err := DownloadShell(df.Dir, df.File); err != nil {
	//		log.Error(err)
	//		return includeStr
	//	} else {
	//		return s
	//	}
	//
	//}

	// path -> /tmp/script/shell ; file -> -f filename
	fpath := path + "/" + file
	fpath = strings.Replace(fpath, "///", "/", -1)
	fpath = strings.Replace(fpath, "//", "/", -1)

	// 如果文件不存在当前目录，则请求gmd server下载文件
	if update == "1" || !utils.IsExist(fpath) {
		fmt.Println("call gmd server")
		//if src, err = DownloadShell(dir, file); err != nil {
		//	log.Error(err)
		//	return
		//}
		//
		//if includeReg, err = regexp.Compile(includeRegStr); err != nil {
		//	log.Error(err)
		//	return
		//}
		//
		//os.MkdirAll(filepath.Dir(fpath), 0777)
		//
		//// 从cli server侧下载的src路径与正则表达式匹配；满足匹配的值的字符串则传入函数DownloadIncludue中
		//// 函数DownloadIncludue则会调用DownloadShell函数下载
		//// 如此逻辑其实调用了两次
		//src = includeReg.ReplaceAllStringFunc(src, DownloadIncludue)
		//
		//this.Util.WriteFile(fpath, src)
	} else {
		src = utils.ReadFile(fpath)
	}

	// 对请求gmd server url http://gmd server/gmd/shell返回的字符串进行处理
	lines := strings.Split(src, "\n")
	is_python := false
	is_shell := false
	is_powershell := false
	if len(lines) > 0 {
		// 判断字符串是否包含 python
		is_python, _ = regexp.MatchString("python", lines[0])
		// 判断字符串是否包含 bash
		is_shell, _ = regexp.MatchString("bash", lines[0])
	}

	// 判断字符串是否包含 ps1 -> powershell
	if strings.HasSuffix(file, ".ps1") {
		is_powershell = true
	}

	os.Chmod(fpath, 0777)
	result := ""

	// 组成执行命令的命令列表
	cmds := []string{
		fpath,
	}
	if is_python {
		cmds = []string{
			"/usr/bin/env",
			"python",
			fpath,
		}
	}
	if is_shell {
		cmds = []string{
			"/bin/bash",
			fpath,
		}
		if debug == "1" {
			cmds = []string{
				"/bin/bash",
				"-x",
				fpath,
			}
		}
	}

	if is_powershell {
		cmds = []string{
			"powershell",
			fpath,
		}
	}

	argvMap := utils.GetArgsMap()

	// 检查执行 cli shell 命令时 位置参数是否有 -a
	// 如果有，则将 -a 后接的所有值使用 “ ”分隔后，再追加到cmds数组中
	if args, ok := argvMap["a"]; ok {
		cmds = append(cmds, strings.Split(args, " ")...)
	} else {
		var args []string
		var tflag bool
		tflag = false
		for i, v := range os.Args {
			if v == "-t" {
				tflag = true
				continue
			}
			if tflag {
				tflag = false
				continue
			}
			// 如果位置参数中不是-x和-u，则将此参数放入列表args中
			if v != "-x" && v != "-u" {
				args = append(args, os.Args[i])
			}
		}
		//fmt.Println("update:",update,"debug:",debug)
		//fmt.Println("args:",args)
		// 把args列表中第6个及之后的值加入命令列表cmds中
		os.Args = args
		cmds = append(cmds, os.Args[6:]...)
		//fmt.Println("cmds", cmds)
	}

	// 本地执行脚本cmds并输出结果
	result, _, _ = this.Exec(cmds, timeout, nil)
	fmt.Println(result)
}

// 如果task_id 为nil, 则在/tmp目录下使用随机数创建一个文件，并打开此文件用来保存执行cmd命令的输出；
// 调用exec.CommandContext来执行cmd
func (this *Cli) Exec(cmd []string, timeout int, kw map[string]string) (string, string, int) {
	var re any
	defer func() {
		if re = recover(); re != nil {
			buffer := debug.Stack()
			corelog.Error("Exec")
			corelog.Error(re)
			corelog.Error(string(buffer))
		}
	}()
	//var out bytes.Buffer

	var fp *os.File
	var err error
	var taskId string
	var fpath string
	var data []byte

	//生成一个任务id
	taskId = time.Now().Format("20060102150405") + fmt.Sprintf("%d", time.Now().Unix())

	// 在tmp目录下创建文件用来存储执行命令生成的输出
	home_path := ""
	if runtime.GOOS == "windows" {
		home_path = config.StoreCmd
	} else {
		home_path = config.StoreCmd
	}

	fpath = home_path + taskId + fmt.Sprintf("_%d", random.New(random.NewSource(time.Now().UnixNano())).Intn(60))
	fp, err = os.OpenFile(fpath, os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		log.Error(err)
		return "", err.Error(), -1
	}
	defer fp.Close()

	// golang 执行操作系统上的脚本或者命令
	duration := time.Duration(timeout) * time.Second
	if timeout == -1 {
		duration = time.Duration(60*60*24*365) * time.Second
	}
	ctx, _ := context.WithTimeout(context.Background(), duration)

	var path string

	// linux 操作系统默认使用"/bin/bash -c " 模式
	var command *exec.Cmd
	command = exec.CommandContext(ctx, cmd[0], cmd[1:]...)
	if "windows" == runtime.GOOS {
		//		command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

		if len(cmd) > 2 {
			cc := strings.Split(cmd[2], " ")
			if cc[0] == "powershell" {
				os.Mkdir(home_path+"/"+"tmp", 0777)
				path = home_path + "/" + "tmp" + "/" + this.GetUUID() + ".ps1"
				ioutil.WriteFile(path, []byte(strings.Join(cc[1:], " ")), 0777)
				command = exec.CommandContext(ctx, "powershell", []string{path}...)
			}
		}
	}
	// 脚本执行后输出到fp中，也就是上面创建的临时文件
	command.Stdin = os.Stdin
	command.Stdout = fp
	command.Stderr = fp

	// 清理创建的fpath文件 和 path文件(windows下执行powershell才会生成)
	RemoveFile := func() {
		fp.Close()
		if path != "" {
			os.Remove(path)
		}
		if fpath != "" {
			os.Remove(fpath)
		}
	}
	_ = RemoveFile
	// 函数退出前，把flag 改为false, 即停止线程读取fpath文件内容
	// 删除fpath 和 path变量指向的文件
	defer RemoveFile()
	// 执行command命令
	err = command.Run()
	// 如果command执行出错，则将命令刷入日志文件中
	// fp文件保存数据
	if err != nil {
		if len(kw) > 0 {
			corelog.Info(kw)
			corelog.Error("error:"+err.Error(), "\ttask_id:"+fpath, "\tcmd:"+cmd[2], "\tcmd:"+strings.Join(cmd, " "))
		} else {
			corelog.Info("task_id:"+fpath, "\tcmd:"+cmd[2], "\tcmd:"+strings.Join(cmd, " "))
		}
		corelog.Flush()
		fp.Sync()
		//fp.Seek(0, 2)
		data, err = ioutil.ReadFile(fpath)
		if err != nil {
			corelog.Error(err)
			corelog.Flush()
			return string(data), err.Error(), -1
		}
		return string(data), "", -1
	}
	status := -1
	// 获取command 命令执行退出状态码并赋值给status，正常退出码赋值为0
	sysTemp := command.ProcessState
	if sysTemp != nil {
		status = sysTemp.Sys().(syscall.WaitStatus).ExitStatus()
	}
	//fp.Seek(0, 2)
	// 将内存中fp的数据输入文件中
	fp.Sync()
	// 读取fpath文件内容
	data, err = ioutil.ReadFile(fpath)
	// 如果操作系统是windows，则将内容使用GBK解码，并最终将执行结果\""\执行状态返回
	if utils.IsWindows() {
		decoder := mahonia.NewDecoder("GBK")
		if decoder != nil {
			if str, ok := decoder.ConvertStringOK(string(data)); ok {
				return str, "", status
			}
		}
	}
	// 如果打开文件失败，则返回data数据
	if err != nil {
		corelog.Error(err, cmd)
		return string(data), err.Error(), -1
	}
	// 最后返回data数据 “” 和command命令退出码
	return string(data), "", status
}
