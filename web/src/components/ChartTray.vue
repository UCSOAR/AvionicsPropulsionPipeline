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
const chartContainer = ref<HTMLDivElement | null>(null);

// This will hold the full data set from your API call
const allData = ref<{ x: number[]; y: number[] }>({ x: [], y: [] });

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
      numRows: 10000, // FOR NOW...
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
    console.log("This is the fetched data: ", data);

    // Validate data structure
    if (data.xColumns && data.yColumns && data.xColumns[colX] && data.yColumns[colY]) {
      const xValues: number[] = data.xColumns[colX].rows;
      const yValues: number[] = data.yColumns[colY].rows;


      if (xValues.length !== yValues.length) {
        throw new Error('Mismatch in lengths of xColumns and yColumns data.');
      }

      // Combine and sort data by x-value
      let chartData: [number, number][] = xValues.map((x, i) => [x, yValues[i]]);

      // Update allData with sorted arrays
      allData.value.x = chartData.map(([x]) => x);
      allData.value.y = chartData.map(([_, y]) => y);

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
    line: { 
      color: '#da0000', // Blue line that works in both themes
      width: 2 
    }
  };

  // Set light/dark mode colors directly in the layout
  const isDark = isDarkMode.value; 
  const layout: Partial<Plotly.Layout> = {
    autosize: true,
    margin: { l: 50, r: 50, t: 30, b: 50 }, // Adjust margins to fit better
    dragmode: 'pan',
    xaxis: {
      range: initialRange,
      title: 'X Axis',
      color: isDark ? '#ccc' : '#333',
      gridcolor: isDark ? '#333' : '#eee',
      zerolinecolor: isDark ? '#444' : '#ddd'
    },
    yaxis: {
      title: 'Y Axis',
      color: isDark ? '#ccc' : '#333',
      gridcolor: isDark ? '#333' : '#eee',
      zerolinecolor: isDark ? '#444' : '#ddd'
    },
    plot_bgcolor: isDark ? '#222' : '#fff',
    paper_bgcolor: isDark ? '#222' : '#fff',
    font: {
      color: isDark ? '#ccc' : '#333'
    }
  };

  const config = {
    responsive: true,
    displayModeBar: true, // Re-enable the mode bar
    modeBarButtonsToRemove: ['lasso2d', 'select2d'], // Remove less commonly used buttons
    displaylogo: false // Remove the plotly logo
  };

  // Destroy existing chart before rendering a new one
  Plotly.purge(chartContainer.value);

  // Create or update the Plotly chart
  Plotly.newPlot(chartContainer.value, [trace], layout, config);
}

// Re-fetch and re-render the chart when metadata changes
watch(
  () => [metadataStore.metadata, metadataStore.colX, metadataStore.colY, isDarkMode.value],
  () => {
    fetchChartData();
  },
  { immediate: true }
);

onMounted(() => {
  fetchChartData();
});
</script>

<style scoped lang="scss">
@import '../styles/variables.scss';


.chart-container {
  width: 100%;
  height: 350px;
  padding: 0;
  border-radius: 10px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  border: 2px solid $soar-red-color;
  transition: background 0.3s ease, color 0.3s ease;
  background-color: #fff;
  overflow: hidden;
  position: relative;
}

.chart-container.dark {
  background-color: #222;
  color: #fff;
  border: 2px solid #444;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.3);
}

/* Make sure all Plotly elements take full width/height */
.chart-container :deep(.js-plotly-plot),
.chart-container :deep(.plot-container),
.chart-container :deep(.svg-container) {
  width: 100% !important;
  height: 100% !important;
}

/* Style the mode bar to better match the theme */
.chart-container.dark :deep(.modebar) {
  background: rgba(30, 30, 30, 0.7) !important;
}

.chart-container.dark :deep(.modebar-btn path) {
  fill: #ccc !important;
}
</style>