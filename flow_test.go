// (c) Copyright 2015-2017 JONNALAGADDA Srinivas
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flow

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var gt *testing.T

// error0 expects only an error value as its argument.
func error0(err error) {
	if err != nil {
		gt.Errorf("%v", err)
	}
}

// error1 expects a value and an error value as its arguments.
func error1(val1 interface{}, err error) interface{} {
	if err != nil {
		gt.Errorf("%v", err)
	}
	return val1
}

// fatal0 expects only an error value as its argument.
func fatal0(err error) {
	if err != nil {
		gt.Fatalf("%v", err)
	}
}

// fatal1 expects a value and an error value as its arguments.
func fatal1(val1 interface{}, err error) interface{} {
	if err != nil {
		gt.Fatalf("%v", err)
	}
	return val1
}

// assertEqual compares the two given values for equality.  In case of
// a difference, it errors with the given message.
func assertEqual(expected, observed interface{}, msg string) {
	if expected == observed {
		return
	}

	gt.Errorf("expected : '%v', observed : '%v'\t%s", expected, observed, msg)
}

// Driver test function.
func TestFlow01(t *testing.T) {
	gt = t

	// Connect to the database.
	driver, connStr := "mysql", "travis@/flow"
	db := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	defer db.Close()
	RegisterDB(db)

	// Test-local state.
	var dtID1, dtID2 DocTypeID
	var dsID1, dsID2, dsID3, dsID4, dsID5 DocStateID
	var daID1, daID2, daID3, daID4, daID5, daID6, daID7, daID8, daID9 DocActionID

	var roleID1, roleID2 RoleID
	var uID1, uID2, uID3, uID4 UserID
	var gID1, gID2, gID3, gID4, gID5, gID6 GroupID

	// Create operations.
	t.Run("Create", func(t *testing.T) {
		t.Run("DocTypes", func(t *testing.T) {
			tx := fatal1(db.Begin()).(*sql.Tx)
			defer tx.Rollback()

			dtID1 = fatal1(DocTypes.New(tx, "STOR_REQ")).(DocTypeID)
			dtID2 = fatal1(DocTypes.New(tx, "COMPUTE_REQ")).(DocTypeID)

			fatal0(tx.Commit())
		})

		t.Run("DocStates", func(t *testing.T) {
			tx := fatal1(db.Begin()).(*sql.Tx)
			defer tx.Rollback()

			dsID1 = fatal1(DocStates.New(tx, dtID1, "INITIAL")).(DocStateID)
			dsID2 = fatal1(DocStates.New(tx, dtID1, "PENDING_APPROVAL")).(DocStateID)
			dsID3 = fatal1(DocStates.New(tx, dtID1, "APPROVED")).(DocStateID)
			dsID4 = fatal1(DocStates.New(tx, dtID1, "REJECTED")).(DocStateID)
			dsID5 = fatal1(DocStates.New(tx, dtID1, "DISCARDED")).(DocStateID)

			fatal0(tx.Commit())
		})

		t.Run("DocActions", func(t *testing.T) {
			tx := fatal1(db.Begin()).(*sql.Tx)
			defer tx.Rollback()

			daID1 = fatal1(DocActions.New(tx, "INITIALISE")).(DocActionID)
			daID2 = fatal1(DocActions.New(tx, "NEW")).(DocActionID)
			daID3 = fatal1(DocActions.New(tx, "GET")).(DocActionID)
			daID4 = fatal1(DocActions.New(tx, "UPDATE")).(DocActionID)
			daID5 = fatal1(DocActions.New(tx, "DELETE")).(DocActionID)
			daID6 = fatal1(DocActions.New(tx, "APPROVE")).(DocActionID)
			daID7 = fatal1(DocActions.New(tx, "REJECT")).(DocActionID)
			daID8 = fatal1(DocActions.New(tx, "RETURN")).(DocActionID)
			daID9 = fatal1(DocActions.New(tx, "DISCARD")).(DocActionID)

			fatal0(tx.Commit())
		})

		t.Run("Roles", func(t *testing.T) {
			tx := fatal1(db.Begin()).(*sql.Tx)
			defer tx.Rollback()

			roleID1 = fatal1(Roles.New(tx, "ADMIN")).(RoleID)
			roleID2 = fatal1(Roles.New(tx, "RESEARCH_ANALYST")).(RoleID)

			fatal0(tx.Commit())
		})

		t.Run("Users", func(t *testing.T) {
			tx := fatal1(db.Begin()).(*sql.Tx)
			defer tx.Rollback()

			res, err := tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
			VALUES('FN 1', 'LN 1', 'email1@example.com', 1)`)
			if err != nil {
				t.Fatalf("%v\n", err)
			}
			uid, _ := res.LastInsertId()
			uID1 = UserID(uid)
			gID1 = fatal1(Groups.NewSingleton(tx, uID1)).(GroupID)

			res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
			VALUES('FN 2', 'LN 2', 'email2@example.com', 1)`)
			if err != nil {
				t.Fatalf("%v\n", err)
			}
			uid, _ = res.LastInsertId()
			uID2 = UserID(uid)
			gID2 = fatal1(Groups.NewSingleton(tx, uID2)).(GroupID)

			res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
			VALUES('FN 3', 'LN 3', 'email3@example.com', 1)`)
			if err != nil {
				t.Errorf("%v\n", err)
			}
			uid, _ = res.LastInsertId()
			uID3 = UserID(uid)
			gID3 = fatal1(Groups.NewSingleton(tx, uID3)).(GroupID)

			res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
			VALUES('FN 4', 'LN 4', 'email4@example.com', 1)`)
			if err != nil {
				t.Fatalf("%v\n", err)
			}
			uid, _ = res.LastInsertId()
			uID4 = UserID(uid)
			gID4 = fatal1(Groups.NewSingleton(tx, uID4)).(GroupID)

			fatal0(tx.Commit())
		})

		t.Run("Groups", func(t *testing.T) {
			tx := fatal1(db.Begin()).(*sql.Tx)
			defer tx.Rollback()

			gID5 = fatal1(Groups.New(tx, "RAs", "G")).(GroupID)
			gID6 = fatal1(Groups.New(tx, "PIs", "G")).(GroupID)

			fatal0(tx.Commit())
		})

		t.Run("GroupsAddUsers", func(t *testing.T) {
			tx := fatal1(db.Begin()).(*sql.Tx)
			defer tx.Rollback()

			fatal0(Groups.AddUser(tx, gID5, uID1))
			fatal0(Groups.AddUser(tx, gID5, uID2))
			fatal0(Groups.AddUser(tx, gID5, uID3))

			fatal0(Groups.AddUser(tx, gID6, uID2))
			fatal0(Groups.AddUser(tx, gID6, uID3))
			fatal0(Groups.AddUser(tx, gID6, uID4))

			fatal0(tx.Commit())
		})
	})

	// Entity update operations.
	t.Run("Update", func(t *testing.T) {
		t.Run("DocTypeRename", func(t *testing.T) {
			tx := fatal1(db.Begin()).(*sql.Tx)
			defer tx.Rollback()

			error0(DocTypes.Rename(tx, dtID1, "STORAGE_REQ"))

			fatal0(tx.Commit())

			obj := error1(DocTypes.Get(dtID1)).(*DocType)
			assertEqual("STORAGE_REQ", obj.Name, "")
		})

		t.Run("DocStateRename", func(t *testing.T) {
			tx := fatal1(db.Begin()).(*sql.Tx)
			defer tx.Rollback()

			error0(DocStates.Rename(tx, dsID1, "DRAFT"))

			fatal0(tx.Commit())

			obj := error1(DocStates.Get(dsID1)).(*DocState)
			assertEqual("DRAFT", obj.Name, "")
		})

		t.Run("DocActionRename", func(t *testing.T) {
			tx := fatal1(db.Begin()).(*sql.Tx)
			defer tx.Rollback()

			error0(DocActions.Rename(tx, daID1, "LIST"))

			fatal0(tx.Commit())

			obj := error1(DocActions.Get(daID1)).(*DocAction)
			assertEqual("LIST", obj.Name, "")
		})

		t.Run("GroupRename", func(t *testing.T) {
			tx := fatal1(db.Begin()).(*sql.Tx)
			defer tx.Rollback()

			error0(Groups.Rename(tx, gID5, "Research Associates"))

			fatal0(tx.Commit())

			obj := error1(Groups.Get(gID5)).(*Group)
			assertEqual("Research Associates", obj.Name, "")
		})
	})

	// Entity deletion operations.
	t.Run("Delete", func(t *testing.T) {
		t.Run("GroupsDeleteUsers", func(t *testing.T) {
			tx := fatal1(db.Begin()).(*sql.Tx)
			defer tx.Rollback()

			error0(Groups.RemoveUser(tx, gID6, uID2))

			fatal0(tx.Commit())
		})
	})

	// Tear down.
	t.Run("TearDown", func(t *testing.T) {
		tx := fatal1(db.Begin()).(*sql.Tx)
		defer tx.Rollback()

		error1(tx.Exec(`DELETE FROM wf_ac_group_roles`))
		error1(tx.Exec(`DELETE FROM wf_ac_group_hierarchy`))
		error1(tx.Exec(`DELETE FROM wf_access_contexts`))

		error1(tx.Exec(`DELETE FROM wf_group_users`))
		error1(tx.Exec(`DELETE FROM wf_groups_master`))
		error1(tx.Exec(`DELETE FROM users_master`))
		error1(tx.Exec(`DELETE FROM wf_roles_master`))

		error1(tx.Exec(`DELETE FROM wf_docactions_master`))
		error1(tx.Exec(`DELETE FROM wf_docstates_master`))
		error1(tx.Exec(`DELETE FROM wf_doctypes_master`))

		fatal0(tx.Commit())
	})
}