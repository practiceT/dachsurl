package main

import "testing"

func Example_No_Argument() {
	goMain([]string{"./dachsurl"})
	// Output:
	// トークンが指定されていません
}

func Example_Token() {
	goMain([]string{"./dachsurl", "--token"})
	// Output:
	// トークンが指定されていません
}

func Example_Clipboard() {
	goMain([]string{"./dachsurl", "--clipboard"})
	// Output:
	// トークンが指定されていません
}

func Example_Delete() {
	goMain([]string{"./dachsurl", "--delete"})
	// Output:
	// トークンが指定されていません
}

func Example_Completion() {
	goMain([]string{"./dachsurl", "--generate-completions"})
	// Output:
	// generate-completions(^^)
	// トークンが指定されていません
}

func Example_Help() {
	goMain([]string{"./dachsurl", "--help"})
	// Output:
	// dachsurl [OPTIONS] [URLs...]
	// OPTIONS
	//     -t, --token <TOKEN>      bit.lyのトークンを指定します. (必須オプション)
	//     -c, --clipboard          短縮URLをクリップボードに出力します.
	//     -d, --delete             指定した短縮URLを削除します.
	//     -h, --help               このメッセージを表示し、終了します.
	//     -v, --version            バージョンを表示し、終了します.
	// ARGUMENT
	//     URL     URLは短縮用のURLを指定します。この引数は複数指定できます.
	//             引数が指定されていない場合、dachsurlは利用可能な短縮URLのリストを表示します.
}

func Test_Main_Help(t *testing.T) {
	if status := goMain([]string{"./dachsurl", "-h"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}

func Test_Main_Version(t *testing.T) {
	if status := goMain([]string{"./dachsurl", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}

func Test_Main_Token(t *testing.T) {
	if status := goMain([]string{"./dachsurl", "-t"}); status != 3 {
		t.Error("Expected 3, got ", status)
	}
}

func Test_Main_Clipboard(t *testing.T) {
	if status := goMain([]string{"./dachsurl", "-c"}); status != 3 {
		t.Error("Expected 3, got ", status)
	}
}

func Test_Main_Delete(t *testing.T) {
	if status := goMain([]string{"./dachsurl", "-d"}); status != 3 {
		t.Error("Expected 3, got ", status)
	}
}
