package model

type ICORatingCompany struct {
	Title            string
	Type             string
	Industry         string
	Description      string
	Features         string
	Country          string
	Website          string
	Whitepaper       string
	InvestmentRating float32
	HypeScore        float32
	RiskScore        float32
	SocialMedia      map[string]string
}
