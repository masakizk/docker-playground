# ベースとなるDockerイメージ指定
FROM golang:latest

# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/app

# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/app

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY . /go/src/app

# コンテナ起動時にアプリケーションを起動
CMD ["go", "run", "main.go"]
