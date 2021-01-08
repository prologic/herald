////////////////////////////////////////////////////////////////////
// HELPERS
// removeOptions will clear a select box
function removeOptions(selectbox) {
    var i
    for (i = selectbox.options.length - 1; i >= 0; i--) {
        selectbox.remove(i)
    }
}

// handleForm will prevent default form action
function handleForm(event) {
    event.preventDefault()
}

////////////////////////////////////////////////////////////////////
// MESSAGES
// printErrorMsg
function printErrorMsg(msg) {
    console.log('error: ', msg)
    $('body').overhang({
        type: 'error',
        message: msg,
        overlay: true,
        closeConfirm: true
    })
}

// printSuccessMsg
function printSuccessMsg(msg) {
    console.log(msg)
    $('body').overhang({
        custom: true,
        primary: '#35cebe',
        accent: '#25beae',
        message: msg
    })
}

// printUpdateMsg
function printUpdateMsg(msg, link) {
    $('body').overhang({
        type: 'confirm',
        custom: true,
        primary: '#35cebe',
        accent: '#25beae',
        html: true,
        overlay: true,
        closeConfirm: true,
        message: msg,
        callback: function(value) {
            if (value) {
                window.location.href = link
            }
        }
    })
}

// updatStatus will print status icon and check time
function updateStatus(id, active) {
    var today = new Date()
    var time =
        today.getHours() +
        ':' +
        (today.getMinutes() < 10 ? '0' : '') +
        today.getMinutes()
    if (active === 'true') {
        document.getElementById(id).innerHTML =
            '<i class="far fa-check-circle" style="color: #35cebe;"></i>'
        document.getElementById(id + '_timestamp').innerText =
            'online since: ' + time
    } else {
        document.getElementById(id).innerHTML =
            '<i class="far fa-times-circle" style="color: red;"></i>'
        document.getElementById(id + '_timestamp').innerText =
            'offline since: ' + time
    }
}

// listen for change in network status
window.addEventListener('online', () => updateStatus('status_network', 'true'))
window.addEventListener('offline', () =>
    updateStatus('status_network', 'false')
)

////////////////////////////////////////////////////////////////////
// BUTTONS
// get the buttons that control the app
const refreshPage = document.getElementById('refreshPage')
const announceSamplesButton = document.getElementById('staging_announce')
const wipeDatabase = document.getElementById('wipeDatabase')

// add an event listener to the refreshPage button
refreshPage.addEventListener('click', async() => {
    console.log('refreshing the app')

    pageRefresh()
    printSuccessMsg('refreshed the app')
})

// add an event listener to the staging_announce button
announceSamplesButton.addEventListener('click', async() => {
    console.log('announcing samples')

    // call the Go announceSamples method
    try {
        await announceSamples()
    } catch (e) {
        printErrorMsg(e)
        return
    }
    console.log('announced')
})

// add an event listener to wipeDatabase button
wipeDatabase.addEventListener('click', async() => {
    console.log('wiping storage')

    // TODO: add a confirm prompt

    // call the Go wipeStorage method
    try {
        await wipeStorage()
    } catch (e) {
        printErrorMsg(e)
        return
    }

    // reset the page and report success
    fullPageRender()
    printSuccessMsg('database wiped')
})

////////////////////////////////////////////////////////////////////
// MODALS
// modal closing
var modalClosers = document.getElementsByClassName('modal-close')
for (let i = 0; i < modalClosers.length; i++) {
    modalClosers[i].addEventListener('click', function() {
        modalClosers[i].closest('.modal').style.display = 'none'
    })
}
window.onclick = function(event) {
    if (event.target.className == 'modal') {
        event.target.style.display = 'none'
    }
}

// addRunModal
// open on button click
const addRunModal = document.getElementById('addRunModal')
const addRunModalOpen = document.getElementById('addRunModalOpen')
addRunModalOpen.addEventListener('click', function() {
    addRunModal.style.display = 'block'
})

// addSampleModal
// open on button click (Go handles the disabled flag)
const addSampleModal = document.getElementById('addSampleModal')
const addSampleModalOpen = document.getElementById('addSampleModalOpen')
addSampleModalOpen.addEventListener('click', function() {
    addSampleModal.style.display = 'block'
})

// sampleDetailsModal
const sampleDetailsModal = document.getElementById('sampleDetailsModal')

////////////////////////////////////////////////////////////////////
// CREATE EXPERIMENT FORM
// get the form and prevent default action
const addRunForm = document.getElementById('addRunForm')
addRunForm.addEventListener('submit', handleForm)

