const app = new Vue({
  el: "#app",
  data: {
    id: "",
    importance: "",
    task: "",
    deadline: "",

    todos: []
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
          this.todos = todoData;
        });
    },

    //削除ボタン
    deleteMode: function (id) {
      let todoData = {
        id: id,
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
