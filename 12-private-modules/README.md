## Add private repository to GOPRIVATE environment variable

export GOPRIVATE=url1,url2,url3

Example: export GOPRIVATE=github.com/rafaelpapastamatiou/goexpert-private-repo


## Configure git to use SSH instead of HTTPS (~/.gitconfig)

[url "ssh://git@github.com/"]
	insteadOf = https://github.com/


## Create vendor directory to keep dependencies safe

go mod vendor