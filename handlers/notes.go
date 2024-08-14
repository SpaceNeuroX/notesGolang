package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Note struct {
	ID        int    `db:"id" json:"id"`
	Title     string `db:"title" json:"title"`
	Content   string `db:"content" json:"content"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

func GetNotes(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		notes := []Note{}
		err := db.Select(&notes, "SELECT * FROM notes ORDER BY created_at DESC")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch notes"})
		}
		return c.JSON(http.StatusOK, notes)
	}
}

func CreateNote(db *sqlx.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		note := new(Note)
		if err := c.Bind(note); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		}

		_, err := db.Exec("INSERT INTO notes (title, content) VALUES ($1, $2)", note.Title, note.Content)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create note"})
		}

		return c.JSON(http.StatusCreated, note)
	}
}
