package contractor
import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"ransomwarecheck/models"

)

func Create() {

	file, err := os.Create(RandomString(10) + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	models.File = "" + file.Name()
	fmt.Println("File " + models.Pathes + "\\" + models.File + " created successfully")
	models.My_slice = append(models.My_slice, models.Pathes+"\\"+models.File)
}

//// Random string creator for file
func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
