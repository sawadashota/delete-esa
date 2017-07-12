# Delete esa posts

指定したIDの範囲の投稿を削除します。
移行をミスった際に使えることがあります。

## Getting Start
```bash
$ go get -u github.com/sawadashota/delete-esa
```

## Usage
```bash
$ delete-esa -e [esa team name] -eToken [esa token] -start-id [Post ID] -end-id [Post ID]
```