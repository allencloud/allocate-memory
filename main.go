package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"runtime"
)

var memoryPool = []([]complex128){}
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
		a := make([]complex128, 1024*64, 1024*64)
		memoryPool = append(memoryPool, a)
		runtime.GC()
		memoryAllocatedInMB++
	}

	message := "Allocated about " + strconv.Itoa(memoryAllocatedInMB) + " MB memory."

	c.String(200, message)
}
