# CLAUDE.md

このファイルは Claude Code がこのリポジトリで作業する際のルールを定義する。

## プロジェクト概要

`sdcd` — Screwdriver.cd 向けのマルチコンテキスト対応 CLI（Go製）。
モジュール: `github.com/suzutan/sdcd-cli`

## 開発サイクル

1. 最新の `main` を起点にブランチを作成する
2. 区切りの良い作業単位で commit・push する
3. **1つ目のコミットを行った後は必ず PR を起票する**
   - `main` とブランチの差分を確認し、PR title・body を作成または更新する
   - PR body は `.github/pull_request_template.md` をベースに書くこと
   - PR の CI job が成功することを確認する

## 開発原則

- **DRY / KISS / YAGNI** を守ること
- フォールバックコードは書かない。使わないコードは後のノイズになる。必要になったら過去のコミットを遡れば良い
- 将来の自分（Claude Code 含む）や他の人が見たときにコンテキストを理解できるコードを書く

## ビルド・テスト

```sh
make build   # bin/sdcd を生成
make test    # go test ./... -v
make lint    # golangci-lint（要インストール）
```

## ディレクトリ構成

```
cmd/                    # cobra コマンド定義（1ファイル1コマンド）
internal/api/           # Screwdriver.cd API クライアント
internal/config/        # 設定ファイルの読み書き
internal/model/         # API レスポンス型定義
internal/output/        # table / JSON / YAML 出力
```

## 設定ファイル

`$XDG_CONFIG_HOME/sdcd-cli/config.yaml`（デフォルト: `~/.config/sdcd-cli/config.yaml`）、パーミッション `0600`。

## 認証

raw API token を `GET /v4/auth/token?api_token=<token>` で JWT に交換。JWT はメモリのみでキャッシュし、ディスクには書かない。

## リリース

`v*.*.*` タグを push すると GitHub Actions (`.github/workflows/release.yml`) が GoReleaser でクロスプラットフォームバイナリをビルドし、GitHub Release を自動作成する。

```sh
git tag v0.1.0
git push origin v0.1.0
```
