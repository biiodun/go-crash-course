package main

import ("net/http"
		"errors"
		"github.com/gin-gonic/gin")

type todo struct{
	Id					string	`json:"id"`
	Item				string	`json:"todo_title"`
	Completed_status	bool	`json:"completed"`

}
var todos= [] todo{
	{"1","Stakeholder meeting", false},
	{"2", "Industry report meeting",false},
	{"3", "Gbemisola scheduling meeting",false},
}

// function to handle GET Todos Request
func getTodos(context *gin.Context){
	 context.IndentedJSON(http.StatusOK,todos)
}

func addTodo (c *gin.Context){
	var newTodo todo
	err := c.BindJSON(&newTodo)
	if(err != nil){
		return
	}
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated,newTodo)
}

func getTodoById (id string) (*todo, error){
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo Not found")
}

func getTodo (c *gin.Context){
	id := c.Param("id")
	todo, err := getTodoById(id)

	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Todo not found"})
		return
	}
	c.IndentedJSON(http.StatusOK,todo)
}

func main(){

	router := gin.Default();
	router.GET("/todos",getTodos)
	router.POST("/new-todo",addTodo)
	router.GET("/todos/:id",getTodo)
	



	router.Run("localhost:8080")
}