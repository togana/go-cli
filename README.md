go-cli
====

# Build

Dockerfile内で mac 用の設定しています。多分環境変数変えたら他の環境でも動くはず。

```
$ make
```

# 使い方

## echo

```
$ ./bin/echo aaa b bb cccc
aaa b bb cccc
```

```
$ ./bin/echo -n aaa b bb cccc
aaa b bb cccc%
```

## wc 

```
$ ./bin/wc ./src/wc/main.go
     108     281    2001 ./src/wc/main.go
```

```
$ ./bin/wc -l ./src/wc/main.go
     108 ./src/wc/main.go
```

```
$ ./bin/wc -w ./src/wc/main.go
     281 ./src/wc/main.go
```

```
$ ./bin/wc -c ./src/wc/main.go
    2001 ./src/wc/main.go
```

# 参考

- https://ja.stackoverflow.com/questions/50938/go%E8%A8%80%E8%AA%9E%E3%81%A7%EF%BC%91%E3%81%A4%E3%81%AE%E3%83%AA%E3%83%9D%E3%82%B8%E3%83%88%E3%83%AA%E3%81%A7%E8%A4%87%E6%95%B0%E3%81%AE%E3%83%90%E3%82%A4%E3%83%8A%E3%83%AA%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%82%92%E3%83%93%E3%83%AB%E3%83%89%E3%81%99%E3%82%8B%E6%96%B9%E6%B3%95
- http://golang.jp/category/%E3%83%89%E3%82%AD%E3%83%A5%E3%83%A1%E3%83%B3%E3%83%88
- https://www.youtube.com/watch?v=7gnpR0Dinqs
- https://www.youtube.com/watch?v=eVw97iUZi5w
- https://takatoshiono.hatenablog.com/entry/2016/09/22/024605
