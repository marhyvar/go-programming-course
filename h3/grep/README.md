# Harjoitus 3e: Grep.go

A go program that scans a file and prints out the lines that contain a search word. The search can be case-insensitive or case-sensitive (default) depending on user's choice.
How to use the program in your terminal (assuming you have [go](https://golang.org/) installed):
```
git clone https://github.com/marhyvar/go-programming-course.git
cd go-programming-course/h3/grep
```

```
go build grep.go
./grep --f <name-of-file> --w <search word> --c <y|n>
```

For example, to search for the word "lorem" in the file test.txt with a case-insensitive search:

```
./grep --f test.txt --w lorem --c y
```

Note to Windows-users: instead of ./grep just use grep in terminal (cmd/PowerShell).

## Links 

[Go by Example: String Functions](https://gobyexample.com/string-functions)

[Package strings](https://golang.org/pkg/strings/)

[Golang Scanner Package](https://linuxhint.com/golang-scanner-package/)
