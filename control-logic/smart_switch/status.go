package smart_switch

type Energy struct {
	Total    float32   `json:"total"`
	ByMinute []float32 `json:"by_minute"`
	MinuteTs int64     `json:"minute_ts"`
}

type Temperature struct {
	Celsius    float32 `json:"tC"`
	Fahrenheit float32 `json:"tF"`
}

type Status struct {
	Id          int         `json:"id"`
	Source      string      `json:"source"`
	Output      bool        `json:"output"`
	APower      float32     `json:"apower"`
	Voltage     float32     `json:"voltage"`
	Current     float32     `json:"current"`
	Freq        float32     `json:"freq"`
	AEnergy     Energy      `json:aenergy`
	RetEnergy   Energy      `json:ret_energy`
	Temperature Temperature `json:temperature`
}
