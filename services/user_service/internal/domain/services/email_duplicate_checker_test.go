package services

import (
	"context"
	"errors"
	"testing"
	aggregate "user_service/internal/domain/aggregates"
	"user_service/internal/domain/models/user"
)

// モックリポジトリの実装
type mockUserRepository struct {
	existsWithEmailFunc func(ctx context.Context, email *user.Email) (bool, error)
	callCount           int
}

func (m *mockUserRepository) ExistsWithEmail(ctx context.Context, email *user.Email) (bool, error) {
	m.callCount++
	return m.existsWithEmailFunc(ctx, email)
}

func (m *mockUserRepository) GetById(ctx context.Context, userId user.UserId) (*aggregate.User, error) {
	return nil, nil
}
func (m *mockUserRepository) Register(ctx context.Context, user *aggregate.User) error { return nil }
func (m *mockUserRepository) Update(ctx context.Context, user *aggregate.User) error   { return nil }

func TestEmailDuplicateService_CheckDuplicate(t *testing.T) {
	tests := []struct {
		name          string
		email         *user.Email
		mockReturn    bool
		mockError     error
		expectedError string
		expectNoError bool
	}{
		{
			name:          "メールアドレスが存在しない場合 - 正常終了",
			email:         &user.Email{},
			mockReturn:    false,
			mockError:     nil,
			expectNoError: true,
		},
		{
			name:          "メールアドレスが既に存在する場合 - 重複エラー",
			email:         &user.Email{},
			mockReturn:    true,
			mockError:     nil,
			expectedError: "このメールアドレスはすでに使用されています。",
		},
		{
			name:          "リポジトリでデータベース接続エラーが発生した場合",
			email:         &user.Email{},
			mockReturn:    false,
			mockError:     errors.New("データベース接続エラー"),
			expectedError: "データベース接続エラー",
		},
		{
			name:          "リポジトリで不明なエラーが発生した場合",
			email:         &user.Email{},
			mockReturn:    false,
			mockError:     errors.New("不明なエラー"),
			expectedError: "不明なエラー",
		},
		{
			name:          "リポジトリがtrueとエラーを同時に返す場合（エラーを優先）",
			email:         &user.Email{},
			mockReturn:    true,
			mockError:     errors.New("データベースエラー"),
			expectedError: "データベースエラー",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックリポジトリの設定
			mockRepo := &mockUserRepository{
				existsWithEmailFunc: func(ctx context.Context, email *user.Email) (bool, error) {
					return tt.mockReturn, tt.mockError
				},
			}

			// サービスの作成
			service := NewEmailDuplicateService(mockRepo)

			// テスト実行
			err := service.CheckDuplicate(context.Background(), tt.email)

			// 結果の検証
			if tt.expectNoError {
				if err != nil {
					t.Errorf("エラーが発生しないことを期待していましたが、エラーが発生しました: %v", err)
				}
			} else {
				if err == nil {
					t.Errorf("エラーが発生することを期待していましたが、エラーが発生しませんでした")
					return
				}
				if err.Error() != tt.expectedError {
					t.Errorf("期待されたエラーメッセージ: %q, 実際のエラーメッセージ: %q", tt.expectedError, err.Error())
				}
			}

			// リポジトリが1回だけ呼ばれることを確認
			if mockRepo.callCount != 1 {
				t.Errorf("ExistsWithEmailは1回だけ呼ばれるべきですが、%d回呼ばれました", mockRepo.callCount)
			}
		})
	}
}

func TestNewEmailDuplicateService(t *testing.T) {
	mockRepo := &mockUserRepository{}
	service := NewEmailDuplicateService(mockRepo)

	// サービスがnilでないことを確認
	if service == nil {
		t.Fatal("NewEmailDuplicateServiceはnilでないサービスを返すべきです")
	}

	// 正しい型が返されることを確認
	concreteService, ok := service.(*EmailDuplicateService)
	if !ok {
		t.Error("NewEmailDuplicateServiceは*emailDuplicateService型を返すべきです")
		return
	}

	// リポジトリが正しく設定されていることを確認
	if concreteService.repository != mockRepo {
		t.Error("リポジトリが正しく設定されていません")
	}
}

func TestEmailDuplicateService_CheckDuplicate_NilEmail(t *testing.T) {
	mockRepo := &mockUserRepository{
		existsWithEmailFunc: func(ctx context.Context, email *user.Email) (bool, error) {
			if email == nil {
				return false, errors.New("emailがnilです")
			}
			return false, nil
		},
	}

	service := NewEmailDuplicateService(mockRepo)

	// nilのemailを渡してテスト
	err := service.CheckDuplicate(context.Background(), nil)

	// エラーが発生することを期待
	if err == nil {
		t.Error("nilのemailを渡した場合にはエラーが発生することを期待しています")
	} else if err.Error() != "emailがnilです" {
		t.Errorf("期待されたエラーメッセージ: %q, 実際のエラーメッセージ: %q", "emailがnilです", err.Error())
	}
}

func TestEmailDuplicateService_InterfaceImplementation(t *testing.T) {
	// インターフェースが正しく実装されていることを確認
	var _ EmailDuplicateChecker = (*EmailDuplicateService)(nil)

	mockRepo := &mockUserRepository{}
	service := NewEmailDuplicateService(mockRepo)

	// 戻り値がEmailDuplicateCheckerインターフェースを実装していることを確認
	var checker EmailDuplicateChecker = service
	if checker == nil {
		t.Error("サービスはEmailDuplicateCheckerインターフェースを実装すべきです")
	}
}

// エッジケースのテスト
func TestEmailDuplicateService_EdgeCases(t *testing.T) {
	t.Run("同じemailで複数回呼び出した場合", func(t *testing.T) {
		callCount := 0
		mockRepo := &mockUserRepository{
			existsWithEmailFunc: func(ctx context.Context, email *user.Email) (bool, error) {
				callCount++
				return false, nil
			},
		}

		service := NewEmailDuplicateService(mockRepo)
		email := &user.Email{} // 実際のEmailオブジェクトに置き換えてください

		// 同じemailで複数回呼び出し
		for i := 0; i < 3; i++ {
			err := service.CheckDuplicate(context.Background(), email)
			if err != nil {
				t.Errorf("呼び出し%d回目でエラーが発生しました: %v", i+1, err)
			}
		}

		// リポジトリが期待される回数呼ばれていることを確認
		if callCount != 3 {
			t.Errorf("ExistsWithEmailは3回呼ばれるべきですが、%d回呼ばれました", callCount)
		}
	})
}
