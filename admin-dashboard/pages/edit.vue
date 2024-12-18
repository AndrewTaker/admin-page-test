<template>
  <div class="container">
    <h1>Редактирование данных</h1>
    <table class="data-table">
      <thead>
        <tr>
          <th>#</th>
          <th>Диапазон от</th>
          <th>Диапазон до</th>
          <th v-for="(col, index) in columns" :key="index">{{ col }}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(row, index) in rows" :key="index">
          <td>{{ index + 1 }}</td>
          <td><input v-model="row.range_from" type="number" class="input-field" /></td>
          <td><input v-model="row.range_to" type="number" class="input-field" /></td>
          <td v-for="(value, col) in row.days_columns" :key="col">
            <input v-model="row.days_columns[col]" type="number" class="input-field" />
          </td>
        </tr>
      </tbody>
    </table>

    <div class="buttons">
      <button @click="addRow" class="btn add-row">Добавить строку</button>
      <button @click="addColumn" class="btn add-column">Добавить колонку</button>
      <button @click="saveData" class="btn save-data">Сохранить</button>
    </div>

    <Notification v-if="notification.visible" :message="notification.message" :type="notification.type"
      :duration="3000" />
  </div>
</template>
<script>
import Notification from "@/components/Notification.vue";
export default {
  components: {
    Notification,
  },
  data() {
    return {
      rows: [],
      columns: [],
      notification: {
        visible: false,
        message: "",
        type: "success",
      },
    };
  },
  methods: {
    addRow() {
      this.rows.push({
        range_from: 0,
        range_to: 0,
        days_columns: Object.fromEntries(this.columns.map((col) => [col, 0])),
      });
    },
    addColumn() {
      const newColumn = prompt("Введите название новой колонки:");
      if (newColumn && newColumn.trim() !== "") {
        const trimmedColumn = newColumn.trim();
        if (this.columns.includes(trimmedColumn)) {
          alert("Имя колонки должно быть уникальным");
          return;
        }
        this.columns.push(trimmedColumn);
        this.rows.forEach((row) => {
          row.days_columns[trimmedColumn] = 0;
        });
      } else {
        alert("Имя колонки не может быть пустым");
      }
    },
    async saveData() {
      const config = useRuntimeConfig();
      console.log(config);
      console.log(config.public.apiUrl);
      try {
        const response = await fetch(`${config.public.apiUrl}/data`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(this.rows),
        });

        if (!response.ok) {
          throw new Error(`Error saving data: ${response.statusText}`);
        }

        this.showNotification("Данные успешно сохранены", "success");
      } catch (error) {
        this.showNotification("Ошибка при сохранении данных", "error");
        console.error("Ошибка при сохранении данных:", error);
      }
    },
    showNotification(message, type = "success") {
      this.notification.visible = true;
      this.notification.message = message;
      this.notification.type = type;

      setTimeout(() => {
        this.notification.visible = false;
      }, 3000);
    },
  },
  mounted() {
    this.addRow();
  },
};
</script>
