package app

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "strings"
	"fmt"
	// "time"
	// "runtime"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

func (app *App) SetupRouter() {

	// app.Router.
	// 	Methods("POST").
	// 	Path("/endpoint").
	// 	HandlerFunc(app.postFunction)

	app.Router.
		Methods("OPTIONS", "POST").
		Path("/login").
		HandlerFunc(app.login)

	app.Router.
		Methods("GET").
		Path("/accounts").
		HandlerFunc(app.getAccounts)

	app.Router.
		Methods("GET").
		Path("/strategylist").
		HandlerFunc(app.getStrategyList)

	app.Router.
		Methods("GET").
		Path("/strategyConfig").
		HandlerFunc(app.getStrategyConfig)

	app.Router.
		Methods("GET").
		Path("/daystats").
		HandlerFunc(app.getDayStats)
	
	app.Router.
		Methods("OPTIONS", "POST").
		Path("/strategyConfig").
		HandlerFunc(app.updateStrategyConfig)
		
	app.Router.
		Methods("GET").
		Path("/contractconf").
		HandlerFunc(app.getContractConf)

	app.Router.
		Methods("OPTIONS", "POST").
		Path("/contractconf").
		HandlerFunc(app.updateContractConf)

	app.Router.
		Methods("GET").
		Path("/commsProfit").
		HandlerFunc(app.getCommsProfit)

	app.Router.
		Methods("GET").
		Path("/ps").
		HandlerFunc(app.getPs)
	
	app.Router.
		Methods("GET").
		Path("/pshistory").
		HandlerFunc(app.getPshistory)

	app.Router.
		Methods("GET").
		Path("/pshistorybasktest").
		HandlerFunc(app.getHisPs)

	app.Router.
		Methods("GET").
		Path("/getPsTimeStamp").
		HandlerFunc(app.getPsTimes)
}

func (app *App) login(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var params map[string]string

    // // 解析参数 存入map
    decoder.Decode(&params)
	// log.Println(r.Body)
	// log.Println(params["username"])

	user := &UserInfo{}

	data := &LoginMsg{}

	err := app.Database.QueryRow("SELECT * FROM `login_account` WHERE account = ? and password = ?", params["username"], params["password"]).Scan(&user.ACCOUNT, &user.PASSWORD)

	switch {
    	case err == sql.ErrNoRows:
			data = &LoginMsg{ false }
			break
		case err == nil:
			data = &LoginMsg{ true }
    	case err != nil:
			log.Println(err)
    }


	// log.Println("login api")
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9528")             //允许访问所有域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    // w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");


	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
// 登录接口

// 获取账户列表
func (app *App) getAccounts(w http.ResponseWriter, r *http.Request) {
	
	dbdata := &DbAccountList{}

	var arg string
	values := r.URL.Query()
    arg = values.Get("type")

	// log.Println(arg)
	s := make([]DbAccountList, 0)
	// rows, err := app.Database.Query("SELECT * FROM `account_list` where type=?;", arg)
	rows, err := app.Database.Query("SELECT * FROM `account_list`")

	if err != nil {
		log.Println("getAccounts: Database SELECT failed", err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&dbdata.TRADEID, &dbdata.TYPE, &dbdata.DAYRETURN, &dbdata.YIELD, &dbdata.DRAWDOWN, &dbdata.MDD, &dbdata.POS, &dbdata.NETPOS, &dbdata.LONGPOS, &dbdata.SHORTPOS, &dbdata.NAV, &dbdata.MAXNAV, &dbdata.SLIPP, &dbdata.BALANCE, &dbdata.SAVING, &dbdata.TOTAL, &dbdata.RRETURN, &dbdata.RYIELD, &dbdata.RTOTAL, &dbdata.RMARGIN, &dbdata.AY, &dbdata.CR, &dbdata.SHARP)
		if err != nil {
			log.Println(err)
		}
		s = append(s, DbAccountList{ dbdata.TRADEID, dbdata.TYPE, dbdata.DAYRETURN, dbdata.YIELD, dbdata.DRAWDOWN, dbdata.MDD, dbdata.POS, dbdata.NETPOS, dbdata.LONGPOS, dbdata.SHORTPOS, dbdata.NAV, dbdata.MAXNAV, dbdata.SLIPP, dbdata.BALANCE, dbdata.SAVING, dbdata.TOTAL, dbdata.RRETURN, dbdata.RYIELD, dbdata.RTOTAL, dbdata.RMARGIN, dbdata.AY, dbdata.CR, dbdata.SHARP  })
	}
	// log.Println("getAccounts api")
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9528")             //允许访问所有域
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");
	w.WriteHeader(http.StatusOK)
	// log.Println(s)

	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}


