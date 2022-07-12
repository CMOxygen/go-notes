package encryption

import (
	"bytes"
	"fmt"
	"go-notes-webapp/main-module/dbmanager"
	"go-notes-webapp/main-module/filemanager"
	"go-notes-webapp/main-module/go_user"
	"strings"
	"testing"
)

//type ByteUser struct {
//	bun.BaseModel `bun:"table:Users"`
//
//	ID       uint32 `json:"id" bun:"userId,pk,autoincrement"`
//	Username string `json:"username" bun:"username,unique,notnull"`
//	Password []byte `json:"password" bun:"password,notnull"`
//}

func TestEncryptSHA(t *testing.T) {
	t.Log("testing encryption sha")
	{
		data := []byte("pass1")
		result, err := EncryptSHA(data)
		if err != nil {
			t.Error(err)
		}
		//fmt.Println(result, len(result))

		testUser := go_user.User{Username: "user1", Password: result}

		url, err := filemanager.ReadFile("/etc/server/c/r")
		var dbm dbmanager.DatabaseManager

		//err = dbm.Connect(string(url[:len(url)-1]))

		err = dbm.Connect(strings.Trim(string(url), "\n"))
		if err != nil {
			t.Error(err)
		}
		//err = dbm.Insert(&testUser)
		//if err != nil {
		//	t.Error(err)
		//}
		result, err = EncryptSHA([]byte("pass1"))
		if err != nil {
			t.Error(err)
		}
		//testUser2 := ByteUser{Username: "user1", Password: result}

		var out go_user.User
		err = dbm.Select(&out, "username=?", "user1")
		if err != nil {
			t.Error(err)
		}
		if bytes.Compare(testUser.Password, out.Password) != 0 {
			t.Errorf("test user password = %v , length = %v \n out password = %v , length = %v", testUser.Password, len(testUser.Password), out.Password, len(out.Password))
		}
		fmt.Printf("test %v length = %v \n out %v length = %v", testUser.Password, len(testUser.Password), out.Password, len(out.Password))
	}
}

//BenchmarkEncryptSHA-6   	 2421812	       484.9 ns/op
func BenchmarkEncryptSHA(b *testing.B) {
	in := []byte("user1")
	for i := 0; i < b.N; i++ {
		EncryptSHA(in)
	}
}
