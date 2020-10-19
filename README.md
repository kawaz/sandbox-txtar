`txtar` っていうテキスト用のマルチパート表現が GoPlayground で使われてて結構便利そうなので試してみた。

# `txtar` とは

- 複数ファイルを1テキストで表現
- mimeマルチパートより遥かに簡単
- ディレクトリ階層も扱える
- 行頭に "-- " を含むファイルは扱えない？
- テストfixtureとか書くのとかに便利そう

##TODO
- vim-precious でいい感じの設定作る


# 使ってみた
golang から `golang.org/x/tools/txtar` で簡単に扱えるぽいから試してみた、ら簡単に使えた。

`go run github.com/kawaz/sandbox-txtar`

```
### コメントという名の無名の1ファイル目的なやつ (96 bytes)
コメント！
コメント！ コメント！
コメント！ コメント！ コメント！
### マルチバイトで「 」空白も入ってるファイル名.txt (254 bytes)
問題ないよね？
↓とか書くとどうなるんだろ？後ろの -- がなければOK？
- あああ
-- いえーい見てる？
お、いけた。
どうやら最後の -- がなければ頭に行頭 -- があっても問題無さそうだ。


### dir/f1.txt (16 bytes)
あいうえお

### dir/f2.txt (16 bytes)
かきくけこ

### 最終ファイルの最終行のファイルは勝手に補完される仕様らしい (104 bytes)
って思ったらvimでファイル作ると勝手に改行入ってたｗ
ので別で試験しよう

### last.txt (4 bytes) <-改行補完されたら4バイトになる筈
ABC
```


