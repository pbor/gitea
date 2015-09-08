# Bindata

To update the current bindata within the resulting binary you need to execute
the following command always if one of these files changes. Afterwards you just
need to add these changes to the git repository and that's it.

```sh
cd ${GOPATH}/src/go-gitea/gitea
go run make.go deps bindata
```
