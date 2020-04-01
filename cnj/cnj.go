package cnj

import (
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

var reCnj *regexp.Regexp
var ErrInvalidNumber error = errors.New("Numero inválido")
var ErrInvalidDigit error = errors.New("Dígito verificador inválido")

func init() {
	reCnj = regexp.MustCompile("[^0-9]")
}

// Valida dígitos de verificação de número de processo no padrão do CNJ.
// O parâmetro num deve conter o número, podendo ou não conter caracteres de
// separação como pontos, traços, espaços, etc.
func ValidaCNJ(num string) error {
	var s string
	var n, d uint64
	s = reCnj.ReplaceAllString(num, "")

	s = strings.TrimLeft(s, "0")
	if len(s) == 0 || len(s) > 20 {
		return ErrInvalidNumber
	}
	if len(s) > 11 {
		var smallnum, _ = new(big.Int).SetString(s[:len(s)-13]+s[len(s)-11:]+"00", 10)
		n = smallnum.Uint64()
		if len(s) == 12 {
			var smallnum, _ = new(big.Int).SetString(s[:1], 10)
			d = smallnum.Uint64()
		} else {
			var smallnum, _ = new(big.Int).SetString(s[len(s)-13:len(s)-11], 10)
			d = smallnum.Uint64()
		}
	} else {
		var smallnum, _ = new(big.Int).SetString(s+"00", 10)
		n = smallnum.Uint64()
		d = 0
	}
	modulo := (98 - (n % 97))
	if modulo == d {
		return nil
	}
	return ErrInvalidDigit
}

// Remove todos os caracteres diferentes de dígitos e insere zeros à esquerda para
// garantir que o número tenha 20 digitos
// Não faz validação do número, assumindo que o número fornecido já esteja validado
func Normaliza(num string) string {
	var n uint
	fmt.Sscanf(reCnj.ReplaceAllString(num, ""), "%d", &n)
	return fmt.Sprintf("%020d", n)
}

// Insere pontos e tracos.
// Assume que num é uma string composta de 20 digitos
func Formata(num string) string {
	return num[:7] + "-" + num[7:9] + "." + num[9:13] + "." + num[13:14] + "." + num[14:16] + "." + num[16:]
}
