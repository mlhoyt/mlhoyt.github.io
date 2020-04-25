// Ref: https://www.thatsoftwaredude.com/content/6125/how-to-paginate-through-a-collection-in-javascript
// Ref: https://www.freecodecamp.org/news/how-to-detect-a-users-preferred-color-scheme-in-javascript-ec8ee514f1ef/
// Ref: https://www.tecmint.com/find-user-account-info-and-login-details-in-linux/

// Build content.js which contains an array of all the content
// ref: https://www.w3schools.com/js/js_classes.asp

class ContentEntry {
  constructor(entryID, title, tags=[]) {
    this._entryID = entryID // Epoch time of entry corresponding to content/<entryID>/index.html
    this._title = title
    this._tags = tags
  }
  // get <FIELD>() { return this._<FIELD> }
  // set <FIELD>(<VALUE>) { this._<FIELD> = <VALUE> }
  TimeStamp() {
    let date = new Date(this._entryID * 1000)
    return [
      date.getFullYear(),
      (date.getMonth()+1).toLocaleString(undefined, {minimumIntegerDigits: 2}),
      date.getDate(),
    ].join("-")
  }
  ToHTMLListItem() {
    // TODO: How to show the list of tags?
    return [
      "<li type=none>",
      this.TimeStamp(),
      this._title,
      "</li>",
    ].join(" ")
  }
}

content = [
  new ContentEntry(1587425974, "TITLE1"),
  new ContentEntry(1587625975, "TITLE2"),
  new ContentEntry(1587825976, "TITLE3"),
]
// let tags = {
//   Keys: function() {
//     return [
//       <TAG>,
//       ...
//     ]
//   },
//   <TAG>: [
//     <ENTRY-ID>,
//     ...
//   ],
// }

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
