package algs

import "caso1/utils"

func Greedy() (bool, int, string) {
	guess := []byte{'a', 'a'}

	cantConsultas := 0

	resp := utils.ProbarPass(string(guess))
	cantConsultas++
	if resp.Status == "SUCCESS" {
		return true, cantConsultas, string(guess)
	}

	if resp.MatchedLetters == 1 {
		guess = []byte{'a', 'b'}
		resp = utils.ProbarPass(string(guess))
		cantConsultas++
		if resp.Status == "SUCCESS" {
			return true, cantConsultas, string(guess)
		}

		if resp.MatchedLetters == 1 {
			guess = []byte{'b', 'a'}
			resp = utils.ProbarPass(string(guess))
			cantConsultas++
			if resp.Status == "SUCCESS" {
				return true, cantConsultas, string(guess)
			}

			if resp.MatchedLetters == 0 {
				guess = []byte{'a', 'a'}
			}
		}
	}
	// Cada posicion es la fase
	i := 0       //Posicion de la letra actual
	letra := 'a' //Letra actual
	matchedLettersActual := int(resp.MatchedLetters)
	for {
		if i >= len(guess) {
			if len(guess) < 4 {
				guess = append(guess, 'a')
			} else {
				i = 0 // Si ya mide 4, vuelve a la primera posición, en caso de que necesite
			}
		}

		guess[i] = byte(letra)

		resp = utils.ProbarPass(string(guess))
		cantConsultas++
		if resp.Status == "SUCCESS" {
			return true, cantConsultas, string(guess)
		}

		if int(resp.MatchedLetters) > matchedLettersActual { //Optimo local de la fase (posicion) es la letra correcta
			matchedLettersActual = int(resp.MatchedLetters)
			i++
			letra = 'a'
		} else {
			letra++
			if letra > 'z' {
				// Si recorrió de la 'a' a la 'z' sin aumentar, reinicia la letra y avanza posición
				letra = 'a'
				i++
			}
		}
	}

}

// ============================================================================
// ALGORITMO: Greedy
// COMPLEJIDAD BIG-O: O(N) donde N es la longitud de la contraseña.
// CRECIMIENTO: Lineal (Luego de pruebas empíricas, en promedio, aumenta en ~15 consultas por letra nueva)
//
// ETAPAS Y ÓPTIMO LOCAL:
// 1. Etapa de prueba de casos limite: Prueba "aa", "ba", "ab" para deducir la posición inicial.
// 2. Cada posición de la contraseña es otra etapa adicional.
// Óptimo Local: En cada posición, elige la letra del abecedario que
//    aumenta en 1 MatchedLetters y pasa a la siguiente posición ignorando el panorama global.
//	  Asume que esa letra es la correcta para esa posición.
// ============================================================================
