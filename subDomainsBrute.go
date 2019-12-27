package main

import (
	"./lib"
	"fmt"
	"time"
)

func init() {
	lib.Parse_args()
}
func main() {
	println("SubDomainsBrute v1.2")
	println("https://github.com/lijiejie/subDomainsBrute")
	tmp_dir := fmt.Sprintf("tmp/%s_%s", lib.Target, time.Now())
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
