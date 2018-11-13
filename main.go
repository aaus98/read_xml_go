package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html/charset"
)

//Invoice Factura
type Invoice struct {
	Supplier Supplier `xml:"AccountingSupplierParty"`
}

//Supplier emisor
type Supplier struct {
	Ruc   string `xml:"CustomerAssignedAccountID"`
	Party Party  `xml:"Party"`
}

//Party contenido del emisor
type Party struct {
	LegalEntity   LegalEntity   `xml:"PartyLegalEntity"`
	PostalAddress PostalAddress `xml:"PostalAddress"`
}

//LegalEntity data de la empresa
type LegalEntity struct {
	NameEnterprise string `xml:"RegistrationName"`
}

//PostalAddress direccion del emisor
type PostalAddress struct {
	AddressTypeCode     int     `xml:"AddressTypeCode"`
	CityName            string  `xml:"CityName"`
	StreetName          string  `xml:"StreetName"`
	CitySubdivisionName string  `xml:"CitySubdivisionName"`
	CountrySubentity    string  `xml:"CountrySubentity"`
	District            string  `xml:"District"`
	Country             Country `xml:"Country"`
}

//Country pais del emisor
type Country struct {
	IdentificationCode string `xml:"IdentificationCode"`
}

var invoice Invoice

func index(w http.ResponseWriter, r *http.Request) {
	xmlFile, err := os.Open("20480072872-01-FB99-00040.xml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)

	decoder.CharsetReader = charset.NewReaderLabel

	if err2 := decoder.Decode(&invoice); err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a, _ := json.Marshal(invoice)

	w.Header().Set("Content-Type", "application/json")
	w.Write(a)
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
