package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

const (
	maxRetries = 3
	retryDelay = 10 * time.Second
)

func pingDB(dsn string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	return conn.Ping(ctx)
}

func fetchCategories(dsn string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return "", err
	}
	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, `SELECT * FROM "catégories"`)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	fields := rows.FieldDescriptions()

	var result string
	count := 0

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return "", err
		}

		result += "• "
		for i, v := range values {
			result += string(fields[i].Name) + "=" + toString(v) + " | "
		}
		result += "\n"
		count++
	}

	if count == 0 {
		return "⚠️ Aucune catégorie trouvée", nil
	}

	return result, nil
}

func toString(v any) string {
	if v == nil {
		return "NULL"
	}
	return fmt.Sprintf("%v", v)
}

func notifySlack(webhook, message string) {
	payload := map[string]string{
		"text": message,
	}

	body, _ := json.Marshal(payload)
	_, err := http.Post(webhook, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("⚠️ Slack notification failed:", err)
	}
}

func main() {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	slackWebhook := os.Getenv("SLACK_WEBHOOK_URL")

	if dsn == "" {
		log.Fatal("❌ DATABASE_URL non définie")
	}

	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := pingDB(dsn)
		if err == nil {
			log.Println("✅ Supabase ping OK")

			categories, err := fetchCategories(dsn)
			if err != nil {
				log.Println("❌ Erreur SELECT categories:", err)

				if slackWebhook != "" {
					notifySlack(
						slackWebhook,
						"⚠️ Ping OK mais SELECT categories FAILED:\n"+err.Error(),
					)
				}
				return
			}

			if slackWebhook != "" {
				notifySlack(
					slackWebhook,
					"✅ Supabase ping OK\n📂 *Categories list*:\n"+categories,
				)
			}
			return
		}

		log.Printf(
			"❌ Ping échoué (tentative %d/%d): %v",
			attempt,
			maxRetries,
			err,
		)

		time.Sleep(retryDelay)
	}

	log.Println("🚨 Supabase ping FAILED après retries")

	if slackWebhook != "" {
		notifySlack(
			slackWebhook,
			"🚨 Supabase ping FAILED après 3 retries (GitHub Actions)",
		)
	}

	os.Exit(1)
}