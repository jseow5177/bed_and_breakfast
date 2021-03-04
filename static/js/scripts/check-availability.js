const DATE_RANGE_INPUTS_ID = 'reservation-dates'
const FORM_ID = 'validate-form'

// HTML for date range inputs form
const html = `
  <form id=${FORM_ID} novalidate>
    <div id=${DATE_RANGE_INPUTS_ID} class="row">
      <div class="col">
        <div class="form-group">
          <input
            type="text"
            id="start_date"
            class="form-control"
            name="start_date"
            autocomplete="off"
            placeholder="Start date"
            required
          >
          <div class="invalid-feedback">
            Start date is required.
          </div>
        </div>
      </div>
      <div class="col">
        <div class="form-group">
          <input
            type="text"
            id="end_date"
            class="form-control"
            name="end_date"
            autocomplete="off"
            placeholder="End date"
            required
          >
          <div class="invalid-feedback">
            End date is required.
          </div>
        </div>
      </div>
    </div>
  </form>
`

const checkAvailabilityBtn = document.getElementById('check-availability')
checkAvailabilityBtn.addEventListener('click', function () {
  FormModal.openFormModal({
    title: 'Pick a date range',
    html,
    onOpen: () => {
      DatePicker.initialiseDateRangePicker(DATE_RANGE_INPUTS_ID) // Initialise date picker after modal is opened
      FormValidation.validateFormBeforeSubmit(FORM_ID) // Listen to submit event to validate form before submission
    },
    onSubmit: () => {
      const form = document.getElementById(FORM_ID)
      const formData = new FormData(form)
      formData.append('csrf_token', CSRFToken)

      fetch('/json-test', {
        method: 'POST',
        body: formData
      })
        .then(res => res.json())
        .then(data => {
          console.log(data)
        })
    },
    confirmButtonText: 'Submit'
  })
})
