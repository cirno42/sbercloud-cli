package evs

import (
	"time"
)

func WaitUntilJobSuccess(projectID, jobID string) ([]string, error) {
	job, err := GetInfoAboutBatchTask(projectID, jobID)
	if err != nil {
		return nil, err
	}
	jobRes := job.Status
	for jobRes != "SUCCESS" && jobRes != "FAIL" {
		job, err = GetInfoAboutBatchTask(projectID, jobID)
		if err != nil {
			return nil, err
		}
		jobRes = job.Status
		time.Sleep(1000 * time.Millisecond)
	}
	if job.Entities.SubJobs == nil {
		job, err := GetInfoAboutSingleTask(projectID, jobID)
		if err != nil {
			return nil, err
		}
		res := make([]string, 1)
		res[0] = job.Entities.VolumeID
		return res, err
	} else {
		res := make([]string, len(job.Entities.SubJobs))
		for i, subJob := range job.Entities.SubJobs {
			res[i] = subJob.Entities.VolumeID
		}
		return res, nil
	}
}
