package main

import (
    "github.com/gin-gonic/gin";    "fmt"; "time";
    "os"

)
func main() {
    router := gin.Default()
	router.GET("/avisaagua/teste", teste)
    router.GET("/avisaagua/leitura", leitura)
    router.Run(":8080")
}
func leitura(c *gin.Context){
	var acao string
	file, err := os.OpenFile("medicoes.txt",os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	currentTime := time.Now()
	s:=	fmt.Sprintf("%d-%d-%d %d:%d:%d",
	currentTime.Year(),
	currentTime.Month(),
	currentTime.Day(),
	currentTime.Hour(),
	currentTime.Minute(),
	currentTime.Second())
	s1 := c.DefaultQuery("sensor1", "null")
	s2 := c.DefaultQuery("sensor2", "null")
	s3 := c.DefaultQuery("sensor3", "null")
	if validar(string(s1)) && validar(string(s2)) && validar(string(s3)){
		if (string(s1) == "ligado" && string(s2) == "ligado") || (string(s1) == "ligado" && string(s2) == "ligado" && string(s3) == "ligado") || (string(s2) == "ligado" && string(s3) == "ligado"){
			c.JSON(200, gin.H{
				"message": "LIGAR MOTOR",
			})
			acao = "LIGOU MOTOR"
		}else if string(s1) == "desligado" && string(s2) == "desligado" && string(s3) == "desligado"{
			c.JSON(200, gin.H{
				"message": "DESLIGAR MOTOR",
			})
			acao = "DESLIGOU MOTOR"
		}else if string(s3) == "ligado" && (string(s2) == "desligado" || string(s1) == "desligado"){
			c.JSON(200, gin.H{
				"message": "SENSOR QUEBROU",
			})
			acao = "SENSOR BUGOU/QUEBROU"
		}else{
			acao = "USUAL"
			c.JSON(200, gin.H{
				"message": "Normal",
			})
			}
		_, err =file.WriteString(s +" Sensor1 = " +s1 +" Sensor2 = " +s2 +" Sensor3 = " +s3 + " Situacao: "+acao +"\n")	

		}else{
			c.JSON(400, gin.H{
				"message": "Error, informacoes faltando ou erradas",
		})
		}
}
func teste(c *gin.Context){

	c.JSON(200, gin.H{
		"message": "teste avisaAgua!",
	})
}
func validar (valor string) bool { 
	padrao := [2]string{"ligado", "desligado"}
	for i:=0;i<len(padrao);i++{
		if padrao[i] == valor{return true}
	}
	return false
}
