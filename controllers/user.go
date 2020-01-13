package controllers

import (
	"encoding/json"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/fogetu/cloud_miner_grpc/models"
	"github.com/fogetu/go_system/system_grpc"
	"github.com/fogetu/miner_service_intf/mine_intf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.Aaode
// @Failure 403 body is empty
//@router /getall [get]
func (u *UserController) GetAll() {
	var conn *grpc.ClientConn
	conn = system_grpc.GetConn("/mine/pool")
	c := mine_intf.NewPoolClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var FilterField mine_intf.PoolFilterField;
	FilterField.Id = "aa2111";
	poolListResponse, err := c.GetList(ctx, &mine_intf.PoolListRequest{FilterField: &FilterField,Page: 1, PageSize: 1})
	// _, err := c.GetList(ctx, &mine_intf.PoolListRequest{Page: 1, PageSize: 50})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// mystruct := models.Aaode{123}
	u.Data["json"] = *poolListResponse
	// u.Data["json"] = mystruct
	u.ServeJSON()
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// mystruct := models.Aaode{123}
	// u.Data["json"] = &mystruct
	// u.ServeJSON()
	//c := mine_intf.NewPoolClient(conn)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	//r, err := c.GetList(ctx, &mine_intf.PoolListRequest{Page: 1, PageSize: 50})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.Data)
	//mystruct := models.Aaode{123}
	//u.Data["json"] = &mystruct
	//u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
