# hello sample

## 概要

hello/binをコンパイルし実行するとでhelloというだけのサンプルですが、  
目的はgo get, go build, go installの動作確認のためのサンプルです

dockerのgolangイメージでテストをしています  
dockerのrunコマンドサンプル
```
docker run --rm --name test01 -it golang /bin/bash
```

## 準備

```
apt-get update
apt-get install tree
tree /go
.
|-- bin
`-- src
```

## go get

### 結果

コンパイルされるのはgetしたbinとそれに関係するパッケージのみ  
importの記述は有効  
githubから取得されるのはリポジトリ単位なので必要最小限ではない  
srcがある場合は、ローカルファイルでコンパイルされる  

getするのがbinでもpackageでも動作に違いはない  

-uを付けるとソースリポジトリが更新されていれば取得するようになる
### 実行

#### ファイルを初期化

```
rm -rf /go/bin /go/pkg /go/src
mkdir -m 777 /go/bin /go/src
mkdir -m 755 /go/pkg
```

#### getを実行
```
go get -u github.com/awisu2/gosample/hello/bin
tree /go
.
|-- bin
|   `-- bin
|-- pkg
|   `-- linux_amd64
|       `-- github.com
|           `-- awisu2
|               |-- golib
|               |   `-- log.a
|               `-- gosample
|                   `-- hello
|                       `-- pac.a
`-- src
    `-- github.com
        `-- awisu2
            |-- golib
            |   |-- README.md
            |   |-- analyse
            |   |   `-- analyse.go
            |   |-- file
            |   |   `-- file.go
            |   `-- log
            |       `-- log.go
            `-- gosample
                |-- gorutine
                |   `-- main.go
                |-- hello
                |   |-- README.md
                |   |-- bin
                |   |   `-- main.go
                |   `-- pac
                |       `-- hello.go
                |-- interface
                |   `-- main.go
                |-- switch
                |   `-- main.go
                `-- test
                    |-- test.go
                    `-- test_test.go
```


## go build

### 結果

カレントディレクトリのmain.goをコンパイルする
結果のバイナリファイルはカレントディレクトリに作成される
go build ./...とすることで、配下のファイルすべてがビルド対象となる

getは行われないので、importに対応するファイルがない場合エラーになる

### 実行
```
go get -u github.com/awisu2/gosample/hello/bin
```

#### binとpkgを削除

```
rm -rf /go/bin /go/pkg
mkdir -m 777 /go/bin
mkdir -m 755 /go/pkg
```

#### buildを実行
```
cd /go/src/github.com/awisu2/gosample/hello/bin/
go build -v
tree /go
.
|-- bin
|-- pkg
`-- src
    `-- github.com
        `-- awisu2
            |-- golib
            |   |-- README.md
            |   |-- analyse
            |   |   `-- analyse.go
            |   |-- file
            |   |   `-- file.go
            |   `-- log
            |       `-- log.go
            `-- gosample
                |-- gorutine
                |   `-- main.go
                |-- hello
                |   |-- README.md
                |   |-- bin
                |   |   |-- bin
                |   |   `-- main.go
                |   `-- pac
                |       `-- hello.go
                |-- interface
                |   `-- main.go
                |-- switch
                |   `-- main.go
                `-- test
                    |-- test.go
                    `-- test_test.go
```

## go install

### 結果

ファイルのコンパイルまではbuildと同じ(引数も)
コンパイル後はバイナリファイルをbin、pkgフォルダ配下に配置(install)する

### 実行

```
go get -u github.com/awisu2/gosample/hello/bin
```

#### binとpkgを削除

```
rm -rf /go/bin /go/pkg
mkdir -m 777 /go/bin
mkdir -m 755 /go/pkg
```

#### installを実行

```
cd /go/src/github.com/awisu2/gosample/hello/bin/
go install -v
tree /go
.
|-- bin
|   `-- bin
|-- pkg
|   `-- linux_amd64
|       `-- github.com
|           `-- awisu2
|               |-- golib
|               |   `-- log.a
|               `-- gosample
|                   `-- hello
|                       `-- pac.a
`-- src
    `-- github.com
        `-- awisu2
            |-- golib
            |   |-- README.md
            |   |-- analyse
            |   |   `-- analyse.go
            |   |-- file
            |   |   `-- file.go
            |   `-- log
            |       `-- log.go
            `-- gosample
                |-- gorutine
                |   `-- main.go
                |-- hello
                |   |-- README.md
                |   |-- bin
                |   |   `-- main.go
                |   `-- pac
                |       `-- hello.go
                |-- interface
                |   `-- main.go
                |-- switch
                |   `-- main.go
                `-- test
                    |-- test.go
                    `-- test_test.go
```

