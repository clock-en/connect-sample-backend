# Golang boilerplate

Golang のシンプルなボイラープレートです。

Air を使用したホットリロード機能も設定しています。

## 新規プロジェクトで使用する

- `go.mod` の module パスを変更する

下記の `connect-sample-backend` 部分を適切な名称に変更してください

```
module github.com/clock-en/connect-sample-backend
```

- `.vscode` ディレクトリの作成

プロジェクトごとに VSCode や Cursor の設定を共有したい場合は適宜 `.vscode` ディレクトリを作成してください。

- [Delve](https://github.com/go-delve/delve) の導入

デバッグツールとして Delve が有名ですが、当リポジトリには含めていません。

使用する場合は適宜 Dockerfile を書き換えたりしてください。

``` Dockerfile
RUN go install github.com/air-verse/air@v1.61.7 && \
    go install github.com/go-delve/delve/cmd/dlv@v1.24.2
```

``` compose.yml
    ports:
      - "8000:8000"
      - "2345:2345"
```

## 環境構築

当プロジェクトでは Dev Container 技術を使用して構成されています。

### 1. 環境変数の設定

`.environments/local/docker/golang/.env.example` を複製して `.env` ファイルを作成してください。

### 3. devcontainer を起動

devcontainer を起動します。

#### VScode の場合
拡張機能で [Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack`
`http://localhost:8080) を使用してください。

#### JetBrains 製品の場合
1. JetBrains 製 の IDE でプロジェクトディレクトリを開く
2. `.devcontainer/devcontainer.json` ファイルを開く
3. 行数表示の 1 行目の横にあるコンテナアイコン (箱のマーク) をクリック  
<img src="./.docs/assets/images/readme-image-01.png" style="max-width: 50%;">
4. 「Docker Container を作成してソースをマウント」を選択し、好みの IDE を選択  
<img src="./.docs/assets/images/readme-image-02.png" style="max-width: 50%;">  
※ パフォーマンス関連の通知が表示されますが無視して OK
5. コンテナが作成されると右下通知が表示されるので「接続」をクリック  
<img src="./.docs/assets/images/readme-image-03.png" style="max-width: 300px;">
6. エディタが開いたらターミナルを表示し、 `app` コマンドを実行する (なぜか Docker に記載したコマンドが実行されない？)

#### 共通の注意点
devcontainer は初めて接続したタイミングのコードでコンテナを作成します。  
.env が変更されたり、 docker 周りの調整が発生した場合はコンテナをリビルドしてください。
