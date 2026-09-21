package algs

import "caso1/utils"

func Backtracking() (bool, int, string) {
	var cantConsultas int
	cantConsultas = 0

	longitud, consultasLongitud, flag := utils.EncontrarLongitudPass()
	if flag != "" {
		return true, consultasLongitud, flag
	}
	cantConsultas += consultasLongitud

	//Hace un barrido inicial para:
	// 	- Descartar letras innecesarias
	//  - Saber la longitiud de la contraseña
	// Esta parte del el algoritmo es la poda
	barridoInicial := make([]byte, longitud)

	var abecedarioFiltrado []byte
	var longitudPass uint8

	for letra := byte('a'); letra <= byte('z'); letra++ {
		//Cambia todas las letras a la letra actual
		for i := range barridoInicial {
			barridoInicial[i] = letra
		}

		//Prueba la pass
		respuesta := utils.ProbarPass(string(barridoInicial))
		cantConsultas++
		if respuesta.Status == "SUCCESS" {
			return true, cantConsultas, string(barridoInicial)
		}

		// Si la letra devuelve matchedLetters, agregela al diccionario filtrado y aumente la longitud de la pass
		if respuesta.MatchedLetters > 0 {
			abecedarioFiltrado = append(abecedarioFiltrado, letra)
			longitudPass += respuesta.MatchedLetters
		}

		//Si ya tiene las suficientes, no busca más
		if longitudPass == longitud {
			break
		}
	}

	guess := make([]byte, longitud)
	for i := range guess {
		guess[i] = abecedarioFiltrado[0]
	}

	var matchedLetters uint8

	respuesta := utils.ProbarPass(string(guess))
	cantConsultas++
	matchedLetters = respuesta.MatchedLetters
	if respuesta.Status == "SUCCESS" {
		return true, cantConsultas, string(guess)
	}

	//probar por cada posicion
	for i := range longitudPass {
		// Guardamos la letra que tenía la posición antes de empezar a probar
		letraAnterior := guess[i]

		for j := range len(abecedarioFiltrado) {
			guess[i] = abecedarioFiltrado[j]
			respuesta = utils.ProbarPass(string(guess))
			cantConsultas++
			if respuesta.Status == "SUCCESS" {
				return true, cantConsultas, string(guess)
			}
			if respuesta.MatchedLetters < matchedLetters {
				//Backtracking ------------------------------
				guess[i] = letraAnterior
				break
			}
			letraAnterior = abecedarioFiltrado[j]
			matchedLetters = respuesta.MatchedLetters
		}
	}
	return false, 0, "ERR"
}

// ============================================================================
// ALGORITMO: Backtracking
// COMPLEJIDAD BIG-O: O(N) donde N es la longitud de la contraseña.
// CRECIMIENTO: Lineal (Luego de pruebas empíricas, en promedio, aumenta en ~6 consultas por letra nueva)
//
// CRITERIO DE PODA:
// Prueba todas las letras primero, para armar un abecedario limitado
// donde solo se prueba especificamente las letras que se sabe con certeza
// que estan en la contraseña, reduciendo los tanteos innecesarios.
// ============================================================================
