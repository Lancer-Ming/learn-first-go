package dbops

import (
	"database/sql"
	"log"
	"http_server/api/defs"
	"github.com/chilts/sid"
	"time"
)


func AddUserCredential(loginName string, pwd string) error {
	var stmtIns, err = dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()

	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("DeleteUser error:%s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddVideoInfo(uid int, name string) (*defs.VideoInfo, error) {
	vid := sid.Id()
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info (id, author_id, name, display_ctime)
									VALUES (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, uid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{
		Id:           vid,
		AuthorId:     uid,
		Name:         name,
		DisplayCtime: ctime,
	}
	defer stmtIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut,err := dbConn.Prepare(`SELECT id, author_id, name, display_ctime FROM video_info WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	var (
		id string
		author_id int
		name string
		display_ctime string
	)
	err = stmtOut.QueryRow(vid).Scan(&id, &author_id, &name, &display_ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()
	res := &defs.VideoInfo{
		Id:           id,
		AuthorId:     author_id,
		Name:         name,
		DisplayCtime: display_ctime,
	}
	return res, nil
}

func DeleteVideo(uid int, vid string) error {
	stmtDel, err := dbConn.Prepare(`DELETE FROM video_info WHERE id = ? AND author_id = ?`)
	if err != nil {
		log.Printf("DeleteVideo error:%s", err)
		return err
	}
	_, err = stmtDel.Exec(vid, uid)
	if err != nil {
		log.Printf("DeleteVideo error:%s", err)
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddComment(uid int, vid string, content string) error {
	id := sid.Id()
	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, vid, uid, content)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content
 								FROM comments INNER JOIN users on comments.author_id = users.id 
 								WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) 
 								AND comments.time <= FROM_UNIXTIME(?)`)
	var res []*defs.Comment
	if err != nil {
		return res, err
	}
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{
			Id:       id,
			VideoId:  vid,
			Author: name,
			Content:  content,
		}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}



