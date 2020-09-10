# docker-composeの使い方
- `docker-compose build`
- `docker-compose run --rm app sh -c "npm install -g create-react-app && create-react-app client --typescript"`
- `docker-compose up`
  - `-d`を付与することで、バックグランド起動`
# build
- client
  - `client`ディレクトリ配下に移動
  - `npm run-script build`、`yarn build`のどちらかを実行（container内で実行）
  - `docker-compose run --rm app sh -c "npm run-script build"`
- golang
  - `go build main.go -o build/main`、`go build && mv main build/`のどちらかを実行
