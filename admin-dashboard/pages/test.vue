<template>
  <div class="table-container">
    <table>
      <!-- Dynamic Header -->
      <thead>
        <tr>
          <th v-for="(header, index) in headers" :key="index">
            {{ header }}
          </th>
        </tr>
      </thead>

      <!-- Table Body -->
      <tbody>
        <!-- Static Rows -->
        <tr v-for="(row, rowIndex) in tableData" :key="rowIndex">
          <td v-for="(cell, cellIndex) in row" :key="cellIndex">
            <input v-model="tableData[rowIndex][cellIndex]"
              :placeholder="`Row ${rowIndex + 1}, Col ${cellIndex + 1}`" />
          </td>
        </tr>
      </tbody>
    </table>

    <!-- Add Column Button -->
    <button @click="addColumn">+ Add Column</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      headers: ["Header 1", "Header 2", "Header 3"], // Initial headers
      tableData: [
        ["Row 1 Col 1", "Row 1 Col 2", "Row 1 Col 3"],
        ["Row 2 Col 1", "Row 2 Col 2", "Row 2 Col 3"],
        ["Row 3 Col 1", "Row 3 Col 2", "Row 3 Col 3"],
      ], // Static data
    };
  },
  methods: {
    addColumn() {
      const newHeader = prompt("Enter the new column header name:");
      if (newHeader && newHeader.trim() !== "") {
        this.headers.push(newHeader.trim()); // Add the new header
        this.tableData.forEach((row) => row.push("")); // Extend each row with an empty cell
      } else {
        alert("Header name cannot be empty.");
      }
    },
  },
};
</script>

<style scoped>
.table-container {
  margin: 20px;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

input {
  width: 100%;
  border: none;
  outline: none;
}

button {
  margin: 10px;
  padding: 5px 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
