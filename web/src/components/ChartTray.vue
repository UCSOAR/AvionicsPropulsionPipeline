<template>
  <!-- The chart container (dark mode classes applied dynamically) -->
  <div :class="['chart-container', { dark: isDarkMode }]" ref="chartContainer"></div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, inject } from 'vue';
import Plotly, { Data } from 'plotly.js-dist-min';
import { useMetadataStore } from '../stores/metadataStore';
import { endpointMapping } from '../utils/constants';

// Inject dark mode from parent (defaulting to false if not provided)
const isDarkMode = inject('isDarkMode', false);

// Access the metadata store (assumed to be set up already)
const metadataStore = useMetadataStore();

// A reference to the chart container DOM element
const chartContainer = ref<Plotly.PlotlyHTMLElement | null>(null);

// This will hold the full data set from your API call
const allData = ref<{ x: number[]; y: number[] }>({ x: [], y: [] });

// Track the start and end of the visible range
let xStart = 0;
let xEnd = 0;

/**
 * Fetch dynamic chart data using a POST request.
 * Expects the API to return a structure with:
 *   data.xColumns[0].rows and data.yColumns[0].rows
 */
async function fetchChartData() {
  const name = metadataStore.name;
  const colX = metadataStore.colX;
  const colY = metadataStore.colY;

  if (!name || !colX || !colY) {
    console.warn('No file name / colX / colY selected yet.');
    return;
  }

  try {
    const payload = {
      name: name,
      startRow: 0, // FOR NOW...
      numRows: 15000, // FOR NOW...
      xColumnNames: [colX],
      yColumnNames: [colY],
    };

    const response = await fetch(endpointMapping.getStaticFireColumnsUrl.toString(), {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    });

    if (!response.ok) {
      throw new Error(`Failed to fetch columns. Status: ${response.status}`);
    }

    const data = await response.json();

    // Validate data structure
    if (data.xColumns && data.yColumns && data.xColumns[colX] && data.yColumns[colY]) {
      const xValues: number[] = data.xColumns[colX].rows;
      const yValues: number[] = data.yColumns[colY].rows;

      if (xValues.length !== yValues.length) {
        throw new Error('Mismatch in lengths of xColumns and yColumns data.');
      }

      // Update allData
      allData.value.x = xValues;
      allData.value.y = yValues;

      // Render the chart with the updated data
      renderChart();
    } else {
      throw new Error('Invalid data structure received from API.');
    }
  } catch (err) {
    console.error('Error fetching chart data:', err);
  }
}

/**
 * Renders (or re-renders) the Plotly chart.
 * The initial view shows only the last 20 data points.
 */
function renderChart() {
  if (!chartContainer.value) return;

  const xData = allData.value.x;
  const yData = allData.value.y;

  if (xData.length === 0 || yData.length === 0) return;

  const totalPoints = xData.length;
  let initialRange: [number, number];

  if (totalPoints >= 20) {
    initialRange = [xData[totalPoints - 20], xData[totalPoints - 1]];
  } else {
    initialRange = [xData[0], xData[totalPoints - 1]];
  }

  const trace: Data = {
    x: xData,
    y: yData,
    mode: 'lines',
    type: 'scattergl', // Use WebGL for better performance with large datasets
    marker: { size: 6 },
  };

  const layout: Partial<Plotly.Layout> = {
    dragmode: 'pan', // Default drag mode enables panning with the mouse
    xaxis: {
      range: initialRange,
      title: 'X Axis'
    },
    yaxis: {
      title: 'Y Axis'
    },
    // Adapt background and font colors based on dark mode
    plot_bgcolor: isDarkMode ? '#222' : '#fff',
    paper_bgcolor: isDarkMode ? '#222' : '#fff',
    font: {
      color: isDarkMode ? '#ccc' : '#333'
    }
  };

  // Create or update the Plotly chart
  Plotly.newPlot(chartContainer.value, [trace], layout, { responsive: true });

  // Update the visible range when the user pans the chart
  chartContainer.value.on('plotly_relayout', (eventData: any) => {
    const rangeStart = eventData["xaxis.range[0]"];
    const rangeEnd = eventData["xaxis.range[1]"];

    // Update the visible range
    if (rangeStart && rangeEnd) {
      xStart = rangeStart;
      xEnd = rangeEnd;
    }
  });
}

// Re-fetch and re-render the chart when metadata changes
watch(
  () => [metadataStore.metadata, metadataStore.colX, metadataStore.colY],
  () => fetchChartData(),
  { immediate: true }
);

onMounted(() => {
  fetchChartData();
});
</script>