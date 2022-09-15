package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"service/yang/internal/config"
	"service/yang/internal/handler"
	"service/yang/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chabai-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	staticFileHandler(server)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

//定义函数
func staticFileHandler(engine *rest.Server) {
	//这里注册
	patern := "web"
	dirpath := "web/assets/"

	rd, err := ioutil.ReadDir(dirpath)
	if err != nil {
		panic(err)
	}

	//添加进路由最后生成 /asset
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/index.html",
				Handler: dirhandler("index.html", patern),
			},
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: dirhandler("/", patern),
			},
		})
	for _, f := range rd {
		filename := f.Name()
		path := "/assets/" + filename
		//最后生成 /asset
		engine.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: dirhandler("/assets/", dirpath),
			})
	}

}

func dirhandler(patern, filedir string) http.HandlerFunc {
	fmt.Println(filedir)
	return func(w http.ResponseWriter, req *http.Request) {
		http.StripPrefix(patern, http.FileServer(http.Dir(filedir))).ServeHTTP(w, req)
	}
}
