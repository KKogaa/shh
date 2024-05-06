package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	messages_schema := `
  CREATE TABLE messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
	  chatroom_id INTEGER NOT NULL,
    payload TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	  FOREIGN KEY (chatroom_id) REFERENCES chatrooms(id)
  );
  `

	chatrooms_schema := `
  CREATE TABLE chatrooms (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    chatroom TEXT NOT NULL UNIQUE
  );
  `

	// chatroom_id INTEGER NOT NULL,
	// FOREIGN KEY (chatroom_id) REFERENCES chatrooms(id),

	_, err = db.Exec(messages_schema)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(chatrooms_schema)
	if err != nil {
		panic(err)
	}

	return db
}
