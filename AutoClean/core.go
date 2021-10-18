package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var workDir = "/data/media/0/Documents/AutoClean/"

func init() {
	logFile, _ := os.OpenFile(workDir+"run.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)
}

func main() {
	config := readConfig()
	for _, path := range config {
		if strings.Contains(path, "#") || path == "" {
			continue
		}
		files, err := filepath.Glob(path)
		if err != nil {
			log.Fatalln(err.Error())
		}
		for _, file := range files {
			removeFileOrDir(file)
		}
	}
}

func readConfig() []string {
	file, err := os.Open(workDir + "config.prop")
	if err != nil {
		log.Fatalln(err.Error())
	}
	var config []string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		config = append(config, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalln(err.Error())
	}
	err = file.Close()
	if err != nil {
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
			log.Println(err.Error())
			return
		}
		log.Println(strings.Join([]string{"删除文件夹", path}, " "))
	} else {
		err := os.Remove(path)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println(strings.Join([]string{"删除文件", path}, " "))
	}
}
