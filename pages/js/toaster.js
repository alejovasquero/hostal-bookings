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

        const { value: result } = await Swal.fire({
            title: title,
            html: html,
            focusConfirm: true,
            showCancelButton: true,
            width: "50em",
            didOpen: () => {
                console.log(config)
                if (config.didOpen !== undefined) {
                    config.didOpen()
                }
            },
            preConfirm: () => {
                return config.preConfirm()
            },

        })

        console.log(config)
        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel) {
                if (config.callback !== undefined) {
                    config.callback(result)
                }
            } else {
                config.callback(false)
            }
        }
    }

    return {
        toastAlert: toastAlert,
        successNotifier: successNotifier,
        errorNotifier: errorNotifier,
        multipleInputNotifier: multipleInputNotifier
    }
}()