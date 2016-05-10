package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	//"github.com/satori/go.uuid"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

var (
	memoryPool              = [](*[1024 * 64]complex128){}
	memoryAllocatedInMB int = 0
	release_interval    int
)

const RELEASE_INTERVAL int = 180

func init() {
	if os.Getenv("RELEASE_INTERVAL") != "" {
		value, err := strconv.Atoi(os.Getenv("RELEASE_INTERVAL"))
		if err != nil {
			panic(err)
		}

		if value <= 0 {
			panic(fmt.Errorf("RELEASE_INTERVAL(%d) should be greater than 0.)", value))
		}

		release_interval = value
	} else {
		release_interval = RELEASE_INTERVAL
	}
}

func main() {
	router := gin.Default()

	// just the ping api
	router.GET("/", func(c *gin.Context) {
		c.String(200, "OK")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.GET("/_ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// consume 256KB memory immediately,
	// and release the memory after 10s
	router.GET("/memory", AllocateQuotaMemory)

	// comsume :size memory of server
	// unit is MB
	router.GET("/memory/:size/action/allocate", AllocateCustomMemory)

	// consume as much cpu as it can
	router.GET("/cpu", ConsumeCPU)

	router.Run(":8080")
}

func AllocateQuotaMemory(c *gin.Context) {
	//id := uuid.NewV4().String()
	go func() {
		cmd := exec.Command("./memory.out")
		if err := cmd.Run(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	time.Sleep(2 * time.Second)

	c.String(http.StatusOK, "OK")
}

func AllocateCustomMemory(c *gin.Context) {
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

func ConsumeCPU(c *gin.Context) {
	cmd := exec.Command("bash", "-c", "awk 'BEGIN{while (i=1) {}}'")
	if err := cmd.Run(); err != nil {
		c.String(500, err.Error())
	}
}
