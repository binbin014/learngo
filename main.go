package main

import (
	"fmt"
	apiRouters "learngin/app/api/routers"
	webRouters "learngin/app/web/routers"
	"learngin/boot"
	"learngin/routers"
)

func main() {

	boot.Initialize()
	routers.Include(
		apiRouters.Routers,
		webRouters.Routers,
	)
	r:=routers.Init()
	if err := r.Run(":8888"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
