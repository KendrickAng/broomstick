"use strict";

function refreshFont() {
  chrome.storage.local.get("font", function (result) {
    if (result.font === "gfont") {
      // Load Google Fonts
      const link = document.createElement("link");
      link.href = "https://fonts.googleapis.com/css2?family=Roboto:wght@400&display=swap";
      link.rel = "stylesheet";
      document.head.appendChild(link);
      document.body.style.fontFamily = "'Roboto', sans-serif";
    } else if (result.font) {
      document.body.style.fontFamily = result.font;
    } else {
      document.body.style.removeProperty('font-family');
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
