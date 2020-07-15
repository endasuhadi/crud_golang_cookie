package main

import(
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Item struct {
	Name string
}

func main() {
	_,_ = fmt.Println("")
	router := gin.Default()
	router.Static("/assets", "./assets")

	router.LoadHTMLGlob("views/*")

	router.GET("/", index)
	router.GET("/form", form)
	router.GET("/del/:index", hapus)
	router.GET("/edit/:index", edit)
	router.POST("/update/:index", update)
	router.POST("/save", save)

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

func edit(c *gin.Context){
	var rItem []Item
	index := c.Param("index")
	i, _ := strconv.Atoi(index)
	cookie, _ := c.Cookie("data_item")
	_ = json.Unmarshal([]byte(cookie), &rItem)
	c.HTML(http.StatusOK, "edit.tmpl", gin.H{
		"title": rItem[i],
		"index": index,
	})
}

func update(c *gin.Context){
	var rItem []Item
	index := c.Param("index")
	item := c.PostForm("item")
	i, _ := strconv.Atoi(index)
	cookie, _ := c.Cookie("data_item")
	_ = json.Unmarshal([]byte(cookie), &rItem)
	
	rItem[i] = Item{Name: item}

	data, _ := json.Marshal(rItem);
	c.SetCookie("data_item", string(data), 3600, "/", "", false, true)
	
	c.HTML(http.StatusOK, "delete.tmpl", gin.H{
		"title": "Delete success",
	})
}

func hapus(c *gin.Context){
	var rItem []Item
	index := c.Param("index")
	i, _ := strconv.Atoi(index)
	cookie, _ := c.Cookie("data_item")
	_ = json.Unmarshal([]byte(cookie), &rItem)

	rItem[i] = rItem[len(rItem)-1]
	rItem[len(rItem)-1] = Item{}
	rItem = rItem[:len(rItem)-1]

	data, _ := json.Marshal(rItem);
	c.SetCookie("data_item", string(data), 3600, "/", "", false, true)
	
	c.HTML(http.StatusOK, "delete.tmpl", gin.H{
		"title": "Delete success",
	})
}

func save(c *gin.Context){
	
	var rItem []Item

	item := c.PostForm("item")

	rItem = append(rItem, Item{Name: item})

	cookie, err := c.Cookie("data_item")

	data, _ := json.Marshal(rItem);

	if err != nil {
		c.SetCookie("data_item", string(data), 3600, "/", "", false, true)
	}else{
		_ = json.Unmarshal([]byte(cookie), &rItem)
		rItem = append(rItem, Item{Name: item})
		data, _ := json.Marshal(rItem);
		c.SetCookie("data_item", string(data), 3600, "/", "", false, true)
	}

	c.Redirect(http.StatusMovedPermanently, "/")

}

func form(c *gin.Context){
	c.HTML(http.StatusOK, "form.tmpl", gin.H{
		"title": "My item",
	})
}

func index(c *gin.Context) {
	var rItem []Item
	for i := 0; i < 10; i++ {
		_ = append(rItem, Item{Name: "Nama: " + strconv.Itoa(i)})
	}

	cookie, _ := c.Cookie("data_item")

	_ = json.Unmarshal([]byte(cookie), &rItem)

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "My item",
		"List": rItem,
	})
}