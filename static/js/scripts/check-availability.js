import { initialiseDateRangePicker } from '../helper/date-range-picker.js'
import { validateFormBeforeSubmit } from '../helper/validate-form.js'

const DATE_RANGE_INPUTS_ID = 'reservation-dates'
const FORM_ID = 'validate-form'

// HTML for date range inputs form
const html = `
  <form id=${FORM_ID} action="" method="GET" novalidate>
    <div id=${DATE_RANGE_INPUTS_ID} class="row">
      <div class="col">
        <div class="form-group">
          <input
            type="text"
            id="start_date"
            class="form-control"
            name="start_date"
            autocomplete="off"
            placeholder="Starting date"
            required
          >
          <div class="invalid-feedback">
            Starting date is required.
          </div>
        </div>
      </div>
      <div class="col">
        <div class="form-group">
          <input
            type="text"
            id="ending_date"
            class="form-control"
            name="ending_date"
            autocomplete="off"
            placeholder="Ending date"
            required
          >
          <div class="invalid-feedback">
            Ending date is required.
          </div>
        </div>
      </div>
    </div>
  </form>
`

const checkAvailabilityBtn = document.getElementById('check-availability')

checkAvailabilityBtn.addEventListener('click', openFormModal)

function openFormModal() {
  Swal.fire({
    title: 'Pick a date range',
    didOpen: () => {
      initialiseDateRangePicker(DATE_RANGE_INPUTS_ID) // Initialise date picker after modal is opened
      validateFormBeforeSubmit(FORM_ID) // Listen to submit event to validate form before submission
    },
    confirmButtonText: 'Done',
    html,
  })
}