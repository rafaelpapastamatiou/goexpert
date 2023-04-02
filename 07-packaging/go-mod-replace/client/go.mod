module github.com/rafaelpapastamatiou/goexpert/07-packaging/go-mod-replace/client

go 1.20

replace github.com/rafaelpapastamatiou/goexpert/07-packaging/go-mod-replace/math => ../math

require github.com/rafaelpapastamatiou/goexpert/07-packaging/go-mod-replace/math v0.0.0-00010101000000-000000000000
