<script setup lang="ts">
import { ref, computed, inject } from 'vue';
import { useMetadataStore } from '../stores/metadataStore';

// Inject dark mode (optional, depending on your implementation)
const isDarkMode = inject('isDarkMode', ref(false));

// Access the metadata store
const metadataStore = useMetadataStore();

// Computed properties for xColumnNames and yColumnNames
const xColumnNames = computed(() => metadataStore.metadata?.xColumnNames ?? []);
const yColumnNames = computed(() => metadataStore.metadata?.yColumnNames ?? []);

// Map xColumnNames and yColumnNames to dropdown options
const xOptions = computed(() =>
  xColumnNames.value.map((name) => ({ value: name, label: name }))
);

const yOptions = computed(() =>
  yColumnNames.value.map((name) => ({ value: name, label: name }))
);

// Default selected values (set to the first available option or empty string)
const selectedXValue = ref(xOptions.value[0]?.value ?? '');
const selectedYValue = ref(yOptions.value[0]?.value ?? '');

// Handle the confirm button click
const confirmOptions = () => {
  console.log('Confirmed Options:', {
    x: selectedXValue.value,
    y: selectedYValue.value,
  });

  // Add logic here to update your chart or fetch new data
};
</script>

<template>
  <div>
    <!-- Column Tray Above the Chart -->
    <div :class="['column-tray', { dark: isDarkMode }]">
      <!-- Dropdown for X Values -->
      <div class="dropdown-group">
        <label for="xValue">X Value:</label>
        <select id="xValue" v-model="selectedXValue">
          <option
            v-for="option in xOptions"
            :key="option.value"
            :value="option.value"
          >
            {{ option.label }}
          </option>
        </select>
      </div>

      <!-- Dropdown for Y Values -->
      <div class="dropdown-group">
        <label for="yValue">Y Value:</label>
        <select id="yValue" v-model="selectedYValue">
          <option
            v-for="option in yOptions"
            :key="option.value"
            :value="option.value"
          >
            {{ option.label }}
          </option>
        </select>
      </div>

      <!-- Confirm Button -->
      <button class="confirm-button" @click="confirmOptions">Confirm</button>
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
