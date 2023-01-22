export const notifier = function () {
    let toastAlert = function (config) {
        const {
            text = "Default text",
            icon = "success",
            position = "top-end",
        } = config;

        const Toast = Swal.mixin({
            toast: true,
            position: position,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({
            icon: icon,
            title: text
        })
    }

    let successNotifier = function (config) {
        const {
            text = "Default text",
            title = "Default title"
        } = config;

        Swal.fire({
            icon: "success",
            title: title,
            text: text,
        })
    }

    let errorNotifier = function (config) {
        const {
            text = "Default text",
            title = "Default title"
        } = config;

        Swal.fire({
            icon: "error",
            title: title,
            text: text,
        })
    }


    let multipleInputNotifier = async function (config) {
        const {
            html = "",
            title = "Date selection"
        } = config

        const { value: formValues } = await Swal.fire({
            title: title,
            html: html,
            focusConfirm: false,
            showCancelButton: true,
            width: "50em",
            didOpen: () => {
                const elem = document.getElementById('date-range-modal');
                const datepicker = new DateRangePicker(elem, {
                    autohide: false,
                    format: "dd/mm/yyyy",
                    orientation: "top"
                });
            },
            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            }
        })

        if (formValues) {
            Swal.fire(JSON.stringify(formValues))
        }
    }

    return {
        toastAlert: toastAlert,
        successNotifier: successNotifier,
        errorNotifier: errorNotifier,
        multipleInputNotifier: multipleInputNotifier
    }
}()