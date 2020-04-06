# Fails to deploy

```
$ gcloud functions deploy HelloSubDir --runtime go113 --trigger-http --allow-unauthenticated --project some-random-project
Deploying function (may take a while - up to 2 minutes)...failed.                                                                                                                           
ERROR: (gcloud.functions.deploy) OperationError: code=3, message=Build failed: srv/gopath/src/p/vendor/hello/hello.go:9:2: cannot find package "github.com/wershlak/hellosubdir/subdir" in any of:
        /tmp/staging/srv/gopath/src/p/vendor/vendor/github.com/wershlak/hellosubdir/subdir (vendor tree)
        /tmp/staging/srv/gopath/src/p/vendor/github.com/wershlak/hellosubdir/subdir
        /usr/local/go/src/github.com/wershlak/hellosubdir/subdir (from $GOROOT)
        /tmp/staging/srv/gopath/src/github.com/wershlak/hellosubdir/subdir (from $GOPATH)
```

So it's not finding the subdirectory

https://cloud.google.com/functions/docs/writing/#functions-writing-file-structuring-go

```Note: You must use Go modules, by providing a go.mod file, to use sub-packages.```

But https://cloud.google.com/functions/docs/writing/specifying-dependencies-go says:

```Warning: If you have both a go.mod file and a vendor directory at the root of your project, the contents of the vendor directory will be ignored when your function is built in the cloud. To ensure that your vendor directory is used, you must exclude the go.mod file from your project's source code prior to deployment.```

So with dependencies from an internal repo it seems you can not structure your functions to use subdirectories.
