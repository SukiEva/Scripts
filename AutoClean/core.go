package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	BlackList []string
}

func init() {
	logFile, _ := os.OpenFile("/data/media/0/Documents/AutoClean/run.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)
}

func main() {
	config := readConfig()
	for _, path := range config.BlackList {
		files, err := filepath.Glob(path)
		if err != nil {
			log.Fatalln(err.Error())
		}
		for _, file := range files {
			removeFileOrDir(file)
		}
	}
}

func readConfig() Config {
	file, err := os.Open("/data/media/0/Documents/AutoClean/config.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
	var c Config
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err.Error())
		}
	}(file)
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return c
}

func removeFileOrDir(path string) {
	s, err := os.Stat(path)
	if err != nil {
		//log.Println(err.Error())
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
