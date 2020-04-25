// Ref: https://www.thatsoftwaredude.com/content/6125/how-to-paginate-through-a-collection-in-javascript

function show() {
  app = document.getElementById("app")

  app.innerHTML = "";
  app.innerHTML += "<ul>";
  // for (r = 0; r < pageList.length; r++) {
  //   app.innerHTML += pageList[r] + "<br/>";
  // }
  app.innerHTML += "<li type=none>YYYY-MM-DD TITLE</li>";
  app.innerHTML += "<li type=none>YYYY-MM-DD TITLE</li>";
  app.innerHTML += "<li type=none>YYYY-MM-DD TITLE</li>";
  app.innerHTML += "</ul>";
}

function load() {
  show();
}

window.onload = load;
