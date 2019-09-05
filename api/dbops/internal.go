package dbops

import (
	"database/sql"
	"http_server/api/defs"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(sid, ttlStr, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	return nil
}

/**
	检索单个session
 */
func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}

	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE id = ?")
	if err != nil {
		return nil, err
	}

	var ttl string
	var uname string

	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}

	defer stmtOut.Close()
	return ss, nil
}

/**
	检索所有的session 以sync map 返回
*/
func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	defer stmtOut.Close()

	for rows.Next() {
		var (
			id        string
			loginName string
			ttlStr    string
		)
		if err := rows.Scan(&id, &ttlStr, &loginName); err != nil {
			log.Printf("retrieve sessions error:%s", err)
			break
		}
		if ttl, err1 := strconv.ParseInt(ttlStr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{Username: loginName, TTL: ttl}
			m.Store(id, ss)
			log.Printf("session id: %s, ttl %d", id, ss.TTL)
		}
	}
	return m, nil
}

/**
 删除session
 */
func DeleteSession(sid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM sessions WHERE id = ?")
	if err != nil {
		return err
	}
	_, err1 := stmtDel.Exec(sid)
	if err1 != nil {
		return err1
	}
	defer stmtDel.Close()
	return nil
}
