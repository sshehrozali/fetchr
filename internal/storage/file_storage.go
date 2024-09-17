package storage

import ("os"
		"filedownloader/internal/models"
)

type FileStorage struct {}

func (*FileStorage) Save(downloadResult models.DownloadResult) (int, error) {
	
}