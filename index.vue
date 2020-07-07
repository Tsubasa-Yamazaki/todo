<template>
  <div>
    <h1>ToDo</h1>

    <v-data-table :headers="headers" :items="todoData" :items-per-page="5" item-key="id">
      <template v-slot:item.importance="{ item }">
        <v-chip :color="getColor(item.importance)" dark></v-chip>
      </template>

      <template v-slot:item.delete="{ item }">
        <v-btn small @click="deleteTodo(item.id)">タスク削除</v-btn>
      </template>
    </v-data-table>

    <v-form>
      <div class="importance">
        <v-chip class="ma-2" color="red">重要度</v-chip>
        <v-radio-group v-model="todo.importance" :mandatory="true">
          <v-radio label="高い" value="high"></v-radio>
          <v-radio label="普通" value="middle"></v-radio>
          <v-radio label="低い" value="low"></v-radio>
        </v-radio-group>
      </div>
      <div class="text">
        <v-col cols="6">
          <v-text-field v-model="todo.task" :counter="255" label="タスク入力" outlined shaped required></v-text-field>
        </v-col>
        <v-menu>
          <template v-slot:activator="{on}">
            <v-btn color="primary" dark v-on="on">
              <v-icon>mdi-calendar</v-icon>カレンダーから日付を入力する
            </v-btn>
          </template>
          <v-date-picker
            locale="ja"
            :day-format="date => new Date(date).getDate()"
            v-model="todo.deadline"
          />
        </v-menu>

        <v-col cols="6">
          <v-text-field
            v-model="todo.deadline"
            :counter="255"
            label="期限入力"
            outlined
            shaped
            required
          ></v-text-field>
        </v-col>
        <v-btn @click="createTodo">登録</v-btn>
      </div>
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
    this.createTodo();
  },
  methods: {
    createTodo() {
      axios
        .post("/todos", this.todo)
        .then(response => {
          this.getTodos();
        })
        .catch(err => {
          console.log(err);
        });
    },
    // 一覧用データ取得
    getTodos() {
      axios
        .get("/todos")
        .then(response => {
          this.todoData = response.data;
        })
        .catch(err => {
          console.log(err);
        });
    },
    // 削除ボタン
    deleteTodo(id) {
      axios
        .delete("/todos", {
          params: { id: id }
        })
        .then(response => {
          this.getTodos();
        })
        .catch(err => {
          console.log(err);
        });
    },
    // 重要度カラー取得
    getColor(importance) {
      if (importance === "high") return "red";
      else if (importance === "middle") return "orange";
      else return "green";
    }
  }
};
</script>

<style>
.text {
  margin: 0;
}

.importance {
  padding-top: 1em;
  padding-left: 1em;
}
</style>