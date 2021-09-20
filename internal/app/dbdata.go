package app

import (
	"time"
)

type DbData struct {
	ID   int       `json:"id"`
	Date time.Time `json:"date"`
	Name string    `json:"name"`
}


type DbAccountList struct {
	TRADEID string    `json:"trade_id"`
	TYPE string    `json:"type"`
	DAYRETURN int   `json:"day_return"`
	YIELD float32   `json:"yield"`
	DRAWDOWN float32   `json:"drawdown"`
	MDD float32   `json:"mdd"`
	POS float32   `json:"pos"`
	NETPOS float32 `json:"net_pos"`
	LONGPOS float32 `json:"long_pos"`
	SHORTPOS float32 `json:"short_pos"`
	NAV float32   `json:"nav"`
	MAXNAV float32   `json:"max_nav"`
	SLIPP float32 `json:"slipp"`
	BALANCE int   `json:"balance"`
	SAVING int   `json:"saving"`
	TOTAL int   `json:"total"`
	RRETURN float32   `json:"r_return"`
	RYIELD float32   `json:"r_yield"`
	RTOTAL int   `json:"r_total"`
	RMARGIN float32 `json:"r_margin"`
	AY float32 `json:"ay"`
	CR float32 `json:"cr"`
	SHARP float32 `json:"sharp"`
}

type DbStrategy struct {
	STRATEGY string `json:"strategy"`
}

type DbContractConf struct {
	COMM string `json:"comm"`
	TIANQIN string `json:"tianqin"`
	JUEJIN string `json:"juejin"`
	MANU string `json:"manu"`
}

type LoginMsg struct {
	SUCCESS bool `json:"success"`
}

type UserInfo struct {
	ACCOUNT string `json:"account"`
	PASSWORD string `json:"password"`
}

type StrategyConfig struct {
	DATE string `json:"date"`
	TRADEID string `json:"trade_id"`
	STRATEGY string `json:"strategy"`
	COEFFICIENT string `json:"coefficient"`
}

type ResponseType struct {
	SUCCESS bool `json:"success"`
	MESSAGE string `json:"message"`
}


type DayStats struct {
	DATE string `json:"date"`
	YIELDS float32 `json:"yields"`
	MARGIN float32 `json:"margin"`
	POS float32 `json:"pos"`
	NETPOS float32 `json:"net_pos"`
	LONGPOS float32 `json:"long_pos"`
	SHORTPOS float32 `json:"short_pos"`
	NETWORTH float32 `json:"net_worth"`
	HIGHESTNETWORTH float32 `json:"highest_net_worth"`
	DRAWDOWN float32 `json:"drawdown"`
	DAYRETURN int `json:"day_return"`
	COMMISSION int `json:"commission"`
	SLIP int `json:"slip"`
	SLIPP float32 `json:"slipp"`
	BALANCE int `json:"balance"`
	DEPOSIT int `json:"deposit"`
	WITHDRAW int `json:"withdraw"`
	SAVING int `json:"saving"`
	TOTAL int `json:"total"`
}


type CommsProfit struct {
	DATE string `json:"date"`
	COMM string `json:"comm"`
	PROFIT string `json:"profit"`
}

type Ps struct {
	TIMESTAMP string `json:"timestamp"`
	COMM string `json:"comm"`
	LOWFREQ string `json:"lowfreq"` 
	MIDFREQ string `json:"midfreq"`
	ACTUAL string `json:"actual"`
}

type PsHistory struct {
	TIMESTAMP string `json:"timestamp"`
	COMM string `json:"comm"`
	PS string `json:"ps"`
}


type PsTime struct {
	TIMESTAMP string `json:"timestamp"`
}