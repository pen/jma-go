# jma-go

## 概要

[気象庁の天気予報ページ](https://www.jma.go.jp/bosai/forecast/#area_type=offices&area_code=130000)の裏で使われている疑似APIの結果をGoで処理しやすくするためのライブラリです。

## サンプル


```shell
go run sample/show-tokyo-forecasts/main.go
```

```shell
go run sample/show-osaka-overview/main.go
```

```shell
go run sample/dump-areas/main.go
```

```shell
curl https://www.jma.go.jp/bosai/common/const/area.json > area.json
go run sample/dump-areas-from-file/main.go area.json
```
