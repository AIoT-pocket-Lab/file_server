package logic

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"

	"file_server/file/internal/svc"
	"file_server/file/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/sigurn/crc16"

)

type BinFileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBinFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BinFileUploadLogic {
	return &BinFileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BinFileUploadLogic) BinFileUpload(req *types.BinFileUploadReq) (resp *types.BinFileUploadResp, err error) {
	fileName := strings.TrimSpace(req.FileName)
	fileDir := strings.TrimSpace(req.FileDir)
	fileContext := req.FileContext
	passKey := strings.TrimSpace(req.PassKey)
	var file_path = ""

	if len(passKey) == 0 {
		return nil, errors.New("上传密钥为空")
	} else if passKey != l.svcCtx.Config.PassKey {
		return nil, errors.New("上传密钥非法")
	} else if (len(fileContext) % 2 != 0) || (len(fileContext) / 2 < 4) {
		return nil, errors.New("文件内容有误")
	}

	if len(fileName) == 0 {
		fileName = fmt.Sprintf("%x%x.bin", fileContext[len(fileContext)-4], fileContext[len(fileContext)-3])
	}

	if len(fileDir) == 0 {
		file_path = fmt.Sprintf("%s%s", l.svcCtx.Config.FileDir, fileName)
	} else {
		file_path = fmt.Sprintf("%s%s", fileDir, fileName)
	}
	
	fileContextByte, _ := hex.DecodeString(fileContext)
	table := crc16.MakeTable(crc16.CRC16_MODBUS)
	crc := crc16.Checksum(fileContextByte[:len(fileContextByte)-2], table)
	crcH := uint16(fileContextByte[len(fileContextByte)-1])
	crcL := uint16(fileContextByte[len(fileContextByte)-2])
	if crc != ((crcH << 8) | crcL) {
		return nil, errors.New("CRC16_Modbus校验失败")
	}

	err = os.WriteFile(file_path, fileContextByte, 0666)
	if err != nil {
		return nil, errors.New("文件保存失败")
	}

	resp = new(types.BinFileUploadResp)
	resp.FilePath = file_path

	return resp, nil
}
