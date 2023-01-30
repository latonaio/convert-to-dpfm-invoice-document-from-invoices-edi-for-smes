package requests

type CodeConversionItemPricingElement struct {
	PricingProcedureCounter   int     `json:"PricingProcedureCounter"`
	ConditionRecord           *int    `json:"ConditionRecord"`
	ConditionSequentialNumber *int    `json:"ConditionSequentialNumber"`
	ConditionType             *string `json:"ConditionType"`
}
