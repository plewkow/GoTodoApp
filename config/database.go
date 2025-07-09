package config

import (
	"draft-zadania-1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB = db
	if err := DB.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatalf("Failed to run automigrate: %v", err)
	}
	seedUsers := `
	INSERT INTO users (id, username, email)
	VALUES
		('11111111-1111-1111-1111-111111111111', 'janek', 'janek@example.com'),
		('22222222-2222-2222-2222-222222222222', 'ala', 'ala@example.com'),
		('33333333-3333-3333-3333-333333333333', 'michal', 'michal@example.com');
    `

	if err := DB.Exec(seedUsers).Error; err != nil {
		log.Fatalf("Failed to seed users: %v", err)
	}

	seedTasks := `
INSERT INTO tasks (id, title, description, due_date, status, user_id) VALUES
  ('aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa', 'Task 1 for Janek', 'Opis zadania 1', NOW() + INTERVAL '3 days', 0, '11111111-1111-1111-1111-111111111111'),
  ('bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb', 'Task 2 for Janek', 'Opis zadania 2', NOW() + INTERVAL '5 days', 1, '11111111-1111-1111-1111-111111111111'),
  ('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Task 3 for Janek', 'Opis zadania 3', NOW() + INTERVAL '7 days', 2, '11111111-1111-1111-1111-111111111111'),

  ('dddddddd-dddd-dddd-dddd-dddddddddddd', 'Task 1 for Ala', 'Opis zadania 1', NOW() + INTERVAL '2 days', 0, '22222222-2222-2222-2222-222222222222'),
  ('eeeeeeee-eeee-eeee-eeee-eeeeeeeeeeee', 'Task 2 for Ala', 'Opis zadania 2', NOW() + INTERVAL '4 days', 1, '22222222-2222-2222-2222-222222222222'),
  ('ffffffff-ffff-ffff-ffff-ffffffffffff', 'Task 3 for Ala', 'Opis zadania 3', NOW() + INTERVAL '6 days', 2, '22222222-2222-2222-2222-222222222222'),

  ('11111111-1111-1111-1111-aaaaaaaaaaaa', 'Task 1 for Michal', 'Opis zadania 1', NOW() + INTERVAL '1 day', 0, '33333333-3333-3333-3333-333333333333'),
  ('22222222-2222-2222-2222-bbbbbbbbbbbb', 'Task 2 for Michal', 'Opis zadania 2', NOW() + INTERVAL '3 days', 1, '33333333-3333-3333-3333-333333333333'),
  ('33333333-3333-3333-3333-cccccccccccc', 'Task 3 for Michal', 'Opis zadania 3', NOW() + INTERVAL '5 days', 2, '33333333-3333-3333-3333-333333333333');
`
	if err := DB.Exec(seedTasks).Error; err != nil {
		log.Fatalf("Failed to seed taks: %v", err)
	}
}