// historic run toggle (set true if an run is entered that already has fast5 data)
var historicExp = false

// get some fields
var expName = document.getElementById('formLabel_runName')
var expOutputLocation = document.getElementById('formLabel_outputLocation')
var fieldset_outputFASTQlocation = document.getElementById(
    'fieldset_outputFASTQlocation'
)
var formLabel_outputFAST5location = document.getElementById(
    'formLabel_outputFAST5location'
)
var formLabel_outputFASTQlocation = document.getElementById(
    'formLabel_outputFASTQlocation'
)
var formLabel_sequence = document.getElementById('formLabel_sequence')
var formLabel_basecall = document.getElementById('formLabel_basecall')
var formLabel_basecallLabel = document.getElementById('formLabel_basecallLabel')
var msgDiv = document.getElementById('addRunValidationMessage')

// reset func to clear the form changes
function addRunFormReset() {
    historicExp = false
    fieldset_outputFASTQlocation.style.display = 'block'
    formLabel_outputFAST5location.value = ''
    formLabel_outputFASTQlocation.value = ''
    formLabel_sequence.checked = true
    formLabel_basecall.checked = true
    formLabel_basecallLabel.style.color = '#d3d3d3'
    formLabel_basecall.disabled = true
    msgDiv.innerHTML = ''
}

// set up the validator
runValidator = {
    validListenter: function(val) {},
    registerListener: function(listener) {
        this.validListenter = listener
    },
    expNameInternal: false,
    set expName(val) {
        this.expNameInternal = val
        this.validListenter(val)
    },
    get expName() {
        return this.expNameInternal
    },
    expLocInternal: false,
    set expLoc(val) {
        this.expLocInternal = val
        this.validListenter(val)
    },
    get expLoc() {
        return this.expLocInternal
    }
}

// the validator listener will adjust form values depending on user input
runValidator.registerListener(async() => {
    // reset if user hasn't input both expName and expOutputLocation
    if (runValidator.expName === false || runValidator.expLoc === false) {
        addRunFormReset()
        return
    }

    // remove spaces from expName
    var expNameDespaced = expName.value.replace(/\s/g, '_')

    // get expected dir names
    var dirName = expOutputLocation.value + '/' + expNameDespaced
    var fast5_dirName = dirName + '/fast5_pass'
    var fastq_dirName = dirName + '/fastq_pass'

    // update the paths
    formLabel_outputFAST5location.value = fast5_dirName
    formLabel_outputFASTQlocation.value = fastq_dirName

    // allow user to change basecalling option
    formLabel_basecall.disabled = false
    formLabel_basecallLabel.style.color = ''

    // check for dirs and print an alert
    try {
        await checkDirExists(dirName)
    } catch (e) {
        msgDiv.innerHTML =
            '<div class="alert background-warning"><i class="fas fa-exclamation-circle"></i> - No existing run directory found, this run will be tagged for sequencing</div>'
        return
    }
    try {
        await checkDirExists(fast5_dirName)
    } catch (e) {
        msgDiv.innerHTML =
            '<div class="alert background-warning"><i class="fas fa-exclamation-circle"></i> - No <em>fast5_pass</em> directory found, this run will be tagged for sequencing</div>'
        return
    }
    formLabel_sequence.checked = true
    historicExp = true

    try {
        await checkDirExists(fastq_dirName)
    } catch (e) {
        msgDiv.innerHTML =
            '<div class="alert background-warning"><i class="fas fa-exclamation-circle"></i> - No <em>fastq_pass</em> directory found, this run will be tagged for base calling (unless you uncheck the box below)</div>'
        return
    }
    msgDiv.innerHTML =
        '<div class="alert background-success"><i class="fas fa-flask"></i> - <em>fast5_pass</em> and <em>fastq_pass</em> found for this run</div>'

    // make sure the fastq path is shown (could be hidden if user has been toggling)
    fieldset_outputFASTQlocation.style.display = 'block'

    // disable basecalling option if fastq_pass exists
    formLabel_basecallLabel.style.color = '#d3d3d3'
    formLabel_basecall.checked = true
    formLabel_basecall.disabled = true
})

// add listener to the runName text box
expName.addEventListener('change', async() => {
    runValidator.expName = false
    if (expName.value.length === 0) {
        return
    }

    // currently Go checks the dir - could do it here instead though
    runValidator.expName = true
})

// add listener to the output location text box so we can check it exists once a user has entered a location
expOutputLocation.addEventListener('change', async() => {
    runValidator.expLoc = false
    if (expOutputLocation.value.length === 0) {
        return
    }
    try {
        await checkDirExists(expOutputLocation.value)
    } catch (e) {
        printErrorMsg(e)
        return
    }
    runValidator.expLoc = true
})

