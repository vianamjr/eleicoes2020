package main

type jsonCidade struct {
	Cidade string
	Ele    string `json:"ele"`
	Tpabr  string `json:"tpabr"`
	Cdabr  string `json:"cdabr"`
	Carper string `json:"carper"`
	T      string `json:"t"`
	F      string `json:"f"`
	Dt     string `json:"dt"`
	Ht     string `json:"ht"`
	Dv     string `json:"dv"`
	Tf     string `json:"tf"`
	V      string `json:"v"`
	Esae   string `json:"esae"`
	Mnae   string `json:"mnae"`
	S      string `json:"s"`
	St     string `json:"st"`
	Pst    string `json:"pst"`
	Snt    string `json:"snt"`
	Psnt   string `json:"psnt"`
	Si     string `json:"si"`
	Psi    string `json:"psi"`
	Sni    string `json:"sni"`
	Psni   string `json:"psni"`
	Sa     string `json:"sa"`
	Psa    string `json:"psa"`
	Sna    string `json:"sna"`
	Psna   string `json:"psna"`
	E      string `json:"e"`
	Ea     string `json:"ea"`
	Pea    string `json:"pea"`
	Ena    string `json:"ena"`
	Pena   string `json:"pena"`
	Esi    string `json:"esi"`
	Pesi   string `json:"pesi"`
	Esni   string `json:"esni"`
	Pesni  string `json:"pesni"`
	C      string `json:"c"`
	Pc     string `json:"pc"`
	A      string `json:"a"`
	Pa     string `json:"pa"`
	Vscv   string `json:"vscv"`
	Vnom   string `json:"vnom"`
	Pvnom  string `json:"pvnom"`
	Vvc    string `json:"vvc"`
	Pvvc   string `json:"pvvc"`
	Vb     string `json:"vb"`
	Pvb    string `json:"pvb"`
	Tvn    string `json:"tvn"`
	Ptvn   string `json:"ptvn"`
	Vn     string `json:"vn"`
	Pvn    string `json:"pvn"`
	Vnt    string `json:"vnt"`
	Pvnt   string `json:"pvnt"`
	Vp     string `json:"vp"`
	Pvp    string `json:"pvp"`
	Vv     string `json:"vv"`
	Pvv    string `json:"pvv"`
	Van    string `json:"van"`
	Pvan   string `json:"pvan"`
	Vansj  string `json:"vansj"`
	Pvansj string `json:"pvansj"`
	Tv     string `json:"tv"`
	Cand   []Cand `json:"cand"`
}
type Cand struct {
	Seq    string `json:"seq"`
	Sqcand string `json:"sqcand"`
	N      string `json:"n"`
	Nm     string `json:"nm"`
	Cc     string `json:"cc"`
	Nv     string `json:"nv"`
	E      string `json:"e"`
	St     string `json:"st"`
	Dvt    string `json:"dvt"`
	Vap    string `json:"vap"`
	Pvap   string `json:"pvap"`
}

func main(){
	pesquisas()
	StartService()
}
