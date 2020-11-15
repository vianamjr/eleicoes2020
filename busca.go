package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func pesquisas()[]jsonCidade{
	c := getCidades()
	a := []jsonCidade{}
	for cidade, codigo := range c{
		fmt.Println(cidade)
		ac := getResultado(codigo)
		ac.Cidade = cidade
		a = append(a, ac)
	}
	return a
}

func getResultado(c string)  jsonCidade{
	url := `https://resultados.tse.jus.br/oficial/ele2020/divulgacao/oficial/426/dados-simplificados/`+c+`-c0011-e000426-r.json`

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)


	var todoStruct jsonCidade
	json.Unmarshal(bodyBytes, &todoStruct)



	//log.Println(string([]byte(bodyBytes)))
	fmt.Println(todoStruct.Pesi)
	for _,v := range todoStruct.Cand{
		candidato(v)
	}
	fmt.Println("-----------")
	return todoStruct
}

func candidato(c Cand){
	fmt.Println("\t", c.Nm, c.Pvap)
}

