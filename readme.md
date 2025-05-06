# utapi-go – UploadThing Go SDK

A Go client for interacting with [UploadThing](https://uploadthing.com) – supports file uploads, listing, deletion, and metadata operations.

---

## 📆 Installation

```bash
go get github.com/r0ld3x/utapi-go
```

---

## 🧠 Features

- ✅ Upload files to UploadThing
- 📄 List uploaded files
- 🔄 Rename files
- 🗑️ Delete files
- 📆 Fetch file info
- 🛡️ Typed API responses

---

## ⚙️ Usage

### 🔐 Initialize the client

- Apikey format: `sk_*************************`

```go
import "github.com/r0ld3x/utapi-go"

cfg := utapi.Config{
	ApiKey: "sk_*************************",
}
client := utapi.NewUtApi(cfg)
```

---

### 📁 Upload a File

```go
fileInfo, err := utapi.GetFileInfo("example.pdf")
if err != nil {
	log.Fatal(err)
}

uploadOpts := utapi.PrepareUploadOpt{
	Files: []utapi.FileRequest{*fileInfo},
}

resp, err := client.PrepareUpload(uploadOpts)
if err != nil {
	log.Fatal(err)
}

err = client.UploadFile(&resp[0], "example.pdf")
if err != nil {
	log.Fatal(err)
}
```

---

### 📃 List Files

```go
files, err := client.ListFiles(utapi.ListFilesOpts{})
if err != nil {
	log.Fatal(err)
}
fmt.Println("Files:", files)
```

---

### ✏️ Rename Files

```go
renameOpts := utapi.RenameRequest{
	Updates: []utapi.RenameUpdate{
		{
			NewName: "renamed.pdf",
			FileKey: "your-file-key",
		},
	},
}

err := client.RenameFiles(renameOpts)
if err != nil {
	log.Fatal(err)
}
```

---

### 🗑️ Delete Files

```go
deleteOpts := utapi.DeleteFilesOpt{
	FileKeys: []string{"your-file-key"},
}

err := client.DeleteFiles(deleteOpts)
if err != nil {
	log.Fatal(err)
}
```

---

## 📄 Documentation

Inline GoDocs are available at [pkg.go.dev/github.com/r0ld3x/utapi-go](https://pkg.go.dev/github.com/r0ld3x/utapi-go)
Or run locally with:

```bash
godoc -http=:6060
```

---

## 🤝 Contributing

Pull requests and issues are welcome! Please ensure code is linted and tested.

---

## 📜 License

MIT © [r0ld3x](https://github.com/r0ld3x)
