package ecs

import "time"

func WaitUntilJobSuccess(projectID, jobID string) ([]string, error) {
	jobRes := ""
	for jobRes != "SUCCESS" && jobRes != "FAIL" {
		job, err := GetInfoAboutTask(projectID, jobID)
		if err != nil {
			return nil, err
		}
		jobRes = job.Status
		time.Sleep(1000 * time.Millisecond)
	}
	job, err := GetInfoAboutTask(projectID, jobID)
	if err != nil {
		return nil, err
	}
	res := make([]string, len(job.Entities.SubJobs))
	for i, subJob := range job.Entities.SubJobs {
		res[i] = subJob.Entities.ServerID
	}
	return res, nil
}

func WaitUntilJobSuccessAndGetStatus(projectID, jobID string) (string, error) {
	jobRes := ""
	for jobRes != "SUCCESS" && jobRes != "FAIL" {
		job, err := GetInfoAboutTask(projectID, jobID)
		if err != nil {
			return "", err
		}
		jobRes = job.Status
		time.Sleep(1000 * time.Millisecond)
	}
	return jobRes, nil
}
