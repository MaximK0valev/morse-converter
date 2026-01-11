package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/SaidHasan-go/morse-converter/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("Ошибка парсинга формы:", err)
		http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Ошибка при получении файла:", err)
		http.Error(w, "Ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Println("Ошибка при чтении файла:", err)
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}
	convert, err := service.AutoConvert(string(data))
	if err != nil {
		log.Println("Ошибка при конвертации данных:", err)
		http.Error(w, "Ошибка при конвертации данных", http.StatusInternalServerError)
		return
	}
	filename := time.Now().UTC().Format("2006-01-02_15-04-05") + filepath.Ext(handler.Filename)
	outputPath := filepath.Join("outputs", filename)

	err = os.WriteFile(outputPath, []byte(convert), 0644)
	if err != nil {
		log.Println("Ошибка записи файла:", err)
		http.Error(w, "Ошибка записи файла", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(convert))
}
