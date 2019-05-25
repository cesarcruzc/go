package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Solo ejecuta la funcion init del paquete
	"github.com/jinzhu/gorm"
)

//Product is a model to create table in the databes
type Product struct {
	gorm.Model // Has fields: id, createdAt, updatedAt, deletedAt
	BarCode string
	Price   uint
}

func main() {
	db, err := gorm.Open("mysql", "root:development@/golang?charset=utf8&parseTime=true&loc=Local")

	if err != nil {
		panic("Database connection error")
	}

	defer db.Close()
	fmt.Println("Success full connection to database")

	/*db.CreateTable(&Product{})
	fmt.Println("Table product has been created in the database")*/

	/*db.Create(&Product{
	BarCode: "0202020202",
	Price:   5000})*/

	//Query where id
	var p Product
	db.First(&p, 2)
	fmt.Println(p)

	//Update
	p.Price = 369
	db.Save(&p)

	fmt.Println(p.Price)

}