// strategy_lists策略列表
func (app *App) getStrategyList(w http.ResponseWriter, r *http.Request) {
	
	dbdata := &DbStrategy{}

	s := make([]string, 0)
	rows, err := app.Database.Query("SELECT * FROM `strategy_list`")

	if err != nil {
		log.Println("getStrategyList: Database SELECT failed", err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&dbdata.STRATEGY)
		if err != nil {
			log.Println(err)
		}
		s = append(s, dbdata.STRATEGY)
	}
	// log.Println("getStrategyList api")
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

// 获取历史数据
func (app *App) getDayStats(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");


	dbdata := &DayStats{}
	s := make([]DayStats, 0)

	var arg string
	var startT string
	var endT string
	values := r.URL.Query()
    arg = values.Get("trade_id")
	startT = values.Get("startTimes")
	endT = values.Get("endTimes")
	
	sqlRaw := fmt.Sprintf(`SELECT * from day_stats_%s where date between "%s" and "%s";`, arg, startT, endT)
	

	// log.Println(sqlRaw)
	rows, err := app.Database.Query(sqlRaw)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&dbdata.DATE, &dbdata.YIELDS, &dbdata.MARGIN, &dbdata.POS, &dbdata.NETPOS, &dbdata.LONGPOS, &dbdata.SHORTPOS, &dbdata.NETWORTH, &dbdata.HIGHESTNETWORTH, &dbdata.DRAWDOWN, &dbdata.DAYRETURN, &dbdata.COMMISSION, &dbdata.SLIP, &dbdata.SLIPP, &dbdata.BALANCE, &dbdata.DEPOSIT, &dbdata.WITHDRAW, &dbdata.SAVING, &dbdata.TOTAL)
		if err != nil {
			log.Println(err)
		}
		s = append(s, DayStats{ dbdata.DATE, dbdata.YIELDS, dbdata.MARGIN, dbdata.POS, dbdata.NETPOS, dbdata.LONGPOS, dbdata.SHORTPOS, dbdata.NETWORTH, dbdata.HIGHESTNETWORTH, dbdata.DRAWDOWN, dbdata.DAYRETURN, dbdata.COMMISSION, dbdata.SLIP, dbdata.SLIPP, dbdata.BALANCE, dbdata.DEPOSIT, dbdata.WITHDRAW, dbdata.SAVING, dbdata.TOTAL })
	}
	if err != nil {
		log.Println("getDayStats: Database SELECT failed", err)
	}
	// log.Println(sqlRaw)
	// log.Println("getDayStats api")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

// strategy_config仓位配置
func (app *App) getStrategyConfig(w http.ResponseWriter, r *http.Request) {
	
	dbdata := &StrategyConfig{}

	var arg string
	values := r.URL.Query()
    arg = values.Get("trade_id")

	// log.Println(arg)
	s := make([]StrategyConfig, 0)

	var sqlRaw string

	if arg != "" {
		sqlRaw = fmt.Sprintf(`SELECT * FROM strategy_conf where trade_id = "%s";`, arg)
	} else {
		sqlRaw = "SELECT * FROM `strategy_conf`"
	}
	rows, err := app.Database.Query(sqlRaw)

	if err != nil {
		log.Println("getStrategyConfig: Database SELECT failed", err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&dbdata.DATE, &dbdata.TRADEID, &dbdata.STRATEGY, &dbdata.COEFFICIENT)
		if err != nil {
			log.Println(err)
		}
		s = append(s, StrategyConfig{ dbdata.DATE, dbdata.TRADEID, dbdata.STRATEGY, dbdata.COEFFICIENT })
	}
	// log.Println("getStrategyConfig api")
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}


