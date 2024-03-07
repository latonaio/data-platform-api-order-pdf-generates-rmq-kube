package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
	OrderPdf            string      `json:"order_pdf"`
	MountPath           *string     `json:"mount_path"`
}

type Message struct {
	Header []Header `json:"Header"`
}

type Header struct {
	OrderID               string `json:"OrderID"`
	OrderStatus           string `json:"OrderStatus"`
	OrderDate             string `json:"OrderDate"`
	OrderType             string `json:"OrderType"`
	Buyer                 int    `json:"Buyer"`
	BuyerName             string `json:"BuyerName"`
	Seller                int    `json:"Seller"`
	SellerName            string `json:"SellerName"`
	RequestedDeliveryDate string `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime string `json:"RequestedDeliveryTime"`
	TotalGrossAmount      string `json:"TotalGrossAmount"`
	//TotalGrossAmount					float32	`json:"TotalGrossAmount"`
	Contract                        *int                       `json:"Contract"`
	ContractItem                    *int                       `json:"ContractItem"`
	Project                         *int                       `json:"Project"`
	ProjectDescription              *string                    `json:"ProjectDescription"`
	WBSElement                      *int                       `json:"WBSElement"`
	WBSElementResponsiblePersonName *string                    `json:"WBSElementResponsiblePersonName"`
	Incoterms                       *string                    `json:"Incoterms"`
	IncotermsName                   *string                    `json:"IncotermsName"`
	PricingDate                     string                     `json:"PricingDate"`
	PaymentTerms                    string                     `json:"PaymentTerms"`
	PaymentTermsName                string                     `json:"PaymentTermsName"`
	PaymentMethod                   string                     `json:"PaymentMethod"`
	TransactionCurrency             string                     `json:"TransactionCurrency"`
	HeaderText                      string                     `json:"HeaderText"`
	OrdersItem                      []OrdersItem               `json:"OrdersItem"`
	OrdersItemPricingElement        []OrdersItemPricingElement `json:"OrdersItemPricingElement"`
	OrdersPartner                   []OrdersPartner            `json:"OrdersPartner"`
	TotalOrderQuantityInBaseUnit    string                     `json:"TotalOrderQuantityInBaseUnit"`
	// PDF 質量合計
	TotalOrderQuantityInDeliveryUnit string `json:"TotalOrderQuantityInDeliveryUnit"`
	// PDF 員数合計
}

type OrdersItem struct {
	OrderID   int `json:"OrderID"`
	OrderItem int `json:"OrderItem"`
	//	OrderStatus                 string  `json:"OrderStatus"`
	//	OrderItemCategory           string  `json:"OrderItemCategory"`
	DeliverToPlant       string  `json:"DeliverToPlant"`
	DeliverToPlantName   string  `json:"DeliverToPlantName"`
	DeliverFromPlant     string  `json:"DeliverFromPlant"`
	DeliverFromPlantName string  `json:"DeliverFromPlantName"`
	Product              string  `json:"Product"`
	ProductSpecification *string `json:"ProductSpecification"`
	SizeOrDimensionText  *string `json:"SizeOrDimensionText"`
	OrderItemText        string  `json:"OrderItemText"`
	//	OrderItemTextByBuyer		string  `json:"OrderItemTextByBuyer"`
	//	OrderItemTextBySeller		string  `json:"OrderItemTextBySeller"`
	OrderQuantityInBaseUnit     float32 `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit float32 `json:"OrderQuantityInDeliveryUnit"`
	BaseUnit                    string  `json:"BaseUnit"`
	DeliveryUnit                string  `json:"DeliveryUnit"`
	RequestedDeliveryDate       string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime       string  `json:"RequestedDeliveryTime"`
	NetAmount                   string  `json:"NetAmount"`
	TaxAmount                   string  `json:"TaxAmount"`
	GrossAmount                 string  `json:"GrossAmount"`
	ProductNetWeight            string  `json:"ProductNetWeight"`
}

type OrdersItemPricingElement struct {
	OrderID                 int     `json:"OrderID"`
	OrderItem               int     `json:"OrderItem"`
	PricingProcedureCounter int     `json:"PricingProcedureCounter"`
	ConditionType           string  `json:"ConditionType"`
	ConditionRateValue      float32 `json:"ConditionRateValue"`
}

type OrdersPartner struct {
	OrderID                 int     `json:"OrderID"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
}
