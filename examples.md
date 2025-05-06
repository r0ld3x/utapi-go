# 📦 utapi-go Usage Examples

This document provides examples of how to use the `utapi-go` package to interact with the UploadThing API.

---

## 🛠️ Setup

Import the package:

```go
import (
    ut "github.com/r0ld3x/utapi-go/api"
)
```

Create a new instance:

```go
utapi := ut.NewUtApi("YOUR_UPLOADTHING_API_KEY")
```

---

## 📤 Uploading a File

```go
fileInfo, err := ut.GetFileInfo("abc.pdf")
if err != nil {
    panic(err)
}

uploadOptions := ut.PrepareUploadOpt{
    Files:        []ut.FileRequest{*fileInfo},
    CallbackURL:  "https://example.com/callback",
    CallbackSlug: fileInfo.Name,
    RouteConfig:  []string{"pdf"},
}

prepared, err := utapi.PrepareUpload(uploadOptions)
if err != nil {
    panic(err)
}

err = utapi.UploadFile(prepared, "abc.pdf")
if err != nil {
    panic(err)
}

println("Uploaded to:", prepared.UfsURL)
```

---

## 🗑️ Deleting Files

```go
keys := []string{"your-file-key.md"}

deleted, err := utapi.DeleteFiles(ut.DeleteFilesOpt{FileKeys: keys})
if err != nil {
    panic(err)
}

println("Files deleted:", deleted.Success)
```

---

## ✏️ Renaming a File

```go
rename := ut.RenameFilesOpt{
    NewName: "renamed.md",
    FileKey: "your-file-key.md",
}

renamed, err := utapi.RenameFiles(ut.RenameFilesOpts{rename})
if err != nil {
    panic(err)
}

println("Files renamed:", renamed.RenamedCount)
```

---

## 📄 Listing Files

```go
files, err := utapi.ListFiles(ut.ListFilesOpts{Limit: 10, Offset: 0})
if err != nil {
    panic(err)
}

for _, file := range files.Files {
    println(file.ID, file.Key, file.Status)
}
```

---

## 📊 Getting Usage Info

```go
usage, err := utapi.GetUsageInfo()
if err != nil {
    panic(err)
}

println("Total bytes used:", usage.AppTotalBytes)
```

---

## 🧪 Notes

- Make sure to set your API key using `NewUtApi("your-api-key")`
- Use error handling appropriately in production environments

---

Enjoy using `utapi-go`! 🎉
