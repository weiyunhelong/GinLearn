package models

import (
	"log"
	"fmt"
 db "GinLearn/GinLearn/database"
)
//表结构
type Person struct {
   Id        int    `json:"id" form:"id"`
   FirstName string `json:"first_name" form:"first_name"`
   LastName  string `json:"last_name" form:"last_name"`
}

//新增记录
func (p *Person) AddPerson() bool {
   rs, err := db.SqlDB.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
   if err != nil {
     return false
   }
   id, err := rs.LastInsertId()
   fmt.Println(id)
   if err!=nil{
     return false
   }else{
     return true
   }
}

//修改记录
func (p *Person) EditPerson() bool {
  rs, err := db.SqlDB.Exec("UPDATE person set first_name=?,last_name=? where id=?", p.FirstName, p.LastName,p.Id)
  if err != nil {
    return false
  }
  id, err := rs.RowsAffected()
  fmt.Println(id)
  if err!=nil{
    return false
  }else{
    return true
  }
}

//删除记录
func DeletePerson(Id int) bool {
  rs, err := db.SqlDB.Exec("Delete From person where id=?", Id)
  if err != nil {
    return false
  }
  id, err := rs.RowsAffected()
  fmt.Println(id)
  if err!=nil{
    return false
  }else{
    return true
  }
}

//得到记录列表
func GetPersonList(pageno,pagesize int,search string) (persons []Person) {

    fmt.Println("搜索参数:"+search)
    persons = make([]Person, 0)
    //SQL查询分页语句
    if search!=""{
      rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person where 1=1 and last_name like '%"+search+"%' or first_name like '%"+search+"%' limit ?,?",(pageno-1)*pagesize,pagesize)
      if err != nil {
        return nil
       }
       defer rows.Close()
    
       //数据添加到数据集中
       for rows.Next() {
         var person Person
         rows.Scan(&person.Id, &person.FirstName, &person.LastName)
         persons = append(persons, person)
       }
       if err = rows.Err(); err != nil {
         return nil
       }
     
    }else{
      rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person where 1=1  limit ?,?",(pageno-1)*pagesize,pagesize)
      if err != nil {
        return nil
      }
      defer rows.Close()

    
     //数据添加到数据集中
     for rows.Next() {
       var person Person
       rows.Scan(&person.Id, &person.FirstName, &person.LastName)
       persons = append(persons, person)
      }
      if err = rows.Err(); err != nil {
       return nil
     }
  }
  return persons
}
//得到记录数
func GetRecordNum(search string) int {
   num:=0;

  //SQL查询分页语句
  if search!=""{
    rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person where 1=1 and first_name like '%?%' or last_name '%?%'",search,search)
    if err != nil {
      return 0
     }
     defer rows.Close()
  
     //数据添加到数据集中
     for rows.Next() {
       num++;
     }
   
  }else{
    rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person where 1=1")
    if err != nil {
      return 0
    }
    defer rows.Close()

  
   //数据添加到数据集中
  //数据添加到数据集中
  for rows.Next() {
    num++;
  }
  
}
return num
}
//得到用户数据
func  GetPersonById(Id int) (p *Person) {
 
  var person Person
  //根据ID查询得到对象
  err := db.SqlDB.QueryRow("SELECT id, first_name, last_name FROM person WHERE id=?", Id).Scan(
    &person.Id, &person.FirstName, &person.LastName,
   )
 
   //打印错误
   if err != nil {
    log.Println(err)
   }
   //返回对象
   return &person
}