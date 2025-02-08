<script setup lang="ts">
import { computed, ref, inject } from 'vue';
import { useMetadataStore } from '../stores/metadataStore';
import VueApexCharts from 'vue3-apexcharts';

// Inject dark mode from the parent component
const isDarkMode = inject('isDarkMode', ref(false));
const metadataStore = useMetadataStore();

console.log(JSON.stringify(metadataStore.metadata, null, 2));



// --- Column Tray State and Options ---
const selectedXValue = ref('');
const selectedYValue = ref('');

// Sample dropdown options (adjust these to match your needs)
const xOptions = ref([
  { value: 'date', label: 'Date' },
  { value: 'time', label: 'Time' },
  { value: 'category', label: 'Category' },
]);

const yOptions = ref([
  { value: 'value', label: 'Value' },
  { value: 'count', label: 'Count' },
  { value: 'total', label: 'Total' },
]);

// This function will be called when the Confirm button is clicked.
const confirmOptions = () => {
  console.log('Confirmed Options:', {
    x: selectedXValue.value,
    y: selectedYValue.value,
  });
  // Add logic here to update your chart options, fetch new data, etc.
};

// --- Chart Data and Options ---
const series = ref([
  {
    name: 'Example Data',
    data: [
      [new Date('2023-11-01').getTime(), 8100],
      [new Date('2023-11-02').getTime(), 8200],
      [new Date('2023-11-03').getTime(), 8300],
      [new Date('2023-11-04').getTime(), 8500],
      [new Date('2023-11-05').getTime(), 8600],
      [new Date('2023-11-06').getTime(), 8700],
      [new Date('2023-11-07').getTime(), 8800],
      [new Date('2023-11-08').getTime(), 8900],
      [new Date('2023-11-09').getTime(), 9000],
      [new Date('2023-11-10').getTime(), 9100],
      [new Date('2023-11-11').getTime(), 9200],
      [new Date('2023-11-12').getTime(), 9300],
      [new Date('2023-11-13').getTime(), 9400],
      [new Date('2023-11-14').getTime(), 9500],
      [new Date('2023-11-15').getTime(), 9600],
      [new Date('2023-11-16').getTime(), 9700],
      [new Date('2023-11-17').getTime(), 9800],
    ],
  },
]);


</script>

<template>
  <div>
    <!-- Column Tray Above the Chart -->
    <div :class="['column-tray', { dark: isDarkMode }]">
      <div class="dropdown-group">
        <label for="xValue">X Value:</label>
        <select id="xValue" v-model="selectedXValue">
          <option v-for="option in xOptions" :key="option.value" :value="option.value">
            {{ option.label }}
          </option>
        </select>
      </div>
      <div class="dropdown-group">
        <label for="yValue">Y Value:</label>
        <select id="yValue" v-model="selectedYValue">
          <option v-for="option in yOptions" :key="option.value" :value="option.value">
            {{ option.label }}
          </option>
        </select>
      </div>
      <button class="confirm-button" @click="confirmOptions">
        Confirm
      </button>
    </div>

  </div>
</template>

<style scoped>
/* Column Tray Styling */
.column-tray {
  display: flex;
  justify-content: space-around;
  align-items: center;
  gap: 1rem;
  padding: 10px;
  background: #f9f9f9;
  border: 2px solid #007bff;
  border-radius: 8px;
  margin-bottom: 10px;
}

.column-tray.dark {
  background: #333;
  border: 2px solid #555;
}

/* Dropdown Group */
.dropdown-group {
  display: flex;
  flex-direction: row;
}

.dropdown-group label {
  font-weight: 500;
  padding-top: 5px;
  padding-right: 5px;
  font-size: 0.9rem;
  margin-bottom: 4px;
  color: #333;
}

.column-tray.dark .dropdown-group label {
  color: #ccc;
}

.dropdown-group select {
  padding: 5px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.column-tray.dark .dropdown-group select {
  background: #555;
  border: 1px solid #777;
  color: #fff;
}

/* Confirm Button */
.confirm-button {
  background: #007bff;
  color: #fff;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  transition: background 0.3s ease;
}

.confirm-button:hover {
  background: #005dc1;
}

.column-tray.dark .confirm-button {
  background: #555;
}

.column-tray.dark .confirm-button:hover {
  background: #777;
}
</style>