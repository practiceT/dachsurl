package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/practiceT/dachsurl"
	flag "github.com/spf13/pflag"
)

const VERSION = "0.1.6"

func versionString(args []string) string {
	prog := "dachsurl"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf("%s version %s", prog, VERSION)
}

/*
helpMessage prints the help message.
This function is used in the small tests, so it may be called with a zero-length slice.
*/
func helpMessage(args []string) string {
	prog := "dachsurl"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf(`%s [OPTIONS] [URLs...]
OPTIONS
    -t, --token <TOKEN>      bit.lyのトークンを指定します. (必須オプション)
    -c, --clipboard          短縮URLをクリップボードに出力します.
    -d, --delete             指定した短縮URLを削除します.
    -h, --help               このメッセージを表示し、終了します.
    -v, --version            バージョンを表示し、終了します.
ARGUMENT
    URL     URLは短縮用のURLを指定します。この引数は複数指定できます.
            引数が指定されていない場合、dachsurlは利用可能な短縮URLのリストを表示します.`, prog)
}

type DachsurlError struct {
	statusCode int
	message    string
}

func (e DachsurlError) Error() string {
	return e.message
}

type flags struct {
	deleteFlag    bool
	helpFlag      bool
	versionFlag   bool
	clipboardFlag bool
}

type runOpts struct {
	token string
}

/*
This struct holds the values of the options.
*/
type options struct {
	runOpt  *runOpts
	flagSet *flags
}

func newOptions() *options {
	return &options{runOpt: &runOpts{}, flagSet: &flags{}}
}

func (opts *options) mode(args []string) dachsurl.Mode {
	switch {
	case len(args) == 0:
		return dachsurl.List
	case opts.flagSet.deleteFlag:
		return dachsurl.Delete
	default:
		return dachsurl.Shorten
	}
}

/*
Define the options and return the pointer to the options and the pointer to the flagset.
*/
func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := newOptions()
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.StringVarP(&opts.runOpt.token, "token", "t", "", "bit.lyのトークンを指定します. (必須オプション)")
	flags.BoolVarP(&opts.flagSet.deleteFlag, "delete", "d", false, "指定した短縮URLを削除します.")
	flags.BoolVarP(&opts.flagSet.clipboardFlag, "clipboard", "c", false, "短縮URLをクリップボードに出力します.")
	flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "このメッセージを表示し、終了します.")
	flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "バージョンを表示し、終了します.")
	return opts, flags
}

/*
parseOptions parses options from the given command line arguments.
*/
func parseOptions(args []string) (*options, []string, *DachsurlError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.flagSet.helpFlag {
		fmt.Println(helpMessage(args))
		return nil, nil, &DachsurlError{statusCode: 0, message: ""}
	}
	if opts.flagSet.versionFlag {
		fmt.Println(versionString(args))
		return nil, nil, &DachsurlError{statusCode: 0, message: ""}
	}
	if opts.runOpt.token == "" {
		return nil, nil, &DachsurlError{statusCode: 3, message: "トークンが指定されていません"}
	}
	return opts, flags.Args(), nil
}

func shortenEach(opts *options, bitly *dachsurl.Bitly, config *dachsurl.Config, url string) error {
	result, err := bitly.Shorten(config, url)
	if err != nil {
		return err
	}
	fmt.Println(result)
	if opts.flagSet.clipboardFlag {
		clipboard.WriteAll(result.Shorten)
	}
	return nil
}

func deleteEach(bitly *dachsurl.Bitly, config *dachsurl.Config, url string) error {
	return bitly.Delete(config, url)
}

func listUrls(bitly *dachsurl.Bitly, config *dachsurl.Config) error {
	urls, err := bitly.List(config)
	if err != nil {
		return err
	}
	for _, url := range urls {
		fmt.Println(url)
	}
	return nil
}

func performImpl(args []string, executor func(url string) error) *DachsurlError {
	for _, url := range args {
		err := executor(url)
		if err != nil {
			return makeError(err, 3)
		}
	}
	return nil
}

func perform(opts *options, args []string) *DachsurlError {
	bitly := dachsurl.NewBitly("")
	config := dachsurl.NewConfig("", opts.mode(args))
	config.Token = opts.runOpt.token
	switch config.RunMode {
	case dachsurl.List:
		err := listUrls(bitly, config)
		return makeError(err, 1)
	case dachsurl.Delete:
		return performImpl(args, func(url string) error {
			return deleteEach(bitly, config, url)
		})
	case dachsurl.Shorten:
		return performImpl(args, func(url string) error {
			return shortenEach(opts, bitly, config, url)
		})
	}
	return nil
}

func makeError(err error, status int) *DachsurlError {
	if err == nil {
		return nil
	}
	ue, ok := err.(*DachsurlError)
	if ok {
		return ue
	}
	return &DachsurlError{statusCode: status, message: err.Error()}
}

func goMain(args []string) int {
	opts, args, err := parseOptions(args)
	if err != nil {
		if err.statusCode != 0 {
			fmt.Println(err.Error())
		}
		return err.statusCode
	}
	if err := perform(opts, args); err != nil {
		fmt.Println(err.Error())
		return err.statusCode
	}
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
