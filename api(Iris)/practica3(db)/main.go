package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // con el guion bajo solo se ejecuta la funcion init de la lib
	"github.com/jinzhu/gorm"
)

// Producto contiene los datos de un articulo
type Producto struct { // para que la estructura funcione bien, se necesita gorm.Model, que contiene los siguiente campos
	gorm.Model   // ID, CreatedAt, UpdatedAt, DeleteAt
	CodigoBarras string
	Precio       uint
}

func main() {
	// coneccion a la base de datos -> ("base de datos","user_db:password_db@/nombre_db? lo que sigue todo igual")
	db, err := gorm.Open("mysql",
		"root:@/prueba?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	fmt.Println("Coneccion con exito!")

	// para crear la base de datos con la structura Productos
	// db.CreateTable(&Producto{}) // gorm usa los nombres de las tablas en plural
	// fmt.Println("Se creo la tabla Productos en la bd")

	// para agregar reguistros en una tabla
	// db.Create(&Producto{CodigoBarras: "01010125454", Precio: 1200})

}
