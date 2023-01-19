// declarar meu pacote principal
package main
// importar função fmt
import "fmt"

// declarar variáveis
const ebulicaoF = 212.0

// função principal
func main() {

	tempF := ebulicaoF // varialvel do valor da temperatura em fahrenheit
	tempC := (tempF - 32) * 5 / 9 // variavel do valor da temperatura em celsius
	// aparecer na tela o valor da temperatura em fahrenheit e celsius
	fmt.Println("A temperatura de ebulição da água é °F", tempF)
	fmt.Println("A temperatura de ebulição da água é °C", tempC)
	fmt.Printf("A temperatura de ebulição da água é °F %g ou °C %g", tempF, tempC)
	//Esperamos que apareça na tela o valor da temperatura de eboliçao da agua em fahrenheit e celsius
}