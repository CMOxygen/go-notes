package dbmanager

import (
	"go-notes-webapp/main-module/filemanager"
	"strings"
	"testing"
)

func TestDatabaseManager_Update(t *testing.T) {

	t.Log("testing database manager update basic")
	{
		//p, err := encryption.EncryptSHA([]byte("go_password1"))
		//if err != nil {
		//	t.Error(err)
		//}
		//testUserToUpdate := go_user.User{ID: 3, Username: "go_username1", Password: p}
		var dbm DatabaseManager

		url, err := filemanager.ReadFile("/etc/server/c/r")
		if err != nil {
			t.Error(err)
		}
		err = dbm.Connect(strings.Trim(string(url), "\n"))
		if err != nil {
			t.Error(err)
		}
		//p, err = encryption.EncryptSHA([]byte("updated_pass2"))
		//if err != nil {
		//	t.Error(err)
		//}
		//err = dbm.Update(&testUserToUpdate, "username=?, password=?", "updated2", p)
		//if err != nil {
		//	t.Error(err)
		//}
	}
}

func TestDatabaseManager_Delete(t *testing.T) {
	t.Log("testing database manager delete basic")
	{
		//p, err := encryption.EncryptSHA([]byte("go_pass2"))
		//testUserToDelete := go_user.User{ID: 4, Username: "go_username2", Password: p}
		var dbm DatabaseManager
		url, err := filemanager.ReadFile("/etc/server/c/r")
		if err != nil {
			t.Error(err)
		}
		err = dbm.Connect(strings.Trim(string(url), "\n"))
		if err != nil {
			t.Error(err)
		}
		//err = dbm.Delete(&testUserToDelete)
		//if err != nil {
		//	t.Error(err)
		//}
	}
}

func TestDatabaseManager_Insert(t *testing.T) {

	t.Log("testing database manager insert")
	{
	}
}
