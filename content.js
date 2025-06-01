"use strict";

const BROOMSTICK_STYLE_TAG_ID = "broomstick-style";

function createOrReplaceBroomstickStyle(fontFamily) {
  let style = document.getElementById(BROOMSTICK_STYLE_TAG_ID);
  if (!style) {
    style = document.createElement("style");
    style.setAttribute("id", BROOMSTICK_STYLE_TAG_ID);
    style.setAttribute("type", "text/css");
    document.head.appendChild(style);
  }
  style.textContent = `* { font-family: "${fontFamily}"; }`;
}

function refreshFont() {
  chrome.storage.local.get("font", function (result) {
    if (result.font === "gfont") {
      // TODO: Load Google Fonts
      const link = document.createElement("link");
      link.href = "https://fonts.googleapis.com/css2?family=Roboto:wght@400&display=swap";
      link.rel = "stylesheet";
      document.head.appendChild(link);

      createOrReplaceBroomstickStyle("'Roboto', sans-serif");
    } else if (result.font) {
      createOrReplaceBroomstickStyle(result.font);
    } else {
      document.head.removeChild(document.getElementById(BROOMSTICK_STYLE_TAG_ID));
    }
  });
}

// Load the font when the extension is loaded
refreshFont();

// Listen for changes in the storage and update the font accordingly
chrome.storage.onChanged.addListener(function (changes, namespace) {
  for (const [key, { oldValue, newValue }] of Object.entries(changes)) {
    console.debug(
      `Storage key "${key}" in namespace "${namespace}" changed.`,
      `Old value was "${oldValue}", new value is "${newValue}".`
    );
  }

  if (namespace === "local") {
    refreshFont();
  } else {
    console.error(`Invalid namespace: ${namespace}`);
  }
});
