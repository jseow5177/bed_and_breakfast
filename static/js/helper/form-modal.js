const FormModal = (function () {
  async function openFormModal({
    title,
    html,
    onOpen = undefined,
    onSubmit = undefined,
    callback = undefined,
    confirmButtonText = 'Ok',
    cancelButtonText = 'Cancel',
    showCancelButton = true,
    showCloseButton = false,
    allowOutsideClick = true,
  }) {
    const result = await Swal.fire({
      title,
      html,
      didOpen: () => {
        if (onOpen !== undefined) {
          onOpen()
        }
      },
      preConfirm: () => {
        if (onSubmit !== undefined) {
          onSubmit()
        }
      },
      confirmButtonText,
      cancelButtonText,
      showCancelButton,
      showCloseButton,
      allowOutsideClick,
    })

    if (result && callback !== undefined) {
      callback(result)
    }
  }

  return {
    openFormModal
  }
})()