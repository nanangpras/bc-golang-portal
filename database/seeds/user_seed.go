package seeds

import (
	"golangnews/internal/core/domain/model"
	"golangnews/lib/conv"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating hash password")
	}

	admin := model.User{
		Name: "Admin",
		Email: "admin@admin.com",
		Password: string(bytes),
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: admin.Email}).Error; err != nil {
		log.Fatal().Err(err).Msg("Error creating admin role");
	}else{
		log.Info().Msg("Admin role created successfully")
	}
}