## README

### How to use?

```shell
go get github.com/parocube/backroom@latest
```

### How to test locally?

#### Start test

* Replace the backrooms path with your local path

```shell
go mod edit -replace=github.com/parocube/backroom=../backroom

go mod tidy
```

* Test your code
* Restore reference

```shell
go mod edit -dropreplace=github.com/parocube/backroom

go mod tidy
```