// 更新仓位配置
func (app *App) updateStrategyConfig(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	// w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");
	// dbdata := &StrategyConfig{}

	res := &ResponseType{}

	decoder := json.NewDecoder(r.Body)
	var params map[string]interface{}
    decoder.Decode(&params)


	trade_id := params["trade_id"]
	

	
	if rec, ok := params["data"].(map[string]interface{}); ok {
		for key, val := range rec {
			// str := "INSERT INFO `strategy_conf` SET coefficient=" + val + " WHERE trade_id=" + trade_id + " strategy=" + key

			// str := strings.Join([]string{"INSERT INFO `strategy_conf` SET coefficient=", val.(string), " WHERE trade_id=", trade_id.(string), " strategy=", key}, "")
			sqlRaw := fmt.Sprintf(`UPDATE strategy_conf SET coefficient="%s"  WHERE trade_id="%s" AND  strategy="%s";`, val, trade_id, key)
			// log.Println(sqlRaw)
			_, err := app.Database.Exec(sqlRaw)
			if err != nil {
				res = &ResponseType{ false, err.Error() }
			} else {
				res = &ResponseType{ true, "" }
			}

		}
		// log.Println("updateStrategyConfig api")
		// w.WriteHeader(http.StatusInternalServerError)
		
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
		
	}
	
}
// contract_conf 合约配置信息

