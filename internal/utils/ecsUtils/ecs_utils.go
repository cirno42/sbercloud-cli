package ecsUtils

import (
	"errors"
	"sbercloud-cli/api/ecs"
)

func GetEcsId(projectID, id, name string) (string, error) {
	if id != "" {
		return id, nil
	} else {
		servers, err := ecs.GetECSList(projectID, 0, 1000)
		if err != nil {
			return "", err
		}
		for _, server := range servers {
			if server.Name == name {
				return server.ID, nil
			}
		}
	}
	return "", errors.New("{\"error\" : \"No ECS with specified name\"}")
}

func GetEcsIds(projectID string, ids, name []string) ([]string, error) {
	if ids != nil && len(ids) > 0 {
		return ids, nil
	} else {
		servers, err := ecs.GetECSList(projectID, 0, 1000)
		if err != nil {
			return nil, err
		}
		nameIdx := 0
		selectedIds := make([]string, len(name))
		for i, server := range servers {
			if server.Name == name[nameIdx] {
				selectedIds[i] = server.ID //FIXME: так не будет работать
				nameIdx++
			}
		}
		return selectedIds, nil
	}
}
