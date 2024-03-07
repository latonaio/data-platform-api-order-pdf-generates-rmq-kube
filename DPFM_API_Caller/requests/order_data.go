package requests

type Header struct {
	OrderID               			int     `json:"OrderID"`
	OrderStatus           			string  `json:"OrderStatus"`
	OrderDate             			string  `json:"OrderDate"`
	OrderType             			string  `json:"OrderType"`
	Buyer                 			int     `json:"Buyer"`
	BuyerName             			string  `json:"BuyerName"`
	Seller                			int     `json:"Seller"`
	SellerName            			string  `json:"SellerName"`
	RequestedDeliveryDate 			string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime 			string  `json:"RequestedDeliveryTime"`
	TotalGrossAmount      			float32 `json:"TotalGrossAmount"`
	Contract              			*int    `json:"Contract"`
	ContractItem          			*int    `json:"ContractItem"`
	Project               			*int    `json:"Project"`
	ProjectDescription    			*string `json:"ProjectDescription"`
	WBSElement            			*int    `json:"WBSElement"`
	WBSElementResponsiblePersonName	*string `json:"WBSElementResponsiblePersonName"`
	Incoterms             			*string `json:"Incoterms"`
	IncotermsName             		*string `json:"IncotermsName"`
	PricingDate           			string  `json:"PricingDate"`
	PaymentTerms          			string  `json:"PaymentTerms"`
	PaymentTermsName     			string  `json:"PaymentTermsName"`
	PaymentMethod         			string  `json:"PaymentMethod"`
	TransactionCurrency   			string  `json:"TransactionCurrency"`
	HeaderText   					string  `json:"HeaderText"`
	Items				  			[]Items `json:"Items"`
	ItemPricingElements				[]ItemPricingElements `json:"ItemPricingElements"`
	Partners			  			[]Partners `json:"Partners"`
}

type Items struct {
	OrderID               		int     `json:"OrderID"`
	OrderItem                   int     `json:"OrderItem"`
	DeliverToPlant              string  `json:"DeliverToPlant"`
	DeliverToPlantName          string  `json:"DeliverToPlantName"`
	DeliverFromPlant            string  `json:"DeliverFromPlant"`
	DeliverFromPlantName        string  `json:"DeliverFromPlantName"`
	Product                     string  `json:"Product"`
	ProductSpecification		*string `json:"ProductSpecification"`
	SizeOrDimensionText			*string `json:"SizeOrDimensionText"`
	OrderItemText               string  `json:"OrderItemText"`
	OrderQuantityInBaseUnit     float32 `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit float32 `json:"OrderQuantityInDeliveryUnit"`
	BaseUnit                    string  `json:"BaseUnit"`
	DeliveryUnit                string  `json:"DeliveryUnit"`
	RequestedDeliveryDate       string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime       string  `json:"RequestedDeliveryTime"`	
	NetAmount                   float32 `json:"NetAmount"`
	TaxAmount                   float32 `json:"TaxAmount"`
	GrossAmount                 float32 `json:"GrossAmount"`
	ProductNetWeight            float32 `json:"ProductNetWeight"`
}

type ItemPricingElements struct {
	OrderID                 int     `json:"OrderID"`
	OrderItem               int     `json:"OrderItem"`
	PricingProcedureCounter int     `json:"PricingProcedureCounter"`
	ConditionType           string  `json:"ConditionType"`
	ConditionRateValue      float32 `json:"ConditionRateValue"`
}

type Partners struct {
	OrderID                 int     `json:"OrderID"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
}
