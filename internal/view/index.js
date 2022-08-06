function showWeightCreateBox() {
  Swal.fire({
    title: "Create Weight",
    html:
      '<input id="id" type="hidden">' +
      '<input id="maxweight" type="number" min="0" class="swal2-input" placeholder="Max" required>' +
      '<input id="minweight" type="number" min="0" class="swal2-input" placeholder="Min" required>' +
      '<input id="createdat" type="date" class="swal2-input" placeholder="Date" required>',
    focusConfirm: false,
    preConfirm: () => {
      weightCreate();
    },
  });
}

function weightCreate() {
  const maxWeight = document.getElementById("maxweight").value;
  const minWeight = document.getElementById("minweight").value;
  const createdAt = document.getElementById("createdat").value;
  if (parseInt(maxWeight) < parseInt(minWeight)) {
    Swal.fire("Max lebih kecil dari min");
    return false;
  }
  const xhttp = new XMLHttpRequest();
  xhttp.open("POST", "http://localhost:8000/api/create");
  xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
  xhttp.send(
    JSON.stringify({
      max_weight: parseInt(maxWeight),
      min_weight: parseInt(minWeight),
      created_at: createdAt,
    })
  );
  xhttp.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      const objects = JSON.parse(this.responseText);
      Swal.fire(objects["message"]);
      loadTable();
    }
  };
}

function showWeightEditBox(id) {
  console.log(id);
  const xhttp = new XMLHttpRequest();
  xhttp.open("GET", "http://localhost:8000/api/weight/" + id);
  xhttp.send();
  xhttp.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      const objects = JSON.parse(this.responseText);
      Swal.fire({
        title: "Edit Weight",
        html:
          '<input id="id" type="hidden" value=' +
          objects["id"] +
          ">" +
          '<input id="maxweight" type="number" class="swal2-input" placeholder="Max" value="' +
          objects["max_weight"] +
          '">' +
          '<input id="minweight" type="number" class="swal2-input" placeholder="Min" value="' +
          objects["min_weight"] +
          '">' +
          '<input id="createdat" type="date" class="swal2-input" placeholder="Date" value="' +
          objects["created_at"].slice(0, 10) +
          '">',
        focusConfirm: false,
        preConfirm: () => {
          weightEdit();
        },
      });
    }
  };
}

function weightEdit() {
  const id = document.getElementById("id").value;
  const maxWeight = document.getElementById("maxweight").value;
  const minWeight = document.getElementById("minweight").value;
  const createdAt = document.getElementById("createdat").value;

  const xhttp = new XMLHttpRequest();
  xhttp.open("PUT", "http://localhost:8000/api/weight/" + id);
  xhttp.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
  xhttp.send(
    JSON.stringify({
      max_weight: parseInt(maxWeight),
      min_weight: parseInt(minWeight),
      created_at: createdAt,
    })
  );
  xhttp.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      const objects = JSON.parse(this.responseText);
      Swal.fire(objects["message"]);
      loadTable();
    }
  };
}

function weightDelete(id) {
  const xhttp = new XMLHttpRequest();
  xhttp.open("DELETE", "http://localhost:8000/api/weight/" + id);
  xhttp.send();
  xhttp.onreadystatechange = function () {
    if (this.readyState == 4) {
      const objects = JSON.parse(this.responseText);
      Swal.fire(objects["message"]);
      loadTable();
    }
  };
}

function showWeight(id) {
  console.log(id);
  const xhttp = new XMLHttpRequest();
  xhttp.open("GET", "http://localhost:8000/api/weight/" + id);
  xhttp.send();
  xhttp.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      var trHTML = "";
      const object = JSON.parse(this.responseText);
      trHTML += "<tr>";
      trHTML += "<th>Tanggal</th>";
      trHTML += "<td>" + object["created_at"].slice(0, 10) + "</td>";
      trHTML += "</tr>";
      trHTML += "<tr>";
      trHTML += "<th>Max</th>";
      trHTML += "<td>" + object["max_weight"] + "</td>";
      trHTML += "</tr>";
      trHTML += "<tr>";
      trHTML += "<th>Min</th>";
      trHTML += "<td>" + object["min_weight"] + "</td>";
      trHTML += "</tr>";
      trHTML += "<tr>";
      trHTML += "<th>Perbedaan</th>";
      trHTML += "<td>" + object["difference"] + "</td>";
      trHTML += "</tr>";
      document.getElementById("mytable2").innerHTML = trHTML;
    }
  };
}

function loadTable() {
  const xhttp = new XMLHttpRequest();
  xhttp.open("GET", "http://localhost:8000/api/weights");
  xhttp.send();
  xhttp.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      console.log(this.responseText);
      var trHTML = "";
      let sumMax = 0,
        sumMin = 0,
        sumDif = 0;
      const objects = JSON.parse(this.responseText);
      for (let object of objects) {
        trHTML += "<tr>";
        trHTML +=
          '<td><buttontype="button" class="btn btn-primary" data-bs-toggle="modal" data-bs-target="#exampleModal" onclick="showWeight(' +
          object["id"] +
          ')">' +
          object["created_at"].slice(0, 10) +
          "</button></td>";
        trHTML += "<td>" + object["max_weight"] + "</td>";
        trHTML += "<td>" + object["min_weight"] + "</td>";
        trHTML += "<td>" + object["difference"] + "</td>";
        trHTML +=
          '<td><button type="button" class="btn btn-outline-secondary" onclick="showWeightEditBox(' +
          object["id"] +
          ')">Edit</button>';
        trHTML +=
          '<button type="button" class="btn btn-outline-danger" onclick="weightDelete(' +
          object["id"] +
          ')">Del</button></td>';
        trHTML += "</tr>";
        sumMax += object["max_weight"];
        sumMin += object["min_weight"];
        sumDif += object["difference"];
      }
      trHTML += "<tr>";
      trHTML += "<th>Rata-rata</th>";
      trHTML += "<td>" + (sumMax / objects.length).toFixed(2) + "</td>";
      trHTML += "<td>" + (sumMin / objects.length).toFixed(2) + "</td>";
      trHTML += "<td>" + (sumDif / objects.length).toFixed(2) + "</td>";

      document.getElementById("mytable").innerHTML = trHTML;
    }
  };
}

loadTable();
