<script setup lang="ts">
import { computed, ref, inject, watch } from "vue";
import { useMetadataStore } from "../stores/metadataStore";
import { endpointMapping } from "../utils/constants";

// Inject dark mode from the parent component
const isDarkMode = inject("isDarkMode", ref(false));
const metadataStore = useMetadataStore();

// This will hold our fetched data for the chart
const series = ref<Array<{ name: string; data: [number, number][] }>>([]);

// Chart options with annotations
const chartOptions = computed(() => ({
  chart: {
    height: 350,
    type: "line",
    background: "transparent",
    toolbar: { show: false },
  },
  annotations: {
    yaxis: [
      {
        y: 8200,
        borderColor: "#00E396",
        label: {
          borderColor: "#00E396",
          style: {
            color: "#fff",
            background: "#00E396",
          },
          text: "Support",
        },
      },
    ],
  },
  dataLabels: {
    enabled: false,
  },
  stroke: {
    curve: "smooth",
    width: 2,
  },
  grid: {
    borderColor: isDarkMode.value ? "#555" : "#ddd",
    padding: { right: 30, left: 20 },
  },
  xaxis: {
    type: "datetime",
    labels: {
      style: {
        colors: isDarkMode.value ? "#ccc" : "#333",
      },
    },
  },
  yaxis: {
    labels: {
      style: {
        colors: isDarkMode.value ? "#ccc" : "#333",
      },
    },
  },
  tooltip: {
    theme: isDarkMode.value ? "dark" : "light",
  },
}));

// Fetch chart data via POST request
async function fetchChartData() {
  const name = metadataStore.name;
  const colX = metadataStore.colX;
  const colY = metadataStore.colY;

  // Check if file name and columns are set
  if (!name || !colX || !colY) {
    console.warn("No file name / colX / colY selected yet.");
    return;
  }

  try {
    const payload = {
      name: name,
      xColumnNames: [colX],
      yColumnNames: [colY],
    };

    console.log("This is the payload", payload);

    const result = await fetch(endpointMapping.getStaticFireColumnsUrl.toString(), {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    });

    if (!result.ok) {
      throw new Error(`Failed to fetch columns. Status: ${result.status} ${result.statusText}`);
    }

    const data = await result.json();
    console.log("This is the fetched data", data);

    // Ensure the fetched data structure matches the expected format
    if (data.xColumns && data.yColumns && data.xColumns[colX] && data.yColumns[colY]) {
      const xValues = data.xColumns[colX].rows;
      const yValues = data.yColumns[colY].rows;

      if (xValues.length !== yValues.length) {
        throw new Error("Mismatch in lengths of xColumns and yColumns data.");
      }

      // Map data to the format expected by ApexCharts
      const chartData = xValues.map((x: number, index: number) => [x, yValues[index]]);
      series.value = [
        {
          name: `${name} - (${colX} vs. ${colY})`,
          data: chartData,
        },
      ];
      console.log("Fetched chart data:", series.value);
    } else {
      throw new Error("Invalid data structure received from API.");
    }
  } catch (err) {
    console.error("Error fetching chart data:", err);
  }
}

// Watch for changes in metadata and columns
watch(
  () => [metadataStore.metadata, metadataStore.colX, metadataStore.colY],
  () => {
    void fetchChartData();
  },
  { immediate: true }
);
</script>

<template>
  <div :class="['chart-container', { dark: isDarkMode }]">
    <apexchart :options="chartOptions" :series="series" type="line" height="350" width="1000" />
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