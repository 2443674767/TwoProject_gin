package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	//"github.com/casbin/casbin/v2/model"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//e, err := casbin.NewEnforcer("./casbin_class/model.conf", "./casbin_class/policy.csv")

	// Initialize a Xorm adapter with MySQL database.
	a, _ := xormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/casbin", true)

	e, _ := casbin.NewEnforcer("./casbin_class/model.conf", a)

	//if err != nil {
	//	log.Fatalf("error: adapter: %s", err)
	//}
	//
	//if err != nil {
	//	fmt.Printf("加载 Casbin 出错: %v\n", err)
	//	return
	//}

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	//增
	//added, err := e.AddPolicy("alice", "data1", "read")
	//fmt.Println(added, err)

	//查
	filteredPolicy, _ := e.GetFilteredPolicy(0, "alice")

	fmt.Println(filteredPolicy)

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// handle err
		fmt.Printf("%s", err)
	}

	if ok == true {
		// permit alice to read data1
		fmt.Println("ok")
	} else {
		// deny the request, show an error
		fmt.Println("not ok")
	}
}
