// Package herald acts as an interface between the storage and service packages
package herald

import (
	"container/list"
	"fmt"
	"sync"

	"github.com/will-rowe/herald/src/services"
	"github.com/will-rowe/herald/src/storage"
)

// Herald is the struct for holding runtime data
type Herald struct {
	sync.Mutex                         // to make the UI binding thread safe
	config            *services.Config // a copy of the config being used by the current Herald instance
	store             *storage.Storage // the key-value store for the samples
	announcementQueue *list.List       // a FIFO queue for announcements

	// runtime count info for JS:
	runCount            int // the number of runs currently in the store
	sampleCount         int // the number of samples currently in the store
	untaggedRecordCount int // the number of samples in the store that are untagged
	taggedRecordCount   int // the number of samples in the store that are tagged with at least one process
	announcementCount   int // the number of samples in the store that have been announced

	// easy access label holders for JS
	sampleDetails [][]string // used to store all the sample labels, creation dates and corresponding run in memory (for JS to access)
	runLabels     []string   // used to store all the run names in memory (for JS to access)
	storeLocation string     // where the store is located on disk
}

// InitHerald will initiate the Herald instance
func InitHerald(storeLocation string) (*Herald, error) {

	// load a copy of the config
	config, err := services.InitConfig(storeLocation)
	if err != nil {
		return nil, err
	}

	// load the store
	var store *storage.Storage
	if store, err = storage.OpenStorage(storeLocation); err != nil {
		return nil, err
	}

	// get a new instance
	heraldObj := &Herald{
		config:            config,
		store:             store,
		announcementQueue: list.New(),
		sampleDetails:     make([][]string, 3),
		storeLocation:     storeLocation,
	}

	// populate runtime info
	if err := heraldObj.GetRuntimeInfo(); err != nil {
		heraldObj.Destroy()
		return nil, err
	}
	return heraldObj, nil
}

// Destroy will properly close down the Herald instance and sync the store to disk
func (herald *Herald) Destroy() error {
	herald.Lock()
	defer herald.Unlock()
	return herald.store.CloseStorage()
}

// WipeStorage will clear all samples and runs from storage and reset the runtime info
func (herald *Herald) WipeStorage() error {
	herald.Lock()
	defer herald.Unlock()
	if err := herald.store.Wipe(); err != nil {
		return err
	}
	return nil
}

// EditConfig will edit the config file with the provided data.
func (herald *Herald) EditConfig(userName, emailAddress string) error {
	herald.Lock()
	defer herald.Unlock()

	// TODO:
	// this is a bit hacky, I'd like to add some config methods
	// and validation of inputs etc. but for now:
	// update the in-memory config
	herald.config.User = &services.User{
		Name:  userName,
		Email: emailAddress,
	}

	// write the in-memory config back to disk
	return herald.config.Write()
}

// GetRuntimeInfo makes a pass of the run and sample stores before populating the Herald instance with data:
// - how many samples are in the storage
// - notes any samples with tags
// - loads all sample labels into a slice (for JS to access)
func (herald *Herald) GetRuntimeInfo() error {
	herald.Lock()
	defer herald.Unlock()

	// reset the runtime data
	herald.runCount = 0
	herald.sampleCount = 0
	herald.untaggedRecordCount = 0
	herald.taggedRecordCount = 0
	herald.announcementCount = 0

	// get the run and sample counts from the store
	baselineExpCount := herald.store.GetNumRuns()
	baselineSampleCount := herald.store.GetNumSamples()

	// restart the queue
	herald.announcementQueue.Init()

	// create run label holder
	herald.runLabels = make([]string, baselineExpCount)

	// range over the runs in storage
	runIterator := 0
	for label := range herald.store.GetRunLabels() {

		// get the full run from storage
		run, err := herald.store.GetRun(string(label))
		if err != nil {
			return err
		}

		// update the relevant counts
		if err := herald.updateCounts(run, true); err != nil {
			return err
		}

		// add the run label to the holder (for display in app)
		herald.runLabels[runIterator] = run.Metadata.GetLabel()

		// increment the iterator
		runIterator++
	}
	if (baselineExpCount != runIterator) || (baselineExpCount != herald.runCount) {
		return fmt.Errorf("run mistmatch between db and in-memory store: %d vs %d", baselineExpCount, runIterator)
	}

	// setup the sample details holder
	herald.sampleDetails = make([][]string, 3)
	for i := 0; i < 3; i++ {
		herald.sampleDetails[i] = make([]string, baselineSampleCount)
	}

	// range over the samples in storage
	sampleIterator := 0
	for label := range herald.store.GetSampleLabels() {

		// get the full sample from storage
		sample, err := herald.store.GetSample(string(label))
		if err != nil {
			return err
		}

		// update the relevant counts
		if err := herald.updateCounts(sample, true); err != nil {
			return err
		}

		// add the details to the holders (for display in app)
		herald.sampleDetails[0][sampleIterator] = sample.Metadata.GetLabel()
		herald.sampleDetails[1][sampleIterator] = sample.Metadata.Created.String()
		herald.sampleDetails[2][sampleIterator] = sample.GetParentRun()

		// increment the iterator
		sampleIterator++
	}
	if (baselineSampleCount != sampleIterator) || (baselineSampleCount != herald.sampleCount) {
		return fmt.Errorf("sample mistmatch between db and in-memory store: %d vs %d", baselineSampleCount, sampleIterator)
	}
	return nil
}

