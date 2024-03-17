package db

import (
	"fmt"
	"go-replica-db/pkg/util"
)

const (
	LogsFilePath = "data/logs/"
)

func Insert(dbId string, key string, value string) error {
	err := util.WriteToFile(LogsFilePath+dbId, fmt.Sprintf("%s,%s", key, value))
	if err != nil {
		return err
	}
	return nil
}

func Read(filePath, key string) (string, error) {
	val, err := util.FindLatestValueForKey(filePath, key)
	if err != nil {
		return "", err
	}
	return val, nil
}
