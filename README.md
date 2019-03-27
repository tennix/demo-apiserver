This is a demo apiserver project implemented with kubernetes [apiserver-builder](https://github.com/kubernetes-incubator/apiserver-builder-alpha)

``` shell
# We must set GOROOT environment variable explictly
export GOROOT=/usr/local/go

# Initialize the project
apiserver-boot init repo --domain pingcap.com

# Add new API resources
apiserver-boot create group version resource --group core --version v1alpha1 --kind TidbInstance
apiserver-boot create group version resource --group core --version v1alpha1 --kind TidbBackup

# Build the project
apiserver-boot build generated
apiserver-boot build executables

# Run locally
apiserver-boot run local # make sure etcd executable is in $PATH

kubectl --kubeconfig kubeconfig api-versions
kubectl --kubeconfig kubeconfig apply -f sample/tidbinstance.yaml
kubectl --kubeconfig kubeconfig apply -f sample/tidbbackup.yaml
```
