// Validate form before submission
// Provide styled responses with Bootstrap classes
export function validateFormBeforeSubmit(formId) {
  const form = document.getElementById(formId)
  form.addEventListener('submit', function (e) {
    // checkValidity method checks whether the element has any constraints and whether it satisfies them
    // if the element fails its constraints, the browser fires a cancelable "invalid" event and returns false
    if (!form.checkValidity()) {
      e.preventDefault()
      e.stopPropagation()
    }
    form.classList.add('was-validated')
  }, false)
}