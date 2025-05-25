"use strict";

const formRadioButtons = document.getElementById('font-form').querySelectorAll('input[type="radio"]');
const noneRadioButton = document.getElementById('none');

// This must be updated when adding support for new fonts
// TODO: refactor this messy code
const radioButtonValue = {
    "courier-new": "Courier New",
    "arial": "Arial",
    "times-new-roman": "Times New Roman",
    "comic-sans-ms": "Comic Sans MS",
    "gfont": "gfont", // TODO REMOVE
}
const radioButtonValueReverse = {
    "Courier New": "courier-new",
    "Arial": "arial",
    "Times New Roman": "times-new-roman",
    "Comic Sans MS": "comic-sans-ms",
    "gfont": "gfont", // TODO REMOVE
}

// Load the current font setting from storage and set the radio button accordingly
chrome.storage.local.get("font", function (result) {
    console.debug("Current font setting:", result.font);

    if (!result.font) {
        noneRadioButton.checked = true;
        return;
    }

    formRadioButtons.forEach(radio => {
        radio.checked = radio.id === radioButtonValueReverse[result.font];
    });
})

// Add event listeners to each radio button to save the selected font to storage
formRadioButtons.forEach((radio) => {
    radio.addEventListener('change', function (event) {
        const selectedRadioButton = event.target;
        if (selectedRadioButton.id in radioButtonValue) {
            chrome.storage.local.set({ "font": radioButtonValue[selectedRadioButton.id] });
        } else {
            chrome.storage.local.remove('font');
        }
    });
});
