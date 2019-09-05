package dbops

import (
	"fmt"
	"http_server/api/defs"
	"strconv"
	"testing"
	"time"
)

var vid string

func clearTables() {
	dbConn.Exec("TRUNCATE users")
	dbConn.Exec("TRUNCATE video_info")
	dbConn.Exec("TRUNCATE comments")
	dbConn.Exec("TRUNCATE sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAdd)
	t.Run("Get", testGet)
	t.Run("Del", testDel)
}

func testAdd(t *testing.T) {
	err := AddUserCredential("Lancer", "123456")
	if err != nil {
		t.Errorf("Error of AddUser:%v", err)
	}
}

func testGet(t *testing.T) {
	pwd, err := GetUserCredential("Lancer")
	if pwd != "123456" || err != nil {
		t.Errorf("Error of GetUser:%v", err)
	}
}

func testDel(t *testing.T) {
	err := DeleteUser("Lancer", "123456")
	if err != nil {
		t.Errorf("Error of DeleteUser:%v", err)
	}
	// 查询检验是否还有数据
	pwd, err := GetUserCredential("Lancer")
	if err != nil {
		t.Errorf("Error of CheckDel:%v", err)
	}
	if pwd != "" {
		t.Errorf("Error of CheckDel:%v", err)
	}
}

///---------------------------------------------------------

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideo)
}

func testAddVideoInfo(t *testing.T) {
	videoInfo, err := AddVideoInfo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo:%v", err)
	}
	vid = videoInfo.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(vid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo:%v", err)
	}
}

func testDeleteVideo(t *testing.T) {
	// 测试删除
	err := DeleteVideo(1, vid)
	if err != nil {
		t.Errorf("Error of DeleteVideoError:%v", err)
	}

	// 查询是否还存在
	var videoInfo *defs.VideoInfo
	videoInfo, err = GetVideoInfo(vid)
	if err != nil || videoInfo != nil {
		t.Errorf("Error of CheckDeleteVideoError:%v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAdd)
	t.Run("AddComments", testAddComment)
	t.Run("ListComment", testListComment)
}

func testAddComment(t *testing.T) {
	vid := "12345"
	uid := 1
	content := "好好看"

	err := AddComment(uid, vid, content)
	if err != nil {
		t.Errorf("Error of AddComment:%v", err)
	}
}

func testListComment(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	fmt.Println(to)
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComment:%v", err)
	}
	for i, ele := range res {
		fmt.Printf("comments: %d, %v \n", i, ele)
	}

}


