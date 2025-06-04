# User Service 技術選定ドキュメント

**プロジェクト**: 分散タスク管理システム  
**サービス**: User Service

---

## 🎯 サービス概要

### 責務・機能要件

- ユーザー登録・認証（JWT）
- プロフィール管理（CRUD）
- 権限・ロール管理（RBAC）
- パスワードリセット機能
- OAuth 連携（Google, GitHub 等）
- セッション管理
- ユーザー活動ログ

### 非機能要件

- **可用性**: 99.9%
- **レスポンス時間**: P95 < 200ms
- **同時接続数**: 1,000 ユーザー
- **データ整合性**: 強一貫性
- **セキュリティ**: OAuth2.0, JWT, 暗号化

---

## 🏗 アーキテクチャ設計

### システム構成図

```
[Client] → [API Gateway] → [User Service] → [MySQL]
                                ↓
                           [Redis Cache]
                                ↓
                           [Kafka Events]
```

### データベース設計

**Primary Store**: MySQL
**Cache Layer**: Redis  
**Event Stream**: Kafka

#### 主要テーブル

- `users` - ユーザー基本情報
- `roles` - ロール定義
- `permissions` - 権限定義
- `user_roles` - ユーザー・ロール関連
- `user_sessions` - セッション管理
- `password_resets` - パスワードリセット

---

## 🛠 技術選定

### プログラミング言語・フレームワーク

#### 選定候補

| 技術                  | 評価       | 理由                                         |
| --------------------- | ---------- | -------------------------------------------- |
| **Go + Gin**          | ⭐⭐⭐⭐⭐ | 軽量かつ高速であるため。拡張性よりシンプルさ |
| **Go + Echo**         | ⭐⭐⭐⭐   |                                              |
| **Go + Fiber**        | ⭐⭐⭐     |                                              |
| **Node.js + Express** | ⭐⭐       |                                              |

#### 選定結果: **Go + Gin**

**選定理由:**

- パフォーマンス:
- 開発効率:
- エコシステム:
- チーム習熟度:
- 運用性:

**メリット:**

- 高性能・低レイテンシ
- 並行処理のサポート
- 豊富なミドルウェア
- シンプルな API 設計

**デメリット:**

- 学習コスト
- エラーハンドリングの冗長性

---

### データストア

#### Primary Database

| データベース   | 評価       | 理由               |
| -------------- | ---------- | ------------------ |
| **MySQL**      | ⭐⭐⭐⭐⭐ | シェア率が高いから |
| **PostgreSQL** | ⭐⭐⭐⭐   |                    |
| **MongoDB**    | ⭐⭐       |                    |

#### 選定結果: **PostgreSQL 15**

**選定理由:**

- ACID 準拠:
- JSON 対応:
- 拡張性:
- 運用実績:

#### Cache Layer

| キャッシュ    | 評価       | 理由 |
| ------------- | ---------- | ---- |
| **Redis**     | ⭐⭐⭐⭐⭐ |      |
| **Memcached** | ⭐⭐⭐     |      |

#### 選定結果: **Redis 7**

**選定理由:**

- セッション管理:
- データ構造:
- Pub/Sub:
- 高可用性:

---

### 認証・認可

#### 認証方式

| 方式               | 評価       | 理由 |
| ------------------ | ---------- | ---- |
| **JWT**            | ⭐⭐⭐⭐⭐ |      |
| **Session Cookie** | ⭐⭐⭐     |      |
| **OAuth2.0**       | ⭐⭐⭐⭐   |      |

#### 選定結果: **JWT + OAuth2.0**

**選定理由:**

- ステートレス:
- スケーラビリティ:
- マイクロサービス対応:
- 標準準拠:

#### 実装ライブラリ

- **JWT**: `golang-jwt/jwt`
- **OAuth**: `oauth2` (Go 標準拡張)
- **暗号化**: `bcrypt`
- **バリデーション**: `go-playground/validator`

---

### API 設計

#### API 仕様

| 項目              | 選定        | 理由 |
| ----------------- | ----------- | ---- |
| **API Style**     | REST        |      |
| **Documentation** | OpenAPI 3.0 |      |
| **Versioning**    | URL Path    |      |
| **Content Type**  | JSON        |      |

#### gRPC Interface

```protobuf
service UserService {
  rpc GetUser(GetUserRequest) returns (GetUserResponse);
  rpc ValidateToken(ValidateTokenRequest) returns (ValidateTokenResponse);
  rpc CheckPermission(CheckPermissionRequest) returns (CheckPermissionResponse);
}
```

---

### 監視・ログ

#### メトリクス

| 項目               | 技術         | 理由 |
| ------------------ | ------------ | ---- |
| **メトリクス収集** | Prometheus   |      |
| **可視化**         | Grafana      |      |
| **アラート**       | AlertManager |      |

#### ログ

