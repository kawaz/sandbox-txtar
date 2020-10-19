package main

import (
	"fmt"
	"golang.org/x/tools/txtar"
)

func main() {
	archive, err := txtar.ParseFile("test.txtar")
	if err != nil {
		panic(err)
	}
	fmt.Printf("### コメントという名の無名の1ファイル目的なやつ (%d bytes)\n%s", len(archive.Comment), string(archive.Comment))
	for _, f := range archive.Files {
		fmt.Printf("### %s (%d bytes)\n", f.Name, len(f.Data))
		fmt.Println(string(f.Data))
	}

	// 最終ファイルの最終行に改行が無いと勝手に改行が補完されるらしいことの確認
	archive = txtar.Parse([]byte("-- last.txt --\nABC"))
	fmt.Printf("### %s (%d bytes) <-改行補完されたら4バイトになる筈\n%s", archive.Files[0].Name, len(archive.Files[0].Data), archive.Files[0].Data)
}
