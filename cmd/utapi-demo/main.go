package main

import (
	ut "github.com/r0ld3x/utapi-go"
)

func main() {
	utapi := ut.NewUtApi("")

	fileInfo, err := ut.GetFileInfo("abc.pdf")
	if err != nil {
		panic(err)
	}
	data := ut.PrepareUploadOpt{
		Files:        []ut.FileRequest{*fileInfo},
		CallbackURL:  "https://nw-apiolakshdkj.com/",
		CallbackSlug: fileInfo.Name,
		RouteConfig:  []string{"pdf"},
	}
	result, err := utapi.PrepareUpload(data)
	if err != nil {
		panic(err)
	}
	err = utapi.UploadFile(result, "abc.pdf")
	if err != nil {
		panic(err)
	}
	println(result.UfsURL)

	// keys := []string{"92cb0884-7d43-4e53-a0a3-358adb99b3d7-9w82os.md"}
	// result, err := utapi.DeleteFiles(ut.DeleteFilesOpt{FileKeys: keys})
	// if err != nil {
	// 	panic(err)
	// }
	// keys := ut.RenameFilesOpt{
	// 	NewName: "renamed.md",
	// 	FileKey: "92cb0884-7d43-4e53-a0a3-358adb99b3d7-9w82os.md",
	// }
	// result, err := utapi.RenameFiles(ut.RenameFilesOpts{keys})
	// if err != nil {
	// 	panic(err)
	// }
	// println(result.RenamedCount)
	// result, err := utapi.ListFiles(ut.ListFilesOpts{Limit: 10, Offset: 0})
	// if err != nil {
	// 	panic(err)
	// }
	// for _, file := range result.Files {
	// 	println(file.ID, file.Key, file.Status)
	// }

	// result, err := utapi.GetUsageInfo()
	// if err != nil {
	// 	panic(err)
	// }
	// println(result.AppTotalBytes)
	// result, err := utapi.ListFiles(ut.ListFilesOpts{Limit: 10, Offset: 0})

}
