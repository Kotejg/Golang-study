package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// gerar vai retornar a app de lina de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "app de linha de comando"
	app.Usage = "busca ip e nome de servires na internet"

	flags :=
		[]cli.Flag{
			cli.StringFlag{
				Name:  "host",
				Value: "devbook.com.br",
			},
		}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS de enderecos na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidor",
			Usage:  "busca servidores",
			Flags:  flags,
			Action: buscaServidores,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	//net
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscaServidores(c *cli.Context) {
	host := c.String("host")

	//serv
	servidores, erro := net.LookupNS(host) // name service = nome do servidor

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
