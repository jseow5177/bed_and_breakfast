const DatePicker = (function () {
  // Initialise date picker for date range inputs
  function initialiseDateRangePicker(dateRangeInputsId, format = 'yyyy-mm-dd') {
    const dateRangeInput = document.getElementById(dateRangeInputsId)

    if (!dateRangeInput) {
      throw new Error('Please check if you have defined date range inputs in your HTML')
    }

    new DateRangePicker(dateRangeInput, { format })
  }

  return {
    initialiseDateRangePicker
  }
})()