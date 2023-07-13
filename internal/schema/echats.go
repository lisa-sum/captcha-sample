package schema

type Echarts struct {
	tableName struct{} `pg:"e_2020_sample"`
	Shirt     int8     `json:"shirt" form:"shirt" pg:"shirt"`
	Cardigan  int8     `json:"cardigan" form:"cardigan" pg:"cardigan"`
	Chiffon   int8     `json:"chiffon" form:"chiffon" pg:"chiffon"`
	Pants     int8     `json:"pants" form:"pants" pg:"pants"`
	Heels     int8     `json:"heels" form:"heels" pg:"heels"`
	Socks     int8     `json:"socks" form:"socks" pg:"socks"`
}
