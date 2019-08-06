# study-google-wire

https://github.com/google/wire/blob/master/docs/guide.md#defining-providers を参考に実際に動かしてみる。

## 手順
* `foobarbaz/foobarbaz.go`作成
* `wire.go`作成
* `wire gen`コマンド実行　-> `wire_gen.go`が作成される
* `main.go` でDIされていることを確認

## 補足
* `wire`コマンドを使うときは、`go mod init`しておく必要がある
