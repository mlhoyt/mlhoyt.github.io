class ContentEntry {
  constructor(entryID, title, tags=[]) {
    this._entryID = entryID // Epoch time of entry corresponding to content/<entryID>/index.html
    this._title = title
    this._tags = tags
  }

  TimeStamp() {
    let date = new Date(this._entryID * 1000)
    return [
      date.getFullYear(),
      (date.getMonth()+1).toLocaleString(undefined, {minimumIntegerDigits: 2}),
      date.getDate().toLocaleString(undefined, {minimumIntegerDigits: 2}),
    ].join("-")
  }

  ToHTMLListItem() {
    // TODO: How to show the list of tags?
    return [
      "<li type=none>",
      this.TimeStamp(),
      "<a href=\"content/"+this._entryID+"/entry.md\">"+this._title+"</a>",
      "</li>",
    ].join(" ")
  }
}
