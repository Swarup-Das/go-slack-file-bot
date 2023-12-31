package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main(){
	os.Setenv("SLACK_BOT_TOKEN","xoxb-5441341290615-5461751522691-irzpj0zU6bMsSiGIRjCU1tJP")
	os.Setenv("CHANNEL_ID","C05DDNFU0AX")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"SwarupDasResume.pdf"}

	for i := 0; i < len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil{
			fmt.Printf("%s\n",err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}