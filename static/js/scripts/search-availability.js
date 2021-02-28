import { initialiseDateRangePicker } from '../helper/date-range-picker.js'
import { validateFormBeforeSubmit } from '../helper/validate-form.js'

// Initialise date picker inputs
initialiseDateRangePicker('date-range-input')

// Add submit event listener to form to run validation before submission
validateFormBeforeSubmit('validate-form')
