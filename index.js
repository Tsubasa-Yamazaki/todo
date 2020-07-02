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
      const todoData = {
        importance: this.importance,
        task: this.task,
        deadline: this.deadline,
      };
      fetch("/todos", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(todoData),
      }).then((response) => {
        this.display();
      }).catch(error => console.error('Error:', error));
      alert("登録しました");
    },
    //一覧用データ取得
    display: function () {
      fetch("/todos", {
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
        }).catch(error => console.error('Error:', error));
    },
    //削除ボタン
    deleteMode: function (id) {
      const todoData = {
        id: id,
      };
      fetch("/todos", {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(todoData),
      }).then((response) => {
        this.display();
      }).catch(error => console.error('Error:', error));
      alert("削除しました");
    },
  },
});
