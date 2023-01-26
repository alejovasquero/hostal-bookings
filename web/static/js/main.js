import { notifier } from "./toaster.js"

(function () {
    'use strict'

    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    var forms = document.querySelectorAll('.needs-validation')

    // Loop over them and prevent submission
    Array.prototype.slice.call(forms)
        .forEach(function (form) {
            form.addEventListener('submit', function (event) {
                if (!form.checkValidity()) {
                    event.preventDefault()
                    event.stopPropagation()
                }

                form.classList.add('was-validated')
            }, false)
        })
})()

const elem = document.getElementById('new-date-range');
const datepicker = new DateRangePicker(elem, {
    autohide: false,
    format: "dd/mm/yyyy"
});

const alertButton = document.getElementById("alert-button")
alertButton.addEventListener("click", function () {
    //toaster.toastAlert({})
    //notifier.successNotifier({text: "Success in the operation", title: "Success"})
    //notifier.errorNotifier({text: "Error in the operation", title: "ERROR"})

    let html = `
        <form class="needs-validation" method="POST" action="" novalidate  style="width: 95%;">
            <div id="date-range-modal" class="form-row">
                <div class="col">
                    <div class="row">
                        <div class="col">
                            <input required class="form-control" type="text" name="start" id="start" placeholder="Arrival date">
                        </div>
                        <div class="col">
                            <input required class="form-control" type="text" name="end" id="end" placeholder="Departure date">
                        </div>
                    </div>
                </div>
            </div>
        </form>
    `

    notifier.multipleInputNotifier({html: html, title: "Choose your dates"})
})


function notify(message, messageType) {
    window.notie.alert({
        type: messageType, // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
        text: message,
        stay: false, // optional, default = false
        time: 1, // optional, default = 3, minimum = 1,
        position: "bottom" // optional, default = 'top', enum: ['top', 'bottom']
    })
}