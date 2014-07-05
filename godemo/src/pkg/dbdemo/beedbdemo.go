package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
)

type Users struct {
	Id       int `PK` //如果表的主键不是 id，那么需要加上 pk 注释，显式的说这个字段是主键
	Login    string
	Password string
	Role_id  int
	Email    string
	First    string
	Last     string

	Locale                 string
	Default_testproject_id string
	Active                 int
	Script_key             string
}

type Components struct {
	components_id  int
	name           string
	title          string
	itype          int
	native         int
	component_type int
	symbol         string
	desc           string
}

func main() {
	db, err := sql.Open("mymysql", "dz:123456@tcp(192.168.11.110:3306)/testlink?charset=utf8")

	if err != nil {
		panic(err)
	}

	orm := beedb.New(db)
	beedb.OnDebug = true

	var user Users

	//Where 接受两个参数，支持整形参数
	orm.Where("id=?", 18).Find(&user)

	fmt.Println(user.Password)

}
