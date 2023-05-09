package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/toolkits/pkg/file"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
)

type Config struct {
	File string `json:"file" form:"file"`
}

func getInputList() []string {
	config := make([]string, 0)
	confPath := fmt.Sprintf("%s/%s", file.SelfDir(), "conf")
	files, err := ioutil.ReadDir(confPath)
	if err != nil{
		log.Println("E! input list failed, err:", err)
		return config
	}
	for _, file := range files{
		if file.IsDir() && strings.HasPrefix(file.Name(), "input"){
			config = append(config, strings.SplitN(file.Name(), ".", 2)[1])
		}
	}
	return config
}


func input(c *gin.Context)  {
	config := getInputList()
	c.HTML(200, "index.html", gin.H{"config": config, "name":"", "file":""})
}

func getInputConfig(c *gin.Context)  {
	inputName := c.Param("input")
	path := fmt.Sprintf("conf/input.%s/%s.toml", inputName, inputName)
	file, _ := ioutil.ReadFile(path)
	c.HTML(200, "config.html", gin.H{"name":inputName, "file":string(file)})
}

func updateInputConfig(c *gin.Context)  {
	var config Config
	c.BindJSON(&config)

	inputName := c.Param("input")
	path := fmt.Sprintf("conf/input.%s/%s.toml", inputName, inputName)

	if err := ioutil.WriteFile(path, []byte(config.File), 0666); err != nil{
		log.Println("E! input config update failed, err:", err)
		c.Status(500)
		return
	}

	//发送信号，Reload配置
	self, _ := os.FindProcess(os.Getpid())
	self.Signal(syscall.SIGHUP)

	c.HTML(200, "config.html", gin.H{"name":inputName, "file":""})
}