| 項目           | 技術          | 理由 |
| -------------- | ------------- | ---- |
| **構造化ログ** | Logrus/Zap    |      |
| **ログ収集**   | Fluent Bit    |      |
| **ログ検索**   | Elasticsearch |      |

#### トレーシング

- **分散トレーシング**: Jaeger
- **OpenTelemetry**: Go SDK

---

### セキュリティ

#### セキュリティ要件

- **データ暗号化**: 保存時・転送時
- **入力検証**: SQL インジェクション対策
- **レート制限**: API 呼び出し制限
- **監査ログ**: ユーザー操作記録
- **HTTPS 強制**: TLS 1.3
- **CORS 対応**: オリジン制限

#### 実装方針

```go
// セキュリティミドルウェア例
- CORS設定
- Rate Limiting
- Input Validation
- SQL Injection対策
- XSS対策
- JWT検証
```

---

### 開発・デプロイ

#### 開発環境

| 項目             | 技術           | 理由 |
| ---------------- | -------------- | ---- |
| **コンテナ**     | Docker         |      |
| **ローカル開発** | Docker Compose |      |
| **テスト**       | Testify        |      |
| **モック**       | GoMock         |      |

#### CI/CD

| 項目             | 技術           | 理由 |
| ---------------- | -------------- | ---- |
| **CI/CD**        | GitHub Actions |      |
| **コード品質**   | SonarQube      |      |
| **セキュリティ** | Snyk           |      |
| **デプロイ**     | Kubernetes     |      |

---

## 📊 パフォーマンス要件

### レスポンス時間目標

- **ユーザー認証**: < 100ms
- **ユーザー情報取得**: < 50ms
- **権限チェック**: < 30ms
- **ユーザー登録**: < 200ms

### スループット目標

- **認証 API**: 1,000 RPS
- **ユーザー情報 API**: 500 RPS
- **権限 API**: 2,000 RPS

### リソース使用量

- **CPU**: 0.5-1 core
- **Memory**: 512MB-1GB
- **Storage**: 10GB

---

## 🧪 テスト戦略

### テスト構成

- **Unit Test**: 80%以上のカバレッジ
- **Integration Test**: DB・外部 API 連携
- **E2E Test**: 主要ユーザーフロー
- **Load Test**: パフォーマンス検証

### テストツール

- **Unit**: Go 標準 testing + Testify
- **Mock**: GoMock
- **Integration**: Testcontainers
- **Load**: k6

---

## 🚀 運用・監視

### SLI/SLO

- **Availability**: 99.9%
- **Latency**: P95 < 200ms
- **Error Rate**: < 0.1%
- **Saturation**: CPU < 80%

### アラート設定

- **Error Rate**: > 1%
- **Response Time**: P95 > 500ms
- **Service Down**: Health Check 失敗
- **High CPU**: > 90%

### 運用ツール

- **Health Check**: `/health`, `/ready`
- **Metrics**: `/metrics` (Prometheus)
- **Profiling**: `pprof`
- **Graceful Shutdown**: SIGTERM 処理

---

## 📋 実装計画

### Phase 1: 基盤実装（1 週間）

- プロジェクト構造作成
- DB 接続・マイグレーション
- 基本的な CRUD API
- Docker 環境構築

### Phase 2: 認証実装（1 週間）

- JWT 認証機能
- パスワードハッシュ化
- Login/Logout API
- セッション管理

### Phase 3: 権限管理（1 週間）

- RBAC 実装
- 権限チェックミドルウェア
- Role/Permission 管理
- gRPC API 実装

### Phase 4: 運用対応（1 週間）

- 監視・ログ実装
- ヘルスチェック
- テスト実装
- ドキュメント作成

---

## 🔍 リスク・課題

### 技術的リスク

| リスク           | 影響度 | 対策 |
| ---------------- | ------ | ---- |
| **JWT 漏洩**     | 高     |      |
| **DB 性能劣化**  | 中     |      |
| **メモリリーク** | 中     |      |

### 運用リスク

| リスク               | 影響度 | 対策 |
| -------------------- | ------ | ---- |
| **大量ユーザー登録** | 高     |      |
| **認証サーバー停止** | 高     |      |
| **DB 接続エラー**    | 中     |      |

---

## 📚 参考資料

### 技術ドキュメント

- Go 公式ドキュメント
- Gin Web フレームワーク
- PostgreSQL 公式ドキュメント
- JWT RFC 7519
- OAuth 2.0 RFC 6749

### ベストプラクティス

- 12-Factor App
- RESTful API 設計ガイド
- マイクロサービスセキュリティ
- Go 言語ベストプラクティス

---

## ✅ レビュー・承認

### レビュー項目

- アーキテクチャ設計
- 技術選定根拠
- セキュリティ要件
- パフォーマンス要件
- 運用・監視設計

### 承認者

- **テックリード**:
- **セキュリティ**:
- **インフラ**:
- **プロダクト**:

**最終更新**: 2025 年 6 月 5 日  
**次回レビュー**: 実装開始前
