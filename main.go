package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"runtime"
	"strconv"
)

var memoryPool = [](*[1024 * 64]complex128){}
var memoryAllocatedInMB int = 0

func main() {
	router := gin.Default()
	router.GET("/memory/:size/action/allocate", AllocateMemory)
	router.Run(":8080")
}

func AllocateMemory(c *gin.Context) {
	size, err := strconv.Atoi(c.Param("size"))
	if err != nil {
		c.String(http.StatusBadRequest, "memory allocate input should be an interger.")
		return
	}

	if size <= 0 {
		c.String(http.StatusBadRequest, "memory allocate input should be larger than 0.")
		return
	}

	for i := 0; i < size; i++ {
		memoryPool = append(memoryPool, &([1024 * 64]complex128{}))
		//runtime.GC()
		memoryAllocatedInMB++
	}

	message := "Allocated about " + strconv.Itoa(memoryAllocatedInMB) + " MB memory."

	c.String(200, message)
}
