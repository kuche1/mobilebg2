
# Run

```
go run .
```

# Cross-Compile for Windows

```
sudo pacman -S mingw-w64-gcc
CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build .
```
