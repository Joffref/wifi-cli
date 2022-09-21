This pkg is a fork from this repository : https://github.com/MonkeyBuisness/golang-iwlist
You can't use the original package because there is a go.mod issue.
```bash
➜  WIFI-cli go get github.com/MonkeyBuisness/golang-iwlist                                          

go: downloading github.com/MonkeyBuisness/golang-iwlist v0.0.0-20220920142351-7b59e0d73297
go: github.com/MonkeyBuisness/golang-iwlist@v0.0.0-20220920142351-7b59e0d73297: parsing go.mod:
        module declares its path as: github.com/MonkeyBusiness/golang-iwlist
                but was required as: github.com/MonkeyBuisness/golang-iwlist

```

After the class, I will make a PR to correct this. Then, I might remove this package and use the original one.
