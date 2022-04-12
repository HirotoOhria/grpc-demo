package handler

import (
    "bytes"
    "io"
    "net/http"

    "image.uploader/api/gen/pb"
)

type ImageUploadHandler struct {
    files map[string][]byte
}

func NewImageUploadHandler() *ImageUploadHandler {
    return &ImageUploadHandler{
        files: make(map[string][]byte),
    }
}

func (h *ImageUploadHandler) Upload(stream pb.ImageUploadService_UploadServer) error {
    req, err := stream.Recv()
    if err != nil {
        return err
    }

    // In the specifications, it was initially meta information.
    fileName := req.GetFileMeta().GetFileName()
    buf := &bytes.Buffer{}

    for {
        r, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }

        chunk := r.GetData()
        if _, err := buf.Write(chunk); err != nil {
            return err
        }
    }

    // Process file

    uuid := "lazy_uuid"
    data := buf.Bytes()
    mimeType := http.DetectContentType(data)

    err = stream.SendAndClose(&pb.ImageUploadResponse{
        Uuid:        uuid,
        Size:        int32(len(data)),
        ContentType: mimeType,
        Filename:    fileName,
    })
    if err != nil {
        return err
    }

    return nil
}
