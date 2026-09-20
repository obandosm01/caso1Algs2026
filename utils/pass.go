package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

func ProbarPass(pass string) RespuestaServer {
	//Abre la conexión
	conexion, err := net.Dial("tcp", "localhost:4000")
	if err != nil {
		panic(err)
	}
	defer conexion.Close() //Cierra la conexión al tener un error

	//Arma la solicitud
	solicitud := fmt.Sprintf(`{"password":"%s"`+"\n", pass)

	//Manda la solicitud
	_, err = conexion.Write([]byte(solicitud))
	if err != nil {
		panic(err)
	}

	//Lee la respuesta
	reader := bufio.NewReader(conexion)
	respuestaStr, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	var resp RespuestaServer
	err = json.Unmarshal([]byte(respuestaStr), &resp) //la parsea a al objeto RespuestaServer
	if err != nil {
		panic(err)
	}

	return resp
}
func EncontrarLongitudPass() (uint8, int) {
	pass := make([]byte, 2)
	cantConsultas := 0
	for letra := byte('a'); letra <= byte('z'); letra++ {
		for i := range pass {
			pass[i] = letra
		}
		resp := ProbarPass(string(pass))
		cantConsultas++
		if resp.Status == "SUCCESS" {
			return 2, cantConsultas
		}

		if resp.MatchedLetters == 0 && resp.Distance > 0 {
			return resp.Distance, cantConsultas
		}
	}
	return 4, cantConsultas
}