// show/hide the fastq_pass path depending on basecall checkbox
formLabel_basecall.addEventListener('click', async() => {
    if (formLabel_basecall.checked === true) {
        fieldset_outputFASTQlocation.style.display = 'block'
    } else {
        fieldset_outputFASTQlocation.style.display = 'none'
    }
})

// add an event listener to the addRunForm submit button
addRunForm.addEventListener('submit', async() => {
    console.log('creating run')

    // create sequence and basecall tags
    var tags = []
    if (formLabel_sequence.checked === true) {
        tags.push('sequence')
    }
    if (formLabel_basecall.checked === true) {
        tags.push('basecall')
    }

    // create an run and add it to the store
    try {
        await addRun(
            expName.value,
            expOutputLocation.value,
            formLabel_outputFAST5location.value,
            formLabel_outputFASTQlocation.value,
            document.getElementById('formLabel_runComment').value,
            tags,
            historicExp
        )
    } catch (e) {
        printErrorMsg(e)
        return
    }

    // reset the form, refresh the page, close the modal and report success
    addRunForm.reset()
    addRunFormReset()
    pageRefresh()
    addRunModal.style.display = 'none'
    printSuccessMsg('run created')
})

////////////////////////////////////////////////////////////////////
// SAMPLE SUBMISSION FORM
// get the form
const addSampleForm = document.getElementById('addSampleForm')
addSampleForm.addEventListener('submit', handleForm)

// updateRunDropDown is a function to update the runs in the sample submission form
const expDropDown = document.getElementById('formLabel_sampleRun')
const updateRunDropDown = async() => {
    // get the current run count so that we can iterate over the runs
    var expCount = `${await window.getRunCount()}`

    // if there are no runs, leave the default blank option
    if (expCount === '0') {
        return
    }

    // wipe the current options
    removeOptions(expDropDown)

    // add each name to the drop down
    for (var i = 0; i < expCount; i++) {
        var expName = `${await window.getRunName(i)}`
        var newOpt = document.createElement('option')
        newOpt.text = expName
        expDropDown.options.add(newOpt)
    }
}

// add an event listener to the addSampleForm submit button
addSampleForm.addEventListener('submit', async() => {
    console.log('adding sample to storage')

    var elements = addSampleForm.elements

    // grab the tags from the form
    var tags = []
    for (var i = 0, element;
        (element = elements[i++]);) {
        if (element.type === 'checkbox' && element.checked) {
            tags.push(element.value)
        }
    }

    // create a sample and add it to the storage
    try {
        // TODO: try reading form straight into protobuf and then send a serialised stream to Go
        await createSample(
            elements['formLabel_sampleLabel'].value,
            elements['formLabel_sampleRun'].value,
            parseInt(elements['formLabel_sampleBarcode'].value, 10),
            elements['formLabel_sampleComment'].value,
            tags
        )
    } catch (e) {
        printErrorMsg(e)
        return
    }

    // update the table
    $('#sampleTable')
        .DataTable()
        .row.add([
            elements['formLabel_sampleLabel'].value,
            elements['formLabel_sampleRun'].value
        ])
        .draw(true)

    // reset the form, refresh the page and report success
    addSampleForm.reset()
    pageRefresh()
    addSampleModal.style.display = 'none'
    printSuccessMsg('sample added')
})

////////////////////////////////////////////////////////////////////
// TABLES
// set up the table
var table = $('#sampleTable').DataTable({
    columnDefs: [{
        targets: 2,
        data: null,
        searchable: true,
        orderable: true,
        defaultContent: '<button class="button button-outline">Manage</button>'
    }]
})

// getSampleJSONdump returns a stringified protobuf dump of a sample from the storage
const getSampleJSONdump = async function(sampleLabel) {
    var sampleJSONdump = `${await window.printSampleToJSONstring(sampleLabel)}`
    return sampleJSONdump
}

// set up the manage button
$('#sampleTable tbody').on('click', 'button', function() {
    var row = table.row($(this).parents('tr'))
    var data = row.data()
    var sampleLabel = data[0]
    document.getElementById('sampleModal_samplename').innerHTML = sampleLabel
    getSampleJSONdump(sampleLabel).then(sampleProtobufDump => {
        // get the sample protobuf dump
        document.getElementById('sampleModal_content').innerHTML =
            '<pre>' + sampleProtobufDump + '</pre>'

        // display modal
        document.getElementById('sampleDetailsModal').style.display = 'block'

        // set up delete button
        document
            .getElementById('sampleModal_delete')
            .addEventListener('click', async() => {
                console.log('deleting sample')

                // delete from the db
                try {
                    await deleteSample(sampleLabel)
                } catch (e) {
                    printErrorMsg(e)
                    return
                }

                // remove from the table
                row.remove().draw(true)

                // reset the runtime info and report success
                await pageRefresh()
                document.getElementById('sampleDetailsModal').style.display = 'none'
                printSuccessMsg('sample deleted')
            })
    })
})

