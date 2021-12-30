package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	workDir    = "/data/media/0/Documents/AutoClean/"
	ignoreList = []string{
		"Android",
		"Android/data",
		"Android/media",
		"Android/obb",
		"DCIM",
		"Documents",
		"Download",
		"Movies",
		"Music",
		"Pictures",
	}
	sdcardList = []string{
		"/sdcard/",
		"/storage/emulated/0/",
		"/data/media/0/",
	}
)

func main() {
	logFile, _ := os.OpenFile(workDir+"run.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)
	config := readConfig()
	for _, path := range config {
		if strings.Contains(path, "#") || path == "" {
			continue
		}
		if !hasPrefix(sdcardList, path) {
			log.Println(strings.Join([]string{"保护：前缀", path}, " "))
			continue
		}
		if strings.HasSuffix(path, "/") {
			path = path[:len(path)-1]
		}
		if ignore(ignoreList, path) {
			log.Println(strings.Join([]string{"保护：忽略", path}, " "))
			continue
		}
		files, err := filepath.Glob(path)
		if err != nil {
			log.Fatalln(strings.Join([]string{"异常：", err.Error()}, " "))
		}
		for _, file := range files {
			removeFileOrDir(file)
		}
	}
}

func readConfig() []string {
	file, err := os.Open(workDir + "config.prop")
	defer file.Close()
	if err != nil {
		log.Fatalln(strings.Join([]string{"异常：", err.Error()}, " "))
	}
	var config []string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		config = append(config, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalln(err.Error())
	}
	return config
}

func removeFileOrDir(path string) {
	s, err := os.Stat(path)
	if err != nil {
		return
	}
	if s.IsDir() {
		err := os.RemoveAll(path)
		if err != nil {
			log.Println(strings.Join([]string{"错误：文件夹", err.Error()}, " "))
			return
		}
		log.Println(strings.Join([]string{"删除：文件夹", path}, " "))
	} else {
		err := os.Remove(path)
		if err != nil {
			log.Println(strings.Join([]string{"错误：文件", err.Error()}, " "))
			return
		}
		log.Println(strings.Join([]string{"删除：文件", path}, " "))
	}
}

func ignore(str []string, tmp string) bool {
	for _, s := range str {
		for _, prefix := range sdcardList {
			if prefix+s == tmp {
				return true
			}
		}
	}
	return false
}

func hasPrefix(str []string, tmp string) bool {
	for _, s := range str {
		if strings.HasPrefix(tmp, s) {
			return true
		}
	}
	return false
}
