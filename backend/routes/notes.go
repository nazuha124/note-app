package routes

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"notes-backend/database"
	"notes-backend/middleware"

	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func saveAndResizeImage(file *multipart.FileHeader) (string, error) {
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, os.ModePerm)
	}

	// Nama file unik
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filePath := filepath.Join(uploadDir, filename)

	// Simpan file sementara
	if err := saveFile(file, filePath); err != nil {
		return "", err
	}

	// ✅ Baca gambar
	img, err := imaging.Open(filePath)
	if err != nil {
		return "", err
	}

	// Resize gambar — maksimal 800px lebar, tinggi menyesuaikan
	resized := imaging.Resize(img, 800, 0, imaging.Lanczos)

	// ✅ Kompres hasil (kualitas 80%)
	outPath := filepath.Join(uploadDir, "resized_"+filename)
	err = imaging.Save(resized, outPath, imaging.JPEGQuality(80))
	if err != nil {
		return "", err
	}

	// Hapus file asli biar gak numpuk
	os.Remove(filePath)

	return outPath, nil
}

func NoteRoutes(app *fiber.App) {
	api := app.Group("/api/notes", middleware.RequireAuth())
	api.Get("/", getNotes)
	api.Post("/", createNote)
	api.Delete("/:id", deleteNote)
}

func saveFile(file *multipart.FileHeader, dest string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.ReadFrom(src)
	return err
}

// ✅ Contoh handler untuk tambah note
func createNote(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int64(claims["user_id"].(float64))

	title := c.FormValue("title")
	content := c.FormValue("content")

	file, err := c.FormFile("image")
	var imageURL string

	if file != nil {
		path, err := saveAndResizeImage(file)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan gambar"})
		}
		imageURL = path
	}

	_, err = database.DB.Exec(
		"INSERT INTO notes (title, content, image_url, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())",
		title, content, imageURL, userID,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Catatan berhasil ditambahkan!"})
}

// ✅ Ambil semua catatan
func getNotes(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int64(claims["user_id"].(float64))

	rows, err := database.DB.Query(
		"SELECT id, title, content, image_url, created_at FROM notes WHERE user_id=$1 ORDER BY created_at DESC",
		userID,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	type Note struct {
		ID        int64     `json:"id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		ImageURL  string    `json:"image_url"`
		CreatedAt time.Time `json:"created_at"`
	}

	var notes []Note
	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.ImageURL, &note.CreatedAt); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		notes = append(notes, note)
	}

	return c.JSON(notes)
}

// ✅ Hapus catatan berdasarkan ID
func deleteNote(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := int64(claims["user_id"].(float64))
	id := c.Params("id")

	_, err := database.DB.Exec("DELETE FROM notes WHERE id=$1 AND user_id=$2", id, userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Catatan berhasil dihapus!"})
}
