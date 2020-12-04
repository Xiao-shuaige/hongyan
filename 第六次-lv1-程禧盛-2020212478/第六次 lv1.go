package main


import (
"database/sql"
"fmt"
"log"
"os"
)
import _ "github.com/go-sql-driver/mysql"

func insertDB(p *sql.DB,apon string)  {
	stmt, err := p.Prepare("insert into u_1 (c_3) apon (?)")
	if err != nil{
		log.Fatal(err)
	}
	stmt.Exec(apon)
}

func query(p *sql.DB,name string,age int)  {
	var sql string
	sql="select name,age"
	err:=p.QueryRow(sql,1).Scan(&name,&age)
	if err!=nil{
		os.Exit(1)

	}
	fmt.Printf("name:%s,age:%d",name,age)
}
func anotherinsertDB(p *sql.DB,a,b int)  {
	p.Exec("insert u_1 (c_1,c_2) values (?,?)",a,b)
}

func selectDB(db *sql.DB)  {
	stmt, err := db.Query("select * from test;")
	if err != nil {
		log.Fatal(err)

	}
	defer stmt.Close()
	for stmt.Next(){
		var id int
		var name string
		err := stmt.Scan(&id,&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id,name)
	}

}

func deleteDB(db *sql.DB,name string)  {
	stmt, err := db.Prepare("delete from test where name = ?",name)
	if err != nil{
		log.Fatal(err);
	}
	stmt.Exec();
}

func updateDB(db *sql.DB,name string,id int)  {
	stmt, err := db.Prepare("update test set name = ? where id = ?",name,id)
	if err != nil{
		log.Fatal(err)
	}
	stmt.Exec();
}

func main()  {
	a,err:=sql.Open("mysql", "mysql.infoschema:cxs20030416@/databasename")
	if err!=nil {
		panic(err.Error())
	}
	defer a.Close()
	var p *sql.DB
	var apon string
	insertDB(p,apon)
	var th int
	a.QueryRow("select a name  ").Scan(&th)
	fmt.Println("it is",th)
	var f float64
	a.QueryRow("select another").Scan(&f)
	fmt.Println("it is",a)
	var t int
	var d int
	anotherinsertDB(p,t,d)
	selectDB(p)
	tx1,err:=a.Begin(){
		if err!=nil{
			os.Exit(1)
		}
	}
	defer func(){
		if err:=recover();err!=nil{
			fmt.Printf("%+v\n",err)
			fmt.Println("======")
		}
	}
	var str string
	str="hello"
	deleteDB(p,str)
	var id int
	id=2
	updateDB(p,str,id)
}

