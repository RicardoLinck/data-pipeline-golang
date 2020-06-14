package main

import "time"

type saveResult struct {
	idsSaved  []string
	timestamp int64
}

func saveData(ic <-chan externalData) <-chan saveResult {
	oc := make(chan saveResult)

	go func() {
		const batchSize = 7
		batch := make([]string, 0)
		for input := range ic {
			if len(batch) < batchSize {
				batch = append(batch, input.inputData.id)
			} else {
				oc <- persistBatch(batch)
				batch = []string{input.inputData.id}
			}
		}

		if len(batch) > 0 {
			oc <- persistBatch(batch)
		}

		close(oc)
	}()

	return oc
}

func persistBatch(batch []string) saveResult {
	return saveResult{batch, time.Now().UnixNano()}
}
