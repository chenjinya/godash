# godash

Golang 小工具合集

- [x] JSON
- [x] Variable Translation
- [x] Slice utils
- [x] HTTP utils
- [x] Mail
- [x] MD5


## Usage 

### string

godash.Template
```go
godash.Template("my name is {{.name}}", map[string]interface{}{
		"name": "John",
	})
```