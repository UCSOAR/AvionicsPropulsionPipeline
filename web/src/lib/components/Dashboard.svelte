<script lang="ts">
  import { onMount } from "svelte";
  import type { PreviewMetadata } from "$lib/models/cacheTreeModels";
  import type { Config, Data, Layout } from "plotly.js";
  import Dropdown from "./Dropdown.svelte";

  type SelectedFile = {
    name: string;
    metadata: PreviewMetadata;
  };

  export let selectedFile: SelectedFile | null = null;

  let selectedXColumnIndex = 0;
  let selectedYColumnIndex = 0;

  let data: Data[] = [];

  const style = {
    margin: 50,
    bgColor: "#1f1f1f",
    txtColor: "#e1e1e1",
  };
  const config: Partial<Config> = { responsive: true };
  const layout: Partial<Layout> = {
    autosize: true,
    margin: {
      l: style.margin,
      r: style.margin,
      t: style.margin,
      b: style.margin,
    },
    paper_bgcolor: style.bgColor,
    plot_bgcolor: style.bgColor,
    font: {
      family: "Inter",
      color: "white",
    },
    xaxis: { color: style.txtColor },
    yaxis: { color: style.txtColor },
  };

  let plotlyChartDiv: HTMLDivElement;

  onMount(async () => {
    const Plotly = await import("plotly.js-dist-min");
    Plotly.newPlot(plotlyChartDiv, data, layout, config);
  });
</script>

<div class="container">
  {#if selectedFile}
    <div class="content-header">
      <div class="title">
        <h1>Dashboard for <i>{selectedFile.name}</i></h1>
        <p>
          Visualizing data for <i
            >{selectedFile.metadata.xColumnNames[selectedXColumnIndex]}</i
          >
          and
          <i>{selectedFile.metadata.yColumnNames[selectedYColumnIndex]}</i>
        </p>
      </div>
      <div class="column-select">
        <Dropdown
          onChange={(index) => (selectedXColumnIndex = index)}
          label="X Column"
          id="x-column"
          options={selectedFile.metadata.xColumnNames}
        />
        <Dropdown
          onChange={(index) => (selectedYColumnIndex = index)}
          label="Y Column"
          id="y-column"
          options={selectedFile.metadata.yColumnNames}
        />
      </div>
    </div>
    <div class="content-container">
      <div class="chart-pod pod">
        <h2>Static Fire Chart</h2>
        <div class="chart-wrapper">
          <div bind:this={plotlyChartDiv} class="chart"></div>
        </div>
      </div>
      <div class="value-pods">
        <div class="min-val-pod pod">
          <label for="min-val">Minimum Value</label>
          <div class="value" id="min-val">0.00</div>
        </div>
        <div class="max-val-pod pod">
          <label for="max-val">Maximum Value</label>
          <div class="value" id="max-val">0.00</div>
        </div>
        <div class="avg-val-pod pod">
          <label for="avg-val">Average Value</label>
          <div class="value" id="avg-val">0.00</div>
        </div>
      </div>
    </div>
  {:else}
    <h1>Dashboard</h1>
    <div class="message-container">
      <p>Select a file from the sidebar to view its data.</p>
    </div>
  {/if}
</div>

<style scoped lang="scss">
  @use "../styles/variables.scss" as *;

  div.content-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex-direction: row;
    gap: 1rem;

    div.column-select {
      display: flex;
      flex-direction: column;
      gap: 1.5rem;
      width: 25%;
      min-width: 9rem;
      max-width: 13rem;
    }
  }

  div.pod {
    background-color: $bg-color-4;
    border: 1px solid $outline-color-1;
    border-radius: $border-radius-1;
    padding: 1rem;

    h2 {
      margin: 0;
      margin-bottom: 0.5em;
    }

    div.value {
      margin-top: 0.5rem;
      font-size: 1.5rem;
      font-weight: bold;
    }
  }

  div.content-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    padding-bottom: 1rem;

    div.value-pods {
      display: flex;
      flex-direction: row;
      gap: 1rem;

      div.pod {
        flex-grow: 1;
      }
    }

    div.chart-pod {
      flex-grow: 1;

      div.chart-wrapper {
        border-radius: $border-radius-1;
        overflow: hidden;
      }
    }
  }

  div.container {
    padding: 1rem;
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    gap: 1rem;

    h1 {
      margin: 0;
      margin-bottom: 0.3rem;

      i {
        color: $txt-color-1;
      }
    }

    div.message-container {
      display: flex;
      justify-content: center;
      align-items: center;
      flex-grow: 1;
      text-align: center;
    }
  }
</style>
