package protheus

import "fmt"

type pedido struct {
	Fecha string `json:"fecha"`
	Items []struct {
		Oferta   bool   `json:"oferta"`
		ProdDesc string `json:"prod_desc"`
		ProdCod  string `json:"prod_cod"`
		Cantidad struct {
			Liberada  int `json:"liberada"`
			Pedido    int `json:"pedido"`
			Entregada int `json:"entregada"`
		} `json:"cantidad"`
		Deposito string `json:"deposito"`
		Item     string `json:"item"`
		Precio   struct {
			Venta    float64 `json:"venta"`
			Lista    float64 `json:"lista"`
			Total    float64 `json:"total"`
			Unitario float64 `json:"unitario"`
		} `json:"precio"`
	} `json:"items"`
	TipoID    string `json:"tipo_id"`
	FluigID   string `json:"fluig_id"`
	Pedido    string `json:"pedido"`
	CliCod    string `json:"cli_cod"`
	TranspID  string `json:"transp_id"`
	CliLoj    string `json:"cli_loj"`
	Condicion string `json:"condicion"`
}

type PedidoResponse struct {
	Consulta struct {
		ID string `json:"id"`
	} `json:"consulta"`
	Data     pedido `json:"data"`
	Endpoint string `json:"endpoint"`
}

// GetPedido retorna un pedido
// Devuelve información de un pedido de venta a partir de su ID.
func (g *Protheus) GetPedido(numero string) (*PedidoResponse, error) {
	pedido := PedidoResponse{}
	err := g.get("/pedidos/"+numero, nil, &pedido)
	if err != nil {
		return nil, err
	}

	return &pedido, nil
}

type RequestNewPedido struct {
	Data []struct {
		Nro        int     `json:"nro"`
		Cod        string  `json:"cod"`
		Cant       int     `json:"cant"`
		Dto        float64 `json:"dto"`
		CliID      string  `json:"cli_id"`
		CliLj      string  `json:"cli_lj"`
		Fecha      string  `json:"fecha"`
		Cond       string  `json:"cond"`
		Entrega    string  `json:"entrega"`
		Vendedor   string  `json:"vendedor"`
		TipoRemito string  `json:"tipoRemito"`
		TipoVenta  string  `json:"tipoVenta"`
		Precio     float64 `json:"precio"`
		Deposito   int     `json:"deposito"`
	} `json:"data"`
}

type NewPedidoResponse struct {
	Data struct {
		Pedidos []string `json:"pedidos"`
	} `json:"data"`
}

// Genera un Pedido de Venta.
// Devuelve el numero de PV generado.
func (g *Protheus) CreatePedido(pedido *RequestNewPedido) (*interface{}, error) {
	var pedidoResponse interface{}
	err := g.post("/pedidos", &pedido, &pedidoResponse)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return &pedidoResponse, nil
}

type LiberarPedidoResponse struct {
	Consulta struct {
		ID string `json:"id"`
	} `json:"consulta"`
	Data struct {
		Found bool `json:"found"`
	} `json:"data"`
	Endpoint string `json:"endpoint"`
}

// Libera un pedido de venta completo por credito y por stock.
func (g *Protheus) LiberarPedido(numeroPedido string) (*LiberarPedidoResponse, error) {
	liberarPedidoResponse := LiberarPedidoResponse{}
	err := g.put("/pedidos/liberate/"+numeroPedido, nil, liberarPedidoResponse)
	if err != nil {
		return nil, err
	}

	return &liberarPedidoResponse, nil
}
