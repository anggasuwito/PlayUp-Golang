package helper

import (
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//CopyImageAndRename CopyImageAndRename
func CopyImageAndRename(photo multipart.File, handler *multipart.FileHeader, username string) string {
	dir, err := os.Getwd()
	fmt.Println("dir = ", dir)
	fmt.Println("dir err = ", err)

	rand.Seed(time.Now().UnixNano())
	min := 11111111111
	max := 99999999999
	//ganti "userPhoto-" dengan username dari parameter agar tiap user beda newPhotoName prefix sesuai nama usernamenya
	newPhotoName := "userPhoto-" + strconv.Itoa(rand.Intn(max-min+1)+min) + filepath.Ext(handler.Filename)
	fmt.Println("new Photo Name = ", newPhotoName)
	fmt.Println(filepath.Dir)
	fileLocation := filepath.Join("uploads", newPhotoName)
	fmt.Println("File location = ", fileLocation)

	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	fmt.Println("targetFile = ", targetFile)
	fmt.Println("targetFile err = ", err)

	copyFile, err := io.Copy(targetFile, photo)
	fmt.Println("copyFile = ", copyFile)
	fmt.Println("copyFile err = ", err)

	return newPhotoName
}
