# CLAUDE.md

このリポジトリで作業する際のルール。

Module: `github.com/suzutan/sdcd-cli`

## 開発サイクル

1. 最新の `main` を起点にブランチを作成する
2. 区切りの良い作業単位で commit・push する
3. **1つ目のコミットを行った後は必ず PR を起票する**
   - `main` とブランチの差分を確認し、PR title・body を作成または更新する
   - PR body は `.github/pull_request_template.md` をベースに書くこと
4. **CI job が成功したら PR をマージする**
   - ユーザーからタスクを依頼されたら、コミット・push・PR 作成・マージまで一気に進める
   - マージ後は `main` ブランチに戻る

## 開発原則

- **DRY / KISS / YAGNI** を守ること
- フォールバックコードは書かない。必要になったら過去のコミットを遡れば良い
- 将来の自分（Claude Code 含む）や他の人がコンテキストを理解できるコードを書く

## コマンド

```sh
make build   # bin/sdcd を生成（ldflags でバージョン埋め込み）
make test    # go test ./... -v
make lint    # golangci-lint run ./...
make install # $GOPATH/bin にインストール
```

## アーキテクチャ上の規約

### `cmd/` — コマンド定義

- ファイル命名: `<resource>.go`（グループ登録のみ）と `<resource>_<action>.go`（実装）
- `cfg *config.Config` と `client *api.Client` は `cmd` パッケージグローバル変数。`PersistentPreRunE` で初期化される。コマンド関数に引数として渡さない
- 出力には必ず `printer()` を使う（`cmd/root.go` のファクトリ関数）
- **新コマンドが API を使わない場合は `noClientNeeded()` に追加する**（`cmd/root.go:108`）。現在対象: `version`, `completion`, `help`, `auth` 配下すべて

### `internal/api/` — API クライアント

- `client.go` の `do()` / `doWithHeaders()` を経由して全リクエストを送る。直接 `http.Client` を使わない
- ログページネーション: クエリパラメータは `from`（行番号オフセット、0始まり）。`X-More-Data: true` ヘッダーで次ページあり。`X-Next-Page` があればその値、なければレスポンスの最終行の `N+1` を使う
- `model.LogLine.T` はミリ秒 Unix タイムスタンプ（秒ではない）

### `internal/config/`

- 設定ファイルパーミッションは `0600`（`Save()` が強制）
- `$XDG_CONFIG_HOME/sdcd-cli/config.yaml`、未設定時は `~/.config/sdcd-cli/config.yaml`

### テスト

- `internal/api` のテストは `NewMockServer(t, routes)` を使う（`internal/api/testutil.go`）
- `internal/config` のテストは `TempConfig(t, cfg)` を使う（`internal/config/testutil.go`）
- これらのヘルパーを再実装しないこと

## リリース

バージョニングは **release-please** で自動化されている。

```
feat:/fix: コミットが main に積まれる
  ↓
release-please が Release PR を自動作成・更新
（feat: → minor bump / fix: → patch bump / feat!: → major bump）
  ↓
Release PR をマージ
  ↓
git tag v*.*.* が自動付与 → GoReleaser が GitHub Release を作成
```

- コミットメッセージは必ず Conventional Commits 形式にする（release-please がバージョン決定に使用）
- `.release-please-manifest.json` は release-please が自動管理するため **手動編集禁止**
