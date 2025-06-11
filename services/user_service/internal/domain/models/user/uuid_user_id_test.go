package user

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestNewUuidUserId(t *testing.T) {
	t.Run("正常にUUIDUserIdが生成される", func(t *testing.T) {
		userId, err := NewUuidUserId()

		if err != nil {
			t.Fatalf("エラーが発生しました: %v", err)
		}

		if userId.value == (uuid.UUID{}) {
			t.Error("UUIDが空です")
		}

		if userId.value.Version() != 7 {
			t.Errorf("期待するバージョン: 7, 実際: %d", userId.value.Version())
		}
	})

	t.Run("複数回呼び出すと異なるUUIDが生成される", func(t *testing.T) {
		userId1, err1 := NewUuidUserId()
		userId2, err2 := NewUuidUserId()

		if err1 != nil {
			t.Fatalf("1回目の生成でエラー: %v", err1)
		}
		if err2 != nil {
			t.Fatalf("2回目の生成でエラー: %v", err2)
		}

		if userId1.value == userId2.value {
			t.Error("異なるUUIDが生成されるべきですが、同じUUIDが生成されました")
		}
	})

	t.Run("生成されるUUIDの形式が正しい", func(t *testing.T) {
		userId, err := NewUuidUserId()

		if err != nil {
			t.Fatalf("エラーが発生しました: %v", err)
		}

		uuidStr := userId.value.String()
		if len(uuidStr) != 36 {
			t.Errorf("UUIDの文字列長が不正: 期待値36, 実際%d", len(uuidStr))
		}

		if len(uuidStr) == 36 && uuidStr[14] != '7' {
			t.Error("UUID v7の形式が正しくありません")
		}
	})
}

func BenchmarkNewUuidUserId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := NewUuidUserId()
		if err != nil {
			b.Fatalf("ベンチマーク中にエラー: %v", err)
		}
	}
}

func TestNewUuidUserId_Performance(t *testing.T) {
	const count = 10000
	uuids := make(map[uuid.UUID]bool, count)

	for i := 0; i < count; i++ {
		userId, err := NewUuidUserId()
		if err != nil {
			t.Fatalf("UUID生成エラー（%d回目）: %v", i+1, err)
		}

		if uuids[userId.value] {
			t.Errorf("重複するUUIDが生成されました: %s", userId.value.String())
		}
		uuids[userId.value] = true
	}

	if len(uuids) != count {
		t.Errorf("期待する一意UUID数: %d, 実際: %d", count, len(uuids))
	}
}

func TestNewUuidUserIdWithGenerator_Error(t *testing.T) {
	t.Run("UUID生成でエラーが発生した場合", func(t *testing.T) {
		errorGenerator := func() (uuid.UUID, error) {
			return uuid.UUID{}, errors.New("mock error")
		}

		userId, err := NewUuidUserIdWithGenerator(errorGenerator)

		if err == nil {
			t.Fatal("エラーが期待されましたが、nilが返されました")
		}

		if err.Error() != "UUID v7の生成に失敗しました。" {
			t.Errorf("期待するエラーメッセージ: 'UUID v7の生成に失敗しました。', 実際: '%s'", err.Error())
		}

		if userId.value != (uuid.UUID{}) {
			t.Error("エラー時は空のUUIDUserIdが返されるべきです")
		}
	})
}
