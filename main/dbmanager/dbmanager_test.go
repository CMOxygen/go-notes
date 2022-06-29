package dbmanager

import (
	"go-notes-webapp/main-module/go_user"
	"testing"
)

func TestDatabaseManager_Update(t *testing.T) {

	t.Log("testing database manager update basic")
	{
		testUserToUpdate := go_user.User{ID: 3, Username: "go_username1", Password: "go_pass1"}
		var dbm DatabaseManager

		err := dbm.Connect("go_notes:AlSkDjFhG_2@/go_notes")
		if err != nil {
			t.Error(err)
		}
		err = dbm.Update(&testUserToUpdate, "username=?, password=?", "updated2", "updated_pass2")
		if err != nil {
			t.Error(err)
		}
	}
}

func TestDatabaseManager_Delete(t *testing.T) {
	t.Log("testing database manager delete basic")
	{
		testUserToDelete := go_user.User{ID: 4, Username: "go_username2", Password: "go_pass2"}
		var dbm DatabaseManager

		err := dbm.Connect("go_notes:AlSkDjFhG_2@/go_notes")
		if err != nil {
			t.Error(err)
		}
		err = dbm.Delete(&testUserToDelete)
		if err != nil {
			t.Error(err)
		}
	}
}