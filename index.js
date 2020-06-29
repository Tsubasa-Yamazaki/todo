function entry() {
  let impElement = document.getElementById("post");
  let impRadio = impElement.imp;
  let importance = impRadio;
  let task = document.getElementById("task");
  let deadline = document.getElementById("deadline");

  let todoData = {
    importance: importance.value,
    task: task.value,
    deadline: deadline.value,
  };

  fetch("/post", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(todoData),
  }).then((response) => {
    display();
  });
  alert("登録しました");
}

function display() {
  fetch("/todoList", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  })
    .then((response) => {
      return response.json();
    })
    .then((todoData) => {
      tableDelete();
      displayTable(todoData);
    });
}

function tableDelete() {
  let table = document.getElementById(`table`);
  let loopCount = table.rows.length;

  for (let i = 0; i < loopCount; i++) {
    if (i === 0) {
      continue;
    }
    table.deleteRow(1);
  }
}

function displayTable(todoData) {
  todoData.forEach(function (element) {
    let table = document.getElementById(`table`);
    let id = element.id;
    let imp = element.importance;
    let task = element.task;
    let limit = element.deadline;

    let newRow = table.insertRow(-1);
    let cell1 = newRow.insertCell(-1);
    let cell2 = newRow.insertCell(-1);
    let cell3 = newRow.insertCell(-1);
    let cell4 = newRow.insertCell(-1);
    let impText = document.createTextNode(imp);
    let taskText = document.createTextNode(task);
    let limitText = document.createTextNode(limit);

    cell1.appendChild(impText);
    cell2.appendChild(taskText);
    cell3.appendChild(limitText);

    let dButton = document.createElement("button");
    dButton.type = "button";
    dButton.value = id;
    dButton.innerText = "削除";
    dButton.addEventListener("click", deleteMode, false);
    cell4.appendChild(dButton);
  });
}

function deleteMode() {
  let todoData = {
    id: parseInt(this.value, 10),
  };

  fetch("/todoDelete", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(todoData),
  }).then((response) => {
    display();
  });
  alert("削除しました");
}
