package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	workDir = "/data/media/0/Documents/AutoMove/"
	cmd *exec.Cmd
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
		if !strings.Contains(path, "&") {
			log.Println(strings.Join([]string{"忽略：缺少 &", path}, " "))
			continue
		}
		dirs := strings.Split(path, "&")
		for i := 0; i < len(dirs); i++ {
			if strings.HasSuffix(dirs[i], "/") {
				dirs[i] = dirs[i][:len(dirs[i])-1]
			}
		}
		moveFileOrDir(dirs[0], dirs[1])
	}
}

func moveFileOrDir(sourcePath, destinationPath string) {
	if !exists(sourcePath) {
		log.Println(strings.Join([]string{"忽略：源路径不存在", sourcePath}, " "))
		return
	}
	if !exists(destinationPath) {
		err := os.Mkdir(destinationPath, os.ModePerm)
		if err != nil {
			log.Println(strings.Join([]string{"错误：", err.Error()}, " "))
			return
		}
	}
	rmList := make([]string, 0)
	err := filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if strings.HasPrefix(info.Name(), ".") {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if path != sourcePath {
			rmList = append(rmList, path)
		}
		if info.IsDir() {
			return nil
		}
		newDestinationPath := strings.Join([]string{destinationPath, "/", info.Name()}, "")
		moveFile(path, newDestinationPath)
		log.Println(strings.Join([]string{"移动：", path, "=>", newDestinationPath}, " "))
		return nil
	})
	if err != nil {
		log.Println(strings.Join([]string{"异常：", err.Error()}, " "))
	}
	for _, rm := range rmList {
		removeFileOrDir(rm)
	}
}

func removeFileOrDir(path string) {
	s, err := os.Stat(path)
	if err != nil {
		return
	}
	if s.IsDir() {
		err := os.RemoveAll(path)
		if err != nil {
			return
		}
	} else {
		err := os.Remove(path)
		if err != nil {
			return
		}
	}
}

func moveFile(sourcePath, destinationPath string) {
	cmd = exec.Command("mv", sourcePath, destinationPath)
	_, err := cmd.Output()
	if err != nil {
		log.Println(strings.Join([]string{"错误：移动失败", err.Error(), sourcePath, "=>", destinationPath}, " "))
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
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
