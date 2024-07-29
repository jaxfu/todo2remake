package config

import (
	"fmt"
	"furrj/todo_2_remake/internal/types"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(filepath string) (types.EnvVars, error) {
	var vars types.EnvVars

	if os.Getenv(MODE) != "PROD" {
		if err := godotenv.Load(filepath); err != nil {
			return vars, err
		}
	}

	vars.MODE = os.Getenv(MODE)
	if vars.MODE == "" {
		return vars, fmt.Errorf("no env variable %s", MODE)
	}

	vars.DB_URL = os.Getenv(DB_URL)
	if vars.DB_URL == "" {
		return vars, fmt.Errorf("no env variable %s", DB_URL)
	}

	vars.PORT = os.Getenv(PORT)
	if vars.PORT == "" {
		return vars, fmt.Errorf("no env variable %s", PORT)
	}

	return vars, nil
}
