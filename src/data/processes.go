// Package data adds some wrapper functions to the protobuf messages
package data

import (
	"fmt"
)

// ProcessRegister is used to register all the available processes
var	ProcessRegister map[string]*Process

// init the process definitions at runtime
func init() {

	// init the process checker
	ProcessRegister = make(map[string]*Process)

	// create the process definitions
	createProcessDefinition("sequence", nil, true, false)
	createProcessDefinition("basecall", nil, true, false)
	createProcessDefinition("rampart", nil, false, true)
	createProcessDefinition("pipelineA", []string{"sequence", "basecall"}, false, true)

}

// createProcessDefinition will init a process
func createProcessDefinition(pName string, pDependsOn []string, availToExp, availToSamples bool) {

	// check the process does not already exist
	if _, exists := ProcessRegister[pName]; exists {
		panic(fmt.Sprintf("process already exists: %v", pName))
	}

	// init the process
	newProcess := &Process{
		Complete: false,
		Name: pName,
		DependsOn: []*Process{},
		History:   []*Comment{},
		AvailableToExperiments: availToExp,
		AvailableToSamples: availToSamples,
	}

	// check the dependencies
	if len(pDependsOn) != 0 {
		for _, depName := range pDependsOn {

			// can't depend on itself
			if depName == pName {
				panic("process can't depend on itself")
			}

			// dependency must already be registered
			dependency, ok := ProcessRegister[depName]
			if !ok {
				panic(fmt.Sprintf("process dependency not registered: %v", depName))
			}

			// copy the dependency to the process
			newProcess.DependsOn = append(newProcess.DependsOn, dependency.copyProcess())
		}
	}

	// register the process
	ProcessRegister[pName] = newProcess
	return
}

// copyProcess is a helper function to create a new instance of a process
func (process *Process) copyProcess() *Process {
	return &Process{
		Complete: false,
		Name: process.Name,
		DependsOn: process.DependsOn,
		History:   process.History,
	}
}