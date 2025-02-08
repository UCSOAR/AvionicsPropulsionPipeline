<script setup lang="ts">
import { computed, ref, inject } from 'vue';
import VueApexCharts from 'vue3-apexcharts';

// Inject dark mode from the parent component
const isDarkMode = inject('isDarkMode', ref(false));

// Example data for the chart with more entries
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

// Chart options with annotations
const chartOptions = computed(() => ({
  chart: {
    height: 350,
    type: 'line',
    id: 'areachart-2',
    background: 'transparent',
    toolbar: { show: false },
  },
  annotations: {
    yaxis: [
      {
        y: 8200,
        borderColor: '#00E396',
        label: {
          borderColor: '#00E396',
          style: {
            color: '#fff',
            background: '#00E396',
          },
          text: 'Support',
        },
      },
      {
        y: 8600,
        y2: 9000,
        borderColor: '#000',
        fillColor: '#FEB019',
        opacity: 0.2,
        label: {
          borderColor: '#333',
          style: {
            fontSize: '10px',
            color: '#333',
            background: '#FEB019',
          },
          text: 'Y-axis range',
        },
      },
    ],
    xaxis: [
      {
        x: new Date('2023-11-05').getTime(),
        strokeDashArray: 0,
        borderColor: '#775DD0',
        label: {
          borderColor: '#775DD0',
          style: {
            color: '#fff',
            background: '#775DD0',
          },
          text: 'Key Event',
        },
      },
      {
        x: new Date('2023-11-10').getTime(),
        x2: new Date('2023-11-15').getTime(),
        fillColor: '#B3F7CA',
        opacity: 0.4,
        label: {
          borderColor: '#B3F7CA',
          style: {
            fontSize: '10px',
            color: '#fff',
            background: '#00E396',
          },
          offsetY: -10,
          text: 'X-axis range',
        },
      },
    ],
    points: [
      {
        x: new Date('2023-11-07').getTime(),
        y: 8800,
        marker: {
          size: 8,
          fillColor: '#fff',
          strokeColor: 'red',
          radius: 2,
          cssClass: 'apexcharts-custom-class',
        },
        label: {
          borderColor: '#FF4560',
          offsetY: 0,
          style: {
            color: '#fff',
            background: '#FF4560',
          },
          text: 'Peak Point',
        },
      },
    ],
  },
  dataLabels: {
    enabled: false,
  },
  stroke: {
    curve: 'smooth',
    width: 2,
  },
  grid: {
    borderColor: isDarkMode.value ? '#555' : '#ddd',
    padding: { right: 30, left: 20 },
  },
  title: {
    text: '',
    align: 'left',
    style: {
      fontSize: '16px',
      fontWeight: 'bold',
      color: isDarkMode.value ? '#fff' : '#333',
    },
  },
  xaxis: {
    type: 'datetime',
    labels: {
      style: {
        colors: isDarkMode.value ? '#ccc' : '#333',
      },
    },
  },
  yaxis: {
    labels: {
      style: {
        colors: isDarkMode.value ? '#ccc' : '#333',
      },
    },
  },
  tooltip: {
    theme: isDarkMode.value ? 'dark' : 'light',
  },
}));
</script>

<template>
  <div :class="['chart-container', { dark: isDarkMode }]">
    <apexchart :options="chartOptions" :series="series" type="line" height="350" width="1000"></apexchart>
  </div>
</template>

<style scoped>
.chart-container {
  width: 96%;
  padding: 20px;
  background: var(--chart-bg, #fff);
  border-radius: 10px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  transition: background 0.3s ease, color 0.3s ease;
}

/* Dark Mode */
.chart-container.dark {
  background: #222;
  color: #fff;
}
</style>