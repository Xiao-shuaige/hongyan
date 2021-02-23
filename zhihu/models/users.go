package models

import (
	"fmt"
)

type Users struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Id         int64  `form:"_id" json:"_id"`           // id
	NickName   string `form:"nickName" json:"nickName"` // 用户昵称
	Email      string `form:"email" json:"email"`       // 邮件地址
	Password   string `form:"password" json:"password"` // 登录密码
	Role       string `form:"role" json:"role"`         // 角色 [admin 超级管理员] [normal 普通用户]
	Avatar     string `form:"avatar" json:"avatar"`     // 头像
	Status     int    `form:"status" json:"status"`     // 状态 [0 未激活] [1 激活]
	CreateTime string `json:"createTime"`               // 创建时间
}

func CreateUser(user Users, now string) (int64, error) {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return 0, err
	}

	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO users (`nickName`, `email`,`password`,`role`,`avatar`,`status`,`createTime`)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return 0, err
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(user.NickName, user.Email, user.Password, user.Role, user.Avatar, user.Status, now)
	if err != nil {
		fmt.Println("Exec fail")
		fmt.Println(err)
		return 0, err
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	id, _ := res.LastInsertId()
	return id, nil
}

func CheckUser(user Users) (bool, Users) {
	tmpPassWd := user.Password
	err := DB.QueryRow("SELECT * FROM users WHERE email = ?", user.Email).Scan(&user.Id, &user.NickName, &user.Email, &user.Password, &user.Role, &user.Avatar, &user.Status, &user.CreateTime)
	if err != nil {
		fmt.Println("MySql Query err:", err)
		return false, user
	}
	if tmpPassWd == user.Password {
		return true, user
	} else {
		return false, user
	}
}

func GetUserList() []Users {
	//执行查询语句
	rows, err := DB.Query("SELECT * from users")
	if err != nil {
		fmt.Println("Mysql Query err:", err)
	}
	var users []Users
	//循环读取结果
	for rows.Next() {
		var user Users
		//将每一行的结果都赋值到一个user对象中
		err := rows.Scan(&user.Id, &user.NickName, &user.Email, &user.Password, &user.Role, &user.Avatar, &user.Status, &user.CreateTime)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		users = append(users, user)
	}
	return users
}

func GetUserById(id string) (ok bool, user Users) {
	err := DB.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.NickName, &user.Email, &user.Password, &user.Role, &user.Avatar, &user.Status, &user.CreateTime)
	if err != nil {
		fmt.Println("MySql Query err:", err)
		return false, user
	}
	return true, user
}

func UpdateUserById(id string, p2 Users) (ok bool, users Users) {
	//开启事务
	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE users SET nickName = ?, status = ?, role = ?, avatar = ? WHERE id = ?")
	if err != nil{
		fmt.Println("Prepare fail")
		return false,users
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(p2.NickName, p2.Status, p2.Role, p2.Avatar, id)
	if err != nil{
		fmt.Println("Exec fail")
		return false,users
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return GetUserById(id)
}

func DeleteUserById(id string) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil{
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(id)
	if err != nil{
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	//获得上一个insert的id
	fmt.Println(res.LastInsertId())
	return true
}

func ChangePasswd(id int64, oldPass string, newPass string) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE users SET password = ? WHERE id = ? AND password = ?")
	if err != nil{
		fmt.Println("Prepare fail")
		return false
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(newPass, id , oldPass)
	if err != nil{
		fmt.Println("Exec fail")
		return false
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return true
}
