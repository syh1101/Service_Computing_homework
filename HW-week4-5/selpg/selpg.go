package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"github.com/spf13/pflag"
)

type selpgArgs struct {
	startPage  int
	endPage    int
	inFileName string
	pageLength    int
	pageType   bool
	printDest  string
}

func main() {
	var args selpgArgs
	AnalyArgs(&args)
	checkArgs(&args)
	excute(&args)
}

func AnalyArgs(args *selpgArgs) {
	pflag.IntVarP(&(args.startPage), "startPage", "s", -1, "Define startPage")
	pflag.IntVarP(&(args.endPage), "endPage", "e", -1, "Define endPage")
	pflag.IntVarP(&(args.pageLength), "pageLength", "l", 8, "Define pageLength")
	pflag.StringVarP(&(args.printDest), "printDest", "d", "", "Define printDest")
	pflag.BoolVarP(&(args.pageType), "pageType", "f", false, "Define pageType")
	pflag.Parse()
	argLeft := pflag.Args()
	if len(argLeft) > 0 {
		args.inFileName = string(argLeft[0])
	} else {
		args.inFileName = ""
	}
}

func checkArgs(args *selpgArgs) {

	if (args.startPage == -1) || (args.endPage == -1) {
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage and endPage can't be empty! \n")
		os.Exit(2)
	} else if (args.startPage <= 0) || (args.endPage <= 0) {
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage and endPage can't be negative! \n")
		os.Exit(3)
	} else if args.startPage > args.endPage {
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage can't be bigger than the endPage! \n")
		os.Exit(4)
	} else if (args.pageType == true) && (args.pageLength != 8) {
		fmt.Fprintf(os.Stderr, "\n[Error]The command -l and -f are confilct! \n")
		os.Exit(5)
	} else if args.pageLength <= 0 {
		fmt.Fprintf(os.Stderr, "\n[Error]The pageLength can't be less than 1 ! \n")
		os.Exit(6)
	} 
}

func checkError(err error, object string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n[Error]%s:", object)
		panic(err)
	}
}

func excute(args *selpgArgs) {
	var fin *os.File
	if args.inFileName == "" {
		fin = os.Stdin
	} else {
		checkfile(args.inFileName)
		var err error
		fin, err = os.Open(args.inFileName)
		checkError(err, "File input")
	}

	if len(args.printDest) == 0 {
		outDes(os.Stdout, fin, args.startPage, args.endPage, args.pageLength, args.pageType)
	} else {
		outDes(cmdExec(args.printDest), fin, args.startPage, args.endPage, args.pageLength, args.pageType)
	}
}

func checkfile(filename string) {
	_, errFileExits := os.Stat(filename)
	if os.IsNotExist(errFileExits) {
		fmt.Fprintf(os.Stderr, "\n[Error]: file \"%s\" does not exist\n", filename)
		os.Exit(7)
	}
}

func cmdExec(printDest string) io.WriteCloser {
	cmd := exec.Command("lp", "-d"+printDest)
	fout, err := cmd.StdinPipe()
	checkError(err, "StdinPipe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	errStart := cmd.Start()
	checkError(errStart, "CMD Start")
	return fout
}

func outDes(fout interface{}, fin *os.File, pageStart int, pageEnd int, pageLength int, pageType bool) {

	lineCount := 0
	pageCount := 1
	buf := bufio.NewReader(fin)
	for true {

		var line string
		var err error
		if pageType {
			// command -f
			line, err = buf.ReadString('\f')
			pageCount++
		} else {
			// command  -l
			line, err = buf.ReadString('\n')
			lineCount++
			if lineCount > pageLength {
				pageCount++
				lineCount = 1
			}
		}

		if err == io.EOF {
			break
		}
		checkError(err, "file read in")

		if (pageCount >= pageStart) && (pageCount <= pageEnd) {
			var outputErr error
			if stdOutput, ok := fout.(*os.File); ok {
				_, outputErr = fmt.Fprintf(stdOutput, "%s", line)
			} else if pipeOutput, ok := fout.(io.WriteCloser); ok {
				_, outputErr = pipeOutput.Write([]byte(line))
			} else {
				fmt.Fprintf(os.Stderr, "\n[Error]:fout type error. ")
				os.Exit(8)
			}
			checkError(outputErr, "Error happend when output the pages.")
		}
	}
	if pageCount < pageStart {
		fmt.Fprintf(os.Stderr, "\n[Error]: startPage (%d) greater than total pages (%d), no output written\n", pageStart, pageCount)
		os.Exit(9)
	} else if pageCount < pageEnd {
		fmt.Fprintf(os.Stderr, "\n[Error]: endPage (%d) greater than total pages (%d), less output than expected\n", pageEnd, pageCount)
		os.Exit(10)
	}
}
