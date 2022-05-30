const go = new Go();
WebAssembly.instantiateStreaming(fetch("zippy.wasm"), go.importObject)
  .then((result) => {
    go.run(result.instance);
    document.getElementById("zip").disabled = false;
  })
  .catch((err) => {
    document.getElementById(
      "status"
    ).innerHTML = `<article class="message mb-3  is-danger"><div class="message-header"><p>WASM not loaded</p></div></article>`;
  });
document.getElementById("files").addEventListener("change", function (e) {
  var files = e.target.files;
  var filenames = document.getElementById("filenames");
  if (filenames?.innerHTML) {
    filenames.innerHTML = "";
  }
  for (var i = 0; i < files.length; i++) {
    filenames.innerHTML += `<span class="file-name">${files[i].name}</span>`;
  }
});

document.getElementById("zip").addEventListener("click", function (e) {
  var files = document.getElementById("files").files;
  var count = files.length;
  if (count === 0) {
    alert("No files selected");
    return;
  }
  var filesArr = [];
  for (var i = 0; i < files.length; i++) {
    var file = files[i];
    var fileName = file.name;
    var reader = new FileReader();
    reader.onload = function (e) {
      var fileArray = new Uint8Array(e.target.result);
      var fileBase64 = "";
      for (var i = 0; i < fileArray.length; i++) {
        fileBase64 += String.fromCharCode(fileArray[i]);
      }
      filesArr.push({
        name: e.target.fileName,
        data: btoa(fileBase64),
      });
      count--;
      if (count === 0) {
        var archiveType = document.getElementById("archiveType").value;
        Zippy(JSON.stringify(filesArr), archiveType);
      }
    };
    reader.fileName = fileName;
    reader.readAsArrayBuffer(file);
  }
});
document.addEventListener("DOMContentLoaded", () => {
  const $navbarBurgers = Array.prototype.slice.call(
    document.querySelectorAll(".navbar-burger"),
    0
  );

  $navbarBurgers.forEach((el) => {
    el.addEventListener("click", () => {
      const target = el.dataset.target;
      const $target = document.getElementById(target);

      el.classList.toggle("is-active");
      $target.classList.toggle("is-active");
    });
  });
});
