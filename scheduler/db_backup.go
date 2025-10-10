package scheduler

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *Scheduler) SendDatabaseBackup() {
	if s.adminChatID == 0 {
		log.Println("admin chat ID not configured, skipping backup")
		return
	}

	backupFilePath, err := s.createDatabaseBackup()
	if err != nil {
		s.sendBackupErrorMessage(err)
		return
	}
	defer os.Remove(backupFilePath)

	fileInfo, err := os.Stat(backupFilePath)
	if err != nil {
		s.sendBackupErrorMessage(err)
		return
	}

	err = s.sendBackupToAdmin(backupFilePath, fileInfo.Size())
	if err != nil {
		s.sendBackupErrorMessage(err)
		return
	}

	log.Printf("Database backup sent successfully: %s (%.2f MB)", backupFilePath, float64(fileInfo.Size())/1024/1024)
}

func (s *Scheduler) createDatabaseBackup() (string, error) {
	timestamp := time.Now().Format("2006-01-02")
	backupFileName := fmt.Sprintf("namoz_bot_backup_%s.sql.gz", timestamp)
	backupFilePath := fmt.Sprintf("/tmp/%s", backupFileName)

	// Create pg_dump command with gzip compression
	cmd := exec.Command("sh", "-c",
		fmt.Sprintf("PGPASSWORD='%s' pg_dump -h %s -U %s -d %s | gzip > %s",
			s.dbPassword,
			s.dbHost,
			s.dbUser,
			s.dbName,
			backupFilePath,
		))

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to create backup: %v, output: %s", err, string(output))
	}

	return backupFilePath, nil
}

func (s *Scheduler) sendBackupToAdmin(filePath string, fileSize int64) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open backup file: %v", err)
	}
	defer file.Close()

	caption := fmt.Sprintf(`📦 <b>Ежедневный бекап БД</b>
📅 Дата: %s
📊 Размер: %.2f MB
✅ Статус: Успешно`,
		time.Now().Format("02.01.2006 15:04"),
		float64(fileSize)/1024/1024,
	)

	_, err = s.telegram.Bot.SendDocument(context.Background(), &bot.SendDocumentParams{
		ChatID:    s.adminChatID,
		Document:  &models.InputFileUpload{Filename: filePath, Data: file},
		Caption:   caption,
		ParseMode: models.ParseModeHTML,
	})

	if err != nil {
		return fmt.Errorf("failed to send backup to admin: %v", err)
	}

	return nil
}

func (s *Scheduler) sendBackupErrorMessage(err error) {
	log.Printf("Database backup error: %v", err)

	if s.adminChatID == 0 {
		return
	}

	errorMsg := fmt.Sprintf(`❌ <b>Ошибка при создании бекапа БД</b>
📅 Дата: %s
⚠️ Ошибка: %v`,
		time.Now().Format("02.01.2006 15:04"),
		err,
	)

	_, err = s.telegram.Bot.SendMessage(context.Background(), &bot.SendMessageParams{
		ChatID:    s.adminChatID,
		Text:      errorMsg,
		ParseMode: models.ParseModeHTML,
	})

	if err != nil {
		log.Printf("Failed to send error message to admin: %v", err)
	}
}
