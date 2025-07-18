package database_test

import (
	"testing"
	"user_service/internal/shared/infrastructure/database"
)

func setup(t *testing.T, envs map[string]string) {
	for k, v := range envs {
		t.Setenv(k, v)
	}
}

func TestNewDBConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *database.DBConfig
		wantErr bool
	}{
		{
			name: "test1",
			want: &database.DBConfig{
				Host: "hoge",
				User: "fuga",
				Pass: "piyo",
				Name: "foo",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				t.Fatal("NewDBConfig() succeeded unexpectedly")
			}
			envs := map[string]string{
				"DB_HOST": "hoge",
				"DB_USER": "fuga",
				"DB_PASS": "piyo",
				"DB_NAME": "foo",
			}
			setup(t, envs)
			cfg, err := database.NewDBConfig()
			if err != nil {
				t.Errorf("NewDBConfig() failed: %v", err)
			}

			if host := cfg.GetHost(); host != tt.want.Host {
				t.Errorf("GetHost() = %v, want %v", host, tt.want.Host)
			}
			if user := cfg.GetUser(); user != tt.want.User {
				t.Errorf("GetUser() = %v, want %v", user, tt.want.Host)
			}
			if pass := cfg.GetPass(); pass != tt.want.Pass {
				t.Errorf("GetPass() = %v, want %v", pass, tt.want.Host)
			}
			if name := cfg.GetName(); name != tt.want.Name {
				t.Errorf("GetName() = %v, want %v", name, tt.want.Host)
			}
		})
	}
}

func TestNewDBConfigFailed(t *testing.T) {
	t.Run("failed test", func(t *testing.T) {
		_, err := database.NewDBConfig()
		if err == nil {
			t.Errorf("NewDBConfig() should have returned an error, but got nil")
		}
	})
}
