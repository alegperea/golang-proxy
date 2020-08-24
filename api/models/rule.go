package model

type IP struct {
	IP 		 string  `json:"IP"`
}

type Rules struct {
	Rules        []Rule `json:"Rules"`
}

type Rule struct {
	RuleName string `json:"RuleName"`
	IP string `json:"RuleName"`
	Path string `json:"RuleName"`
	RuleType int16 `json:"RuleType"`
}

func GetRulesByIP(IP) ([]Rule) {

	// Reglas obtenidas desde NoSQL DB a patir IP recibida
	rules := []Rule {
		{
			RuleName: "Deny localhost",
			IP: "127.0.0.1",
			Path: "",
			RuleType: 1,
		},
		{
			RuleName: "Deny localhost",
			IP: "127.0.0.1",
			Path: "categories/MLA1271",
			RuleType: 2,
		},
		
	}		
	
	return rules

}
