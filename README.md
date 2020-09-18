# instalation
1. `git clone [repository]`
2. `cd [repository name]`
3. `docker-compose build`
4. `docker-compose run --rm app sh -c "rm -rf node_modules && cd client && npm install -g && npm install -g create-react-app`
5. `docker-compose up`


# docker-composeの使い方
- `docker-compose build`
- **※react用のディレクトリが存在しない場合**
  - `docker-compose run --rm app sh -c "npm install -g create-react-app && create-react-app client --typescript"`
- **※react用のディレクトリが既に存在する場合**
  - `docker-compose run --rm app sh -c "rm -rf node_modules && cd client && npm install -g && npm install -g create-react-app`
- `docker-compose up"`
  - `-d`を付与することで、バックグランド起動`
# build
- client
  - `client`ディレクトリ配下に移動
    - `npm run-script build`、`yarn build`のどちらかを実行（container内で実行）
    - `docker-compose run --rm app sh -c "npm run-script build"`
- golang
  - `go build main.go -o build/main`、`go build && mv main build/`のどちらかを実行
    - `../client/build` の静的ファイルを配信することで、Reactのシングルページアプリケーションを配信できる
