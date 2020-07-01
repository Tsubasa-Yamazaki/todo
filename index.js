const app = new Vue({
  el: "#app",
  data: {
    id: "",
    importance: "",
    task: "",
    deadline: "",
  },

  created: function () {
    this.display();
  },

  methods: {
    entry: function () {
      let todoData = {
        importance: this.importance,
        task: this.task,
        deadline: this.deadline,
      };

      fetch("/post", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(todoData),
      }).then((response) => {
        this.display();
      });
      alert("登録しました");
    },

    //一覧用データ取得
    display: function () {
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
          this.tableDelete();
          this.displayTable(todoData);
        });
    },

    //画面テーブル削除
    tableDelete: function () {
      let table = document.getElementById(`table`);
      let loopCount = table.rows.length;

      for (let i = 0; i < loopCount; i++) {
        if (i === 0) {
          continue;
        }
        table.deleteRow(1);
      }
    },

    //画面テーブル生成
    displayTable: function (todoData) {
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
        dButton.addEventListener("@click", this.deleteMode, false);
        cell4.appendChild(dButton);
      });
    },

    //削除ボタン
    deleteMode: function () {
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
        this.display();
      });
      alert("削除しました");
    },
  },
});
