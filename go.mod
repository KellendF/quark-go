module github.com/quarkcms/quark-go

go 1.17

replace (
    github.com/quarkcms/quark-go/bootstrap => ./quark-go/bootstrap
    github.com/quarkcms/quark-go/app/http/controllers => ./quark-go/app/http/controllers
)
