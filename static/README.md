# Bindata

To update the bindata within this project you need to install the go-bindata
utitlity, follow the steps below to get this utility and to update the included
bindata. We are not covering within this steps how to setup an Go environment.


## Prepare

To get the required utility execute the following commands, generally this have
to be done only once.

```
go get -u github.com/jteeuwen/go-bindata
go install github.com/jteeuwen/go-bindata/...
```


## Update

To update the current bindata within the resulting binary you need to execute
the following command always if one of these configuration files changes.
Execute this command always in the root directory of this project.

```
go-bindata -o=modules/bindata/conf/static.go \
  -ignore='\\.DS_Store|README\.md' -prefix=static/ -pkg=conf static/conf/...

go-bindata -o=modules/bindata/gitignore/static.go -tags='!nobindata' \
  -ignore='\\.DS_Store|README\.md' -prefix=static/ -pkg=gitignore static/gitignore/...

go-bindata -o=modules/bindata/license/static.go -tags='!nobindata' \
  -ignore='\\.DS_Store|README\.md' -prefix=static/ -pkg=license static/license/...

go-bindata -o=modules/bindata/locale/static.go -tags='!nobindata' \
  -ignore='\\.DS_Store|README\.md' -prefix=static/ -pkg=locale static/locale/...

go-bindata -o=modules/bindata/public/static.go -tags='!nobindata' \
  -ignore='\\.DS_Store|README\.md' -prefix=static/public/ -pkg=public static/public/...

go-bindata -o=modules/bindata/templates/static.go -tags='!nobindata' \
  -ignore='\\.DS_Store|README\.md' -prefix=static/templates/ -pkg=templates static/templates/...
```
