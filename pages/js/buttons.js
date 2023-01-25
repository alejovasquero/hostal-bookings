import { notifier } from "./toaster.js"

export const buttons = function () {
    let createPopUpForButton = function (buttonName, config) {
        const alertButton = document.getElementById(buttonName)
        alertButton.addEventListener("click", function () {
            //toaster.toastAlert({})
            //notifier.successNotifier({text: "Success in the operation", title: "Success"})
            //notifier.errorNotifier({text: "Error in the operation", title: "ERROR"})

            let html = `
                <div class="container-fluid">
                    <form id="range-date-form" class="needs-validation" method="POST" action="" novalidate  style="width: 95%;">
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
                </div>
            `
            notifier.multipleInputNotifier({ html: html, title: "Choose your dates", callback: config.callback, didOpen: config.didOpen, preConfirm: config.preConfirm })
        })
    }

    return {
        createPopUpForButton: createPopUpForButton
    }
}()

window.buttons = buttons