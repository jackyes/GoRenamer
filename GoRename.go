package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	// variables declaration
	var path string
	var prefix string
	var append string
	var StrToReplace string
	var StrReplacer string
	var RegExPattern string
	var RegExReplace string
	var RegExRemove string
	var FileType string

	var dr bool
	var tp bool
	var ta bool

	flag.StringVar(&path, "path", "", "Select path.")
	flag.StringVar(&prefix, "prefix", "", "Add <string> before filename")
	flag.StringVar(&append, "append", "", "Add <string> at the end of filename")
	flag.StringVar(&StrToReplace, "StrToReplace", "", "String to be replaced on original filename. (Use with -StrReplacer <string>.")
	flag.StringVar(&StrReplacer, "StrReplacer", "", "String to be replaced on renamed filename. (Use with -StrToReplace <string>.")
	flag.StringVar(&RegExPattern, "RegExPattern", "", "RegEx Pattern. (Use with -RegExReplace <string>. ")
	flag.StringVar(&RegExReplace, "RegExReplace", "", "String to replace the matched regex. (Use with -RegExPattern <string>.")
	flag.StringVar(&RegExRemove, "RegExRemove", "", "Remove file if match regex.")
	flag.StringVar(&FileType, "FileType", "", "Rename only file with extension (Example: .dwg).")

	flag.BoolVar(&dr, "dr", false, "Add -dr to modify also Directory name")
	flag.BoolVar(&tp, "tp", false, "Add YYYY-MM-DD as prefix.")
	flag.BoolVar(&ta, "ta", false, "Append YYYY-MM-DD to filename.")

	flag.Parse()

	if FileType != "" && !strings.HasPrefix(FileType, ".") {
		fmt.Println("Insert File ext with a starting .")
		return
	}

	if (StrToReplace != "" && StrReplacer == "") || (StrReplacer != "" && StrToReplace == "") {
		fmt.Println("If you use -StrToReplace and/or -StrReplacer, they both must be provided.")
		return
	}

	if path == "" {
		fmt.Println("Insert path. Use -h for help")
		return
	}

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		var originalpath string = filepath.Join(path, v.Name())
		ext := filepath.Ext(originalpath)
		if FileType != "" && FileType != ext {
			return //skip file if extension mismatch
		}
		if prefix != "" {
			fmt.Printf("prefix")
			if !(dr) && v.IsDir() {
				fmt.Println(v.Name() + " is a directory. Use -dr to apply change also to directory.")
			} else {
				e := os.Rename(originalpath, path+prefix+v.Name())
				if e != nil {
					fmt.Println(e)
				}
			}
		} else if append != "" {
			fmt.Printf("append")
			if !(dr) && v.IsDir() {
				fmt.Println(v.Name() + " is a directory. Use -dr to aplly change also to directory.")
			} else {
				ext := filepath.Ext(originalpath)
				e := os.Rename(originalpath, originalpath[0:len(originalpath)-len(ext)]+append+ext)
				if e != nil {
					fmt.Println(e)
				}
			}
		} else if tp {
			if !(dr) && v.IsDir() {
				fmt.Println(v.Name() + " is a directory. Use -dr to aplly change also to directory.")
			} else {
				now := time.Now()
				e := os.Rename(originalpath, filepath.Join(path, now.Format("2006-01-02")+v.Name()))
				if e != nil {
					fmt.Println(e)
				}
			}
		} else if ta {
			if !(dr) && v.IsDir() {
				fmt.Println(v.Name() + " is a directory. Use -dr to aplly change also to directory.")
			} else {
				now := time.Now()
				ext := filepath.Ext(originalpath)
				e := os.Rename(originalpath, originalpath[0:len(originalpath)-len(ext)]+now.Format("2006-01-02")+ext)
				if e != nil {
					fmt.Println(e)
				}
			}
		} else if StrToReplace != "" || StrReplacer != "" {
			if !(StrToReplace == "") && !(StrReplacer == "") {
				if !(dr) && v.IsDir() {
					fmt.Println(v.Name() + " is a directory. Use -dr to aplly change also to directory.")
				} else {
					e := os.Rename(originalpath, filepath.Join(path, strings.Replace(v.Name(), StrToReplace, StrReplacer, -1)))
					if e != nil {
						fmt.Println(e)
					}
				}
			} else {
				fmt.Println("-StrToReplace and -StrReplacer must be used togheter. User -h for help.")
			}

		} else if RegExPattern != "" || RegExReplace != "" {
			if !(RegExPattern == "") || !(RegExReplace == "") {
				if !(dr) && v.IsDir() {
					fmt.Println(v.Name() + " is a directory. Use -dr to aplly change also to directory.")
				} else {
					var re = regexp.MustCompile(RegExPattern)
					s := re.ReplaceAllString(v.Name(), RegExReplace)
					e := os.Rename(originalpath, filepath.Join(path, s))
					if e != nil {
						fmt.Println(e)
					}
				}
			} else {
				fmt.Println("-RegExPattern and -RegExReplace must be used togheter. User -h for help.")
			}
		} else if RegExRemove != "" {
			if !(dr) && v.IsDir() {
				fmt.Println(v.Name() + " is a directory. Use -dr to aplly change also to directory.")
			} else {
				sampleRegexp := regexp.MustCompile(RegExRemove)
				match := sampleRegexp.Match([]byte(v.Name()))
				if match {
					e := os.Remove(originalpath)
					if e != nil {
						fmt.Println(e)
					}
				}
			}
		}
	}
}
