const Notify = (function () {
  function alert({ type, msg }) {
    notie.alert({ type, text: msg })
  }

  return {
    alert
  }
})()