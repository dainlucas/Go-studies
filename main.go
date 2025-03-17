package main

import (
	"github.com/gin-gonic/gin"
)

//type Pizza struct {
//ID int
//nome string
//preco float64
//}

func main() {

	//criação de instância do gin
	router := gin.Default()
	router.GET("/pizzas", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pizzas": "Frango com Catupiry",
		})
	})

	router.Run() 

	// declaração de variável nome da Pizzaria
	//var nomePizzaria string = "Pizzaria do Go_rdo"

	// declaração de variavel pizzas e seus objetos
	//var pizzas = []Pizza{
		//{ID: 1, nome: "Toscana", preco: 49.5},
		//{ID: 2, nome: "Marguerita", preco: 79.5},
		//{ID: 3, nome: "Atum com queijo", preco: 69.5},

	//}

	// impressão de variáveis
	//fmt.Println(nomePizzaria)
	//fmt.Println(pizzas)
}