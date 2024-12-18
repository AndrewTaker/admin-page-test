<template>
  <div class="container">
    <h1>Сохраненные данные</h1>
    <table>

      <thead>
        <tr>
          <th scope="col" rowspan="2">#</th>
          <th scope="col" colspan="2">Диапазон</th>
          <th scope="col" :colspan="columns.length">Дни</th>
        </tr>
        <tr>
          <th scope="col">От</th>
          <th scope="col">До</th>
          <th scope="col" rowspan="2" v-for="header in columns" :key="header">{{ header }}</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="(row, index) in data" :key="index">
          <td>{{ row.id }}</td>
          <td>{{ row.range_from }}</td>
          <td>{{ row.range_to }}</td>
          <td v-for="header in columns" :key="header">
            {{ row.days_columns[header] || '-' }}
          </td>
        </tr>
      </tbody>

    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      data: [],
      columns: [],
    };
  },
  async mounted() {
    const config = useRuntimeConfig();
    try {
      const response = await fetch(`${config.public.apiUrl}/data`);
      const result = await response.json();
      this.data = result;
      if (this.data.length > 0) {
        this.columns = Array.from(
          new Set(this.data.flatMap(item => Object.keys(item.days_columns)))
        );
      }
    } catch (error) {
      console.error('Ошибка получения данных:', error);
    }
  },
};
</script>
