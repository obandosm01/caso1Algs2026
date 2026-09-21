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
func EncontrarLongitudPass() (uint8, int, string) {
	pass := make([]byte, 2)
	cantConsultas := 0

	var longitud uint8
	longitud = 0
	for letra := byte('a'); letra <= byte('e'); letra++ {
		for i := range pass {
			pass[i] = letra
		}

		resp := ProbarPass(string(pass))
		cantConsultas++

		//Si dio SUCCESS, el servidor ya cambió la contraseña
		if resp.Status == "SUCCESS" {
			return 2, cantConsultas, string(pass)
		}

		if resp.MatchedLetters == 0 && resp.Distance > 0 {
			if resp.Distance > longitud {
				longitud = resp.Distance
			}
		}
	}

	return longitud, cantConsultas, ""
}
