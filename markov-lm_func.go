/* Ricevo da linea di comando due interi k e n.
 * Da stdin leggo un testo e genero il modello linguistico markoviano di ordine k (ogni carattere dipende dai k precedenti).
 * Alla fine genero sulla base di questo un testo di lunghezza n.
	./markov-lm k lunghezza <file.txt*/
package main

import "strconv"
import "os"
import "fmt"
import "bufio"
import "math/rand"

func Markov (k int, n int) []rune {


	//Manda stringhe di lunghezza k in stringhe di lunghezza 1. Ad esempio alla chiave "th" saranno associati tutti i caratteri che nel testo compaiono dopo "th"
	modello := make(map[string][]string)

	scanner := bufio.NewScanner(os.Stdin)
	tuttoTesto := ""
	for scanner.Scan() {
		riga := scanner.Text()
		tuttoTesto += riga + "\n" //Mette tutto il testo in un'unica stringa lunghissima.
	}

	inizi := []string{}
	testo := []rune(tuttoTesto)
	for i := 0; i < len(testo) - k - 1; i++ {
		prima := string(testo[i:i+k])
		inizi = append(inizi, prima)
		dopo := string(testo[i+k])
		modello[prima] = append(modello[prima], dopo)
	}
	generato := []rune(inizi[rand.Intn(len(inizi))])
	for len(generato) < n {
		key := string(generato[len(generato) - k:])
		valori := modello[key]
		valore := valori[rand.Intn(len(valori))]
		generato = append(generato, ([]rune(valore))[0])
	}
	return generato
}

func main () {
	k, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[2])
	fmt.Println(string(Markov(k, n)))
}
