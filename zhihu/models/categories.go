package models

import "fmt"

type Categories struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	Id         int64  `form:"_id" json:"_id"`           // id
	ClassName  string `form:"className" json:"className"` // 用户昵称
	Title      string `form:"title" json:"title"`       // 邮件地址
	CreateTime string `json:"createAt"`                 // 创建时间
}

func CreateCategories(cate Categories,now string) (int64, error) {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return 0, err
	}

	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO categories (`className`, `title`, `createTime`)" +
		" VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return 0, err
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(cate.ClassName, cate.Title, now)
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

func GetCategoriesById(id string) (ok bool, cate Categories) {
	err := DB.QueryRow("SELECT * FROM categories WHERE id = ?", id).Scan(&cate.Id, &cate.ClassName, &cate.Title, &cate.CreateTime)
	if err != nil {
		fmt.Println("MySql Query err:", err)
		return false, cate
	}
	return true, cate
}

func GetCategoriesList() []Categories {
	//执行查询语句
	rows, err := DB.Query("SELECT * from categories")
	if err != nil {
		fmt.Println("Mysql Query err:", err)
	}
	var cates []Categories
	//循环读取结果
	for rows.Next() {
		var cate Categories
		//将每一行的结果都赋值到一个user对象中
		err := rows.Scan(&cate.Id, &cate.ClassName, &cate.Title, &cate.CreateTime)
		if err != nil {
			fmt.Println("rows fail")
		}
		//将user追加到users的这个数组中
		cates = append(cates, cate)
	}
	return cates
}

func GetCategoriesCount() (count int) {
	rows, err := DB.Query("SELECT * from categories")
	if err != nil {
		fmt.Println("Mysql Query err:", err)
	}
	count = 0
	//循环读取结果
	for rows.Next() {
		count ++
	}
	return count
}

func UpdateCategoriesById(id string, p2 Categories) (ok bool, cate Categories)  {
	//开启事务
	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("UPDATE categories SET className = ?, title = ? WHERE id = ?")
	if err != nil{
		fmt.Println("Prepare fail")
		return false,cate
	}
	//设置参数以及执行sql语句
	res, err := stmt.Exec(p2.ClassName, p2.Title, id)
	if err != nil{
		fmt.Println("Exec fail")
		return false,cate
	}
	//提交事务
	tx.Commit()
	fmt.Println(res.LastInsertId())
	return GetCategoriesById(id)
}

func DeleteCategoriesById(id string) bool  {
	//开启事务
	tx, err := DB.Begin()
	if err != nil{
		fmt.Println("tx fail")
	}
	//准备sql语句
	stmt, err := tx.Prepare("DELETE FROM categories WHERE id = ?")
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