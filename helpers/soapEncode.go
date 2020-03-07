package helpers

import(
	"encoding/xml"

)
type SoapEnvelope struct {
	XMLName struct{} `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Body    SoapBody
}

type SoapBody struct {
	XMLName  struct{} `xml:"http://www.w3.org/2003/05/soap-envelope Body"`
	Contents []byte   `xml:",innerxml"`
}

func SoapEncode(contents interface{}) ([]byte, error) {
	data, err := xml.MarshalIndent(contents, "    ", "  ")
	if err != nil {
		return nil, err
	}
	data = append([]byte("\n"), data...)
	env := SoapEnvelope{Body: SoapBody{Contents: data}}
	return xml.MarshalIndent(&env, "", "  ")
}