// Code generated by goctl. DO NOT EDIT.
package types

type BinFileUploadReq struct {
	FileName    string `json:"file_name,optional"`
	FileDir     string `json:"file_dir,optional"`
	FileContext string `json:"file_context"`
	PassKey     string `json:"pass_key"`
}

type BinFileUploadResp struct {
	FilePath string `json:"file_path"`
}