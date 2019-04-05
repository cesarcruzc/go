package main

import "fmt"

func main() {
	fmt.Println("Conectando a la base de datos...")
	defer closeDB()

	fmt.Println("Query to info, data set")
	defer closeDataSet()
}

func closeDB() {
	fmt.Println("Close DB")
}

func closeDataSet() {
	fmt.Println("Close dataSet")
}