// buildTable will get the database keys via Go and then populate the table
const buildTable = async() => {
    console.log('building table from sample labels in storage')

    // wipe any existing table
    table.clear().draw(true)

    // get the current sample number so that we can iterate over the samples
    var sampleCount = `${await window.getSampleCount()}`

    // process each sample label
    for (var i = 0; i < sampleCount; i++) {
        var sampleLabel = `${await window.getSampleLabel(i)}`
            //var sampleCreation = `${await window.getSampleCreation(i)}`
        var sampleRun = `${await window.getSampleRun(i)}`

        // create the table entry
        table.row.add([sampleLabel, sampleRun]).draw(true)
    }
}

////////////////////////////////////////////////////////////////////
// PAGE RENDERING
// setup the time stamps
const dateStamp = document.getElementById('dateStamp')
const refreshStamp = document.getElementById('refreshStamp')

// printTimeStamps will add the date and the refresh time to the app
function printTimeStamps() {
    var today = new Date()
    var options = {
        year: 'numeric',
        month: 'long',
        day: 'numeric'
    }
    dateStamp.innerHTML = today.toLocaleDateString('en-US', options)
    refreshStamp.innerHTML =
        'last refreshed: ' +
        today.getHours() +
        ':' +
        (today.getMinutes() < 10 ? '0' : '') +
        today.getMinutes()
}

// set up the empty pie chart
var pieCanvas = document.getElementById('pieChart')
var pieData = {
    labels: ['Announcements Made', 'Tagged Samples', 'Untagged Samples'],
    datasets: [{
        label: 'entry point',
        data: [0, 0, 0],
        backgroundColor: ['#35cebe', '#a0a0a0', '#dfdfdf'],
        hoverBackgroundColor: ['#25beae', '#999999', '#cccccc']
    }]
}
var pieOptions = {
    responsive: true,
    segmentShowStroke: false,
    legend: false
}
var myPieChart = new Chart(pieCanvas, {
    type: 'pie',
    data: pieData,
    options: pieOptions
})

// updatePieChart will refresh the pie chart with current data
const updatePieChart = async() => {
    // get counts
    var untaggedRecordCount = `${await window.getUntaggedSampleCount()}`
    var taggedRecordCount = `${await window.getTaggedSampleCount()}`
    var announcementCount = `${await window.getAnnouncementCount()}`

    // update the chart data
    myPieChart.data.datasets[0].data[0] = announcementCount
    myPieChart.data.datasets[0].data[1] = taggedRecordCount
    myPieChart.data.datasets[0].data[2] = untaggedRecordCount

    // update the chart
    myPieChart.update()
}

// pageRefresh will refresh the Herald runtime info in Go and then freshen up the page (does not rebuild the table)
const pageRefresh = async() => {
    console.log('refreshing runtime info and re-rendering the page')

    // reload the Go Herald instance and repopulate the page data
    try {
        await loadRuntimeInfo()
    } catch (e) {
        printErrorMsg(e)
        return
    }

    // update the run drop down
    await updateRunDropDown()

    // update the pie chart
    await updatePieChart()

    // check the minKNOW status TODO: do this in Go via routine
    var minknowStatus = `${await checkAPIstatus()}`
    updateStatus('status_minknow', minknowStatus)

    // print a new timestamp
    printTimeStamps()
}

// fullPageRender will insert various bits of runtime info from JS and Go into the app
// TODO: this is virtually the same as pageRefresh - so just combine and add flag?
const fullPageRender = async() => {
    console.log('starting Go Herald instance and rendering the page')

    // load the Go Herald instance and populate the page data
    try {
        await loadRuntimeInfo()
    } catch (e) {
        printErrorMsg(e)
        return
    }

    // update the run drop down
    await updateRunDropDown()

    // print the pie chart
    await updatePieChart()

    // check the minKNOW status TODO: do this in Go via routine
    var minknowStatus = `${await checkAPIstatus()}`
    updateStatus('status_minknow', minknowStatus)

    // print the table
    await buildTable()

    // print a new timestamp
    printTimeStamps()
}