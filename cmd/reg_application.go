package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"gopkg.in/guregu/null.v3"
	"log"
	"os"
	"send-email/model/entity"
	"send-email/model/repository"
	"send-email/utils"
	"strings"
)

func main() {
	app := &cli.App{
		Name: "reg_application",
		Usage: "create oauth2 application",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "name",
				Usage: "アプリケーション名称",
			},
			&cli.StringFlag{
				Name: "domain",
				Usage: "ドメイン",
			},
		},
		Action: func(c *cli.Context) error {
			reader := bufio.NewReader(os.Stdin)
			name := c.String("name")
			domain := c.String("domain")
			if name == "" {
				name = rawInput(reader, "アプリケーション名称：")
			}
			if domain == "" {
				domain = rawInput(reader, "ドメイン：")
			}
			repository.Init()
			appId := registerApplication(name, domain)
			fmt.Printf("アプリケーションが作成しました。id:%d", appId)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func rawInput(r *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	name, _ := r.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	return name
}

func registerApplication(name string, domain string) int64 {
	clientId, _ := utils.MakeRandomStr(40)
	clientSecret, _ := utils.MakeRandomStr(128)
	appId, err := repository.CreateApplication(entity.ApplicationEntity{
		ClientId: clientId,
		ClientSecret: clientSecret,
		Domain: domain,
		Name: null.NewString(name, true),
	})
	if err != nil {
		log.Fatal(err)
	}
	return appId
}