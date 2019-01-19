package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"bufio"
	flag "github.com/spf13/pflag"
)

type selpg_args struct {
	start_page  int
	end_page    int
	page_len    int
	page_type   bool
	filename string
	destination  string
}

var program_name string

var sa = new(selpg_args)

func main() {
	program_name = os.Args[0]
	var args selpg_args
	Init(&args)
	ProcessArgs(&args)
	ProcessInput(&args)
}

func Init(args *selpg_args) {
	flag.IntVarP(&args.start_page,"start_page", "s", -1, "start page.")
	flag.IntVarP(&args.end_page,"end_page", "e", -1, "end page.")
	flag.IntVarP(&args.page_len,"page_len", "l", 72, "page length.")
	flag.BoolVarP(&args.page_type,"page_type", "f", false, "page type")
	flag.StringVarP(&args.destination,"destination", "d", "", "destination")
    flag.Usage = Usage
	flag.Parse()
}

func Usage() {
	fmt.Fprintf(os.Stderr, "\nUSAGE: %s [-s start_page] [-e end_page] [-l page_len ] [-f page_type] [-d destination ] [filename]\n", program_name)
}

func ProcessArgs(args *selpg_args) {
	if flag.NFlag() < 2 {
		fmt.Fprintf(os.Stderr, "%s: not enough args\n\n", program_name)
		flag.Usage()
		os.Exit(1)
	}

	if args.start_page > args.end_page || args.start_page < 1 || args.end_page < 1 {
		fmt.Fprintln(os.Stderr, "Invalid page")
		flag.Usage()
		os.Exit(1)
	}
}

func ProcessInput(args *selpg_args) {
	var err error
	var cmd *exec.Cmd
	var stdin io.WriteCloser

	if args.destination != "" {
		cmd = exec.Command("cat", "-n")
		stdin, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		stdin = nil
	}

	if flag.NArg() > 0 {
		args.filename = flag.Arg(0)
		output, err := os.Open(args.filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		reader := bufio.NewReader(output)
		if args.page_type {
			for page_num := 1; page_num <= args.end_page; page_num++ {
				line, err := reader.ReadString('\f')
				if err != io.EOF && err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if err == io.EOF {
					break
				}
				if page_num >= args.start_page {
	          		if args.destination != "" {
	            		stdin.Write([]byte(string(line) + "\n"))
		        	} else {
			        	fmt.Print(string(line))
		        	}
	        	}
	        	page_num ++
			}
		} else {
			line_num := 0
	      	page_num := 1
			for {
				line, _, err := reader.ReadLine()
				if err != io.EOF && err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if err == io.EOF {
					break
				}
	        	if page_num >= args.start_page && page_num <= args.end_page {
	          		if args.destination != "" {
	            		stdin.Write([]byte(string(line) + "\n"))
	          		} else {
	            		fmt.Println(string(line))
	          		}
	        	}
	        	line_num ++
	        	if line_num == args.page_len {
	          		page_num ++
	          		line_num = 0
	        	}
	        	if page_num > args.end_page {
	          		break
	        	}
			}
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		line_num := 0
    	page_num := 1
		out_string := ""
		for scanner.Scan() {
			line := scanner.Text()
			line += "\n"
      		if page_num >= args.start_page && page_num <= args.end_page {
        		out_string += line
      		}
      		line_num ++
      		if line_num == args.page_len {
        		line_num = 0
        		page_num ++
      		}
		}
		if args.destination != "" {
      		stdin.Write([]byte(string(out_string) + "\n"))
    	} else {
      		fmt.Println(string(out_string))
    	}
	}

	if args.destination != "" {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
