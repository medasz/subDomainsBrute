package main

import (
	"os"
	"github.com/urfave/cli"
)

func main() {
	//lib.Parse_args()
	app := cli.NewApp()
	app.Name = "subDomainsBrute"
	app.Usage = "域名信息探测和枚举"
	app.UsageText = "subDomainsBrute [options] target.com"
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		println("hello")
		return nil
	}
	_ = app.Run(os.Args)
	//config,err:=dns.ClientConfigFromFile("./dns_server.txt")
	//if err!=nil{
	//	log.Fatal(err)
	//}
	//构造客户端
	//c:=new(dns.Client)
	//构造dns报文
	//m:=new(dns.Msg)
	//m.SetQuestion(dns.Fqdn("www.imun.edu.cn"), dns.TypeA)
	//m.RecursionDesired = true
	//r,_,err:=c.Exchange(m,net.JoinHostPort(config.Servers[0],config.Port))
	//r,_,err:=c.Exchange(m,"8.8.8.8:53")
	// if err!=nil{
	// 	log.Fatal(err)
	// }
	// fmt.Println(r.Answer[0].(*dns.A).A)
}
