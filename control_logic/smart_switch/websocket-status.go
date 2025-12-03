package smart_switch

type Connected struct {
	Connected bool `json:"connected"`
}

type Input struct {
	Id    string `json:"id"`
	State bool   `json:"state"`
}

type Matter struct {
	NumFabrics     int  `json:"num_fabrics"`
	Commissionable bool `json:"commissionable"`
}

type AvailableUpdates struct {
	Stable struct {
		Version string `json:"version"`
	}
}

type Stable struct {
	Version string `json:"version"`
	BuildId string `json:"build_id"`
}

type Alt struct {
	S1PMG4ZB struct {
		Name   string `json:"name"`
		Desc   string `json:"desc"`
		Stable Stable `json:"state"`
	}
}

type Wifi struct {
	StaIp  string `json:"sta_ip"`
	Status string `json:"status"`
	Ssid   string `json:"ssid"`
	Rssi   string `json:"rssi"`
}

type Sys struct {
	Mac              string           `json:"mac"`
	RestartRequired  bool             `json:"restart_required"`
	Time             string           `json:"time"`
	UnixTime         int64            `json:"unix_time"`
	LastSyncTs       int64            `json:"last_sync_ts"`
	Uptime           int64            `json:"uptime"`
	RamSize          int              `json:"ram_size"`
	RamFree          int              `json:"ram_free"`
	RamMinFree       int              `json:"ram_min_free"`
	FsSize           int              `json:"fs_size"`
	FsFree           int              `json:"fs_free"`
	CfgRev           int              `json:"cfg_rev"`
	KvsRev           int              `json:"kvs_rev"`
	ScheduleRev      int              `json:"schedule_rev"`
	WebhookRev       int              `json:"webhook_rev"`
	BtRelay          int              `json:"bt_relay"`
	AvailableUpdates AvailableUpdates `json:"available_updates"`
	Alt              Alt              `json:"alt"`
	ResetReason      int              `json:"reset_reason"`
	UtcOffset        int              `json:"utc_offset"`
}

type Result struct {
	Ble        []string  `json:"ble"`
	Bthome     []string  `json:"bthome"`
	Cloud      Connected `json:"cloud"`
	InputZero  Input     `json:"input:0"`
	Knx        []string  `json:"knx"`
	Matter     Matter    `json:"matter"`
	Mqtt       Connected `json:"mqtt"`
	SwitchZero Status    `json:"switch:0"`
	Wifi       Wifi      `json:"wifi"`
	Ws         Connected `json:"ws"`
}

type WebSocketResponse struct {
	Id     string `json:"id"`
	Src    string `json:"src"`
	Dst    string `json:"dst"`
	Result Result `json:"result"`
}
