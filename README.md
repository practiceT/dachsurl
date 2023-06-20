# dachsurl
URL短縮ツール(開発中) :pleading_face:

<!-- https://www.apache.org/licenses/LICENSE-2.0 -->

[![build](https://github.com/practiceT/dachsurl/actions/workflows/build.yml/badge.svg)](https://github.com/practiceT/dachsurl/actions/workflows/build.yml)
[![Coverage Status](https://coveralls.io/repos/github/practiceT/dachsurl/badge.svg)](https://coveralls.io/github/practiceT/dachsurl)
[![codebeat badge](https://codebeat.co/badges/518faef1-bda5-47e1-9083-0414106885b8)](https://codebeat.co/projects/github-com-practicet-dachsurl-main)
[![Go Report Card](https://goreportcard.com/badge/github.com/practiceT/dachsurl)](https://goreportcard.com/report/github.com/practiceT/dachsurl)

![GitHub](https://img.shields.io/github/license/practiceT/dachsurl?style=plastic)
![Version](https://img.shields.io/badge/Version-0.1.23-informational)


## 概要
CLIによるシンプルなURL短縮ツール。
入力した冗長なURLをダックスフントの足のように短くします。
URLや足を長くしたいというご要望にはお応えできません。


## 使用方法
```
Usage:
  dachsurl [OPTIONS] [URLs...]

OPTIONS
    -t, --token <TOKEN>      bit.lyのトークンを指定します. (必須オプション)
    -c, --clipboard          短縮URLをクリップボードに出力します.
    -d, --delete             指定した短縮URLを削除します.
    -h, --help               このメッセージを表示し、終了します.
    -v, --version            バージョンを表示し、終了します.
ARGUMENT
    URL     URLは短縮用のURLを指定します。この引数は複数指定できます.
            引数が指定されていない場合、dachsurlは利用可能な短縮URLのリストを表示します.
```
※ 「-c, --clipboard」オプションを指定した状態で、複数URLを指定すると、一番最後のURLがクリップボードにコピーされます.


<!-- ## インストール方法

### Homebrew

```
brew install practicet/brew/dachshurl
```

### Docker

```
docker run -it --rm practicet/dachshurl:latest -t <token> <url...>
``` -->


## プロジェクトについて

### 開発者
HAYASHI Tomokazu

### ライセンス
[Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0)

| Permissions | |
| :--- | :---: |
| Commercial use | :white_check_mark: |
| Modification | :white_check_mark: |
| Distribution | :white_check_mark: |
| Patent use | :white_check_mark: |
| Private use | :white_check_mark: |

| Limitations | |
| :--- | :---: |
| Trademark use | :x: |
| Liability | :x: |
| Warranty | :x: |

| Conditions | |
| :--- | :---: |
| License and copyright notice | :information_source: |
| State changes | :information_source: |

### アイコン
| ![アイコン](./docs/static/images/dachsurl.svg) | ![アイコン](./docs/static/images/dachsurl.svg) |
| :---: | :---: |
| ![アイコン](./docs/static/images/dachsurl.svg) | ![アイコン](./docs/static/images/dachsurl.svg) |
| ![アイコン](./docs/static/images/dachsurl.svg) | ![アイコン](./docs/static/images/dachsurl.svg) |


### 名前の由来
URLをダックスフント(Dachshund)の足のように短くするという由来です。
