package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// coment
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação de Teste"
	app.Usage = "Busca IPs e Nomes de Servidores na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "ip",
			ShortName: "i",
			Usage:     "Busca IPs de endereços na Internet",
			Flags:     flags,
			Action:    BuscarIps,
		},
		{
			Name:      "servidores",
			ShortName: "s",
			Usage:     "Busca o nome de servidores na Internet",
			Flags:     flags,
			Action:    BuscarServidores,
		},
	}
	return app
}

func BuscarIps(c *cli.Context) error {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}

func BuscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor)
	}
}
