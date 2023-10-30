package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

var currentImage *imageupload.Image

func UploadImage(c *gin.Context) {
	fmt.Println("Hola")

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 10*1024*1024)

	file, handler, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Fallo al leer imagen...",
		})
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded file name: %+v\n", handler.Filename)
	fmt.Printf("Uploaded file size %+v\n", handler.Size)
	fmt.Printf("File mime type %+v\n", handler.Header)

	// Get the file content type and access the file extension
	fileType := strings.Split(handler.Header.Get("Content-Type"), "/")[1]

	// Create the temporary file name
	fileName := fmt.Sprintf("upload-*.%s", fileType)
	// Create a temporary file with a dir folder
	tempFile, err := os.CreateTemp("pictures", fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)
	c.JSON(http.StatusOK, gin.H{})
}
