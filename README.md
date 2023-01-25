# Go-batch-sample


# 構成

```
- main.go
- batch/
    - hello/
        - hello.go
       
    - morgning/
        - morning.go
        
    - countryIp/
        - countryIp.go
```

# 実行

## ソースコード
```
go run main.go -s=NAME morning
```

```
go run main.go -s=NAME hello
```

```
go run main.go countryIp
```

## バイナリファイル

```
./go-scraping -s=NAME morning
```

```
./go-scraping -s=NAME hello
```

```
./go-scraping countryIp
```
