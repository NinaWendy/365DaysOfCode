
Remove previous installations

```sh
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
```

Add /usr/local/go/bin to the PATH environment variable.

```sh
sudo nano.bashrc
```

```yaml
export GOROOT=/usr/local/go
export GOPATH=/usr/local/gofmt
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

Verify installation

```sh
go version
```
