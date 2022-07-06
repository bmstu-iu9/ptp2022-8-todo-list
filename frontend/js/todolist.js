'use strict'

window.onload = function () {
    /* bootlint.showLintReportForCurrentDocument([], {
        hasProblems: false,
        problemFree: false,
    }) */
    const inpts = document.getElementsByClassName('edit-todo-input')
    for (var i = 0; i < inpts.length; i++) {
        inpts[i].addEventListener('click', () => {
            inpts[0].readOnly = false
        })
        inpts[i].addEventListener('mouseout', () => {
            inpts[0].classList.add('bg-transparent')
        })
    }
}