// AddRun creates an run record, updates the runtime info and adds the record to storage
// TODO: this might be bypassed later and instead get JS to encode the form to protobuf directly
func (herald *Herald) AddRun(runLabel, outDir, fast5Dir, fastqDir, comment string, tags []string, existingRun bool) error {
	herald.Lock()
	defer herald.Unlock()

	// create the run
	newRun := services.InitRun(runLabel, outDir, fast5Dir, fastqDir)

	// add any comment
	if len(comment) != 0 {
		if err := newRun.Metadata.AddComment(comment); err != nil {
			return err
		}
	}

	// tag the run and update its status
	if len(tags) != 0 {
		if err := newRun.Metadata.AddTags(tags); err != nil {
			return err
		}
	}

	// if it's an existing run, sequencing and basecalling have been done so mark these as completed tags
	//
	// NOTE: this does not take into account existing runs that have not been basecalled or that request
	// re-basecalling but this is not an option yet
	//
	if existingRun {

		// add a comment to the history
		if err := newRun.Metadata.AddComment("existing sample - fast5 and fastq found"); err != nil {
			return err
		}

		// add the tags for sequencing and basecalling and mark them complete
		if err := newRun.Metadata.AddTags([]string{"sequence", "basecall"}); err != nil {
			return err
		}
		if err := newRun.Metadata.SetTag("sequence", true); err != nil {
			return err
		}
		if err := newRun.Metadata.SetTag("basecall", true); err != nil {
			return err
		}
	}

	// check and update the tag status (all tags might now be set to complete)
	if err := newRun.Metadata.CheckStatus(); err != nil {
		return err
	}

	// add the run to the store
	if err := herald.store.AddRun(newRun); err != nil {
		return err
	}

	// update the runtime info (grow the label slice, update counts, add to announcement queue etc.)
	herald.runLabels = append(herald.runLabels, runLabel)
	return herald.updateCounts(newRun, true)
}

// CreateSample creates a sample record, updates the runtime info and adds the record to storage
// TODO: this might be bypassed later and instead get JS to encode the form to protobuf directly
func (herald *Herald) CreateSample(label string, runName string, barcode int32, comment string, tags []string) error {
	herald.Lock()
	defer herald.Unlock()

	// get the run from storage
	exp, err := herald.store.GetRun(runName)
	if err != nil {
		return err
	}

	// TODO: copy the tag history over from the run to the samples (sequence and basecall)?
	//tags = append(exp.Metadata.GetRequestOrder(), tags...)

	// create the sample
	sample := services.InitSample(label, exp.Metadata.GetLabel(), barcode)
	if len(comment) != 0 {
		if err := sample.Metadata.AddComment(comment); err != nil {
			return err
		}
	}

	// tag the sample and update its status
	if len(tags) != 0 {
		if err := sample.Metadata.AddTags(tags); err != nil {
			return err
		}
	}

	// add the sample to the store
	if err := herald.store.AddSample(sample); err != nil {
		return err
	}

	// update the runtime info (grow the label slice, update counts, add to announcement queue etc.)
	herald.sampleDetails[0] = append(herald.sampleDetails[0], label)
	herald.sampleDetails[1] = append(herald.sampleDetails[1], sample.Metadata.GetCreated().String())
	herald.sampleDetails[2] = append(herald.sampleDetails[2], runName)
	return herald.updateCounts(sample, true)
}

// DeleteSample removes a sample record from storage and updates the counts
func (herald *Herald) DeleteSample(sampleLabel string) error {
	herald.Lock()
	defer herald.Unlock()

	// get the sample from storage
	sample, err := herald.store.GetSample(sampleLabel)
	if err != nil {
		return err
	}

	// delete the sample from the store
	if err := herald.store.DeleteSample(sampleLabel); err != nil {
		return err
	}

	// update the counts etc.
	return herald.updateCounts(sample, false)
}

// updateCounts takes an element and a bool to indicate if it is being added (true) or removed (false)
// from the storage.
// it will check the provided element is either a run or sample
// it will then increment/decrement the appropriate counters.
// it will also add/remove it from the queue if needed.
func (herald *Herald) updateCounts(element interface{}, add bool) error {
	value := -1
	if add {
		value = 1
	}

	// check for run or sample
	var status string
	switch element.(type) {
	case *services.Run:
		status = element.(*services.Run).Metadata.GetStatus().String()
		herald.runCount += value
	case *services.Sample:
		status = element.(*services.Sample).Metadata.GetStatus().String()
		herald.sampleCount += value
	default:
		return fmt.Errorf("unsupported type provided to updateCounts")
	}

	// process the status
	switch status {

	case "untagged":
		herald.untaggedRecordCount += value
		return nil

	case "tagged":
		herald.taggedRecordCount += value

		// handle the queue
		if add {
			herald.announcementQueue.PushBack(element)
		} else {
			herald.announcementQueue.Remove(&list.Element{Value: element})
		}
		return nil

	case "announced":
		herald.announcementCount += value
		return nil

	default:
		return fmt.Errorf("unrecognised status: %v", status)
	}
}