func (app *App) getContractConf(w http.ResponseWriter, r *http.Request) {
	
	dbdata := &DbContractConf{}

	s := make([]DbContractConf, 0)
	rows, err := app.Database.Query("SELECT IFNULL(comm, ''), IFNULL(tianqin, ''), IFNULL(juejin, ''), IFNULL(manu, '') FROM `contract_conf`")

	if err != nil {
		log.Println("getContractConf: Database SELECT failed", err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&dbdata.COMM, &dbdata.TIANQIN, &dbdata.JUEJIN, &dbdata.MANU)
		if err != nil {
			log.Println(err)
		}
		s = append(s, DbContractConf{ dbdata.COMM, dbdata.TIANQIN, dbdata.JUEJIN, dbdata.MANU })
	}
	// log.Println("getContractConf api")

	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

// 更新合约配置
func (app *App) updateContractConf(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	// w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");
	// dbdata := &StrategyConfig{}

	res := &ResponseType{}

	decoder := json.NewDecoder(r.Body)
	var params map[string]string
    decoder.Decode(&params)

	comm := params["comm"]
	value := params["manu"]
	sqlRaw := fmt.Sprintf(`UPDATE contract_conf SET manu="%s"  WHERE comm="%s";`, value, comm)
	_, err := app.Database.Exec(sqlRaw)
	if err != nil {
		res = &ResponseType{ false, err.Error() }
	} else {
		res = &ResponseType{ true, "" }
	}
	
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
	
}


// 品种盈亏
func (app *App) getCommsProfit(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");

	dbdata := &CommsProfit{}

	var start string
	var end string
	var trade_id string
	values := r.URL.Query()
    start = values.Get("startTimes")
	end = values.Get("endTimes")
	trade_id = values.Get("trade_id")
	s := make([]CommsProfit, 0)

	var sqlRaw string

	sqlRaw = fmt.Sprintf(`SELECT * FROM comms_profit_%s WHERE date >= "%s" and date <= "%s";`, trade_id, start, end)
	// log.Println(sqlRaw)

	rows, err := app.Database.Query(sqlRaw)

	if err != nil {
		log.Println("getCommsProfit: Database SELECT failed", err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&dbdata.DATE, &dbdata.COMM, &dbdata.PROFIT)
		if err != nil {
			log.Println(err)
		}
		s = append(s, CommsProfit{ dbdata.DATE, dbdata.COMM, dbdata.PROFIT })
	}
	// log.Println("getCommsProfit api")
	
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}
// 获取仓位最新时间
func (app *App) getPsTimes(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");

	dbdata := &PsTime{}

	var trade_id_type string
	values := r.URL.Query()
	trade_id_type = values.Get("trade_id")
	s := make([]PsTime, 0)
	var sqlRaw string

	// var sqlRaw1 string
	sqlRaw = fmt.Sprintf(`select timestamp from pos_%s order by timestamp desc limit 1`, trade_id_type)
	// log.Println(sqlRaw)
	rows, err := app.Database.Query(sqlRaw)

	if err != nil {
		log.Println("getPsTimes: Database SELECT failed", err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&dbdata.TIMESTAMP)
		if err != nil {
			log.Println(err)
		}
		s = append(s, PsTime{ dbdata.TIMESTAMP })
	}
	// log.Println("getPsTimes api")
	
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

// 实时仓位
func (app *App) getPs(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");

	dbdata := &Ps{}

	var trade_id_type string
	var timestamp string
	values := r.URL.Query()
	trade_id_type = values.Get("trade_id")
	timestamp = values.Get("timestamp")
	s := make([]Ps, 0)

	var sqlRaw string

	sqlRaw = fmt.Sprintf(`select timestamp, comm,lowfreq,midfreq,actual from pos_%s where timestamp ="%s";`, trade_id_type, timestamp)

	// log.Println(sqlRaw)
	rows, err := app.Database.Query(sqlRaw)

	if err != nil {
		log.Println("getPs: Database SELECT failed", err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&dbdata.TIMESTAMP, &dbdata.COMM, &dbdata.LOWFREQ, &dbdata.MIDFREQ, &dbdata.ACTUAL)
		if err != nil {
			log.Println(err)
		}
		s = append(s, Ps{ dbdata.TIMESTAMP, dbdata.COMM, dbdata.LOWFREQ, dbdata.MIDFREQ, dbdata.ACTUAL })
	}
	// log.Println("getPs api")
	
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}

// 回测实时仓位
// 实时仓位
func (app *App) getHisPs(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");

	dbdata := &PsHistory{}

	var trade_id_type string
	var timestamp string
	values := r.URL.Query()
	trade_id_type = values.Get("trade_id")
	timestamp = values.Get("timestamp")
	s := make([]PsHistory, 0)

	var sqlRaw string

	sqlRaw = fmt.Sprintf(`select timestamp, comm, ps from pos_%s where timestamp ="%s";`, trade_id_type, timestamp)

	// log.Println(sqlRaw)
	rows, err := app.Database.Query(sqlRaw)

	if err != nil {
		log.Println("PsHistory: Database SELECT failed", err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&dbdata.TIMESTAMP, &dbdata.COMM, &dbdata.PS)
		if err != nil {
			log.Println(err)
		}
		s = append(s, PsHistory{ dbdata.TIMESTAMP, dbdata.COMM, dbdata.PS })
	}
	// log.Println("PsHistory api")
	
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}


// 获取仓位历史
// 实时仓位
func (app *App) getPshistory(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods","GET, POST, PUT, PATCH, POST, DELETE, OPTIONS");

	dbdata := &Ps{}

	var trade_id string
	var limit string
	values := r.URL.Query()
	trade_id = values.Get("trade_id")
	limit = values.Get("limit")
	s := make([]Ps, 0)

	var sqlRaw string

	sqlRaw = fmt.Sprintf(`select * from pos_%s where timestamp like "%s" or timestamp like '%s' order by timestamp desc limit %s`, trade_id, "%29%", "%59%", limit)

	// log.Println(sqlRaw)
	rows, err := app.Database.Query(sqlRaw)

	if err != nil {
		log.Println("getPshistory: Database SELECT failed", err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&dbdata.TIMESTAMP, &dbdata.LOWFREQ, &dbdata.MIDFREQ, &dbdata.ACTUAL, &dbdata.COMM)
		if err != nil {
			log.Println(err)
		}
		s = append(s, Ps{ dbdata.TIMESTAMP, dbdata.LOWFREQ, dbdata.MIDFREQ, dbdata.ACTUAL, dbdata.COMM })
	}
	// log.Println("getPshistory api")
	
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}