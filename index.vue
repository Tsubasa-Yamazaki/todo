<template>
  <div>
    <h1>ToDo</h1>

    <v-data-table :headers="headers" :items="todoData" :items-per-page="5" item-key="id">
      <template v-slot:item.delete="{ item }">
        <v-btn small @click="deleteMode(item.id)">タスク削除</v-btn>
      </template>
    </v-data-table>

    <v-form>
      <div>
        重要度
        <v-radio-group v-model="todo.importance" :mandatory="true">
          <v-radio label="高い" value="重"></v-radio>
          <v-radio label="普通" value="普"></v-radio>
          <v-radio label="低い" value="低"></v-radio>
        </v-radio-group>
      </div>
      <v-col>
        <v-text-field v-model="todo.task" :counter="255" label="タスク" required></v-text-field>
      </v-col>
      <v-col>
        <v-text-field v-model="todo.deadline" :counter="255" label="期限" required></v-text-field>
      </v-col>
      <v-btn @click="entry">登録</v-btn>
    </v-form>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      headers: [
        {
          text: "重要度",
          value: "importance"
        },
        {
          text: "タスク",
          value: "task"
        },
        {
          text: "期限",
          value: "deadline"
        },
        {
          text: "削除",
          value: "delete",
          sortable: false
        }
      ],
      todo: {},
      todoData: []
    };
  },

  created: function() {
    this.display();
  },
  methods: {
    entry: function() {
      axios
        .post("/todos", {
          importance: this.todo.importance,
          task: this.todo.task,
          deadline: this.todo.deadline
        })
        .then(response => {
          this.display();
        })
        .catch(err => {
          console.log(err);
        });
    },
    //一覧用データ取得
    display: function() {
      axios
        .get("/todos")
        .then(response => {
          this.todoData = response.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    //削除ボタン
    deleteMode: function(id) {
      axios
        .delete("/todos", {
          data: { id: id }
        })
        .then(response => {
          this.display();
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>

<style>
</style>