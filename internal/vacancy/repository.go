package vacancy

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

// Определение структуры репозитория
type VacancyRepository struct {
	Dbpool       *pgxpool.Pool
	CustomLogger *zerolog.Logger
}

// Инициализация репозитория
func NewVacancyRepository(dbpool *pgxpool.Pool, customLogger *zerolog.Logger) *VacancyRepository {
	return &VacancyRepository{
		Dbpool:       dbpool,
		CustomLogger: customLogger,
	}
}

// Метод добавления вакансии
func (r *VacancyRepository) addVacancy(form VacancyCreateForm) error {
	query := `INSERT INTO vacancies (email, role, company, salary, type, location, created_at) VALUES (@email, @role, @company, @salary, @type, @location, @created_at)`
	args := pgx.NamedArgs{
		"email":    form.Email,
		"role":     form.Role,
		"company":  form.Company,
		"salary":   form.Salary,
		"type":     form.Type,
		"location": form.Location,
		"created_at": time.Now(),
	}
	_, err := r.Dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("невозможно создать вакансию: %w", err)
	}
	return nil
}
