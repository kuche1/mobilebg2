package configuration

type Config struct {
	// price (EUR)

	PriceMin int
	PriceMax int
}

func NewConfig() *Config {
	return &Config{
		PriceMin: 10_000, //1_852 // bgn 3_600
		PriceMax: 15_000, //3_852 // 4_632 // bgn 9_000
	}
}
