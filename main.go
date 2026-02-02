package main

import (
	"bytes"
	"context"
	"encoding/json"
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
	// 🔹 Charge .env si présent (local), ignore si absent (GitHub Actions)
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

			if slackWebhook != "" {
				notifySlack(slackWebhook, "✅ Supabase pings OK (GitHub Actions)")
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
