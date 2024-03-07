package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-order-pdf-generates-rmq-kube/DPFM_API_Input_Formatter"
	"fmt"
)

func SetToPdfData(
	headerData *dpfm_api_input_reader.Header,
	inputItems []dpfm_api_input_reader.Items,
	inputItemPricingElements []dpfm_api_input_reader.ItemPricingElements,
	inputPartners []dpfm_api_input_reader.Partners,
) *Header {
	pm := &Header{}

	var ordersItem []OrdersItem
	for _, dataItems := range inputItems {
		if headerData.OrderID == dataItems.OrderID {
			ordersItem = append(ordersItem, OrdersItem{
				OrderID:                     dataItems.OrderID,
				OrderItem:                   dataItems.OrderItem,
				DeliverToPlant:              dataItems.DeliverToPlant,
				DeliverToPlantName:          dataItems.DeliverToPlantName,
				DeliverFromPlant:            dataItems.DeliverFromPlant,
				DeliverFromPlantName:        dataItems.DeliverFromPlantName,
				Product:                     dataItems.Product,
				ProductSpecification:        dataItems.ProductSpecification,
				SizeOrDimensionText:         dataItems.SizeOrDimensionText,
				OrderItemText:               dataItems.OrderItemText,
				OrderQuantityInBaseUnit:     dataItems.OrderQuantityInBaseUnit,
				OrderQuantityInDeliveryUnit: dataItems.OrderQuantityInDeliveryUnit,
				BaseUnit:                    dataItems.BaseUnit,
				DeliveryUnit:                dataItems.DeliveryUnit,
				RequestedDeliveryDate:       dataItems.RequestedDeliveryDate,
				RequestedDeliveryTime:       dataItems.RequestedDeliveryTime,
				NetAmount:                   fmt.Sprintf("%.0f", dataItems.NetAmount),
				TaxAmount:                   fmt.Sprintf("%.0f", dataItems.TaxAmount),
				GrossAmount:                 fmt.Sprintf("%.0f", dataItems.GrossAmount),
				ProductNetWeight:            fmt.Sprintf("%.0f", dataItems.ProductNetWeight),
			})
		}
	}

	var ordersItemPricingElement []OrdersItemPricingElement
	for _, dataItemPricingElements := range inputItemPricingElements {
		if headerData.OrderID == dataItemPricingElements.OrderID {
			ordersItemPricingElement = append(ordersItemPricingElement, OrdersItemPricingElement{
				OrderID:                 dataItemPricingElements.OrderID,
				OrderItem:               dataItemPricingElements.OrderItem,
				PricingProcedureCounter: dataItemPricingElements.PricingProcedureCounter,
				ConditionType:           dataItemPricingElements.ConditionType,
				ConditionRateValue:      dataItemPricingElements.ConditionRateValue,
			})
		}
	}

	var ordersPartner []OrdersPartner
	for _, dataPartners := range inputPartners {
		if headerData.OrderID == dataPartners.OrderID {
			ordersPartner = append(ordersPartner, OrdersPartner{
				OrderID:                 dataPartners.OrderID,
				PartnerFunction:         dataPartners.PartnerFunction,
				BusinessPartner:         dataPartners.BusinessPartner,
				BusinessPartnerFullName: dataPartners.BusinessPartnerFullName,
				BusinessPartnerName:     dataPartners.BusinessPartnerName,
			})
		}
	}

	pm = &Header{
		OrderID:                          fmt.Sprintf("%d", headerData.OrderID),
		OrderStatus:                      headerData.OrderStatus,
		OrderDate:                        headerData.OrderDate,
		OrderType:                        headerData.OrderType,
		Buyer:                            headerData.Buyer,
		BuyerName:                        headerData.BuyerName,
		Seller:                           headerData.Seller,
		SellerName:                       headerData.SellerName,
		RequestedDeliveryDate:            headerData.RequestedDeliveryDate,
		RequestedDeliveryTime:            headerData.RequestedDeliveryTime,
		TotalGrossAmount:                 fmt.Sprintf("%.0f", headerData.TotalGrossAmount),
		Contract:                         headerData.Contract,
		ContractItem:                     headerData.ContractItem,
		Project:                          headerData.Project,
		ProjectDescription:               headerData.ProjectDescription,
		WBSElement:                       headerData.WBSElement,
		WBSElementResponsiblePersonName:  headerData.WBSElementResponsiblePersonName,
		Incoterms:                        headerData.Incoterms,
		IncotermsName:                    headerData.IncotermsName,
		PricingDate:                      headerData.PricingDate,
		PaymentTerms:                     headerData.PaymentTerms,
		PaymentTermsName:                 headerData.PaymentTermsName,
		PaymentMethod:                    headerData.PaymentMethod,
		TransactionCurrency:              headerData.TransactionCurrency,
		HeaderText:                       headerData.HeaderText,
		OrdersItem:                       ordersItem,
		OrdersPartner:                    ordersPartner,
		TotalOrderQuantityInBaseUnit:     calculateTotalOrderQuantityInBaseUnit(inputItems),
		TotalOrderQuantityInDeliveryUnit: calculateTotalOrderQuantityInDeliverUnit(inputItems),
	}

	return pm
}

// Order Header が1件しかないという想定での計算
func calculateTotalOrderQuantityInBaseUnit(
	inputItems []dpfm_api_input_reader.Items,
) string {
	var totalOrderQuantityInBaseUnit float32

	for _, dataItems := range inputItems {
		totalOrderQuantityInBaseUnit += dataItems.OrderQuantityInBaseUnit
	}

	return fmt.Sprint(totalOrderQuantityInBaseUnit)
}

func calculateTotalOrderQuantityInDeliverUnit(
	inputItems []dpfm_api_input_reader.Items,
) string {
	var totalOrderQuantityInDeliverUnit float32 = 0

	for _, dataItems := range inputItems {
		totalOrderQuantityInDeliverUnit += dataItems.OrderQuantityInDeliveryUnit
	}

	return fmt.Sprint(totalOrderQuantityInDeliverUnit)
}
