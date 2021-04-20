package taobaomodels

type Item struct {
	NumIid        int64          `json:"num_iid"`
	Title         string         `json:"title"`
	Nick          string         `json:"nick"`
	Type          string         `json:"type"`
	Cid           int64          `json:"cid"`
	SellerCids    string         `json:"seller_cids"`
	Props         string         `json:"props"`
	PicUrl        string         `json:"pic_url"`
	Num           int            `json:"num"`
	ValidThru     int            `json:"valid_thru"`
	ListTime      TaobaoDatetime `json:"list_time"`
	DelistTime    TaobaoDatetime `json:"delist_time"`
	Price         string         `json:"price"`
	HasDiscount   bool           `json:"has_discount"`
	HasInvoice    bool           `json:"has_invoice"`
	HasWarranty   bool           `json:"has_warranty"`
	Modified      TaobaoDatetime `json:"modified"`
	ApproveStatus string         `json:"approve_status"`
	PostageId     int            `json:"postage_id"`
	OuterId       string         `json:"outer_id"`
	IsVirtual     bool           `json:"is_virtual"`
	IsTaobao      bool           `json:"is_taobao"`
	IsEx          bool           `json:"is_ex"`

	SoldQuantity int64 `json:"sold_quantity"` // onsale才有
	IsCspu       bool  `json:"is_cspu"`       // onsale才有
}
