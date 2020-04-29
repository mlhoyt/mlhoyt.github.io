// Ref: https://www.thatsoftwaredude.com/content/6125/how-to-paginate-through-a-collection-in-javascript
// Ref: https://www.freecodecamp.org/news/how-to-detect-a-users-preferred-color-scheme-in-javascript-ec8ee514f1ef/
// Ref: https://www.tecmint.com/find-user-account-info-and-login-details-in-linux/

// Build content.js which contains an array of all the content
// ref: https://www.w3schools.com/js/js_classes.asp

function show() {
  app = document.getElementById("app")

  app.innerHTML = ""
  app.innerHTML += "<ul>"
  for (i = 0; i < content.length; i++) {
    app.innerHTML += content[i].ToHTMLListItem()
  }
  app.innerHTML += "</ul>"
}

function load() {
  show();
}

window.onload = load;
